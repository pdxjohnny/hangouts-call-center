package api

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/pdxjohnny/go-json-rest-middleware-jwt"
	"github.com/spf13/viper"
	"golang.org/x/net/websocket"

	"github.com/pdxjohnny/hangouts-call-center/variables"
)

// CreateAuthMiddleware creates the middleware for authtication
func CreateAuthMiddleware() (*jwt.Middleware, error) {
	err := variables.LoadTokenKeys()
	if err != nil {
		return nil, err
	}

	authMiddleware := &jwt.Middleware{
		Realm:            "hangouts-call-center",
		SigningAlgorithm: variables.SigningAlgorithm,
		Key:              variables.TokenSignKey,
		VerifyKey:        variables.TokenVerifyKey,
		// Ten year refresh
		Timeout:    time.Hour * 24 * 365 * 10,
		MaxRefresh: time.Hour * 24 * 365 * 10,
		Authenticator: func(username string, password string) error {
			log.Println("Attempting login for", username)
			// FIXME: Make these passed in from viper
			if username != viper.GetString("username") ||
				password != viper.GetString("password") {
				return errors.New("Incorrect username and password")
			}
			return nil
		},
	}
	return authMiddleware, nil
}

// MakeHandler creates the api request handler
func MakeHandler() *http.Handler {
	api := rest.NewApi()

	authMiddleware, err := CreateAuthMiddleware()
	if err != nil {
		panic(err)
	}

	wsHandler := websocket.Handler(CallerHandler)

	api.Use(&rest.IfMiddleware{
		// Only authenticate non login requests
		Condition: func(request *rest.Request) bool {
			return (request.URL.Path != variables.APIPathLoginServer)
		},
		IfTrue: authMiddleware,
	})
	api.Use(rest.DefaultProdStack...)
	router, err := rest.MakeRouter(
		// For login and refresh
		rest.Post(variables.APIPathLoginServer, authMiddleware.LoginHandler),
		rest.Get(variables.APIPathRefreshServer, authMiddleware.RefreshHandler),
		// For placing a call
		rest.Get(variables.APIPathCallServer, GetCall),
		// For ending a call
		rest.Get(variables.APIPathEndServer, GetEnd),
		// For caller nodes
		rest.Get(variables.APIPathCallerServer, func(w rest.ResponseWriter, r *rest.Request) {
			wsHandler.ServeHTTP(w.(http.ResponseWriter), r.Request)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	handler := api.MakeHandler()
	return &handler
}
