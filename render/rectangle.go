package render

import "github.com/QuestScreen/api/common"

// Rectangle describes a positioned rectangle.
// You can use it to center things inside some frame.
//
// This can also be used as intermediate structure for positioning stuff,
// to avoid having to deal with Transform's float64 values too much.
type Rectangle struct {
	// coordinate of the lower left corner
	X, Y          int32
	Width, Height int32
}

// Translation returns the transformation needed to move an object centered on
// the origin to the center of the rectangle.
func (r *Rectangle) Translation() Transform {
	ret := Identity()
	ret.Translate(float64(r.X)+float64(r.Width)/2.0,
		float64(r.Y)+float64(r.Height)/2.0)
	return ret
}

// Move moves the rectangle by the given delta
func (r *Rectangle) Move(dx, dy int32) {
	r.X += dx
	r.Y += dy
}

// Shrink removes dw from the rectangles width and dh from its height,
// repositioning it so that the center stays the same.
func (r *Rectangle) Shrink(dw, dh int32) {
	r.X += dw / 2
	r.Y += dh / 2
	r.Width -= dw
	r.Height -= dh
}

// Fill fills the rectangle with the given color.
func (r *Rectangle) Fill(renderer Renderer, color common.RGBAColor) {
	renderer.FillRect(r.Width, r.Height, r.Translation(), color)
}
