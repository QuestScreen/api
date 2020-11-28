package resources

// CollectionIndex indexes all resource collections of a module.
type CollectionIndex int

// Resource describes a selectable resource file.
type Resource interface {
	// Name of the resource as it should be presented to the user.
	Name() string
	// Absolute path to the file
	Path() string
}

// Provider is the interface for querying resources.
type Provider interface {
	// GetResources queries the list of available resources of the given
	// resource collection index.
	//
	// The resources are filtered by the currently active system, group and scene.
	// Each Resource object is read-only and may be freely shared between threads.
	GetResources(index CollectionIndex) []Resource
	// GetTextures queries the list of available textures.
	GetTextures() []Resource
}

// Names generates a list of resource names from a list of resources.
func Names(resources []Resource) []string {
	ret := make([]string, len(resources))
	for i := range resources {
		ret[i] = resources[i].Name()
	}
	return ret
}

// Selector defines where a module finds resource files.
// A selector can either be used to find one specific file (in which case
// Name must be non-empty) or a list of files (in which case Name must be empty
// and Suffixes may be set).
//
// The file(s) are searched in the Subdirectory inside the module's directory in
// the current scene, group, system and base directory (in that order).
type Selector struct {
	// may be empty, in which case resource files are searched directly
	// in the module directories.
	Subdirectory string
	// Name of the file including suffix. If not empty, only the first file
	// matching the name is returned.
	Name string
	// filters files by suffix. If empty or nil, no filter will be applied
	// (note however that files starting with a dot will always be filtered out).
	// ignored if Name is not empty.
	Suffixes []string
}
