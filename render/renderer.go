package render

import "github.com/QuestScreen/api"

// Image is a rectangular image stored as OpenGL texture
// `id` is the name of an OpenGL texture iff width!=0.
// any Image width width==0 is considered the empty image.
//
// In the context of this API, this type should be considered opaque with only
// Width and Height being read-accessible. The fields are public so that plugins
// can use the data for custom OpenGL processing.
type Image struct {
	TextureID     uint32
	Width, Height int32
	// Flipped describes the row order in the image data.
	//   true  -> first datapoint in image data is lower left corner.
	//   false -> first datapoint in image data is upper left corner.
	// This is used because loading image files returns the top row first,
	// while internally OpenGL places the bottom row first (e.g. when rendering
	// to a texture).
	Flipped bool
	// true iff the texture has an alpha channel.
	HasAlpha bool
}

// EmptyImage returns an image that has no linked OpenGL texture.
func EmptyImage() Image {
	return Image{}
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
	FillRect(t Transform, color api.RGBA)
	// DrawImage renders the given image if it is not empty on a square with
	// edge length of 1.0 centered around the origin, transformed with the given
	// transformation. alpha modifies the image's opacity.
	//
	// For the high-level API, use Image's Draw() instead.
	DrawImage(image Image, t Transform, alpha uint8)
	// RenderText renders the given text with the given font into an image with
	// transparent background.
	// Returns an empty image if it wasn't able to create the texture.
	RenderText(text string, fontFamily int, size api.FontSize,
		style api.FontStyle, color api.RGBA) Image
	// CreateCanvas creates a canvas to draw content into, and fills it with the
	// given background. The returned content rectangle is the canvas area minus
	// the borders.
	//
	// Borders are added in each given direction. Border width/height is added to
	// the given innerWidth / innerHeight values.
	//
	// The texture created by the canvas will have an alpha channel only if the
	// primary color has an alpha value other than 255, or if a mask is set and
	// the secondary color has an alpha value other than 255.
	CreateCanvas(innerWidth, innerHeight int32,
		primaryColor, secondaryColor api.RGBA, textureIndex int,
		borders Directions) (canvas Canvas, content Rectangle)
	// LoadImageFile loads an image file from the specified path.
	// if an error is returned, the returned image is empty.
	//
	// if scaleDownToOutput is true, the image is scaled down to the output
	// context's dimensions so that it completely fits into the display while
	// preserving aspect ratio.
	//
	// The image will always be scaled down to GL_MAX_TEXTURE_SIZE if its
	// dimensions exceed it.
	LoadImageFile(path string, scaleDownToOutput bool) (Image, error)
	// LoadImageMem loads an image from data in memory.
	// if an error is returned, the returned image is empty.
	//
	// scaleDownToOutput work like for LoadImageFile.
	LoadImageMem(data []byte, scaleDownToOutput bool) (Image, error)
	// FreeImage destroys the texture associated with the image (if one exists)
	// and sets i to be the empty image. Does nothing on empty images.
	FreeImage(i *Image)
}
