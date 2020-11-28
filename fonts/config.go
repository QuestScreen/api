package fonts

import (
	"encoding/json"
	"log"

	"github.com/QuestScreen/api/colors"
	"github.com/QuestScreen/api/comms"
	"github.com/QuestScreen/api/server"
	"gopkg.in/yaml.v3"
)

// Config is a config.Item that allows the user to select a font family, size,
// style and color.
type Config struct {
	FamilyIndex int         `json:"familyIndex"`
	Size        Size        `json:"size"`
	Style       Style       `json:"style"`
	Color       colors.RGBA `json:"color"`
}

type persistedFont struct {
	Family string      `yaml:"family"`
	Size   Size        `yaml:"size"`
	Style  Style       `yaml:"style"`
	Color  colors.RGBA `yaml:"color"`
}

type webFont struct {
	FamilyIndex comms.ValidatedInt `json:"familyIndex"`
	Size        comms.ValidatedInt `json:"size"`
	Style       comms.ValidatedInt `json:"style"`
	Color       colors.RGBA        `json:"color"`
}

// Name returns "Font"
func (f *Config) Name() string {
	return "Font"
}

// LoadWeb loads a font from a json input
// `{"familyIndex": <number>, "size": <number>, "style": <number>}`
func (f *Config) LoadWeb(
	input json.RawMessage, ctx server.Context) error {
	tmp := webFont{
		FamilyIndex: comms.ValidatedInt{Min: 0, Max: ctx.NumFontFamilies() - 1},
		Size:        comms.ValidatedInt{Min: 0, Max: int(Huge)},
		Style:       comms.ValidatedInt{Min: 0, Max: int(BoldItalic)},
	}
	if err := comms.ReceiveData(input, &tmp); err != nil {
		return err
	}
	*f = Config{FamilyIndex: tmp.FamilyIndex.Value,
		Size:  Size(tmp.Size.Value),
		Style: Style(tmp.Style.Value),
		Color: tmp.Color}
	return nil
}

// LoadPersisted loads a selectable font from a YAML input
// `{family: <string>, size: <number>, style: <number>}`
func (f *Config) LoadPersisted(
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
func (f *Config) WebView(ctx server.Context) interface{} {
	return f
}

// PersistingView returns a view that gives the family name as string.
func (f *Config) PersistingView(ctx server.Context) interface{} {
	return &persistedFont{
		Family: ctx.FontFamilyName(f.FamilyIndex),
		Size:   f.Size,
		Style:  f.Style,
		Color:  f.Color,
	}
}
