package caller

import (
	"bufio"
	"fmt"
	"log"
	"os"

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

	callerData, err := api.Caller(
		// The call center host
		viper.GetString("host"),
		// Our login token
		token,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(callerData)
	go callerData.Listen()
	go func() {
		for {
			select {
			// End this loop on close
			case <-callerData.Close:
				callerData.Close <- true
				fmt.Println("Main closing")
				return
			// Print errors
			case err := <-callerData.Err:
				fmt.Println("Main got err:", err)
			// Print recved
			case recv := <-callerData.Recv:
				fmt.Println("Main got recv:", recv)
			}
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Main got stdin:", text)
		callerData.Write([]byte(text))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
