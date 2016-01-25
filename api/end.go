package api

import (
	"strings"

	"github.com/pdxjohnny/easyreq"

	"github.com/pdxjohnny/hangouts-call-center/variables"
)

// End ends a call that has been started provided you have the lock
func End(host, token, lock string) (*map[string]interface{}, error) {
	path := variables.APIPathEnd
	path = strings.Replace(path, ":lock", lock, 1)
	return easyreq.GenericRequest(host, path, token, nil)
}
