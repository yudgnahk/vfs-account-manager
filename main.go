package main

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yudgnahk/vfs-account-manager/chrome"
	"github.com/yudgnahk/vfs-account-manager/configs"
	"github.com/yudgnahk/vfs-account-manager/vfs"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	configs.GetAllAccountPasswords()
	path, err := downloadChromeDriver()
	if err != nil {
		fmt.Println("Error downloading ChromeDriver:", err)
		return
	}

	configs.SetChromeDriverPath(path)

	mux := http.NewServeMux()
	mux.HandleFunc("/", GetUserData)

	// Start the HTTP server
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// return path of chromedriver
func downloadChromeDriver() (string, error) {
	macURL := "https://storage.googleapis.com/chrome-for-testing-public/123.0.6312.86/mac-arm64/chromedriver-mac-arm64.zip"
	linuxUrl := "https://storage.googleapis.com/chrome-for-testing-public/123.0.6312.86/linux64/chromedriver-linux64.zip"
	url := linuxUrl

	if runtime.GOOS == "darwin" {
		url = macURL
	}

	// Destination file path to save the downloaded ZIP file
	zipFilePath := "chromedriver.zip"

	// Download the ZIP file
	if err := downloadFile(url, zipFilePath); err != nil {
		fmt.Println("Error downloading file:", err)
		return "", err
	}

	// Unzip the downloaded file
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	unzipPath, err := unzip(zipFilePath, pwd, pwd)
	if err != nil {
		fmt.Println("Error unzipping file:", err)
		return "", err
	}

	fmt.Println("ChromeDriver downloaded and unzipped successfully.")

	return fmt.Sprintf("%s/chromedriver", unzipPath), nil
}

// Function to download a file from a URL
func downloadFile(url, filePath string) error {
	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func unzip(zipFilePath, dst, dest string) (string, error) {
	archive, err := zip.OpenReader(zipFilePath)
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	var unzipFolder string

	for _, f := range archive.File {
		filePath := filepath.Join(dst, f.Name)

		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			fmt.Println("invalid file path")
			return "", errors.New("invalid file path")
		}
		if f.FileInfo().IsDir() {
			fmt.Println("creating directory...")
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return "", err
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return "", err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			return "", err
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return "", err
		}

		dstFile.Close()
		fileInArchive.Close()

		// get dir of dstFile
		unzipFolder = filepath.Dir(filePath)
	}

	//	return unzip folder
	return unzipFolder, nil
}

func init() {
	//	load configs
	fmt.Println("Loading configs...")
	err := configs.NewConfig()
	if err != nil {
		panic(err)
	}
}

func GetUserData(w http.ResponseWriter, r *http.Request) {
	// get the username and password from the request
	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// check if the username and password are correct
	if username == configs.AppConfig.BasicAuthUser && password == configs.AppConfig.BasicAuthPass {
		//	get :id from the request
		id := strings.TrimPrefix(r.URL.Path, "/")
		if len(id) == 0 {
			http.Error(w, "Missing id", http.StatusBadRequest)
			return
		}

		//	check if id is in the list of accounts
		found := false
		for _, account := range configs.AppConfig.Accounts {
			if account == id {
				found = true
				break
			}
		}

		if !found {
			http.Error(w, "Account not found", http.StatusNotFound)
			return
		}

		// get session from the configs.AppConfig.AccountSessions
		session := configs.GetAccountSession(id)
		if len(session) == 0 {
			// check if the password exists
			if len(configs.GetAccountPassword(id)) == 0 {
				http.Error(w, "Password not found", http.StatusNotFound)
				return
			}

			// go get session from the VFS site
			// create a new client
			client := chrome.NewChromeBrowser(configs.AppConfig.ChromeDriverPath, 4444)
			sess, err := client.LoginVFS(id, configs.GetAccountPassword(id))
			if err != nil {
				fmt.Println("Error logging in:", err)
				http.Error(w, "Error logging in", http.StatusInternalServerError)
				return
			}

			// save the session to the configs
			configs.AddAccountSession(id, sess)
			session = sess
		}

		//	get portfolio
		//	get today transaction
		//	return the response
		client := vfs.NewClient(session)
		portfolio, err := client.GetPortfolio()
		if err != nil {
			http.Error(w, "Error getting portfolio", http.StatusInternalServerError)
			return
		}

		todayTransaction, err := client.GetTodayTransaction(id)
		if err != nil {
			http.Error(w, "Error getting today transaction", http.StatusInternalServerError)
			return
		}

		response := struct {
			Portfolio        *vfs.GetPortfolioResponse        `json:"portfolio"`
			TodayTransaction *vfs.GetTodayTransactionResponse `json:"todayTransaction"`
		}{
			Portfolio:        portfolio,
			TodayTransaction: todayTransaction,
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Error marshaling JSON response", http.StatusInternalServerError)
			return
		}

		// Set content type header
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the response body
		_, err = w.Write(jsonResponse)
		if err != nil {
			http.Error(w, "Error writing JSON response", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
