package render

import "github.com/QuestScreen/api/common"

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
	DisplaySize() (width int32, height int32)
	// FillRect fills the rectangle with the specified dimensions with the
	// specified color. The rectangle is positions via the given transformation.
	FillRect(width, height int32, t Transform, color common.RGBAColor)
	// LoadImage loads an image file from the specified path.
	// if an error is returned, the returned image is empty.
	LoadImage(path string) (Image, error)
	// FreeImage destroys the texture associated with the image (if one exists)
	// and sets i to be the empty image. Does nothing on empty images.
	FreeImage(i *Image)
	// DrawImage renders the given image if it is not empty, using the given
	// transformation. alpha modifies the image's opacity.
	DrawImage(image Image, t Transform, alpha uint8)
	// Center returns a transformation that moves a rectangle with given width and
	// height to the center of the screen
	Center(width int32, height int32) Transform
}
