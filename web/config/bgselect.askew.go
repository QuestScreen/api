package config

// Code generated by askew. DO NOT EDIT.

import (
	"syscall/js"

	"github.com/QuestScreen/api"
	"github.com/QuestScreen/api/web/controls"
	askew "github.com/flyx/askew/runtime"
)

var αBackgroundSelectTemplate = js.Global().Get("document").Call("createElement", "template")

func init() {
	αBackgroundSelectTemplate.Set("innerHTML", `
	<!--data-->
	<!--handlers-->
	<table class="qs-config-item-table">
		<thead>
			<tr>
				<th></th><th>Primary</th><th>Secondary</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<th>Color</th>
				<td><input type="color" name="primary-color" required=""/></td>
				<td><input type="color" name="secondary-color" required=""/></td>
			</tr>
			<tr>
				<th>Opacity</th>
				<td><input type="range" name="primary-opacity" min="0" max="255" step="1" required=""/></td>
				<td><input type="range" name="secondary-opacity" min="0" max="255" step="1" required=""/></td>
			</tr>
		</tbody>
	</table>
	<div class="qs-config-item-fragment">
		<label for="texture">Texture</label>
		<!--embed(texture)-->
	</div>
`)
}

// BackgroundSelect is a DOM component autogenerated by Askew
type BackgroundSelect struct {
	αcd              askew.ComponentData
	primaryColor     askew.StringValue
	pcDisabled       askew.BoolValue
	secondaryColor   askew.StringValue
	scDisabled       askew.BoolValue
	primaryOpacity   askew.IntValue
	poDisabled       askew.BoolValue
	secondaryOpacity askew.IntValue
	soDisabled       askew.BoolValue
	data             api.Background
	editHandler      EditHandler
	texture          controls.Dropdown
}

// FirstNode returns the first DOM node of this component.
// It implements the askew.Component interface.
func (o *BackgroundSelect) FirstNode() js.Value {
	return o.αcd.First()
}

// askewInit initializes the component, discarding all previous information.
// The component is initially a DocumentFragment until it gets inserted into
// the main document. It can be manipulated both before and after insertion.
func (o *BackgroundSelect) askewInit() {
	o.αcd.Init(αBackgroundSelectTemplate.Get("content").Call("cloneNode", true))

	o.primaryColor.BoundValue = askew.NewBoundProperty(&o.αcd, "value", 5, 3, 1, 3, 0)
	o.pcDisabled.BoundValue = askew.NewBoundProperty(&o.αcd, "disabled", 5, 3, 1, 3, 0)
	o.secondaryColor.BoundValue = askew.NewBoundProperty(&o.αcd, "value", 5, 3, 1, 5, 0)
	o.scDisabled.BoundValue = askew.NewBoundProperty(&o.αcd, "disabled", 5, 3, 1, 5, 0)
	o.primaryOpacity.BoundValue = askew.NewBoundProperty(&o.αcd, "value", 5, 3, 3, 3, 0)
	o.poDisabled.BoundValue = askew.NewBoundProperty(&o.αcd, "disabled", 5, 3, 3, 3, 0)
	o.secondaryOpacity.BoundValue = askew.NewBoundProperty(&o.αcd, "value", 5, 3, 3, 5, 0)
	o.soDisabled.BoundValue = askew.NewBoundProperty(&o.αcd, "disabled", 5, 3, 3, 5, 0)
	{
		src := o.αcd.Walk(5, 3, 1, 3, 0)
		{
			wrapper := js.FuncOf(func(this js.Value, arguments []js.Value) interface{} {
				o.αcalledited()
				return nil
			})
			src.Call("addEventListener", "input", wrapper)
		}
	}
	{
		src := o.αcd.Walk(5, 3, 1, 5, 0)
		{
			wrapper := js.FuncOf(func(this js.Value, arguments []js.Value) interface{} {
				o.αcalledited()
				return nil
			})
			src.Call("addEventListener", "input", wrapper)
		}
	}
	{
		src := o.αcd.Walk(5, 3, 3, 3, 0)
		{
			wrapper := js.FuncOf(func(this js.Value, arguments []js.Value) interface{} {
				o.αcalledited()
				return nil
			})
			src.Call("addEventListener", "input", wrapper)
		}
	}
	{
		src := o.αcd.Walk(5, 3, 3, 5, 0)
		{
			wrapper := js.FuncOf(func(this js.Value, arguments []js.Value) interface{} {
				o.αcalledited()
				return nil
			})
			src.Call("addEventListener", "input", wrapper)
		}
	}
	{
		container := o.αcd.Walk(7)
		o.texture.Init(controls.SelectAtMostOne, controls.SelectionIndicator, ``)
		o.texture.InsertInto(container, container.Get("childNodes").Index(3))
		o.texture.Controller = o
	}
}

