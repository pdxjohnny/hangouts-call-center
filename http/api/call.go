package api

import (
	"log"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/pdxjohnny/hangouts-call-center/api"
)

// ClientMakeCall is passed to the operator which will make the call then
// send the lock back through the lock channel once it has been made
type ClientMakeCall struct {
	Number string
	Lock   chan string
}

// GetCall requests a call be made
func GetCall(w rest.ResponseWriter, r *rest.Request) {
	// Get the number the client wants to call
	number := r.PathParam("number")
	log.Println("Will initiate call with number:", number)
	// Create the ClientMakeCall struct which hold the number to call and
	// the channel which will return the lock on the call
	makeCall := &ClientMakeCall{
		Lock:   make(chan string, 1),
		Number: number,
	}
	// Request that the call be made
	Op.MakeCall <- makeCall
	// Once the call is made we will have the lock
	lock := <-makeCall.Lock
	// Send the client the lock so they can end teh call if they wish
	w.WriteJson(api.ClientInCall{
		Number: number,
		Lock:   lock,
	})
	return
}
