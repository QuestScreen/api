package render

import "github.com/QuestScreen/api/colors"

// Rectangle describes a positioned rectangle.
//
// This can be used as intermediate structure for positioning stuff,
// to avoid having to deal with Transform's float64 values too much.
type Rectangle struct {
	// coordinate of the lower left corner
	X, Y          int32
	Width, Height int32
}

// Translation returns the transformation needed to move an object centered on
// the origin to the center of the rectangle.
func (r Rectangle) Translation() Transform {
	return Identity().Translate(float64(r.X)+float64(r.Width),
		float64(r.Y)+float64(r.Height))
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
		r.Height -= length
		return r, Rectangle{X: r.X, Y: r.Y + r.Height, Width: r.Width, Height: length}
	case East:
		r.Width -= length
		return r, Rectangle{X: r.X + r.Width, Y: r.Y, Width: length, Height: r.Height}
	case South:
		r.Height -= length
		ret := Rectangle{X: r.X, Y: r.Y, Width: r.Width, Height: length}
		r.Y += length
		return r, ret
	case West:
		r.Width -= length
		ret := Rectangle{X: r.X, Y: r.Y, Width: length, Height: r.Height}
		r.X += length
		return r, ret
	default:
		panic("illegal edge (must be North, East, South or West)")
	}
}

// Fill fills the rectangle with the given color.
func (r Rectangle) Fill(renderer Renderer, color colors.RGBA) {
	renderer.FillRect(r.Width, r.Height, r.Translation(), color)
}
