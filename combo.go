package gig

import (
	"github.com/leonsal/gig/imgui"
)

type Combo struct {
	Widget
	label    imgui.CString
	preview  imgui.CString
	flags    imgui.ComboFlags
	onchange func(*Combo)
}

func NewCombo(label string) *Combo {

	c := new(Combo)
	c.Init(c)
	c.label.SetString(label)
	c.SetRender(func() {
		if imgui.BeginComboCS(c.label, c.preview, c.flags) {
			c.RenderChildren()
			imgui.EndCombo()
		}
	})
	return c
}

func (c *Combo) Save(pc **Combo) *Combo {

	*pc = c
	return c
}

func (c *Combo) ClearChildren() *Combo {

	c.Widget.ClearChildren()
	c.preview.SetString("")
	return c
}

func (c *Combo) AddChildren(children ...*Selectable) *Combo {

	for _, child := range children {
		child.SetOnClick(c.itemSelected)
		c.Widget.AddChildren(child)
	}
	return c
}

func (c *Combo) AddStrings(item ...string) *Combo {

	for _, it := range item {
		sel := NewSelectable(it)
		sel.SetOnClick(c.itemSelected)
		c.Widget.AddChildren(sel)
	}
	return c
}

func (c *Combo) Selected() *Selectable {

	for _, w := range c.Children() {
		sel := w.Super().(*Selectable)
		if sel.Selected() {
			return sel
		}
	}
	return nil
}

func (c *Combo) SelectedIndex() int {

	for idx, w := range c.Children() {
		sel := w.Super().(*Selectable)
		if sel.Selected() {
			return idx
		}
	}
	return -1
}

func (c *Combo) SetSelected(sel *Selectable) *Combo {

	for _, w := range c.Children() {
		curr := w.Super().(*Selectable)
		if curr == sel {
			c.preview.SetString(curr.Label())
			curr.SetSelected(true)
		} else {
			curr.SetSelected(false)
		}
	}
	return c
}

func (c *Combo) SetSelectedIndex(sel int) *Combo {

	if sel < 0 || sel > len(c.Children()) {
		return c
	}
	for idx, w := range c.Children() {
		curr := w.Super().(*Selectable)
		if idx == sel {
			c.preview.SetString(curr.Label())
			curr.SetSelected(true)
		} else {
			curr.SetSelected(false)
		}
	}
	return c
}

func (c *Combo) Flags() imgui.ComboFlags {

	return c.flags
}

func (c *Combo) SetFlags(flags imgui.ComboFlags) *Combo {

	c.flags = flags
	return c
}

func (c *Combo) SetOnChange(f func(*Combo)) *Combo {

	c.onchange = f
	return c
}

func (c *Combo) SetPreview(text string) *Combo {

	c.preview.SetString(text)
	return c
}

func (c *Combo) itemSelected(it *Selectable) {

	c.preview.SetString(it.Label())
	it.SetSelected(true)
	for _, w := range c.Children() {
		curr := w.Super().(*Selectable)
		if curr != it {
			curr.SetSelected(false)
		}
	}
	if c.onchange != nil {
		c.onchange(c)
	}
}
