package chrome

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type ChromeBrowser struct {
	// ChromeDriverPath is the path to the chromedriver executable
	ChromeDriverPath string
	// Port is the port number to start the ChromeDriver server
	Port int
}

// NewChromeBrowser creates a new instance of ChromeBrowser
func NewChromeBrowser(path string, port int) *ChromeBrowser {
	return &ChromeBrowser{
		ChromeDriverPath: path,
		Port:             port,
	}
}

// LoginVFS logs into the VFS site and returns the session cookie
func (b *ChromeBrowser) LoginVFS(username, password string) (string, error) {
	service, err := selenium.NewChromeDriverService(b.ChromeDriverPath, b.Port)
	if err != nil {
		fmt.Println("Error creating ChromeDriver service:", err)
		return "", err
	}
	defer service.Stop()

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{"--headless"}})

	// create a new remote client with the specified options
	driver, err := selenium.NewRemote(caps, "")

	if err != nil {
		return "", err
	}

	// maximize the current window to avoid responsive rendering
	err = driver.MaximizeWindow("")
	if err != nil {
		return "", err
	}

	err = driver.Get("https://trading.vfs.com.vn/login#portfolio")
	if err != nil {
		return "", err
	}

	// wait for the page to load
	time.Sleep(5 * time.Second)

	// find the username input field and type in the username
	usernameField, err := driver.FindElement(selenium.ByName, "ot_username")
	if err != nil {
		return "", err
	}
	usernameField.SendKeys(username)

	// find the password input field and type in the password
	passwordField, err := driver.FindElement(selenium.ByName, "ot_password")
	if err != nil {
		return "", err
	}

	passwordField.SendKeys(password)

	// find the login button and click it
	loginButton, err := driver.FindElement(selenium.ByCSSSelector, "#login_submit")
	if err != nil {
		return "", err
	}
	loginButton.Click()

	// get the cookies
	cookies, err := driver.GetCookies()
	if err != nil {
		log.Fatal("Get cookies error:", err)
	}

	// extract the session cookie
	var sessionCookie string
	for _, cookie := range cookies {
		if cookie.Name == "SESSION" {
			sessionCookie = cookie.Value
			break
		}
	}

	return sessionCookie, nil
}
