package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/pdxjohnny/hangouts-call-center/api"
	"github.com/pdxjohnny/hangouts-call-center/variables"
)

// GetCall requests a call be made
func GetCall(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	doc, err := api.GetCall(variables.ServiceDBURL, r.Env["JWT_RAW"].(string), id)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if doc == nil {
		w.(http.ResponseWriter).Write(variables.BlankResponse)
		return
	}
	w.WriteJson(doc)
}
