package http

import (
	"fmt"
	"net/http"

	"github.com/pdxjohnny/hangouts-call-center/http/api"
)

// NewServeMux creates the main request multiplexer
func NewServeMux(static string) *http.ServeMux {
	mux := http.NewServeMux()
	staticDir := http.FileServer(http.Dir(static))
	mux.Handle("/", staticDir)
	mux.Handle("/api/", http.StripPrefix("/api", *api.MakeHandler()))
	return mux
}

// ServeMux starts a server as http or https if a cert and key are
// provided
func ServeMux(mux *http.ServeMux, address, port, cert, key string) error {
	listen := fmt.Sprintf("%s:%s", address, port)
	if cert == "false" || key == "false" {
		fmt.Printf("About to listen on http://%s/\n", listen)
		err := http.ListenAndServe(listen, mux)
		if err != nil {
			return err
		}
	} else {
		fmt.Printf("About to listen on https://%s/\n", listen)
		err := http.ListenAndServeTLS(listen, cert, key, mux)
		if err != nil {
			return err
		}
	}
	return nil
}