// InsertInto inserts this component into the given object.
// The component will be in inserted state afterwards.
//
// The component will be inserted in front of 'before', or at the end if 'before' is 'js.Undefined()'.
func (o *BackgroundSelect) InsertInto(parent js.Value, before js.Value) {
	o.αcd.DoInsert(parent, before)
}

// Extract removes this component from its current parent.
// The component will be in initial state afterwards.
func (o *BackgroundSelect) Extract() {
	o.αcd.DoExtract()
}

// Destroy destroys this element (and all contained components). If it is
// currently inserted anywhere, it gets removed before.
func (o *BackgroundSelect) Destroy() {
	o.texture.Destroy()
	o.αcd.DoDestroy()
}

func (o *BackgroundSelect) αcalledited() {
	o.edited()
}

// BackgroundSelectList is a list of BackgroundSelect whose manipulation methods auto-update
// the corresponding nodes in the document.
type BackgroundSelectList struct {
	αmgr   askew.ListManager
	αitems []*BackgroundSelect
}

// Init initializes the list, discarding previous data.
// The list's items will be placed in the given container, starting at the
// given index.
func (l *BackgroundSelectList) Init(container js.Value, index int) {
	l.αmgr = askew.CreateListManager(container, index)
	l.αitems = nil
}

// Len returns the number of items in the list.
func (l *BackgroundSelectList) Len() int {
	return len(l.αitems)
}

// Item returns the item at the current index.
func (l *BackgroundSelectList) Item(index int) *BackgroundSelect {
	return l.αitems[index]
}

// Append appends the given item to the list.
func (l *BackgroundSelectList) Append(item *BackgroundSelect) {
	if item == nil {
		panic("cannot append nil to list")
	}
	l.αmgr.Append(item)
	l.αitems = append(l.αitems, item)
	return
}

// Insert inserts the given item at the given index into the list.
func (l *BackgroundSelectList) Insert(index int, item *BackgroundSelect) {
	var prev js.Value
	if index < len(l.αitems) {
		prev = l.αitems[index].αcd.First()
	}
	if item == nil {
		panic("cannot insert nil into list")
	}
	l.αmgr.Insert(item, prev)
	l.αitems = append(l.αitems, nil)
	copy(l.αitems[index+1:], l.αitems[index:])
	l.αitems[index] = item
	return
}

// Remove removes the item at the given index from the list and returns it.
func (l *BackgroundSelectList) Remove(index int) *BackgroundSelect {
	item := l.αitems[index]
	item.Extract()
	copy(l.αitems[index:], l.αitems[index+1:])
	l.αitems = l.αitems[:len(l.αitems)-1]
	return item
}

// Destroy destroys the item at the given index and removes it from the list.
func (l *BackgroundSelectList) Destroy(index int) {
	item := l.αitems[index]
	item.Destroy()
	copy(l.αitems[index:], l.αitems[index+1:])
	l.αitems = l.αitems[:len(l.αitems)-1]
}

// DestroyAll destroys all items in the list and empties it.
func (l *BackgroundSelectList) DestroyAll() {
	for _, item := range l.αitems {
		item.Destroy()
	}
	l.αitems = l.αitems[:0]
}

// OptionalBackgroundSelect is a nillable embeddable container for BackgroundSelect.
type OptionalBackgroundSelect struct {
	αcur *BackgroundSelect
	αmgr askew.ListManager
}

// Init initializes the container to be empty.
// The contained item, if any, will be placed in the given container at the
// given index.
func (o *OptionalBackgroundSelect) Init(container js.Value, index int) {
	o.αmgr = askew.CreateListManager(container, index)
	o.αcur = nil
}

// Item returns the current item, or nil if no item is assigned
func (o *OptionalBackgroundSelect) Item() *BackgroundSelect {
	return o.αcur
}

// Set sets the contained item destroying the current one.
// Give nil as value to simply destroy the current item.
func (o *OptionalBackgroundSelect) Set(value *BackgroundSelect) {
	if o.αcur != nil {
		o.αcur.Destroy()
	}
	o.αcur = value
	if value != nil {
		o.αmgr.Append(value)
	}
}

// Remove removes the current item and returns it.
// Returns nil if there is no current item.
func (o *OptionalBackgroundSelect) Remove() askew.Component {
	if o.αcur != nil {
		ret := o.αcur
		ret.Extract()
		o.αcur = nil
		return ret
	}
	return nil
}
