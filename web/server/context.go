package server

import (
	"github.com/QuestScreen/api/resources"
	"github.com/QuestScreen/api/web/groups"
)

// Context gives access to data available in the server thread.
// This is a read-only view of data required for serialization and state
// initialization.
//
// Details on Fonts and Heroes are available in the display thread via
// render.Context.
type Context interface {
	resources.Provider
	NumFontFamilies() int
	FontFamilyName(index int) string
	ActiveGroup() groups.Group
}
