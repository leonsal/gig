package dialog

import (
	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
)

type MsgDialog struct {
	gig.PopupModal
	text    *gig.TextWrapped
	buttons []*gig.Button
}

const dialogDefWidth = 300

func NewMsgDialog(title string, msg string, buttons ...string) *MsgDialog {

	d := new(MsgDialog)
	d.PopupModal.Init(d, title)
	d.text = gig.NewTextWrapped(msg)

	const spaceY = 10
	d.AddChildren(
		gig.NewDummy(imgui.Vec2{0, spaceY}),
		d.text,
		gig.NewDummy(imgui.Vec2{0, spaceY}),
	)
	for _, label := range buttons {
		b := gig.NewButton(label)
		d.buttons = append(d.buttons, b)
		d.AddChildren(b, gig.NewSameLine())
	}

	d.SetSize(imgui.Vec2{dialogDefWidth, -1}, imgui.Cond_Once)
	return d
}

func (d *MsgDialog) SetMsg(msg string) *MsgDialog {

	d.text.SetText(msg)
	return d
}

func (d *MsgDialog) SetOnClick(f func(*MsgDialog, int)) *MsgDialog {

	for idx, b := range d.buttons {
		index := idx
		b.SetOnClick(func(b *gig.Button) {
			f(d, index)
		})
	}
	return d
}

type InputDialog struct {
	gig.PopupModal
	input   *gig.InputText
	buttons []*gig.Button
}

func NewInputDialog(title string, label string, buttons ...string) *InputDialog {

	d := new(InputDialog)
	d.PopupModal.Init(d, title)

	d.input = gig.NewInputText("").
		SetFlags(imgui.InputTextFlags_EnterReturnsTrue)
	d.input.SetItemWidth(-1)

	const spaceY = 10
	d.AddChildren(
		gig.NewDummy(imgui.Vec2{0, spaceY}),
		gig.NewText(label),
		d.input,
		gig.NewDummy(imgui.Vec2{0, spaceY}),
	)

	for _, label := range buttons {
		b := gig.NewButton(label)
		d.buttons = append(d.buttons, b)
		d.AddChildren(b, gig.NewSameLine())
	}
	d.SetSize(imgui.Vec2{dialogDefWidth, -1}, imgui.Cond_Once)
	return d
}

func (d *InputDialog) SetOnEnter(f func(*InputDialog)) *InputDialog {

	d.input.SetOnChange(func(inp *gig.InputText) {
		f(d)
	})
	return d
}

func (d *InputDialog) SetOnClick(f func(*InputDialog, int)) *InputDialog {

	for idx, b := range d.buttons {
		index := idx
		b.SetOnClick(func(b *gig.Button) {
			f(d, index)
		})
	}
	return d
}

func (d *InputDialog) Text() string {

	return d.input.Text()
}

func (d *InputDialog) SetText(text string) *InputDialog {

	d.input.SetText(text)
	return d
}
