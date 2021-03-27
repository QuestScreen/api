// +build !js

package config

import (
	"log"

	"github.com/QuestScreen/api"
	"github.com/QuestScreen/api/server"
	"gopkg.in/yaml.v3"
)

type persistedFont struct {
	Family string        `yaml:"family"`
	Size   api.FontSize  `yaml:"size"`
	Style  api.FontStyle `yaml:"style"`
	Color  api.RGBA      `yaml:"color"`
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

// Persist returns a view that gives the family name as string.
func (f *FontSelect) Persist(ctx server.Context) interface{} {
	return &persistedFont{
		Family: ctx.FontFamilyName(f.FamilyIndex),
		Size:   f.Size,
		Style:  f.Style,
		Color:  f.Color,
	}
}

// Load loads a background from a YAML input
// `{primary: <rgb>, secondary: <rgb>, texture: <name>}`
func (b *BackgroundSelect) Load(
	input *yaml.Node, ctx server.Context) error {
	var value persistedBackground
	if err := input.Decode(&value); err != nil {
		return err
	}
	b.Primary = value.Primary
	b.Secondary = value.Secondary
	b.TextureIndex = -1
	if value.Texture != "" {
		textures := ctx.GetTextures()
		for i := range textures {
			if textures[i].Name == value.Texture {
				b.TextureIndex = i
				break
			}
		}
		if b.TextureIndex == -1 {
			log.Println("Unknown texture: " + value.Texture)
		}
	}
	return nil
}

// Persist returns a view that gives the texture name as string.
func (b *BackgroundSelect) Persist(ctx server.Context) interface{} {
	ret := &persistedBackground{
		Primary: b.Primary, Secondary: b.Secondary,
	}
	if b.TextureIndex != -1 {
		ret.Texture = ctx.GetTextures()[b.TextureIndex].Name
	}
	return ret
}
