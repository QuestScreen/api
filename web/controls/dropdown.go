package controls

import (
	"strconv"

	"github.com/QuestScreen/api/web"
)

// SelectorKind defines how items in a dropdown menu are selected.
type SelectorKind int

const (
	// SelectAtMostOne is like SelectOne but adds an additional item „None“ to the
	// list of selectable items.
	SelectAtMostOne SelectorKind = iota
	// SelectOne is used when selecting a new item unselects the previous one.
	SelectOne
	// SelectMultiple is used when multiple items can be selected at the same time.
	SelectMultiple
)

// IndicatorKind defines what kind of indicator is displayed in front of a menu
// item depending on its selection status.
type IndicatorKind int

const (
	// SelectionIndicator shows a general „selected“ icon if the item is selected.
	SelectionIndicator IndicatorKind = iota
	// VisibilityIndicator shows a visibility icon if the item is selected.
	VisibilityIndicator
	// InvisibilityIndicator shows an invisibility icon if the item is deselected.
	InvisibilityIndicator
)

// NewDropdown creates a new Dropdown and initializes it.
func NewDropdown(kind SelectorKind, indicator IndicatorKind, caption string) *Dropdown {
	ret := new(Dropdown)
	ret.Init(kind, indicator, caption)
	return ret
}

// Init initializes the Dropdown.
func (d *Dropdown) Init(kind SelectorKind, indicator IndicatorKind, caption string) {
	d.askewInit(kind, indicator)
	switch kind {
	case SelectAtMostOne:
		d.items.Append(newDropdownItem(indicator, true, "None", -1))
	case SelectMultiple:
		d.caption.Set(caption)
	}
}

func (d *Dropdown) click() {
	if !d.Disabled.Get() {
		d.Toggle()
	}
}

// Toggle toggles the state of the dropdown (open/closed)
func (d *Dropdown) Toggle() {
	if d.opened.Get() {
		d.opened.Set(false)
		if web.InSmartphoneMode() {
			d.menuHeight.Set("")
		}
		d.link.Get().Call("blur")
	} else {
		d.opened.Set(true)
		if web.InSmartphoneMode() {
			d.menuHeight.Set(strconv.Itoa(d.items.Len()*2) + "em")
		}
	}
}

// Hide hides the dropdown menu
func (d *Dropdown) Hide() {
	if d.opened.Get() {
		d.Toggle()
	}
}

func (d *Dropdown) clickItem(index int) {
	if d.Controller != nil {
		go func() {
			newVal := d.Controller.ItemClicked(index)
			d.SetItem(index, newVal)
		}()
	} else {
		d.SetItem(index, true)
	}
	if d.kind != SelectMultiple {
		// auto-hide dropdown unless it's multi select
		d.Hide()
	}
}

// SetItem sets the value of an item.
// For single-select dropdowns, this does nothing for value == false (to
// unselect the current item, select another one or -1 for SelectAtMostOne
// dropdowns).
func (d *Dropdown) SetItem(index int, value bool) {
	var itemIndex int
	if d.kind == SelectAtMostOne {
		itemIndex = index + 1
	} else {
		itemIndex = index
	}

	if d.kind == SelectMultiple {
		item := d.items.Item(itemIndex)
		item.Selected.Set(value)
	} else {
		for i := 0; i < d.items.Len(); i++ {
			if i == itemIndex {
				item := d.items.Item(i)
				item.Selected.Set(true)
				d.caption.Set(item.caption.Get())
			} else {
				d.items.Item(i).Selected.Set(false)
			}
		}
		d.emphCaption.Set(index == -1)
		d.CurIndex = index
	}
}

// AddItem adds an item of the given name to the dropdown list.
func (d *Dropdown) AddItem(name string, selected bool) {
	var index int
	if d.kind == SelectAtMostOne {
		index = d.items.Len() - 1
	} else {
		index = d.items.Len()
	}
	item := newDropdownItem(d.indicator, false, name, index)
	item.Selected.Set(selected)
	if selected {
		d.CurIndex = index
	}
	d.items.Append(item)
}
