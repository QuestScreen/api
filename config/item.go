// Package config defines the interface of a data item that is part of a
// module's configuration.
package config

import (
	"encoding/json"

	"github.com/QuestScreen/api/server"
	"gopkg.in/yaml.v3"
)

// Item describes an item in a module's configuration.
// A Item's public fields will be loaded from YAML structure automatically
// via reflection, and JSON serialization will also be done via reflection.
// you may use the tags `json:` and `yaml:` on those fields as documented in
// the json and yaml.v3 packages.
type Item interface {
	server.Serializable
	// LoadWeb loads the item's state from JSON data that as been
	// sent from the web client.
	//
	// Any structural and value error should result in returning a server.Error
	// (typically a *BadRequest) and should not alter the item's state.
	// Implementation should typically use the ReceiveData func, possibly together
	// with the strict ValidatedX types provided by the api package.
	LoadWeb(input json.RawMessage, ctx server.Context) server.Error
	// LoadPersisted loads the item's state from YAML data that has been
	// read from the file system.
	//
	// LoadPersisted should be robust when loading from Persisted layout, handling
	// errors by logging them and setting appropriate default values. An error
	// returned from loading Persisted data will lead to the app to exit.
	LoadPersisted(input *yaml.Node, ctx server.Context) error
}
