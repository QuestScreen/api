package render

import (
	"github.com/QuestScreen/api/config"
)

// Context is the context given to all rendering funcs of a module
// It provides the rendering facilities of Renderer, as well as information
// about the rendering area and utility functions.
type Context interface {
	Renderer
	// Display returns a rectangle that describes the dimensions of the current
	// rendering area.
	Display() Rectangle
	// Unit is the scaled smallest unit in pixels. It is defined as being
	// 1/144 of the screen's width or height, whichever is smaller.
	// Borders typically have the size of one Unit.
	//
	// Use this value as base for calculating sizes to make your renderer output
	// scale depending on the target display size. Not using lengths smaller than
	// Unit avoids making stuff too small for your users to see.
	Unit() int32
	// UpdateMask is a helper function that loads a selected grayscale image into
	// an the given Image so that its alpha channel is set to the image's
	// grayscale channel and the color channels are set to the given secondary
	// color.
	//
	// If *target is non-empty, the previous texture gets destroyed.
	// If not mask texture is selected, *target will be empty.
	UpdateMask(target *Image, bg config.Background)
	// LoadImage loads an image file from the specified path.
	// if an error is returned, the returned image is empty.
	LoadImage(path string) (Image, error)
	// FreeImage destroys the texture associated with the image (if one exists)
	// and sets i to be the empty image. Does nothing on empty images.
	FreeImage(i *Image)
}
