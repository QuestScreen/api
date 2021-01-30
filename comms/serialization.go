package comms

import (
	"encoding/json"

	"github.com/QuestScreen/api/server"
)

// Sender describes an object that can be serialized to JSON and sent via
// HTTP between server and client.
type Sender interface {
	// Send returns a view of the data structure that should be serialized to
	// JSON to communicate between server and client.
	//
	// The returned view will be serialized as JSON, possibly as part of a
	// larger structure. If you need to manually serialize the structure, return
	// a json.RawMessage. Generally, if the returned value implements
	// json.Marshaler, that will be used for serializing.
	Send(ctx server.Context) interface{}
}

// Receiver is the counterpart of Sender, it is able to receive JSON data.
type Receiver interface {
	// Receive invalidates the Receiver's state and overwrites it with the data
	// provided by the given JSON input.
	//
	// An error is returned if there are any problems with the data.
	// If an error is returned, the Receiver's previous state must be preserved.
	Receive(input json.RawMessage, ctx server.Context) error
}
