package api

import (
	"io"
	"log"

	"github.com/pdxjohnny/microsocket/random"
	"github.com/spf13/viper"

	"golang.org/x/net/websocket"
)

const (
	// StateNotReady -
	StateNotReady = "not_ready"
	// StateHasLoginInfo -
	StateHasLoginInfo = "has_login_info"
	// StateLogin -
	StateLogin = "login"
	// StateReady -
	StateReady = "ready"
	// StateMakeCall -
	StateMakeCall = "make_call"
	// StateInCall -
	StateInCall = "in_call"
	// StateEndCall -
	StateEndCall = "end_call"
	// StateError -
	StateError = "error"
)

// CallerMessage holds messages like state updates that get sent to the callers
type CallerMessage struct {
	// For changing state
	State string `json:"state,omitempty"`
	// For updating properties like gmail_username
	Set   string `json:"set,omitempty"`
	Value string `json:"value,omitempty"`
	// For making a call
	Number string `json:"number,omitempty"`
}

// CallerReadyMessage is used to pass an id and websocket to the operator so that the
// caller can be removed if it disconnects
type CallerReadyMessage struct {
	ID string
	Ws *websocket.Conn
}

// CallerHandler handles callers connecting via websocket
func CallerHandler(ws *websocket.Conn) {
	// Generate a random callerID so that it can be removed if it disconnects
	callerID := random.Letters(10)
	log.Println("Caller connected", callerID)
	// Receive message from the caller
	for {
		var message CallerMessage
		err := websocket.JSON.Receive(ws, &message)
		if err == io.EOF {
			ws.Close()
			// Make sure the operator knows that the caller has disconnected
			Op.CallerDisconnected <- callerID
			log.Println("Caller disconnected", callerID)
			return
		} else if err != nil {
			log.Println("Error receiving from callerID", callerID, ":", err)
		} else {
			// Echo back
			websocket.JSON.Send(ws, message)
			// Preform actions based on state
			switch message.State {
			case StateNotReady:
				// Send it the login info
				websocket.JSON.Send(ws, map[string]string{
					"set":   "gmail_username",
					"value": viper.GetString("gmail_username"),
				})
				websocket.JSON.Send(ws, map[string]string{
					"set":   "gmail_password",
					"value": viper.GetString("gmail_password"),
				})
				// Tell it that it has the login info
				websocket.JSON.Send(ws, map[string]string{
					"state": StateHasLoginInfo,
				})
			case StateReady:
				// If its ready then the operator needs to know that
				callerReadyMessage := &CallerReadyMessage{
					ID: callerID,
					Ws: ws,
				}
				Op.CallerReady <- callerReadyMessage
			}
		}
	}
}
