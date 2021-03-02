package controls

import (
	"syscall/js"

	askew "github.com/flyx/askew/runtime"
)

// PopupContent supplies the content of the popup and implements its controller.
type PopupContent interface {
	PopupBaseController
	askew.Component
}

// Show shows the given content inside the popup.
// Calling show does not block; you are responsible for awaiting for user
// confirmation or cancellation via the Cancel() / Confirm() callbacks of the
// content.
func (pb *PopupBase) Show(title string, content PopupContent, confirmCaption, cancelCaption string) {
	pb.Controller = content
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
// Blocks until user dismisses the message, must be called from a goroutine.
func (pb *PopupBase) ErrorMsg(text string) {
	pt := newPopupText(text)
	pb.Show("Error", pt, "OK", "")
	<-pt.val
	pt.Destroy()
}

func (pt *popupText) Confirm() {
	pt.val <- true
}

func (pt *popupText) Cancel() {
	pt.val <- false
}

func (pt *popupText) NeedsDoShow() bool {
	return false
}

func (pt *popupText) DoShow() {}

// Confirm shows the popup and returns true if the user clicks OK, false if
// Cancel. Blocking, must be called from a goroutine.
func (pb *PopupBase) Confirm(text string) bool {
	pt := newPopupText(text)
	pb.Show("Confirm", pt, "OK", "Cancel")
	ret := <-pt.val
	pt.Destroy()
	return ret
}

func (pi *popupInput) Confirm() {
	str := pi.Value.Get()
	pi.val <- &str
}

func (pi *popupInput) Cancel() {
	pi.val <- nil
}

func (pi *popupInput) NeedsDoShow() bool {
	return false
}

func (pi *popupInput) DoShow() {}

// TextInput shows the popup and returns the entered string if the user clicks
// OK, nil if Cancel. Blocking, must be called from a goroutine.
func (pb *PopupBase) TextInput(title, label string) *string {
	pi := newPopupInput(label)
	pb.Show(title, pi, "OK", "Cancel")
	ret := <-pi.val
	pi.Destroy()
	return ret
}
