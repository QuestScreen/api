package common

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)

// RGBColor represents a color in RGB color space.
type RGBColor struct {
	Red   uint8 `yaml:"r"`
	Green uint8 `yaml:"g"`
	Blue  uint8 `yaml:"b"`
}

// RGBAColor represents a color in RGB color space with alpha channel
type RGBAColor struct {
	Red   uint8 `yaml:"r"`
	Green uint8 `yaml:"g"`
	Blue  uint8 `yaml:"b"`
	Alpha uint8 `yaml:"a"`
}

// WithAlpha generates an RGBAColor with the given alpha value.
func (c *RGBColor) WithAlpha(alpha uint8) RGBAColor {
	return RGBAColor{Red: c.Red, Green: c.Green, Blue: c.Blue, Alpha: alpha}
}

// UnmarshalJSON loads a JSON string as HTML hexcode into RGBColor
func (c *RGBColor) UnmarshalJSON(data []byte) error {
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
	c.Red = bytes[0]
	c.Green = bytes[1]
	c.Blue = bytes[2]
	return nil
}

// MarshalJSON represents the color as JSON string containing a HTML hexcode
func (c *RGBColor) MarshalJSON() ([]byte, error) {
	bytes := [3]byte{c.Red, c.Green, c.Blue}
	s := "#" + hex.EncodeToString(bytes[:])
	return json.Marshal(&s)
}

// UnmarshalJSON loads a JSON string as HTML hexcode into RGBAColor
func (c *RGBAColor) UnmarshalJSON(data []byte) error {
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
	c.Red = bytes[0]
	c.Green = bytes[1]
	c.Blue = bytes[2]
	c.Alpha = bytes[3]
	return nil
}

// MarshalJSON represents the color as JSON string containing a HTML hexcode
func (c *RGBAColor) MarshalJSON() ([]byte, error) {
	bytes := [4]byte{c.Red, c.Green, c.Blue, c.Alpha}
	s := "#" + hex.EncodeToString(bytes[:])
	return json.Marshal(&s)
}
