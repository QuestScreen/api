package config

import (
	"encoding/json"
	"log"

	"github.com/QuestScreen/api"
	"github.com/QuestScreen/api/comms"
	"github.com/QuestScreen/api/server"
	"gopkg.in/yaml.v3"
)

// FontSelect is an Item that allows the user to select a font family, size,
// style and color.
type FontSelect struct {
	api.Font
}

type persistedFont struct {
	Family string        `yaml:"family"`
	Size   api.FontSize  `yaml:"size"`
	Style  api.FontStyle `yaml:"style"`
	Color  api.RGBA      `yaml:"color"`
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

// Load loads a selectable font from a YAML input
// `{family: <string>, size: <number>, style: <number>}`
func (f *FontSelect) Load(
	input *yaml.Node, ctx server.Context) error {
	var tmp persistedFont
	if err := input.Decode(&tmp); err != nil {
		return err
	}
	f.Size = tmp.Size
	f.Style = tmp.Style
	f.Color = tmp.Color
	for i := 0; i < ctx.NumFontFamilies(); i++ {
		if tmp.Family == ctx.FontFamilyName(i) {
			f.FamilyIndex = i
			return nil
		}
	}
	log.Printf("unknown font \"%s\"\n", tmp.Family)
	f.FamilyIndex = 0
	return nil
}

// Send returns the object itself.
func (f *FontSelect) Send(ctx server.Context) interface{} {
	return f.Font
}

// Persist returns a view that gives the family name as string.
func (f *FontSelect) Persist(ctx server.Context) interface{} {
	return &persistedFont{
		Family: ctx.FontFamilyName(f.FamilyIndex),
		Size:   f.Size,
		Style:  f.Style,
		Color:  f.Color,
	}
}
