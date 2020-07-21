// Package web is the API of the web interface.
package web

import "github.com/flyx/askew/runtime"

// ConfigItem is the API of a configuration item.
type ConfigItem interface {
	// ID returns the identifier of this ConfigItem, which must be unique among
	// all ConfigItems.
	ID() string
	// UI generates the user interface of the ConfigItem.
	// It must create a new component for each call.
	UI() runtime.Component
}

// ModuleState is the API of a module's state.
type ModuleState interface {
	// ID returns the identifier of this ModuleState, which must be unique among
	// all ModuleStates.
	ID() string
	// UI generates the user interface of the ModuleState.
	// It must create a new component for each call.
	UI() runtime.Component
}

// App is the API of the web application, implemented by the QuestScreen core.
type App interface {
	// RegisterConfigItem registers the given config item with the app.
	// returns an error if given nil or if the given item's ID is already in use
	// by another item.
	RegisterConfigItem(item ConfigItem) error
	// RegisterModuleState registers the given module state with the app.
	// returns an error if given nil or if the given item's ID is already in use
	// by another item.
	RegisterModuleState(item ModuleState) error
}
