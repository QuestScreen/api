package modules

import (
	"encoding/json"

	"github.com/QuestScreen/api/web/groups"
	"github.com/QuestScreen/api/web/server"
	"github.com/flyx/askew/runtime"
)

// State represents the module's state.
// A State acts independently, communicating with the server on its own
// to communicate changes. the only interface to the main application is the
// generation of its user interface that will be placed onto the state page.
type State interface {
	// LoadFrom initializes the state from the given JSON data.
	// It acts as constructor and discards any previous state.
	LoadFrom(data json.RawMessage, server server.State, group groups.Group) error
	// UI generates the user interface of the State.
	// This method is called exactly once after a successful LoadFrom
	UI(server server.State) runtime.Component
}
