package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

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
  // FIXME: The username and password should be given to viper
	if recvDoc.Username != "user" && recvDoc.Password != "pass" {
		rest.Error(w, "Incorrect username and password", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.(http.ResponseWriter).Write(variables.BlankResponse)
	return
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
