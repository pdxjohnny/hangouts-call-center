package api

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/pdxjohnny/hangouts-call-center/variables"
)

// GetCall requests a call be made
func GetCall(w rest.ResponseWriter, r *rest.Request) {
	number := r.PathParam("number")
	log.Println("Will initiate call with number:", number)
	w.(http.ResponseWriter).Write(variables.BlankResponse)
	return
}
