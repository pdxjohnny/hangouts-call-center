package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/pdxjohnny/hangouts-call-center/api"
	"github.com/pdxjohnny/hangouts-call-center/variables"
)

// PostLoginUser logs in a user
func PostLoginUser(w rest.ResponseWriter, r *rest.Request) {
	var recvDoc api.LoginData
	err := r.DecodeJsonPayload(&recvDoc)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	doc, err := api.LoginUser(variables.ServiceUserURL, variables.BackendToken, recvDoc)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if doc == nil {
		w.(http.ResponseWriter).Write(variables.BlankResponse)
		return
	}
	w.WriteJson(doc)
}

// PostRefreshUser logs in a user
func PostRefreshUser(w rest.ResponseWriter, r *rest.Request) {
	doc, err := api.RefreshUser(variables.ServiceUserURL, r.Env["JWT_RAW"].(string))
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if doc == nil {
		w.(http.ResponseWriter).Write(variables.BlankResponse)
		return
	}
	w.WriteJson(doc)
}

// PostRegisterUser registers a new user
func PostRegisterUser(w rest.ResponseWriter, r *rest.Request) {
	var recvDoc map[string]interface{}
	err := r.DecodeJsonPayload(&recvDoc)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	doc, err := api.RegisterUser(variables.ServiceUserURL, variables.BackendToken, recvDoc)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if doc == nil {
		w.(http.ResponseWriter).Write(variables.BlankResponse)
		return
	}
	w.WriteJson(doc)
}

// GetUser returns the accounts for an id
func GetUser(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	doc, err := api.GetUser(variables.ServiceUserURL, r.Env["JWT_RAW"].(string), id)
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

// PostUser uses get to retrive a document
func PostUser(w rest.ResponseWriter, r *rest.Request) {
	var recvDoc map[string]interface{}
	id := r.PathParam("id")
	err := r.DecodeJsonPayload(&recvDoc)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	doc, err := api.SaveUser(variables.ServiceUserURL, r.Env["JWT_RAW"].(string), id, recvDoc)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if doc == nil {
		w.(http.ResponseWriter).Write(variables.BlankResponse)
		return
	}
	w.WriteJson(doc)
}
