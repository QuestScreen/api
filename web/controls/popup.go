package controls

import (
	"syscall/js"

	askew "github.com/flyx/askew/runtime"
)

// PopupController is an interface for controlling a popup.
type PopupController interface {
	Confirm()
	Cancel()
	NeedsDoShow() bool
	DoShow()
}

func (pb *PopupBase) show(title string, content askew.Component, confirmCaption, cancelCaption string) {
	pb.Title.Set(title)
	pb.Content.Set(content)
	pb.ConfirmCaption.Set(confirmCaption)
	if cancelCaption == "" {
		pb.cancelVisible.Set("hidden")
	} else {
		pb.cancelVisible.Set("visible")
		pb.CancelCaption.Set(cancelCaption)
	}
	if pb.Controller != nil && pb.Controller.NeedsDoShow() {
		pb.Visibility.Set("hidden")
		pb.Display.Set("flex")
		pb.Controller.DoShow()
		// this is required to avoid flickering. I have no idea why.
		// it doesn't work if the timeout simply removes style.visibility.
		pb.Display.Set("none")
		pb.Visibility.Set("")
		var f js.Func
		f = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			pb.Display.Set("flex")
			f.Release()
			return nil
		})
		js.Global().Call("setTimeout", f, 10)
	} else {
		pb.Display.Set("flex")
	}
}

func (pb *PopupBase) confirm() {
	if pb.Controller != nil {
		pb.Controller.Confirm()
	}
	pb.cleanup()
}

func (pb *PopupBase) cancel() {
	if pb.Controller != nil {
		pb.Controller.Cancel()
	}
	pb.cleanup()
}

func (pb *PopupBase) cleanup() {
	pb.Display.Set("")
	pb.Content.Set(nil)
	pb.Controller = nil
}

// ErrorMsg shows the popup containing the given text titled as 'Error'.
// Does not block.
func (pb *PopupBase) ErrorMsg(text string) {
	pb.show("Error", newPopupText(text), "OK", "")
}

type confirmController struct {
	val chan bool
}

func (cc *confirmController) Confirm() {
	cc.val <- true
}

func (cc *confirmController) Cancel() {
	cc.val <- false
}

func (cc *confirmController) NeedsDoShow() bool {
	return false
}

func (cc *confirmController) DoShow() {}

// Confirm shows the popup and returns true if the user clicks OK, false if
// Cancel. Blocking, must be called from a goroutine.
func (pb *PopupBase) Confirm(text string) bool {
	c := &confirmController{make(chan bool, 1)}
	pb.Controller = c
	pb.show("Confirm", newPopupText(text), "OK", "Cancel")
	ret := <-c.val
	pb.Controller = nil
	return ret
}

type textInputController struct {
	val   chan *string
	input *popupInput
}

func (tic *textInputController) Confirm() {
	str := tic.input.Value.Get()
	tic.val <- &str
}

func (tic *textInputController) Cancel() {
	tic.val <- nil
}

func (tic *textInputController) NeedsDoShow() bool {
	return false
}

func (tic *textInputController) DoShow() {}

// TextInput shows the popup and returns the entered string if the user clicks
// OK, nil if Cancel. Blocking, must be called from a goroutine.
func (pb *PopupBase) TextInput(title, label string) *string {
	tic := &textInputController{make(chan *string, 1), newPopupInput(label)}
	pb.Controller = tic
	pb.show(title, tic.input, "OK", "Cancel")
	ret := <-tic.val
	pb.Controller = nil
	return ret
}
