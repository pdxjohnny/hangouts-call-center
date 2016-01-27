package api

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/pdxjohnny/hangouts-call-center/variables"
)

// GetEnd returns the accounts for an id
func GetEnd(w rest.ResponseWriter, r *rest.Request) {
	lock := r.PathParam("lock")
	log.Println("Will end call with lock:", lock)
	Op.EndCall <- lock
	w.(http.ResponseWriter).Write(variables.OKResponse)
	return
}
