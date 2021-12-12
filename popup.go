package gig

import (
	"strconv"

	"github.com/leonsal/gig/imgui"
)

type Popup struct {
	Widget
	strId      imgui.CString
	flags      imgui.WindowFlags
	popupFlags imgui.PopupFlags
	open       bool
}

func NewPopup() *Popup {

	p := new(Popup)
	p.Init(p)
	p.strId.SetString(strconv.Itoa(p.ID()))
	p.SetRender(func() {
		if p.open {
			imgui.OpenPopupCS(p.strId, p.popupFlags)
			p.open = false
		}
		if imgui.BeginPopupCS(p.strId, p.flags) {
			p.RenderChildren()
			imgui.EndPopup()
		}
	})
	return p
}

func (p *Popup) Save(pp **Popup) *Popup {

	*pp = p
	return p
}

func (p *Popup) SetFlags(flags imgui.WindowFlags) *Popup {

	p.flags = flags
	return p
}

func (p *Popup) SetPopupFlags(flags imgui.PopupFlags) *Popup {

	p.popupFlags = flags
	return p
}

func (p *Popup) Open() *Popup {

	p.open = true
	return p
}

func CloseCurrentPopup() {

	imgui.CloseCurrentPopup()
}

type PopupModal struct {
	Widget
	title      imgui.CString
	flags      imgui.WindowFlags
	popupFlags imgui.PopupFlags
	size       imgui.Vec2
	sizeCond   imgui.Cond
	doOpen     bool
	doClose    bool
	open       bool
	popen      *bool
}

func NewPopupModal(name string) *PopupModal {

	p := new(PopupModal)
	p.Init(p, name)
	return p
}

func (p *PopupModal) Init(control IWidget, name string) {

	p.Widget.Init(control)
	p.title.SetString(name)
	p.SetRender(func() {
		if p.doOpen {
			imgui.OpenPopupCS(p.title, p.popupFlags)
			p.doOpen = false
		}
		if p.size.X != 0 && p.size.Y != 0 {
			imgui.SetNextWindowSize(p.size, p.sizeCond)
		}
		if imgui.BeginPopupModalCS(p.title, p.popen, p.flags) {
			if p.doClose {
				imgui.CloseCurrentPopup()
				p.doClose = false
			}
			p.RenderChildren()
			imgui.EndPopup()
		}
		if p.popen != nil && p.open == false {
			p.open = true
		}
	})
}

func (p *PopupModal) Save(pp **PopupModal) *PopupModal {

	*pp = p
	return p
}

func (p *PopupModal) SetTitle(title string) *PopupModal {

	p.title.SetString(title)
	return p
}

func (p *PopupModal) SetSize(size imgui.Vec2, cond imgui.Cond) *PopupModal {

	p.size = size
	p.sizeCond = cond
	return p
}

func (p *PopupModal) SetCanClose(canclose bool) *PopupModal {

	if canclose {
		p.open = true
		p.popen = &p.open
	} else {
		p.popen = nil
	}
	return p
}

func (p *PopupModal) Open() *PopupModal {

	p.doOpen = true
	return p
}

func (p *PopupModal) Close() *PopupModal {

	p.doClose = true
	return p
}

type Tooltip struct {
	Widget
}

func NewToolTip() *Tooltip {

	t := new(Tooltip)
	t.Init(t)
	t.SetRender(func() {
		imgui.BeginTooltip()
		t.RenderChildren()
		imgui.EndTooltip()
	})
	return t
}
