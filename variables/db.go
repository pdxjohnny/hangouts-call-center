package variables

import (
	"os"
)

const (
	// DBName is the name of the database in the database
	DBName = "hangouts-call-center"
	// DBAccounts is the name of the accounts collection in the database
	DBAccounts = "accounts"
	// DBUsers is the name of the users collection in the database
	DBUsers = "users"
	// EnvDBAddress is the env variable that the database address will be in
	EnvDBAddress = "MONGO_PORT_27017_TCP_ADDR"
)

var (
	// DBAddress is the address of database server
	DBAddress string
)

func init() {
	DBAddress = os.Getenv(EnvDBAddress)
}
