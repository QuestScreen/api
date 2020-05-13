package colors

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)

// RGB represents a color in RGB color space.
type RGB struct {
	R uint8 `yaml:"r"`
	G uint8 `yaml:"g"`
	B uint8 `yaml:"b"`
}

// RGBA represents a color in RGB color space with alpha channel
type RGBA struct {
	R uint8 `yaml:"r"`
	G uint8 `yaml:"g"`
	B uint8 `yaml:"b"`
	A uint8 `yaml:"a"`
}

// WithAlpha generates an RGBAColor with the given alpha value.
func (c RGB) WithAlpha(alpha uint8) RGBA {
	return RGBA{R: c.R, G: c.G, B: c.B, A: alpha}
}

// AsBackground returns a Background with this color as background and no
// texture set.
func (c RGB) AsBackground() Background {
	return Background{
		Primary: c.WithAlpha(255), TextureIndex: -1}
}

// UnmarshalJSON loads a JSON string as HTML hexcode into RGBColor
func (c *RGB) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if len(s) != 7 || s[0] != '#' {
		return fmt.Errorf("\"%s\" is not a valid color hexcode", s)
	}
	bytes, err := hex.DecodeString(s[1:])
	if err != nil {
		return err
	}
	*c = RGB{bytes[0], bytes[1], bytes[2]}
	return nil
}

// MarshalJSON represents the color as JSON string containing a HTML hexcode
func (c *RGB) MarshalJSON() ([]byte, error) {
	bytes := [3]byte{c.R, c.G, c.B}
	s := "#" + hex.EncodeToString(bytes[:])
	return json.Marshal(&s)
}

// AsBackground returns a Background with this color as background and no
// texture set.
func (c RGBA) AsBackground() Background {
	return Background{Primary: c, TextureIndex: -1}
}

// UnmarshalJSON loads a JSON string as HTML hexcode into RGBAColor
func (c *RGBA) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if len(s) != 9 || s[0] != '#' {
		return fmt.Errorf("\"%s\" is not a valid color hexcode", s)
	}
	bytes, err := hex.DecodeString(s[1:])
	if err != nil {
		return err
	}
	*c = RGBA{R: bytes[0], G: bytes[1], B: bytes[2], A: bytes[3]}
	return nil
}

// MarshalJSON represents the color as JSON string containing a HTML hexcode
func (c *RGBA) MarshalJSON() ([]byte, error) {
	bytes := [4]byte{c.R, c.G, c.B, c.A}
	s := "#" + hex.EncodeToString(bytes[:])
	return json.Marshal(&s)
}
