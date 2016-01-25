package http

import (
	"fmt"
	"net/http"
)

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
