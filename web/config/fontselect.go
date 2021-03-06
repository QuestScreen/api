package config

import (
	"encoding/json"

	"github.com/QuestScreen/api"
	"github.com/QuestScreen/api/server"
)

// NewFontSelect creates a new FontSelect widget and initializes it.
func NewFontSelect(ctx server.Context, editHandler EditHandler) Widget {
	ret := new(FontSelect)
	ret.Init(ctx, editHandler)
	return ret
}

// Init initializes the FontSelect widget.
func (fs *FontSelect) Init(ctx server.Context, editHandler EditHandler) {
	families := make([]string, ctx.NumFontFamilies())
	for i := 0; i < ctx.NumFontFamilies(); i++ {
		families[i] = ctx.FontFamilyName((i))
	}
	fs.askewInit(families)
}

// Receive loads the data given via input.
func (fs *FontSelect) Receive(input json.RawMessage, ctx server.Context) error {
	if err := json.Unmarshal(input, &fs.data); err != nil {
		return err
	}
	fs.Reset()
	return nil
}

// Reset resets the UI to the values that have last been queried via Send().
// If the values have never been queried, the UI is reset to the initial
// data the state object was loaded with.
func (fs *FontSelect) Reset() {
	fs.family.Set(fs.data.FamilyIndex)
	fs.size.Set(int(fs.data.Size))
	fs.bold.Set(fs.data.Style == api.BoldFont || fs.data.Style == api.BoldItalicFont)
	fs.italic.Set(fs.data.Style == api.ItalicFont || fs.data.Style == api.BoldItalicFont)
}

// SetEnabled enables or disables the GUI.
func (fs *FontSelect) SetEnabled(value bool) {
	fs.familiesDisabled.Set(!value)
	fs.sizeDisabled.Set(!value)
	fs.boldDisabled.Set(!value)
	fs.italicDisabled.Set(!value)
}

// Send returns an api.Font object containing the currently selected values.
func (fs *FontSelect) Send(ctx server.Context) interface{} {
	fs.data.FamilyIndex = fs.family.Get()
	fs.data.Size = api.FontSize(fs.size.Get())
	fs.data.Style = api.RegularFont
	if fs.bold.Get() {
		fs.data.Style++
	}
	if fs.italic.Get() {
		fs.data.Style += 2
	}
	return &fs.data
}

func (fs *FontSelect) toggleBold() {
	fs.bold.Set(!fs.bold.Get())
}

func (fs *FontSelect) toggleItalic() {
	fs.italic.Set(!fs.italic.Get())
}
