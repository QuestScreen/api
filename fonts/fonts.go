package fonts

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// Style describes possible styles of a font
type Style int

const (
	// Regular is the default font style
	Regular Style = iota
	// Bold is the bold font style
	Bold
	// Italic is the italic font style
	Italic
	// BoldItalic is the bold and italic font style
	BoldItalic
	// NumStyles is not a valid Style, but used for iterating.
	NumStyles
)

// Size describes the size of a font.
// Font sizes are relative to the screen size.
type Size int

const (
	// Small is the smallest size available
	Small Size = iota
	// Content is the size used for content text by default.
	Content
	// Medium is a size between Content and Heading.
	Medium
	// Heading is the size used for heading text by default.
	Heading
	// Large is a size larger than Heading.
	Large
	// Huge is the largest font; usually used for displaying a single word
	// on the screen.
	Huge
	// NumSizes is not a valid Size, but used for iterating
	NumSizes
)

// UnmarshalYAML sets the font style from a YAML scalar
func (fs *Style) UnmarshalYAML(value *yaml.Node) error {
	var name string
	if err := value.Decode(&name); err != nil {
		return err
	}
	switch name {
	case "Regular":
		*fs = Regular
	case "Bold":
		*fs = Bold
	case "Italic":
		*fs = Italic
	case "BoldItalic":
		*fs = BoldItalic
	default:
		return fmt.Errorf("Unknown font style: %s", name)
	}
	return nil
}

// MarshalYAML maps the given font style to a string
func (fs Style) MarshalYAML() (interface{}, error) {
	switch fs {
	case Regular:
		return "Regular", nil
	case Bold:
		return "Bold", nil
	case Italic:
		return "Italic", nil
	case BoldItalic:
		return "BoldItalic", nil
	default:
		return nil, fmt.Errorf("Unknown font style: %v", fs)
	}
}

// UnmarshalYAML sets the font size from a YAML scalar
func (fs *Size) UnmarshalYAML(value *yaml.Node) error {
	var name string
	if err := value.Decode(&name); err != nil {
		return err
	}
	switch name {
	case "Small":
		*fs = Small
	case "Content":
		*fs = Content
	case "Medium":
		*fs = Medium
	case "Heading":
		*fs = Heading
	case "Large":
		*fs = Large
	case "Huge":
		*fs = Huge
	default:
		return fmt.Errorf("Unknown font size: %s", name)
	}
	return nil
}

// MarshalYAML maps the given font size to a string
func (fs Size) MarshalYAML() (interface{}, error) {
	switch fs {
	case Small:
		return "Small", nil
	case Content:
		return "Content", nil
	case Medium:
		return "Medium", nil
	case Heading:
		return "Heading", nil
	case Large:
		return "Large", nil
	case Huge:
		return "Huge", nil
	default:
		return nil, fmt.Errorf("Unknown font size: %v", fs)
	}
}
