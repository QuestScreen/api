package server

// Serializable describes a server item that can be serialized.
// Serialization happens both when the item is sent to a client via the Web API,
// and when the item is persisted to the file system.
//
// A Serializable item provides two views: One for communicating the item to the
// web client, and one for persisting it to the file system. Both
// implementations may trivially return a pointer to the item itself, if no
// special handling is required.
//
// On the client, all items received from and sent to the server are directly
// (de)serialized using the json package; this interface is only for server
// items that may provide a modified view for the client and/or for persisting.
type Serializable interface {
	// WebView returns a view of the data structure as it should be sent to the
	// web client.
	//
	// The returned view will be serialized as JSON, possibly as part of a
	// larger structure. If you need to manually serialize the structure, return
	// a json.RawMessage.
	WebView(ctx Context) interface{}

	// PersistingView returns a view of the data structure that can be
	// communicated to the client.
	//
	// The returned view will be serialized as YAML as part of a larger structure.
	// If you need to manually serialize the structure, return a *yaml.Node.
	PersistingView(ctx Context) interface{}
}
