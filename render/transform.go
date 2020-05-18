package render

import "math"

// Transform encodes a 2D affine transformation matrix
type Transform [6]float32

// Identity returns the identity transformation
func Identity() Transform {
	return Transform{1.0, 0.0, 0.0, 1.0, 0.0, 0.0}
}

// Translate adds a translation to the transformation.
// A translation moves by (x,y).
func (t Transform) Translate(x float32, y float32) Transform {
	t[4] += x
	t[5] += y
	return t
}

// Rotate adds a rotation to the transformation.
// The given angle is in radian.
func (t Transform) Rotate(angle float32) Transform {
	s := float32(math.Sin(float64(angle)))
	c := float32(math.Cos(float64(angle)))
	return Transform{
		c*t[0] + s*t[2], s*t[3] + c*t[1],
		c*t[2] - s*t[0], c*t[3] - s*t[1], t[4], t[5]}
}

// Scale adds scaling to the transformation.
func (t Transform) Scale(x float32, y float32) Transform {
	return Transform{
		t[0] * x, t[1] * x, t[2] * y, t[3] * y, t[4], t[5]}
}

// Invert inverts the transformation
func (t Transform) Invert() Transform {
	d := t[0]*t[3] - t[1]*t[2]
	return Transform{
		t[3] / d, -t[1] / d, -t[2] / d, t[0] / d,
		(t[2]*t[5] - t[3]*t[4]) / d, (t[1]*t[4] - t[0]*t[5]) / d}
}

// Compose adds the given transformation to the current transformation
func (t Transform) Compose(other Transform) Transform {
	return Transform{
		other[0]*t[0] + other[1]*t[2], other[1]*t[3] + other[0]*t[1],
		other[2]*t[0] + other[3]*t[2], other[3]*t[3] + other[2]*t[1],
		other[4]*t[0] + other[5]*t[2] + t[4], other[5]*t[3] + other[4]*t[1] + t[5]}
}
