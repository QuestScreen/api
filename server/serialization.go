// +build !js

package server

import "gopkg.in/yaml.v3"

// Persister describes an object whose state can be persisted into a YAML
// stream.
type Persister interface {
	// Persist returns a view of the data that can be written to persistent
	// storage in YAML format, possibly as part of a larger structure.
	//
	// If you need to manually serialize the structure, return a *yaml.Node.
	// Generally, if the returned value implements yaml.Marshaler, that will be
	// used for serialization.
	Persist(ctx Context) interface{}
}

// Loader describes an object whose state can be loaded from YAML data.
type Loader interface {
	// Load invalidates the Loader's state and replaces it with the given YAML
	// data.
	//
	// If the YAML data is invalid, an error is returned and the Loader's previous
	// state must be preserved.
	Load(node *yaml.Node, ctx Context) error
}
