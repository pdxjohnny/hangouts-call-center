package api

import (
	"strings"

	"github.com/pdxjohnny/easyreq"

	"github.com/pdxjohnny/hangouts-call-center/variables"
)

// Call makes a call to a number and returns a lock
func Call(host, token, number string) (*map[string]interface{}, error) {
	path := variables.APIPathCall
	path = strings.Replace(path, ":number", number, 1)
	return easyreq.GenericRequest(host, path, token, nil)
}
