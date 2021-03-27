// +build !js

package api

import (
	"fmt"

	"gopkg.in/yaml.v3"
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
		return fmt.Errorf("unknown font style: %s", name)
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
		return nil, fmt.Errorf("unknown font style: %v", fs)
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
		return fmt.Errorf("unknown font size: %s", name)
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
		return nil, fmt.Errorf("unknown font size: %v", fs)
	}
}
