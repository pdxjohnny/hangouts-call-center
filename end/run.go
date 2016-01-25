package end

import (
	"log"

	"github.com/pdxjohnny/hangouts-call-center/api"
	"github.com/spf13/viper"
)

// Run requests a call from the api server
func Run() {
	tokenData, err := api.Login(
		// Host we are loging in to
		viper.GetString("host"),
		// No token needed to login
		"",
		// Username and password to authenticate
		api.LoginData{
			Username: viper.GetString("username"),
			Password: viper.GetString("password"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	token, ok := (*tokenData)["Token"].(string)
	if ok != true {
		log.Fatal("TokenData had no token")
	}

	api.End(
		// The call center host
		viper.GetString("host"),
		// Our login token
		token,
		// The lock to end a call
		viper.GetString("lock"),
	)
}
