package render

import (
	"github.com/QuestScreen/api/colors"
	"github.com/QuestScreen/api/fonts"
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
func (i Image) IsEmpty() bool {
	return i.Width == 0
}

// Draw draws the image to the given rectangular area.
// The image will be stretched to fit the whole area.
func (i Image) Draw(r Renderer, area Rectangle, alpha uint8) {
	r.DrawImage(i, area.Transformation(), alpha)
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
	// FillRect fills a rectangle with the specified color.
	// The rectangle is a square with edge length of 1.0 centered around the
	// origin, transformed with the given transformation.
	//
	// For the high-level API, use Rectangle's Fill() instead.
	FillRect(t Transform, color colors.RGBA)
	// DrawImage renders the given image if it is not empty on a square with
	// edge length of 1.0 centered around the origin, transformed with the given
	// transformation. alpha modifies the image's opacity.
	//
	// For the high-level API, use Image's Draw() instead.
	DrawImage(image Image, t Transform, alpha uint8)
	// RenderText renders the given text with the given font into an image with
	// transparent background.
	// Returns an empty image if it wasn't able to create the texture.
	RenderText(text string, font fonts.Config) Image
	// CreateCanvas creates a canvas to draw content into, and fills it with the
	// given background.
	//
	// Borders are added in each given direction. Border width/height is added to
	// the given innerWidth / innerHeight values.
	CreateCanvas(innerWidth, innerHeight int32, bg colors.Background,
		borders Directions) Canvas
	// LoadImageFile loads an image file from the specified path.
	// if an error is returned, the returned image is empty.
	LoadImageFile(path string) (Image, error)
	// LoadImageMem loads an image from data in memory.
	// if an error is returned, the returned image is empty.
	LoadImageMem(data []byte) (Image, error)
	// FreeImage destroys the texture associated with the image (if one exists)
	// and sets i to be the empty image. Does nothing on empty images.
	FreeImage(i *Image)
}
