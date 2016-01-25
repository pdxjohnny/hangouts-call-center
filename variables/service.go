package variables

import (
	"os"
	"strings"
)

const (
	// EnvServiceProtocol is the env vaiable which contains
	// either http or https or none for http
	EnvServiceProtocol = "SERVICE_PROTOCOL"
	// EnvServiceDBURL is the env vaiable which contains
	// the url which the db connector service will reside at
	EnvServiceDBURL = "DB_PORT"
	// EnvServiceUserURL is the env vaiable which contains
	// the url which the user connector service will reside at
	EnvServiceUserURL = "USER_PORT"
)

var (
	// ServiceProtocol determines wiether communications are encrypted or not
	// either http or https or none for http
	ServiceProtocol string
	// ServiceDBURL is the url which the db connector service will reside at
	ServiceDBURL string
	// ServiceUserURL is the url which the user connector service will reside at
	ServiceUserURL string
)

func init() {
	ServiceProtocol = os.Getenv(EnvServiceProtocol)
	if ServiceProtocol == "" {
		ServiceProtocol = "https"
	}
	ServiceDBURL = os.Getenv(EnvServiceDBURL)
	ServiceDBURL = strings.Replace(ServiceDBURL, "tcp", ServiceProtocol, 1)
	ServiceUserURL = os.Getenv(EnvServiceUserURL)
	ServiceUserURL = strings.Replace(ServiceUserURL, "tcp", ServiceProtocol, 1)
}
