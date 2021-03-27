package config

import (
	"encoding/json"

	"github.com/QuestScreen/api"
	"github.com/QuestScreen/api/comms"
	"github.com/QuestScreen/api/server"
)

// FontSelect is an Item that allows the user to select a font family, size,
// style and color.
type FontSelect struct {
	api.Font
}

type webFont struct {
	FamilyIndex comms.ValidatedInt `json:"familyIndex"`
	Size        comms.ValidatedInt `json:"size"`
	Style       comms.ValidatedInt `json:"style"`
	Color       api.RGBA           `json:"color"`
}

// NewFontSelect creates a new FontSelect item with the given values
func NewFontSelect(familyIndex int, size api.FontSize, style api.FontStyle,
	color api.RGBA) *FontSelect {
	return &FontSelect{Font: api.Font{FamilyIndex: familyIndex, Size: size,
		Style: style, Color: color}}
}

// Receive loads a font from a json input
// `{"familyIndex": <number>, "size": <number>, "style": <number>}`
func (f *FontSelect) Receive(
	input json.RawMessage, ctx server.Context) error {
	tmp := webFont{
		FamilyIndex: comms.ValidatedInt{Min: 0, Max: ctx.NumFontFamilies() - 1},
		Size:        comms.ValidatedInt{Min: 0, Max: int(api.HugeFont)},
		Style:       comms.ValidatedInt{Min: 0, Max: int(api.BoldItalicFont)},
	}
	if err := comms.ReceiveData(input, &tmp); err != nil {
		return err
	}
	f.Font = api.Font{FamilyIndex: tmp.FamilyIndex.Value,
		Size:  api.FontSize(tmp.Size.Value),
		Style: api.FontStyle(tmp.Style.Value),
		Color: tmp.Color}
	return nil
}

// Send returns the object itself.
func (f *FontSelect) Send(ctx server.Context) interface{} {
	return f.Font
}
