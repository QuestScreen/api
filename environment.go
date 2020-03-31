package api

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// the interfaces declared in this file are implemented by the QuestScreen core.

// MessageSender is used to send warnings and errors issued at startup to the
// client to be shown on the main page.
type MessageSender interface {
	Warning(text string)
	Error(text string)
}

// Resource describes a selectable resource (typically a file).
type Resource interface {
	// Name of the file as it should be presented to the user.
	Name() string
	// Absolute path to the file.
	Path() string
}

// Hero describes a hero (player character).
type Hero interface {
	// Name of the hero
	Name() string
	// ID of the hero
	ID() string
	// Short description (e.g. class/race/etc)
	Description() string
}

// HeroList describes the list of heroes.
type HeroList interface {
	Hero(index int) Hero
	NumHeroes() int
}

// ResourceProvider is the interface to files on the file system that have been
// selected by a module's ResourceSelectors. Resources are read-only and
// available on both server and display thread.
type ResourceProvider interface {
	// GetResources queries the list of available resources of the given
	// resource collection index.
	//
	// The resources are filtered by the currently active system, group and scene.
	// Each Resource object is read-only and may be freely shared between threads.
	GetResources(index ResourceCollectionIndex) []Resource
	// GetTextures queries the list of available textures.
	GetTextures() []Resource
}

// ServerContext gives access to data available in the server thread.
// This is a read-only view of data required for serialization and state
// initialization.
//
// Details on Fonts and Heroes are available in the display thread via
// [Extended]RenderContext.
type ServerContext interface {
	ResourceProvider
	NumFontFamilies() int
	FontFamilyName(index int) string
	NumHeroes() int
	HeroID(index int) string
}

// Canvas is a facility to render content into a rectangular canvas.
// use it to pre-render content into a texture you can later copy to the
// renderer when rendering a scene.
//
// Creation of a Canvas sets the target of the underlying sdl.Renderer to the
// texture this Canvas draws on. Ensure that you either Close or Finish a
// Canvas object after creating it to reset the renderer's target.
//
// To draw into the canvas, use the normal renderer.
type Canvas interface {
	// Finish resets the target of the underlying renderer and returns the
	// rendered picture as texture.
	Finish() *sdl.Texture
	// Close resets the target of the underlying renderer.
	// It is idempotent and does nothing if Finish() was called before.
	Close()
}

// Directions is a bitset of directions.
type Directions uint8

const (
	// Nowhere contains no direction at all.
	Nowhere Directions = 0
	// North is the northern direction
	North Directions = 1 << iota
	// East is the eastern direction
	East
	// South is the southern direction
	South
	// West is the western direction
	West
	// Everywhere contains all directions.
	Everywhere = North | East | South | West
)

// RenderContext is the context given to all rendering funcs of a module
type RenderContext interface {
	ResourceProvider
	Renderer() *sdl.Renderer
	// Font returns the font face of the selected font.
	Font(fontFamily int, style FontStyle, size FontSize) *ttf.Font
	// DefaultBorderWidth returns the default size (in pixels) of a border line.
	DefaultBorderWidth() int32
	// LoadTexture is a helper function that loads a grayscale texture image into
	// an SDL texture which has its alpha channel set to the image's grayscale
	// channel and the color channels set to the given color.
	LoadTexture(textureIndex int, color RGBColor) (*sdl.Texture, error)
	// TextToTexture renders the given text with the given font and the given
	// color into a texture with transparent background.
	// Returns nil if it wasn't able to create the texture.
	TextToTexture(text string, font *ttf.Font, color sdl.Color) *sdl.Texture
	// CreateCanvas creates a canvas to draw content into, and optionally fills it
	// with background color and/or with a repeating tile texture.
	//
	// Borders are added in each given direction. Border width/height is added to
	// the given innerWidth / innerHeight values.
	CreateCanvas(innerWidth, innerHeight int32, background *sdl.Color,
		tile *sdl.Texture, borders Directions) Canvas
}

// ExtendedRenderContext is the context used for rebuilding the whole module
// and may contain additional data depending on the module's description.
type ExtendedRenderContext interface {
	RenderContext
	// Heroes returns a non-null list iff the module's description has UseHeroes set.
	Heroes() HeroList
}

// ResourceNames generates a list of resource names from a list of resources.
func ResourceNames(resources []Resource) []string {
	ret := make([]string, len(resources))
	for i := range resources {
		ret[i] = resources[i].Name()
	}
	return ret
}
