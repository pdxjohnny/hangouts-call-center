package api

import (
	"log"
	"time"

	"github.com/pdxjohnny/microsocket/random"
	"golang.org/x/net/websocket"
)

const (
	opBuffer = 20
	// DefaultCallTimeout is the default time to wait before ending a call
	DefaultCallTimeout = 2 * time.Minute
)

// Op is used if you don't want to create a NewOperator
var Op *Operator

// Operator connects calls that need to be made with callers that make them
type Operator struct {
	Timeout time.Duration
	// Send requests to make a call to this channel
	MakeCall chan *ClientMakeCall
	// Send requests to end a call to this channel
	EndCall chan string
	// Callers who enter the ready state need to be delivered to the op here
	CallerReady chan *CallerReadyMessage
	// Callers that disconnect must be removed from the ready map
	CallerDisconnected chan string
	// These are the calls waiting for a caller to make them
	CallsWaiting []*ClientMakeCall
	// These are the callers that are ready to make calls
	CallersReady map[string]*websocket.Conn
	// Callers maped by locks so their calls can be ended
	CallersInCall map[string]*websocket.Conn
}

func init() {
	Op = NewOperator()
	go Op.Route()
}

// NewOperator creates an operator
func NewOperator() *Operator {
	op := Operator{
		Timeout:            DefaultCallTimeout,
		MakeCall:           make(chan *ClientMakeCall, opBuffer),
		EndCall:            make(chan string, opBuffer),
		CallerReady:        make(chan *CallerReadyMessage, opBuffer),
		CallerDisconnected: make(chan string, opBuffer),
		CallsWaiting:       make([]*ClientMakeCall, 0),
		CallersReady:       make(map[string]*websocket.Conn, opBuffer),
		CallersInCall:      make(map[string]*websocket.Conn, opBuffer),
	}
	return &op
}

// Route starts the operator so that is listens for incoming call requests
// and pairs them with available callers
func (o *Operator) Route() {
	for {
		select {
		case caller := <-o.CallerReady:
			log.Println(caller, "is ready")
			// Apped to the callers that are ready
			o.CallersReady[caller.ID] = caller.Ws
			// Check to see if this new caller can make a call that is waiting to be
			// made
			o.CheckCallsWaiting()
		case callerID := <-o.CallerDisconnected:
			// Take the caller out of the ready map because it disconnected
			delete(o.CallersReady, callerID)
		case makeCall := <-o.MakeCall:
			log.Println("Got request to call", makeCall)
			// Add this number to the numbers that are waiting to be called
			o.CallsWaiting = append(o.CallsWaiting, makeCall)
			// Check to see if there is a caller available to call the number
			o.CheckCallsWaiting()
		case lock := <-o.EndCall:
			// Check to see if there is a caller in the in_call state for this lock
			ws, ok := o.CallersInCall[lock]
			// If there is not caller for this lock then the caller has already gone
			// out of the in_call state
			if ok != true {
				continue
			}
			// End the call because the lock is associated with a caller in a call
			websocket.JSON.Send(ws, CallerMessage{
				State: "end_call",
			})
			// Delete the call from the CallersInCall because it has just been ended
			delete(o.CallersInCall, lock)
			log.Println("Call with lock", lock, "has been ended")
		default:
			continue
		}
	}
}

// CheckCallsWaiting checks to see if there are any callers available to make
// the calls waiting to be called
func (o *Operator) CheckCallsWaiting() {
	log.Println("Checking for calls waiting...")
	log.Println(o.CallsWaiting, o.CallersReady)
	// If there are no calls waiting to be made there is nothing to do
	if len(o.CallsWaiting) < 1 {
		return
	}
	// There are no callers ready so there is nothing we can do
	if len(o.CallersReady) < 1 {
		return
	}
	// Grab a random caller
	var callerID string
	var caller *websocket.Conn
	for id, callerWs := range o.CallersReady {
		callerID = id
		caller = callerWs
		// Grab one and exit the loop
		break
	}
	// The caller should be removed from the ready state becase it will now be
	// in the in_call state
	delete(o.CallersReady, callerID)
	// The call we are about to make will no longer waiting to be made and
	// therefore needs to be removed from the waiting slice
	makeCall, callsWaiting := o.CallsWaiting[0], o.CallsWaiting[1:]
	o.CallsWaiting = callsWaiting
	// Take the caller we just removed from ready and put them in the make_call
	// state for the number we just removed from CallsWaiting
	log.Println("About to have", callerID, "call", makeCall.Number)
	websocket.JSON.Send(caller, CallerMessage{
		State:  "make_call",
		Number: makeCall.Number,
	})

	// We need a way to end the call so take the caller and put them in the
	// CallersInCall map, the key to access it will be a randomly generated
	// lock which the client will get back in case they want to end the call.
	// If the client never ends the call then EndCallTimeout will occor and the
	// call will be ended

	// Create the lock
	log.Println("Creating lock for", makeCall.Number, "...")
	lock := random.Letters(10)
	log.Println("Status make_call has been sent to", callerID,
		"for number", makeCall.Number,
		"with lock", lock,
	)
	// Put the caller in the CallersInCall map by that lock
	o.CallersInCall[lock] = caller
	// Start the timeout in case the client does not end the call
	go o.EndCallTimeout(lock, o.Timeout)
	// Return the lock to the client that requested this call in case they want
	// to end it before the timeout
	makeCall.Lock <- lock
}

// EndCallTimeout will send the end call request when the timeout occurs
// it whoud be called as soon as the call is made
func (o *Operator) EndCallTimeout(lock string, timeout time.Duration) {
	<-time.After(timeout)
	log.Println("Lock:", lock, "timed out")
	o.EndCall <- lock
}
