package render

import (
	"github.com/QuestScreen/api/common"
	"github.com/QuestScreen/api/config"
)

// Context is the context given to all rendering funcs of a module
// Besides being a renderer, it contains several helper functions.
type Context interface {
	Renderer
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
	// RenderText renders the given text with the given font into an image with
	// transparent background.
	// Returns an empty image if it wasn't able to create the texture.
	RenderText(text string, font config.Font) Image
	// CreateCanvas creates a canvas to draw content into, and optionally fills it
	// with background color and/or with a repeating tile texture.
	//
	// Borders are added in each given direction. Border width/height is added to
	// the given innerWidth / innerHeight values.
	CreateCanvas(innerWidth, innerHeight int32, background common.RGBAColor,
		tile Image, borders Directions) Canvas
}
