package http

import (
	"log"

	"github.com/spf13/viper"
)

// Run starts the http(s) server for the cli
func Run() {
	mux := NewServeMux(viper.GetString("static"))
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
