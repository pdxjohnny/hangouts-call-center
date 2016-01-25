package api

import (
	"github.com/pdxjohnny/easyreq"

	"github.com/pdxjohnny/hangouts-call-center/variables"
)

// Login authenticates a client
func Login(host, token string, doc LoginData) (*map[string]interface{}, error) {
	path := variables.APIPathLogin
	return easyreq.GenericRequest(host, path, token, doc)
}

// Refresh gets a new token for the client
func Refresh(host, token string) (*map[string]interface{}, error) {
	path := variables.APIPathRefresh
	return easyreq.GenericRequest(host, path, token, nil)
}
