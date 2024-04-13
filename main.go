package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/yudgnahk/vfs-account-manager/configs"
	"github.com/yudgnahk/vfs-account-manager/vfs"
)

func main() {
	configs.GetAllAccountPasswords()

	mux := http.NewServeMux()
	mux.HandleFunc("/", GetUserData)

	// Start the HTTP server
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
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

			c := vfs.NewClient("")
			sess, err := c.GetNewSession(id, configs.GetAccountPassword(id))
			if err != nil {
				http.Error(w, "Error getting new session", http.StatusInternalServerError)
				return
			}

			// save the session to the configs
			configs.AddAccountSession(id, sess)
			session = sess
		}

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
		//w.Write([]byte(session))
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
