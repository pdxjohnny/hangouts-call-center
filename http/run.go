package http

import (
	"log"
	"net/http"

	"github.com/spf13/viper"

	"github.com/pdxjohnny/hangouts-call-center/http/api"
)

// Run starts the http(s) server for the cli
func Run() {
	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", *api.MakeHandler()))
	err := ServeMux(
		mux,
		viper.GetString("addr"),
		viper.GetString("port"),
		viper.GetString("cert"),
		viper.GetString("key"),
	)
	if err != nil {
		log.Println(err)
		return
	}
}
