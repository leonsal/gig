package gig

import (
	"github.com/leonsal/gig/imgui"
)

type Window struct {
	Widget
	title     imgui.CString
	open      bool
	popen     *bool
	flags     imgui.WindowFlags
	maximize  bool
	pos       imgui.Vec2
	posCond   imgui.Cond
	size      imgui.Vec2
	sizeCond  imgui.Cond
	rendering func(*Window)
}

func NewWindow(title string) *Window {

	w := new(Window)
	w.Init(title)
	return w
}

func (w *Window) Init(title string) *Window {

	w.Widget.Init(w)
	w.title.SetString(title)
	w.open = false
	w.pos = imgui.Vec2{-1, -1}
	w.size = imgui.Vec2{-1, -1}
	w.SetRender(func() {
		if w.maximize {
			vp := imgui.GetMainViewport()
			imgui.SetNextWindowPos(vp.WorkPos(), 0, imgui.Vec2{0, 0})
			imgui.SetNextWindowSize(vp.WorkSize(), 0)
		} else {
			if w.pos.X >= 0 && w.pos.Y >= 0 {
				imgui.SetNextWindowPos(w.pos, w.posCond, imgui.Vec2{0, 0})
			}
			if w.size.X >= 0 && w.size.Y >= 0 {
				imgui.SetNextWindowSize(w.size, w.sizeCond)
			}
		}
		if imgui.BeginCS(w.title, w.popen, w.flags) {
			if w.rendering != nil {
				w.rendering(w)
			}
			w.RenderChildren()
		}
		imgui.End()
		if w.popen != nil && w.open == false {
			w.visible = false
			w.open = true
		}
	})
	return w
}

func (w *Window) AddChildren(c ...IWidget) *Window {

	w.Widget.AddChildren(c...)
	return w
}

func (w *Window) Title() string {

	return w.title.String()
}

func (w *Window) SetTitle(title string) *Window {

	w.title.SetString(title)
	return w
}

func (w *Window) Flags() imgui.WindowFlags {

	return w.flags
}

func (w *Window) SetFlags(flags imgui.WindowFlags) *Window {

	w.flags = flags
	return w
}

func (w *Window) CanClose() bool {

	if w.popen != nil {
		return true
	}
	return false
}

func (w *Window) SetCanClose(canclose bool) *Window {

	if canclose {
		w.open = true
		w.popen = &w.open
	} else {
		w.popen = nil
	}
	return w
}

func (w *Window) SetMaximize(maximize bool) *Window {

	w.maximize = maximize
	return w
}

func (w *Window) SetPos(pos imgui.Vec2, cond imgui.Cond) {

	w.pos = pos
	w.posCond = cond
}

func (w *Window) SetSize(size imgui.Vec2, cond imgui.Cond) {

	w.size = size
	w.sizeCond = cond
}

func (w *Window) SetRendering(f func(w *Window)) *Window {

	w.rendering = f
	return w
}

type ChildWindow struct {
	Widget
	size   imgui.Vec2
	border bool
	flags  imgui.WindowFlags
}

func NewChildWindow() *ChildWindow {

	w := new(ChildWindow)
	w.Init(w)
	return w
}

func (w *ChildWindow) Init(control IWidget) {

	w.Widget.Init(control)
	w.SetRender(func() {
		if imgui.BeginChildID(imgui.ID(w.ID()), w.size, w.border, w.flags) {
			w.RenderChildren()
		}
		imgui.EndChild()
	})
}

func (w *ChildWindow) AddChildren(c ...IWidget) *ChildWindow {

	w.Widget.AddChildren(c...)
	return w
}

func (w *ChildWindow) Size() imgui.Vec2 {

	return w.size
}

func (w *ChildWindow) SetSize(size imgui.Vec2) {

	w.size = size
}

func (w *ChildWindow) Border() bool {

	return w.border
}

func (w *ChildWindow) SetBorder(border bool) *ChildWindow {

	w.border = border
	return w
}

func (w *ChildWindow) Flags() imgui.WindowFlags {

	return w.flags
}

func (w *ChildWindow) SetFlags(flags imgui.WindowFlags) *ChildWindow {

	w.flags = flags
	return w
}

type ImGuiDemoWindow struct {
	Widget
	open bool
}

func NewImGuiDemoWindow() *ImGuiDemoWindow {

	w := new(ImGuiDemoWindow)
	w.Init(w)
	w.open = true
	w.SetRender(func() {
		imgui.ShowDemoWindow(&w.open)
		if w.open == false {
			w.visible = false
			w.open = true
		}
	})
	return w
}

type ImGuiMetricsWindow struct {
	Widget
	open bool
}

func NewImGuiMetricsWindow() *ImGuiMetricsWindow {

	w := new(ImGuiMetricsWindow)
	w.Init(w)
	w.open = true
	w.SetRender(func() {
		imgui.ShowMetricsWindow(&w.open)
		if w.open == false {
			w.visible = false
			w.open = true
		}
	})
	return w
}

type ImPlotDemoWindow struct {
	Widget
	open bool
}

func NewImPlotDemoWindow() *ImPlotDemoWindow {

	w := new(ImPlotDemoWindow)
	w.Init(w)
	w.open = true
	w.SetRender(func() {
		imgui.PlotShowDemoWindow(&w.open)
		if w.open == false {
			w.visible = false
			w.open = true
		}
	})
	return w
}
