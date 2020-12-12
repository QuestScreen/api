package server

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

// State is the web UI's interface to the server's state.
type State interface {
	// Fetch sends an HTTP request to the server with the given method and given
	// payload, at the path created from joining a base path with the given
	// subpath. The result JSON is deserialized into the given target.
	//
	// The base path is defined by the ServerState implementation.
	// subpath must be relative and may not begin with a slash. It may be empty.
	// If it is not empty, it will be joined with the base path using a slash.
	Fetch(method RequestMethod, subpath string, payload interface{}, target interface{})
}
