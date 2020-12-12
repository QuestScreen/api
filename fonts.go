package api

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// FontStyle describes possible styles of a font
type FontStyle int

const (
	// RegularFont is the default font style
	RegularFont FontStyle = iota
	// BoldFont is the bold font style
	BoldFont
	// ItalicFont is the italic font style
	ItalicFont
	// BoldItalicFont is the bold and italic font style
	BoldItalicFont
	// NumFontStyles is not a valid Style, but used for iterating.
	NumFontStyles
)

// FontSize describes the size of a font.
// Font sizes are relative to the screen size.
type FontSize int

const (
	// SmallFont is the smallest size available
	SmallFont FontSize = iota
	// ContentFont is the size used for content text by default.
	ContentFont
	// MediumFont is a size between Content and Heading.
	MediumFont
	// HeadingFont is the size used for heading text by default.
	HeadingFont
	// LargeFont is a size larger than Heading.
	LargeFont
	// HugeFont is the largest font; usually used for displaying a single word
	// on the screen.
	HugeFont
	// NumFontSizes is not a valid Size, but used for iterating
	NumFontSizes
)

// UnmarshalYAML sets the font style from a YAML scalar
func (fs *FontStyle) UnmarshalYAML(value *yaml.Node) error {
	var name string
	if err := value.Decode(&name); err != nil {
		return err
	}
	switch name {
	case "Regular":
		*fs = RegularFont
	case "Bold":
		*fs = BoldFont
	case "Italic":
		*fs = ItalicFont
	case "BoldItalic":
		*fs = BoldItalicFont
	default:
		return fmt.Errorf("Unknown font style: %s", name)
	}
	return nil
}

// MarshalYAML maps the given font style to a string
func (fs FontStyle) MarshalYAML() (interface{}, error) {
	switch fs {
	case RegularFont:
		return "Regular", nil
	case BoldFont:
		return "Bold", nil
	case ItalicFont:
		return "Italic", nil
	case BoldItalicFont:
		return "BoldItalic", nil
	default:
		return nil, fmt.Errorf("Unknown font style: %v", fs)
	}
}

// UnmarshalYAML sets the font size from a YAML scalar
func (fs *FontSize) UnmarshalYAML(value *yaml.Node) error {
	var name string
	if err := value.Decode(&name); err != nil {
		return err
	}
	switch name {
	case "Small":
		*fs = SmallFont
	case "Content":
		*fs = ContentFont
	case "Medium":
		*fs = MediumFont
	case "Heading":
		*fs = HeadingFont
	case "Large":
		*fs = LargeFont
	case "Huge":
		*fs = HugeFont
	default:
		return fmt.Errorf("Unknown font size: %s", name)
	}
	return nil
}

// MarshalYAML maps the given font size to a string
func (fs FontSize) MarshalYAML() (interface{}, error) {
	switch fs {
	case SmallFont:
		return "Small", nil
	case ContentFont:
		return "Content", nil
	case MediumFont:
		return "Medium", nil
	case HeadingFont:
		return "Heading", nil
	case LargeFont:
		return "Large", nil
	case HugeFont:
		return "Huge", nil
	default:
		return nil, fmt.Errorf("Unknown font size: %v", fs)
	}
}
