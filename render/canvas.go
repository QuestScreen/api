package render

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
	Finish() Image
	// Close resets the target of the underlying renderer.
	// It is idempotent and does nothing if Finish() was called before.
	Close()
}
