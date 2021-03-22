package config

import (
	"encoding/json"

	"github.com/QuestScreen/api"
	"github.com/QuestScreen/api/server"
)

// NewBackgroundSelect creates a new BackgroundSelect widget and initializes it.
func NewBackgroundSelect(ctx server.Context) Widget {
	ret := new(BackgroundSelect)
	ret.Init(ctx)
	return ret
}

// Init initializes the BackgroundSelect widget.
func (bg *BackgroundSelect) Init(ctx server.Context) {
	bg.askewInit()
	for _, t := range ctx.GetTextures() {
		bg.texture.AddItem(t.Name, false)
	}
	bg.texture.SetItem(-1, true)
}

func (bg *BackgroundSelect) SetEditHandler(editHandler EditHandler) {
	bg.editHandler = editHandler
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
	bg.primaryColor.Set(bg.data.Primary.WithoutAlpha().HexRepr())
	bg.primaryOpacity.Set(int(bg.data.Primary.A))
	bg.secondaryColor.Set(bg.data.Secondary.WithoutAlpha().HexRepr())
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
	var tmp api.RGB
	if err := tmp.FromHexRepr(bg.primaryColor.Get()); err != nil {
		panic(err)
	}
	bg.data.Primary = tmp.WithAlpha(uint8(bg.primaryOpacity.Get()))

	if err := tmp.FromHexRepr(bg.secondaryColor.Get()); err != nil {
		panic(err)
	}
	bg.data.Secondary = tmp.WithAlpha(uint8(bg.secondaryOpacity.Get()))

	bg.data.TextureIndex = bg.texture.CurIndex
	return &bg.data
}

func (bg *BackgroundSelect) edited() {
	bg.editHandler.Edited()
}

func (bg *BackgroundSelect) ItemClicked(index int) bool {
	bg.editHandler.Edited()
	return true
}
