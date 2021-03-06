package resources

import (
	"encoding/json"
	"net/url"
)

// CollectionIndex indexes all resource collections of a module.
type CollectionIndex int

// Resource describes a selectable resource file.
type Resource struct {
	// Name of the resource as it should be presented to the user.
	Name string
	// Location of the resource. For files from the file system, the URL will have
	// file:// schema and is relative to the server, i.e. must not be resolved in
	// the web UI.
	//
	// Is a pointer because of package url's interface, must not be nil.
	Location *url.URL
}

// MarshalJSON marshals the resource's location as a simple string, to be parsed
// as URL again.
func (r *Resource) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name     string
		Location string
	}{r.Name, r.Location.String()})
}

// UnmarshalJSON loads the location into a string and then parses it as URL.
func (r *Resource) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Name     string
		Location string
	}
	var err error
	if err = json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	if r.Location, err = url.Parse(tmp.Location); err != nil {
		return err
	}
	r.Name = tmp.Name
	return nil
}

// Provider is the interface for querying resources.
type Provider interface {
	// GetResources returns the list of available resources of the given
	// resource collection index.
	//
	// The resources are filtered by the currently active system, group and scene.
	// Each Resource object is read-only and may be freely shared between threads.
	GetResources(index CollectionIndex) []Resource
	// GetTextures returns the list of available textures.
	GetTextures() []Resource
	// NumFontFamilies returns the number of available font families.
	NumFontFamilies() int
	// FontFamilyName returns the name of the font family at the given index.
	FontFamilyName(index int) string
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
