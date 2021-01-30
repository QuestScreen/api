package modules

import (
	"encoding/json"

	"github.com/QuestScreen/api/groups"
	"github.com/QuestScreen/api/web"
	"github.com/flyx/askew/runtime"
)

// State represents the module's state.
// A State acts independently, communicating with the server on its own
// to communicate changes. the only interface to the main application is the
// generation of its user interface that will be placed onto the state page.
type State interface {
	// UI generates the user interface of the State.
	// This method is called exactly once after a successful LoadFrom
	UI(server web.Server) runtime.Component
}

// Constructor is a function that constructs a state from given JSON input and
// server/group environment.
type Constructor func(data json.RawMessage, server web.Server,
	group groups.Group) (State, error)
