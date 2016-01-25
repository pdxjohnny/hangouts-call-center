package end

import (
	"fmt"
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

	token, ok := (*tokenData)["token"].(string)
	if ok != true {
		fmt.Println(tokenData)
		log.Fatal("TokenData had no token")
	}

	response, err := api.End(
		// The call center host
		viper.GetString("host"),
		// Our login token
		token,
		// The lock to end a call
		viper.GetString("lock"),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}
