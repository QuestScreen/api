package config

import (
	"github.com/QuestScreen/api/comms"
	"github.com/QuestScreen/api/server"
)

// Item describes an item in a module's configuration.
// A Item's public fields will be loaded from YAML structure automatically
// via reflection, and JSON serialization will also be done via reflection.
// you may use the tags `json:` and `yaml:` on those fields as documented in
// the json and yaml.v3 packages.
type Item interface {
	comms.Sender
	comms.Receiver
	server.Persister
	server.Loader
}
