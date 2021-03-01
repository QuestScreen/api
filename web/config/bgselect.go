package config

import (
	"encoding/json"
	"fmt"

	"github.com/QuestScreen/api/server"
)

// NewBackgroundSelect creates a new BackgroundSelect widget and initializes it.
func NewBackgroundSelect(ctx server.Context, editHandler EditHandler) *BackgroundSelect {
	ret := new(BackgroundSelect)
	ret.Init(ctx, editHandler)
	return ret
}

// Init initializes the BackgroundSelect widget.
func (bg *BackgroundSelect) Init(ctx server.Context, editHandler EditHandler) {
	bg.askewInit()
	for i := 0; i < ctx.NumTextures(); i++ {
		bg.texture.AddItem(ctx.TextureName(i), false)
	}
}

// Receive loads the data in the given JSON input.
func (bg *BackgroundSelect) Receive(input json.RawMessage, ctx server.Context) error {
	if err := json.Unmarshal(input, &bg.data); err != nil {
		return err
	}
	bg.Reset()
	return nil
}

// Reset resets the UI to the values that have last been queried via Data().
// If the values have never been queried, the UI is reset to the initial
// data the state object was loaded with.
func (bg *BackgroundSelect) Reset() {
	bg.primaryColor.Set(fmt.Sprintf(
		"%02x%02x%02x", bg.data.Primary.R, bg.data.Primary.G, bg.data.Primary.B))
	bg.primaryOpacity.Set(int(bg.data.Primary.A))
	bg.secondaryColor.Set(fmt.Sprintf(
		"%02x%02x%02x", bg.data.Secondary.R, bg.data.Secondary.G, bg.data.Secondary.B))
	bg.secondaryOpacity.Set(int(bg.data.Secondary.A))
	bg.texture.SetItem(bg.data.TextureIndex, true)
}

// SetEnabled enables or disables the GUI.
func (bg *BackgroundSelect) SetEnabled(value bool) {
	bg.pcDisabled.Set(!value)
	bg.poDisabled.Set(!value)
	bg.scDisabled.Set(!value)
	bg.soDisabled.Set(!value)
	bg.texture.Disabled.Set(!value)
}

// Send returns an instance of api.Background
func (bg *BackgroundSelect) Send(ctx server.Context) interface{} {
	fmt.Sscanf(bg.primaryColor.Get(), "%02x%02x%02x", &bg.data.Primary.R,
		&bg.data.Primary.G, &bg.data.Primary.B)
	bg.data.Primary.A = uint8(bg.primaryOpacity.Get())
	fmt.Sscanf(bg.secondaryColor.Get(), "%02x%02x%02x", &bg.data.Secondary.R,
		&bg.data.Secondary.G, &bg.data.Secondary.B)
	bg.data.Secondary.A = uint8(bg.secondaryOpacity.Get())
	bg.data.TextureIndex = bg.texture.CurIndex
	return &bg.data
}
