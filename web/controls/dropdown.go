package controls

import (
	"strconv"
	"syscall/js"

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

// Root implements controls.FocusHolder and returns the root list element that
// contains the dropdown list.
func (d *Dropdown) Root() js.Value {
	return d.rootItem.Get()
}

// FocusLeaving implements controls.FocusHolder and is called when the focus
// leaves this dropdown.
func (d *Dropdown) FocusLeaving() {
	if d.opened.Get() {
		d.toggle()
	}
}

func (d *Dropdown) toggle() bool {
	if d.opened.Get() {
		d.opened.Set(false)
		if web.InSmartphoneMode() {
			d.menuHeight.Set("")
		}
		return false
	} else {
		d.opened.Set(true)
		if web.InSmartphoneMode() {
			d.menuHeight.Set(strconv.Itoa(d.items.Len()*2) + "em")
		}
		SetFocusHolder(d)
		return true
	}
}

func (d *Dropdown) click() {
	if !d.Disabled.Get() {
		if !d.toggle() {
			SetFocusHolder(nil)
		}
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
		d.FocusLeaving()
		SetFocusHolder(nil)
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
