package config

import (
	"github.com/QuestScreen/api/comms"
	"github.com/flyx/askew/runtime"
)

// EditHandler is a listener that gets called when an UI element's displayed
// state is changed.
type EditHandler interface {
	Edited()
}

// Widget is the UI element for a ConfigItem.
type Widget interface {
	runtime.Component
	// Receiver initializes the controller's state from JSON sent from the server.
	comms.Receiver
	// Sender sends the data entered by the user back to the server.
	// By sending the data, it consolidates it so that future calls to Reset()
	// will reset the values to the values being sent.
	comms.Sender
	// Reset resets the UI to the values that have last been queried via Data().
	// If the values have never been queried, the UI is reset to the initial
	// data the state object was loaded with.
	Reset()
	// SetEnabled enables or disables the GUI.
	SetEnabled(value bool)
}
