package config

import (
	"encoding/json"

	"github.com/QuestScreen/api"
	"github.com/QuestScreen/api/comms"
	"github.com/QuestScreen/api/server"
)

// BackgroundSelect is an Item that allows the user to define a background
// by setting a primary color and optionally, a secondary color together with a
// texture.
type BackgroundSelect struct {
	api.Background
}

// NewBackgroundSelect creates a new BackgroundSelect having the given primary
// color and no texture. This can be used for the default config of a module,
// since it requires *Background
// (and RGBA.AsBackground does provide Background).
func NewBackgroundSelect(value api.Background) *BackgroundSelect {
	return &BackgroundSelect{Background: value}
}

// Receive loads a background from a json input
// `{"primary": <rgb>, "secondary": <rgb>, "textureIndex": <number>}`
func (b *BackgroundSelect) Receive(
	input json.RawMessage, ctx server.Context) error {
	textures := ctx.GetTextures()
	value := struct {
		Primary      api.RGBA           `json:"primary"`
		Secondary    api.RGBA           `json:"secondary"`
		TextureIndex comms.ValidatedInt `json:"textureIndex"`
	}{TextureIndex: comms.ValidatedInt{Min: -1, Max: len(textures) - 1}}
	if err := comms.ReceiveData(input, &value); err != nil {
		return err
	}
	b.Background = api.Background{Primary: value.Primary,
		Secondary: value.Secondary, TextureIndex: value.TextureIndex.Value}
	return nil
}

type persistedBackground struct {
	Primary, Secondary api.RGBA
	Texture            string
}

// Send returns the object itself.
func (b *BackgroundSelect) Send(ctx server.Context) interface{} {
	return b.Background
}
