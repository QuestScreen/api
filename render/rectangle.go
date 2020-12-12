package render

import "github.com/QuestScreen/api"

// Rectangle describes a rectangle positioned on the screen.
//
// The screen coordinates are (0,0) at the lower left corner and
// (width, height) at the upper right corner.
//
// Rectangle is the high-level API for positioning stuff on the screen.
// It provides functionality to position one rectangle in another one and so on.
// You can use it in conjunction with the lower-level Transform API which can
// do rotations by using Rectangle's Translation() / Transformation() funcs.
type Rectangle struct {
	// coordinate of the lower left corner
	X, Y          int32
	Width, Height int32
}

// Translation returns the transformation needed to move an object centered on
// the origin to the center of the rectangle.
func (r Rectangle) Translation() Transform {
	return Identity().Translate(float32(r.X)+float32(r.Width)/2.0,
		float32(r.Y)+float32(r.Height)/2.0)
}

// Transformation returns the transformation needed to transform a square with
// edge length of 1.0 centered around the origin to the subject rectangle.
func (r Rectangle) Transformation() Transform {
	return r.Translation().Scale(float32(r.Width), float32(r.Height))
}

// Move moves the rectangle by the given delta
func (r Rectangle) Move(dx, dy int32) Rectangle {
	return Rectangle{r.X + dx, r.Y + dy, r.Width, r.Height}
}

// Shrink removes dw from the rectangles width and dh from its height,
// repositioning it so that the center stays the same.
func (r Rectangle) Shrink(dw, dh int32) Rectangle {
	return Rectangle{r.X + dw/2, r.Y + dh/2, r.Width - dw, r.Height - dh}
}

// Scale scales the rectangle's width and height by the given factor,
// repositioning it so that the center stays the same.
func (r Rectangle) Scale(factor float32) Rectangle {
	ret := Rectangle{Width: int32(float32(r.Width) * factor),
		Height: int32(float32(r.Height) * factor)}
	ret.X = r.X + (r.Width-ret.Width)/2
	ret.Y = r.Y + (r.Height-ret.Height)/2
	return ret
}

// HAlign defines horizontal alignment
type HAlign int

const (
	// Left is horizontal alignment to the left
	Left HAlign = iota
	// Center is horizontal alignment in the center
	Center
	// Right is horizontal alignment to the right
	Right
	// HStretch stretches width to the maximum
	HStretch
)

// VAlign defines vertical alignment
type VAlign int

const (
	// Top is vertical alignment to the top
	Top VAlign = iota
	// Middle is vertical alignment in the middle
	Middle
	// Bottom is vertical alignment to the bottom
	Bottom
	// VStretch stretches height to the maximum
	VStretch
)

// Position returns a rectangle with the given width and height, which is
// position in the current rectangle according to the given flags.
//
// giving HStretch and VStretch will override the given width and height
// respectively, the other positioning flags will only set the position.
func (r Rectangle) Position(width, height int32, horiz HAlign,
	vert VAlign) Rectangle {
	ret := Rectangle{Width: width, Height: height}
	switch horiz {
	case Left:
		ret.X = r.X
	case Center:
		ret.X = r.X + (r.Width-width)/2
	case Right:
		ret.X = r.X + r.Width - width
	case HStretch:
		ret.X = r.X
		ret.Width = r.Width
	}
	switch vert {
	case Top:
		ret.Y = r.Y + r.Height - height
	case Middle:
		ret.Y = r.Y + (r.Height-height)/2
	case Bottom:
		ret.Y = r.Y
	case VStretch:
		ret.Y = r.Y
		ret.Height = r.Height
	}
	return ret
}

// Carve removes a rectangle of the given length starting at the given edge
// from the current rectangle.
//
// edge must be North, East, South or West. The carved rectangle is returned
// as `carved`, the remaining rectangle as `rest`
func (r Rectangle) Carve(edge Directions,
	length int32) (carved Rectangle, rest Rectangle) {
	switch edge {
	case North:
		rest = Rectangle{X: r.X, Y: r.Y, Width: r.Width, Height: r.Height - length}
		carved = Rectangle{X: r.X, Y: r.Y + rest.Height, Width: r.Width, Height: length}
	case East:
		rest = Rectangle{X: r.X, Y: r.Y, Width: r.Width - length, Height: r.Height}
		carved = Rectangle{X: r.X + rest.Width, Y: r.Y, Width: length, Height: r.Height}
	case South:
		rest = Rectangle{X: r.X, Y: r.Y + length, Width: r.Width, Height: r.Height - length}
		carved = Rectangle{X: r.X, Y: r.Y, Width: r.Width, Height: length}
	case West:
		rest = Rectangle{X: r.X + length, Y: r.Y, Width: r.Width - length, Height: r.Height}
		carved = Rectangle{X: r.X, Y: r.Y, Width: length, Height: r.Height}
	default:
		panic("illegal edge (must be North, East, South or West)")
	}
	return
}

// Fill fills the rectangle with the given color.
func (r Rectangle) Fill(renderer Renderer, color api.RGBA) {
	renderer.FillRect(r.Transformation(), color)
}
