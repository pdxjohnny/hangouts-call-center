package api

import "golang.org/x/net/websocket"

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

// CallerData contains a recv channel that will be filled with messages
// received via websocket from the call center as well as a send channel
// which is used to send messages to the call center
type CallerData struct {
	Ws    *websocket.Conn
	Close chan bool
	Err   chan error
	Recv  chan interface{}
	Send  chan interface{}
}

// StringData is used for the write method of CallerData so that it can be used
// as an io.Writer
type StringData struct {
	Data string `json:"data"`
}
