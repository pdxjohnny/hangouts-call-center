package api

import (
	"errors"
	"io"
	"log"
	"strings"

	"golang.org/x/net/websocket"

	"github.com/pdxjohnny/hangouts-call-center/variables"
)

const (
	messageBuffer = 20
)

// Caller connects a call maker to the call center
func Caller(host, token string) (*CallerData, error) {
	// Create the websocket url
	url := strings.Replace(host, "http", "ws", 1) + variables.APIPathCaller
	// Create a config so we can add the authorization header
	config, err := websocket.NewConfig(url, host)
	if err != nil {
		return nil, err
	}
	// Set the authorization header to use the login token
	config.Header.Set("Authorization", "Bearer "+token)
	// Connect to the host
	ws, err := websocket.DialConfig(config)
	if err != nil {
		return nil, err
	}
	// Now we have a websocket connection
	callerData := NewCallerData(ws)
	return callerData, nil
}

// NewCallerData reate new chat CallerData.
func NewCallerData(ws *websocket.Conn) *CallerData {
	return &CallerData{
		Ws:    ws,
		Close: make(chan bool, messageBuffer),
		Err:   make(chan error, messageBuffer),
		Recv:  make(chan interface{}, messageBuffer),
		Send:  make(chan interface{}, messageBuffer),
	}
}

// Conn returns the underlying websocket
func (c *CallerData) Conn() *websocket.Conn {
	return c.Ws
}

// Write allows us to use CallerData as an io.Writer
func (c *CallerData) Write(p []byte) (n int, err error) {
	message := StringData{
		Data: string(p),
	}
	select {
	case c.Send <- message:
		return len(message.Data), nil
	default:
		c.Close <- true
		return 0, errors.New("Could not send message websocket disconnected")
	}
}

// Done closes the websocket connection
func (c *CallerData) Done() {
	c.Close <- true
}

// Listen Write and Read request via chanel
func (c *CallerData) Listen() {
	go c.listenWrite()
	c.listenRead()
}

// Listen write request via chanel
func (c *CallerData) listenWrite() {
	log.Println("Listening write to CallerData")
	for {
		select {
		// End this loop on close
		case <-c.Close:
			// Make sure listenRead gets the close
			c.Close <- true
			return
		// Send messages through the websocket on Send channel message received
		case message := <-c.Send:
			log.Println("Send:", message)
			websocket.JSON.Send(c.Ws, message)
		}
	}
}

// Listen read request via chanel
func (c *CallerData) listenRead() {
	log.Println("Listening read from CallerData")
	for {
		select {
		// End this loop on close
		case <-c.Close:
			// Make sure listenWrite gets the close
			c.Close <- true
			return
		// Read from the websocket into the recv channel
		default:
			var message map[string]interface{}
			err := websocket.JSON.Receive(c.Ws, &message)
			if err == io.EOF {
				c.Close <- true
			} else if err != nil {
				log.Println(err)
				c.Err <- err
				log.Println("Error went through channel")
			} else {
				c.Recv <- message
			}
		}
	}
}
