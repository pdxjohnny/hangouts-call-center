package api

// LoginData provides username and password to authenticate with
type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CallData contains the number to call
type CallData struct {
	Number string `json:"number"`
}

// EndData the the lock on the call which allows a client to end that call
type EndData struct {
	Lock string `json:"lock"`
}
