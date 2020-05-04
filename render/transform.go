package render

import "math"

// Transform encodes a 2D affine transformation matrix
type Transform [6]float64

// Identity returns the identity transformation
func Identity() Transform {
	return Transform{1.0, 0.0, 0.0, 1.0, 0.0, 0.0}
}

// Translate adds a translation to the transformation.
// A translation moves by (x,y).
func (t *Transform) Translate(x float64, y float64) {
	t[4] += x
	t[5] += y
}

// Rotate adds a rotation to the transformation.
// The given angle is in radian.
func (t *Transform) Rotate(angle float64) {
	s := math.Sin(angle)
	c := math.Cos(angle)
	old := [4]float64{t[0], t[1], t[2], t[3]}
	t[0] = c*old[0] + s*old[2]
	t[1] = s*old[3] + c*old[1]
	t[2] = c*old[2] - s*old[0]
	t[3] = c*old[3] - s*old[1]
}

// Scale adds scaling to the transformation.
func (t *Transform) Scale(x float64, y float64) {
	t[0] *= x
	t[1] *= x
	t[2] *= y
	t[3] *= y
}

// Invert inverts the transformation
func (t *Transform) Invert() {
	d := t[0]*t[3] - t[1]*t[2]
	old := *t
	t[0] = old[3] / d
	t[1] = -old[1] / d
	t[2] = -old[2] / d
	t[3] = old[0] / d
	t[4] = (old[2]*old[5] - old[3]*old[4]) / d
	t[5] = (old[1]*old[4] - old[0]*old[5]) / d
}

// Compose adds the given transformation to the current transformation
func (t *Transform) Compose(other *Transform) {
	old := *t
	t[0] = other[0]*old[0] + other[1]*old[2]
	t[1] = other[1]*old[3] + other[0]*old[1]
	t[2] = other[2]*old[0] + other[3]*old[2]
	t[3] = other[3]*old[3] + other[2]*old[1]
	t[4] = other[4]*old[0] + other[5]*old[2] + old[4]
	t[5] = other[5]*old[3] + other[4]*old[1] + old[5]
}
