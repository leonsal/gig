package gig

import (
	"strconv"

	"github.com/leonsal/gig/imgui"
)

type TabBar struct {
	Widget
	strId imgui.CString
	flags imgui.TabBarFlags
}

func NewTabBar() *TabBar {

	t := new(TabBar)
	t.Init(t)
	t.strId.SetString(strconv.Itoa(t.ID()))
	t.SetRender(func() {
		if imgui.BeginTabBarCS(t.strId, t.flags) {
			t.RenderChildren()
			imgui.EndTabBar()
		}
	})
	return t
}

func (t *TabBar) Save(pt **TabBar) *TabBar {

	*pt = t
	return t
}

func (t *TabBar) Flags() imgui.TabBarFlags {

	return t.flags
}

func (t *TabBar) SetFlags(flags imgui.TabBarFlags) *TabBar {

	t.flags = flags
	return t
}

type TabItem struct {
	Widget
	label imgui.CString
	flags imgui.TabItemFlags
	open  bool
	popen *bool
}

func NewTabItem(label string) *TabItem {

	ti := new(TabItem)
	ti.Init(ti)
	ti.label.SetString(label)
	ti.SetRender(func() {
		if imgui.BeginTabItemCS(ti.label, ti.popen, ti.flags) {
			ti.RenderChildren()
			imgui.EndTabItem()
		}
		if ti.popen != nil && ti.open == false {
			ti.visible = false
			ti.open = true
		}
	})
	return ti
}

func (ti *TabItem) SetLabel(label string) *TabItem {

	ti.label.SetString(label)
	return ti
}

func (ti *TabItem) Flags() imgui.TabItemFlags {

	return ti.flags
}

func (ti *TabItem) SetFlags(flags imgui.TabItemFlags) *TabItem {

	ti.flags = flags
	return ti
}

func (ti *TabItem) SetCanClose(canclose bool) *TabItem {

	if canclose {
		ti.open = true
		ti.popen = &ti.open
	} else {
		ti.popen = nil
	}
	return ti
}

type TabItemButton struct {
	Widget
	label   imgui.CString
	flags   imgui.TabItemFlags
	onclick func(*TabItemButton)
}

func NewTabItemButton(label string) *TabItemButton {

	ib := new(TabItemButton)
	ib.Init(ib)
	ib.label.SetString(label)
	ib.SetRender(func() {
		if imgui.TabItemButtonCS(ib.label, ib.flags) {
			if ib.onclick != nil {
				ib.onclick(ib)
			}
		}
	})
	return ib
}

func (ib *TabItemButton) SetLabel(label string) *TabItemButton {

	ib.label.SetString(label)
	return ib
}

func (ib *TabItemButton) Flags() imgui.TabItemFlags {

	return ib.flags
}

func (ib *TabItemButton) SetFlags(flags imgui.TabItemFlags) *TabItemButton {

	ib.flags = flags
	return ib
}

func (ib *TabItemButton) SetOnClick(f func(b *TabItemButton)) *TabItemButton {

	ib.onclick = f
	return ib
}
