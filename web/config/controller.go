package config

import (
	"github.com/QuestScreen/api/comms"
	"github.com/QuestScreen/api/server"
	"github.com/flyx/askew/runtime"
)

// EditHandler is a listener that gets called when an UI element's displayed
// state is changed.
type EditHandler interface {
	Edited()
}

// Controller backs the UI for manipulating config data in the web UI.
type Controller interface {
	// Receiver initializes the controller's state from JSON sent from the server.
	comms.Receiver
	// Sender sends the data entered by the user back to the server.
	// By sending the data, it consolidates it so that future calls to Reset()
	// will reset the values to the values being sent.
	comms.Sender
	// Load loads the given JSON data into the controller object.
	// The JSON data
	Load(json []byte, ctx server.Context) error
	// UI creates and returns this controller's user interface.
	// This method is called exactly once on each controller instance.
	UI(editHandler EditHandler) runtime.Component
	// Reset resets the UI to the values that have last been queried via Data().
	// If the values have never been queried, the UI is reset to the initial
	// data the state object was loaded with.
	Reset()
	// SetEnabled enables or disables the GUI.
	SetEnabled(value bool)
}

// Item is a named Controller.
type Item struct {
	Controller
	// Name is the name of this config item and will be displayed in the UI.
	Name string
}
