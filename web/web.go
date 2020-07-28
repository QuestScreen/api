// Package web is the API of the web interface.
package web

import (
	"encoding/json"

	"github.com/flyx/askew/runtime"
)

// EditHandler is a listener that gets called when an UI element's displayed
// state is changed.
type EditHandler interface {
	Edited()
}

// ConfigItemState is the state of a ConfigItem instance.
type ConfigItemState interface {
	// UI creates and returns this state's user interface.
	// This method is called exactly once on each state instance.
	UI(editHandler EditHandler) runtime.Component
	// Reset resets the UI to the values that have last been queried via Data().
	// If the values have never been queried, the UI is resetted to the initial
	// data the state object was loaded with.
	Reset()
	// SetEnabled enables or disables the GUI.
	SetEnabled(value bool)
	// Data returns an object that will be serialized and sent back to the server
	// to update the values of this ConfigItem state on the server side.
	Data() interface{}
}

// ConfigItem describes a ConfigItem type.
type ConfigItem struct {
	// ID is the identifier of this ConfigItem type and must be unique among all
	// ConfigItem instances.
	ID string
	// Load creates a state from the given JSON subtree.
	// data may be nil in which case default values must be loaded.
	Load func(data *json.RawMessage) (ConfigItemState, error)
}

// RequestMethod is an enum of known methods for Fetch.
type RequestMethod int

const (
	// Get is a GET request
	Get RequestMethod = iota
	// Post is a POST request
	Post
	// Put is a PUT request
	Put
	// Delete is a DELETE request
	Delete
)

func (r RequestMethod) String() string {
	switch r {
	case Get:
		return "GET"
	case Post:
		return "POST"
	case Put:
		return "PUT"
	case Delete:
		return "DELETE"
	default:
		panic("unknown request method!")
	}
}

// ServerState is the interface to the server's state.
type ServerState interface {
	// Fetch sends an HTTP request to the server with the given method and given
	// payload, at the path created from joining a base path with the given
	// subpath. The result JSON is deserialized into the given target.
	//
	// The base path is defined by the ServerState implementation.
	// subpath must be relative and may not begin with a slash. It may be empty.
	// If it is not empty, it will be joined with the base path using a slash.
	Fetch(method RequestMethod, subpath string, payload interface{}, target interface{})
}

// GroupData gives access to data of the current group shared between modules.
type GroupData interface {
	NumHeroes() int
	HeroName(index int) string
	HeroID(index int) string
}

// ModuleState is the API of a module's state.
// A ModuleState acts independenly, communicating with the server on its own
// to communicate changes. the only interface to the main application is the
// generation of its user interface that will be placed onto the state page.
type ModuleState interface {
	// UI generates the user interface of the ModuleState.
	// This method is called exactly once per ModuleState instance.
	UI() runtime.Component
}

// Module describes a module type.
type Module struct {
	// ID is the identifier of this Module, unique within the owning plugin.
	ID string
	// Load creates a new ModuleState from the given JSON message, which must not
	// be nil. server will be a ServerState with base path "/<plugin>/<module>"
	Load func(data json.RawMessage, server ServerState, group GroupData) (ModuleState, error)
}

// PluginRegistrator is the interface for registering ConfigItems and Modules
// of a plugin during app initialization.
type PluginRegistrator interface {
	// RegisterConfigItem registers the given config item with the app.
	// returns an error if given nil or if the given item's ID is already in use
	// by another item.
	RegisterConfigItem(item *ConfigItem) error
	// RegisterModule registers the given module with the app.
	// returns an error if given nil or if the given item's ID is already in use
	// by another item.
	RegisterModule(item *Module) error
}
