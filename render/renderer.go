package render

import (
	"github.com/QuestScreen/api/common"
	"github.com/QuestScreen/api/config"
)

// Image is a rectangular image stored as OpenGL texture
// `id` is the name of an OpenGL texture iff width!=0.
// any Image width width==0 is considered the empty image.
type Image struct {
	TextureID     uint32
	Width, Height int32
}

// IsEmpty tests whether the image is empty
// (i.e. does not link to an OpenGL texture)
func (i *Image) IsEmpty() bool {
	return i.Width == 0
}

// Renderer describes an object providing functions for rendering objects.
type Renderer interface {
	// OutputSize returns a rectangle that describes the dimensions in pixels
	// of the current rendering area. X and Y are always 0.
	//
	// This rectangle will be the whole OpenGL surface by default, but if a Canvas
	// is active, it returns the Canvas' size.
	OutputSize() Rectangle
	// Unit is the scaled smallest unit in pixels. It is defined as being
	// 1/144 of the screen's width or height, whichever is smaller.
	// Borders typically have the size of one Unit.
	//
	// Use this value as base for calculating sizes to make your renderer output
	// scale depending on the target display size. Not using lengths smaller than
	// Unit avoids making stuff too small for your users to see.
	//
	// Font sizes depend directly on this size.
	Unit() int32
	// FillRect fills the rectangle with the specified dimensions with the
	// specified color. The rectangle is positions via the given transformation.
	FillRect(width, height int32, t Transform, color common.RGBAColor)
	// DrawImage renders the given image if it is not empty, using the given
	// transformation. alpha modifies the image's opacity.
	DrawImage(image Image, t Transform, alpha uint8)
	// RenderText renders the given text with the given font into an image with
	// transparent background.
	// Returns an empty image if it wasn't able to create the texture.
	RenderText(text string, font config.Font) Image
	// CreateCanvas creates a canvas to draw content into, and fills it with the
	// given background.
	//
	// Borders are added in each given direction. Border width/height is added to
	// the given innerWidth / innerHeight values.
	CreateCanvas(innerWidth, innerHeight int32, bg config.Background,
		borders Directions) Canvas
	// LoadImage loads an image file from the specified path.
	// if an error is returned, the returned image is empty.
	LoadImage(path string) (Image, error)
	// FreeImage destroys the texture associated with the image (if one exists)
	// and sets i to be the empty image. Does nothing on empty images.
	FreeImage(i *Image)
}
