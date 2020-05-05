package api

import (
	"github.com/QuestScreen/api/render"
	"github.com/QuestScreen/api/resources"
)

// the interfaces declared in this file are implemented by the QuestScreen core.

// MessageSender is used to send warnings and errors issued at startup to the
// client to be shown on the main page.
type MessageSender interface {
	Warning(text string)
	Error(text string)
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

// ServerContext gives access to data available in the server thread.
// This is a read-only view of data required for serialization and state
// initialization.
//
// Details on Fonts and Heroes are available in the display thread via
// [Extended]RenderContext.
type ServerContext interface {
	resources.Provider
	NumFontFamilies() int
	FontFamilyName(index int) string
	NumHeroes() int
	HeroID(index int) string
}

// Canvas is a facility to render content into a rectangular canvas.
// use it to pre-render content into a texture you can later copy to the
// renderer when rendering a scene.
//
// Creation of a Canvas sets the target of the underlying render.Renderer to the
// texture this Canvas draws on. Ensure that you either Close or Finish a
// Canvas object after creating it to reset the renderer's target.
//
// To draw into the canvas, use the normal renderer.
type Canvas interface {
	// Finish resets the target of the underlying renderer and returns the
	// rendered image.
	Finish() render.Image
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
	render.Renderer
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
	UpdateMask(target *render.Image, bg SelectableTexturedBackground)
	// RenderText renders the given text with the given font into an image with
	// transparent background.
	// Returns an empty image if it wasn't able to create the texture.
	RenderText(text string, font SelectableFont) render.Image
	// CreateCanvas creates a canvas to draw content into, and optionally fills it
	// with background color and/or with a repeating tile texture.
	//
	// Borders are added in each given direction. Border width/height is added to
	// the given innerWidth / innerHeight values.
	CreateCanvas(innerWidth, innerHeight int32, background render.RGBAColor,
		tile render.Image, borders Directions) Canvas
}

// ExtendedRenderContext is the context used for rebuilding the whole module
// and may contain additional data depending on the module's description.
type ExtendedRenderContext interface {
	RenderContext
	// Heroes returns a non-null list iff the module's description has UseHeroes set.
	Heroes() HeroList
}
