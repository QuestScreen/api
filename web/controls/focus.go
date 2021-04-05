package controls

import "syscall/js"

// FocusHolder is a UI component that needs to react when the focus leaves its
// root element.
type FocusHolder interface {
	// Root returns the root element of the component.
	Root() js.Value
	// FocusLeaving is called when the focus is leaving the root element, but
	// before the focus has actually left.
	FocusLeaving()
}

var curHolder FocusHolder

func SetFocusHolder(holder FocusHolder) {
	if curHolder != nil && holder != nil {
		js.Global().Get("console").Call(
			"log", "[warn] SetHolder called when curHolder was not nil")
	}
	curHolder = holder
}

var cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if curHolder != nil {
		if !curHolder.Root().Call("contains", args[0].Get("relatedTarget")).Bool() {
			curHolder.FocusLeaving()
			curHolder = nil
		}
	}
	return nil
})

func init() {
	js.Global().Get("document").Call("addEventListener", "focusout", cb)
}
