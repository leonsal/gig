package gig

import "github.com/leonsal/gig/imgui"

type ColorEdit3 struct {
	Widget
	label    imgui.CString
	color    imgui.Vec4
	flags    imgui.ColorEditFlags
	onchange func(*ColorEdit3)
}

func NewColorEdit3(label string, color imgui.Vec4) *ColorEdit3 {

	c := new(ColorEdit3)
	c.Init(c)
	c.label.SetString(label)
	c.color = color
	c.SetRender(func() {
		if imgui.ColorEdit3CS(c.label, &c.color.X, c.flags) {
			if c.onchange != nil {
				c.onchange(c)
			}
		}
	})
	return c
}

func (c *ColorEdit3) SetLabel(label string) *ColorEdit3 {

	c.label.SetString(label)
	return c
}

func (c *ColorEdit3) Color() imgui.Vec4 {

	return c.color
}

func (c *ColorEdit3) SetColor(color imgui.Vec4) *ColorEdit3 {

	c.color = color
	return c
}

func (c *ColorEdit3) Flags() imgui.ColorEditFlags {

	return c.flags
}

func (c *ColorEdit3) SetFlags(flags imgui.ColorEditFlags) *ColorEdit3 {

	c.flags = flags
	return c
}

func (c *ColorEdit3) SetOnChange(f func(*ColorEdit3)) *ColorEdit3 {

	c.onchange = f
	return c
}

type ColorEdit4 struct {
	Widget
	label    imgui.CString
	color    imgui.Vec4
	flags    imgui.ColorEditFlags
	onchange func(*ColorEdit4)
}

func NewColorEdit4(label string, color imgui.Vec4) *ColorEdit4 {

	c := new(ColorEdit4)
	c.Init(c)
	c.label.SetString(label)
	c.color = color
	c.SetRender(func() {
		if imgui.ColorEdit4CS(c.label, &c.color.X, c.flags) {
			if c.onchange != nil {
				c.onchange(c)
			}
		}
	})
	return c
}

func (c *ColorEdit4) SetLabel(label string) *ColorEdit4 {

	c.label.SetString(label)
	return c
}

func (c *ColorEdit4) Color() imgui.Vec4 {

	return c.color
}

func (c *ColorEdit4) SetColor(color imgui.Vec4) *ColorEdit4 {

	c.color = color
	return c
}

func (c *ColorEdit4) Flags() imgui.ColorEditFlags {

	return c.flags
}

func (c *ColorEdit4) SetFlags(flags imgui.ColorEditFlags) *ColorEdit4 {

	c.flags = flags
	return c
}

func (c *ColorEdit4) SetOnChange(f func(*ColorEdit4)) *ColorEdit4 {

	c.onchange = f
	return c
}

type ColorPicker3 struct {
	Widget
	label    imgui.CString
	color    imgui.Vec4
	flags    imgui.ColorEditFlags
	onchange func(*ColorPicker3)
}

func NewColorPicker3(label string, color imgui.Vec4) *ColorPicker3 {

	c := new(ColorPicker3)
	c.Init(c)
	c.label.SetString(label)
	c.color = color
	c.SetRender(func() {
		if imgui.ColorPicker3CS(c.label, &c.color.X, c.flags) {
			if c.onchange != nil {
				c.onchange(c)
			}
		}
	})
	return c
}

func (c *ColorPicker3) SetLabel(label string) *ColorPicker3 {

	c.label.SetString(label)
	return c
}

func (c *ColorPicker3) Color() imgui.Vec4 {

	return c.color
}

func (c *ColorPicker3) SetColor(color imgui.Vec4) *ColorPicker3 {

	c.color = color
	return c
}

func (c *ColorPicker3) Flags() imgui.ColorEditFlags {

	return c.flags
}

func (c *ColorPicker3) SetFlags(flags imgui.ColorEditFlags) *ColorPicker3 {

	c.flags = flags
	return c
}

func (c *ColorPicker3) SetOnChange(f func(*ColorPicker3)) *ColorPicker3 {

	c.onchange = f
	return c
}

type ColorPicker4 struct {
	Widget
	label    imgui.CString
	color    imgui.Vec4
	refCol   *imgui.Vec4
	pref     *float32
	flags    imgui.ColorEditFlags
	onchange func(*ColorPicker4)
}

func NewColorPicker4(label string, color imgui.Vec4) *ColorPicker4 {

	c := new(ColorPicker4)
	c.Init(c)
	c.label.SetString(label)
	c.color = color
	c.SetRender(func() {
		if imgui.ColorPicker4CS(c.label, &c.color.X, c.flags, c.pref) {
			if c.onchange != nil {
				c.onchange(c)
			}
		}
	})
	return c
}

func (c *ColorPicker4) SetLabel(label string) *ColorPicker4 {

	c.label.SetString(label)
	return c
}

func (c *ColorPicker4) Color() imgui.Vec4 {

	return c.color
}

func (c *ColorPicker4) SetColor(color imgui.Vec4) *ColorPicker4 {

	c.color = color
	return c
}

func (c *ColorPicker4) SetRefColor(color *imgui.Vec4) *ColorPicker4 {

	c.refCol = color
	if c.refCol != nil {
		c.pref = &c.refCol.X
	} else {
		c.pref = nil
	}
	return c
}

func (c *ColorPicker4) Flags() imgui.ColorEditFlags {

	return c.flags
}

func (c *ColorPicker4) SetFlags(flags imgui.ColorEditFlags) *ColorPicker4 {

	c.flags = flags
	return c
}

func (c *ColorPicker4) SetOnChange(f func(*ColorPicker4)) *ColorPicker4 {

	c.onchange = f
	return c
}

type ColorButton struct {
	Widget
	descId  imgui.CString
	color   imgui.Vec4
	flags   imgui.ColorEditFlags
	size    imgui.Vec2
	onclick func(*ColorButton)
}

func NewColorButton(descId string, color imgui.Vec4) *ColorButton {

	b := new(ColorButton)
	b.Init(b)
	b.descId.SetString(descId)
	b.color = color
	b.SetRender(func() {
		if imgui.ColorButtonCS(b.descId, b.color, b.flags, b.size) {
			if b.onclick != nil {
				b.onclick(b)
			}
		}
	})
	return b
}

func (c *ColorButton) SetDescID(descId string) *ColorButton {

	c.descId.SetString(descId)
	return c
}

func (c *ColorButton) Color() imgui.Vec4 {

	return c.color
}

func (c *ColorButton) SetColor(color imgui.Vec4) *ColorButton {

	c.color = color
	return c
}

func (c *ColorButton) Flags() imgui.ColorEditFlags {

	return c.flags
}

func (c *ColorButton) SetFlags(flags imgui.ColorEditFlags) *ColorButton {

	c.flags = flags
	return c
}

func (c *ColorButton) SetOnClick(f func(*ColorButton)) *ColorButton {

	c.onclick = f
	return c
}
