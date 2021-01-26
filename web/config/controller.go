package config

import (
	"encoding/json"

	"github.com/flyx/askew/runtime"
)

// EditHandler is a listener that gets called when an UI element's displayed
// state is changed.
type EditHandler interface {
	Edited()
}

// Controller backs the UI for manipulating config data in the web UI.
type Controller interface {
	json.Unmarshaler
	// UI creates and returns this controller's user interface.
	// This method is called exactly once on each controller instance.
	UI(editHandler EditHandler) runtime.Component
	// Reset resets the UI to the values that have last been queried via Data().
	// If the values have never been queried, the UI is reset to the initial
	// data the state object was loaded with.
	Reset()
	// SetEnabled enables or disables the GUI.
	SetEnabled(value bool)
	// Data returns an object that will be serialized and sent back to the server
	// to update the values of this ConfigItem state on the server side.
	Data() interface{}
}

// Item is a named Controller.
type Item struct {
	Controller
	// Name is the name of this config item and will be displayed in the UI.
	Name string
}
