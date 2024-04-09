package configs

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"os"
	"strings"
)

type Account struct {
	Session string
}

type Config struct {
	ChromeDriverPath string
	AccountsString   string `envconfig:"ACCOUNTS"`
	Accounts         []string
	AccountSessions  map[string]string
	AccountPasswords map[string]string
	BasicAuthUser    string `envconfig:"BASIC_AUTH_USER"`
	BasicAuthPass    string `envconfig:"BASIC_AUTH_PASS"`
}

var AppConfig = Config{}

func NewConfig() error {
	if err := envconfig.Process("", &AppConfig); err != nil {
		return err
	}

	AppConfig.Accounts = strings.Split(AppConfig.AccountsString, ",")
	AppConfig.AccountSessions = make(map[string]string)
	AppConfig.AccountPasswords = make(map[string]string)

	return nil
}

func SetChromeDriverPath(path string) {
	AppConfig.ChromeDriverPath = path
}

func AddAccountSession(account string, session string) {
	AppConfig.AccountSessions[account] = session
}

func GetAccountSession(account string) string {
	return AppConfig.AccountSessions[account]
}

func GetAccountPassword(account string) string {
	return AppConfig.AccountPasswords[account]
}

func GetAllAccountPasswords() {
	for _, account := range AppConfig.Accounts {
		AppConfig.AccountPasswords[account], _ = os.LookupEnv(fmt.Sprintf("PASSWORD_%s", strings.ToUpper(account)))
	}
}
