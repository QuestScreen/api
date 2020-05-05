package config

import (
	"encoding/json"
	"log"

	"github.com/QuestScreen/api/common"
	"github.com/QuestScreen/api/server"
	"gopkg.in/yaml.v3"
)

// Font is a ConfigItem that allow the user to select a font family.
type Font struct {
	FamilyIndex int              `json:"familyIndex"`
	Size        common.FontSize  `json:"size"`
	Style       common.FontStyle `json:"style"`
	Color       common.RGBAColor `json:"color"`
}

type persistedFont struct {
	Family string           `yaml:"family"`
	Size   common.FontSize  `yaml:"size"`
	Style  common.FontStyle `yaml:"style"`
	Color  common.RGBAColor `yaml:"color"`
}

type webFont struct {
	FamilyIndex server.ValidatedInt `json:"familyIndex"`
	Size        server.ValidatedInt `json:"size"`
	Style       server.ValidatedInt `json:"style"`
	Color       common.RGBAColor    `json:"color"`
}

// LoadWeb loads a font from a json input
// `{"familyIndex": <number>, "size": <number>, "style": <number>}`
func (f *Font) LoadWeb(
	input json.RawMessage, ctx server.Context) common.SendableError {
	tmp := webFont{
		FamilyIndex: server.ValidatedInt{Min: 0, Max: ctx.NumFontFamilies() - 1},
		Size:        server.ValidatedInt{Min: 0, Max: int(common.HugeFont)},
		Style:       server.ValidatedInt{Min: 0, Max: int(common.BoldItalic)},
	}
	if err := server.ReceiveData(input, &tmp); err != nil {
		return err
	}
	*f = Font{FamilyIndex: tmp.FamilyIndex.Value,
		Size:  common.FontSize(tmp.Size.Value),
		Style: common.FontStyle(tmp.Style.Value),
		Color: tmp.Color}
	return nil
}

// LoadPersisted loads a selectable font from a YAML input
// `{family: <string>, size: <number>, style: <number>}`
func (f *Font) LoadPersisted(
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

// WebView returns the object itself.
func (f *Font) WebView(ctx server.Context) interface{} {
	return f
}

// PersistingView returns a view that gives the family name as string.
func (f *Font) PersistingView(ctx server.Context) interface{} {
	return &persistedFont{
		Family: ctx.FontFamilyName(f.FamilyIndex),
		Size:   f.Size,
		Style:  f.Style,
		Color:  f.Color,
	}
}
