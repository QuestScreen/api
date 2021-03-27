package api

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

// Font describes the font used for drawing text.
type Font struct {
	FamilyIndex int       `json:"familyIndex"`
	Size        FontSize  `json:"size"`
	Style       FontStyle `json:"style"`
	Color       RGBA      `json:"color"`
}
