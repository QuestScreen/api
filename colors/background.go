package colors

import (
	"encoding/json"
	"log"

	"github.com/QuestScreen/api/server"
	"gopkg.in/yaml.v3"
)

// Background is a config.Item that allows the user to
// define a background color by setting a primary color and optionally,
// a secondary color together with a texture.
type Background struct {
	Primary      RGBA `json:"primary"`
	Secondary    RGBA `json:"secondary"`
	TextureIndex int  `json:"textureIndex"`
}

// LoadWeb loads a background from a json input
// `{"primary": <rgb>, "secondary": <rgb>, "textureIndex": <number>}`
func (b *Background) LoadWeb(
	input json.RawMessage, ctx server.Context) server.SendableError {
	textures := ctx.GetTextures()
	value := struct {
		Primary      RGBA                `json:"primary"`
		Secondary    RGBA                `json:"secondary"`
		TextureIndex server.ValidatedInt `json:"textureIndex"`
	}{TextureIndex: server.ValidatedInt{Min: -1, Max: len(textures) - 1}}
	if err := server.ReceiveData(input, &value); err != nil {
		return err
	}
	*b = Background{Primary: value.Primary,
		Secondary: value.Secondary, TextureIndex: value.TextureIndex.Value}
	return nil
}

type persistedBackground struct {
	Primary, Secondary RGBA
	Texture            string
}

// LoadPersisted loads a background from a YAML input
// `{primary: <rgb>, secondary: <rgb>, texture: <name>}`
func (b *Background) LoadPersisted(
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
			if textures[i].Name() == value.Texture {
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

// WebView returns the object itself.
func (b *Background) WebView(ctx server.Context) interface{} {
	return b
}

// PersistingView returns a view that gives the texture name as string.
func (b *Background) PersistingView(ctx server.Context) interface{} {
	ret := &persistedBackground{
		Primary: b.Primary, Secondary: b.Secondary,
	}
	if b.TextureIndex != -1 {
		ret.Texture = ctx.GetTextures()[b.TextureIndex].Name()
	}
	return ret
}
