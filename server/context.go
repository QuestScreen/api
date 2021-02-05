package server

import (
	"github.com/QuestScreen/api/groups"
	"github.com/QuestScreen/api/resources"
)

// Context gives access to data available in the server thread.
// This is a read-only view of data required for serialization and state
// initialization.
//
// Details on Fonts and Heroes are available in the display thread via
// render.Context.
type Context interface {
	resources.Provider
	ActiveGroup() groups.Group
}

// MessageSender is used to send warnings and errors issued at startup to the
// client to be shown on the main page.
type MessageSender interface {
	Warning(text string)
	Error(text string)
}
