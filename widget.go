package gig

import (
	"github.com/leonsal/gig/imgui"
)

// IWidget is the interface for all Widgets
type IWidget interface {
	Super() IWidget
	GetWidget() *Widget // GetWidget returns pointer to the base embedded Widget
}

type Widget struct {
	super      IWidget
	id         int
	visible    bool
	preRender  func(IWidget)
	postRender func(IWidget)
	render     func()
	preList    []*Widget
	postList   []*Widget
	children   []*Widget
	userData   interface{}
}

// Next widget ID
var widgetNextID int = 1

// Maps widget reference string to IWidget
var widgetMapRef = map[string]IWidget{}

// Ref returns the widget associated with the specified string reference.
// Returns nil if not found.
func Ref(ref string) IWidget {

	return widgetMapRef[ref]
}

// Init initializes the widget
func (w *Widget) Init(control IWidget) {

	w.super = control
	w.id = widgetNextID
	widgetNextID++
	w.visible = true
	w.children = make([]*Widget, 0)
}

// SetPreRender sets a user callback which will be called right before the Widget is rendered.
// Pass nil to remove the callback.
func (w *Widget) SetPreRender(f func(IWidget)) *Widget {

	w.preRender = f
	return w
}

// SetPostRender sets a user callback which will be called right aftere the Widget is rendered.
// Pass nil to remove the callback.
func (w *Widget) SetPostRender(f func(IWidget)) *Widget {

	w.postRender = f
	return w
}

// SetRender set the callback render function
func (w *Widget) SetRender(f func()) *Widget {

	w.render = f
	return w
}

// Super returns the IWidget which embeds this Widget
func (w *Widget) Super() IWidget {

	return w.super
}

// GetWidget returns the pointer to this Widget
func (w *Widget) GetWidget() *Widget {

	return w
}

func (w *Widget) ID() int {

	return w.id
}

func (w *Widget) SetID(id int) *Widget {

	for _, it := range w.preList {
		pi, ok := it.super.(*PushID)
		if ok {
			pi.SetID(id)
			return w
		}
	}
	w.appendPreList(NewPushID(id))
	return w
}

// SaveRef saves a reference to the widget associated with the specified string.
// The widget can get be obtained later by GetRef(ref)
func (w *Widget) SaveRef(ref string) *Widget {

	widgetMapRef[ref] = w.super
	return w
}

// Visible returns the visibility state of this widget
func (w *Widget) Visible() bool {

	return w.visible
}

// SetVisible sets the visibility state of this widget
func (w *Widget) SetVisible(visible bool) *Widget {

	w.visible = visible
	return w
}

func (w *Widget) SetFont(font *imgui.Font) *Widget {

	for _, it := range w.preList {
		f, ok := it.super.(*PushFont)
		if ok {
			f.font = font
			return w
		}
	}
	w.appendPreList(NewPushFont(font))
	w.appendPostList(NewPopFont())
	return w
}

func (w *Widget) Font() *imgui.Font {

	for _, it := range w.preList {
		f, ok := it.super.(*PushFont)
		if ok {
			return f.font
		}
	}
	return nil
}

func (w *Widget) SetItemWidth(width float32) *Widget {

	for _, it := range w.preList {
		sw, ok := it.super.(*SetNextItemWidth)
		if ok {
			sw.SetWidth(width)
			return w
		}
	}
	w.appendPreList(NewSetNextItemWidth(width))
	return w
}

func (w *Widget) SetDisabled(disabled bool) *Widget {

	for _, it := range w.preList {
		bd, ok := it.super.(*BeginDisabled)
		if ok {
			bd.SetDisabled(disabled)
			return w
		}
	}
	w.appendPreList(NewBeginDisabled(disabled))
	w.appendPostList(NewEndDisabled())
	return w
}

func (w *Widget) SetStyleColor(idx imgui.Col, color imgui.Vec4) *Widget {

	for _, it := range w.preList {
		sc, ok := it.super.(*PushStyleColor)
		if ok && sc.idx == idx {
			sc.color = color
			return w
		}
	}
	w.appendPreList(NewPushStyleColor(idx, color))
	w.appendPostList(NewPopStyleColor(1))
	return w
}

func (w *Widget) SetStyleVar(idx imgui.StyleVar, value float32) *Widget {

	for _, it := range w.preList {
		sv, ok := it.super.(*PushStyleVar)
		if ok && sv.idx == idx {
			sv.value = value
			return w
		}
	}
	w.appendPreList(NewPushStyleVar(idx, value))
	w.appendPostList(NewPopStyleVar(1))
	return w
}

func (w *Widget) SetStyleVarVec2(idx imgui.StyleVar, value imgui.Vec2) *Widget {

	for _, it := range w.preList {
		sv, ok := it.super.(*PushStyleVarVec2)
		if ok && sv.idx == idx {
			sv.value = value
			return w
		}
	}
	w.appendPreList(NewPushStyleVarVec2(idx, value))
	w.appendPostList(NewPopStyleVar(1))
	return w
}

func (w *Widget) SetOnHovered(flags imgui.HoveredFlags, f func(w *Widget)) *Widget {

	for _, it := range w.postList {
		oa, ok := it.super.(*IsItemHovered)
		if ok {
			if f == nil {
				oa.fn = nil
				return w
			}
			oa.flags = flags
			prev := oa.fn
			oa.fn = func(w *Widget) {
				if prev != nil {
					prev(w)
				}
				f(w)
			}
			return w
		}
	}
	w.appendPostList(NewIsItemHovered(flags, f))
	return w
}

func (w *Widget) SetOnContext(flags imgui.HoveredFlags, f func(w *Widget)) *Widget {

	w.SetOnHovered(flags, func(w *Widget) {
		if imgui.IsMouseReleased(imgui.MouseButton_Right) {
			f(w)
		}
	})
	return w
}

// SetOnActivated sets a callback function which will be called when the widget is activated.
// Pass a nil function to remove the callback.
func (w *Widget) SetOnActivated(f func(w *Widget)) *Widget {

	for _, it := range w.postList {
		of, ok := it.super.(*IsItemActivated)
		if ok {
			of.fn = f
			return w
		}
	}
	w.appendPostList(NewIsItemActivated(f))
	return w
}

// SetOnFocused sets a callback function which will be called when a widget gains focus.
// Pass a nil function to remove the callback.
func (w *Widget) SetOnFocused(f func(w *Widget)) *Widget {

	for _, it := range w.postList {
		of, ok := it.super.(*IsItemFocused)
		if ok {
			of.fn = f
			return w
		}
	}
	w.appendPostList(NewIsItemFocused(f))
	return w
}

// SetTooltip sets the tooltip of the widget
func (w *Widget) SetTooltip(text string) *Widget {

	w.SetTooltipFlags(text, imgui.HoveredFlags_None)
	return w
}

// SetTooltipFlags set the tooltip and flags for the widget
func (w *Widget) SetTooltipFlags(text string, flags imgui.HoveredFlags) *Widget {

	if text == "" {
		w.SetOnHovered(0, nil)
		return w
	}

	var ctext imgui.CString
	ctext.SetString(text)
	w.SetOnHovered(flags, func(w *Widget) {
		imgui.SetTooltipCS(ctext)
	})
	return w
}

func (w *Widget) AddChildren(c ...IWidget) *Widget {

	for i := 0; i < len(c); i++ {
		w.children = append(w.children, c[i].GetWidget())
	}
	return w
}

func (w *Widget) ClearChildren() {

	w.children = w.children[:0]
}

func (w *Widget) Children() []*Widget {

	return w.children
}

func (w *Widget) Call(f func(IWidget)) *Widget {

	f(w.super)
	return w
}

func (w *Widget) Userdata() interface{} {

	return w.userData
}

func (w *Widget) SetUserdata(ud interface{}) *Widget {

	w.userData = ud
	return w
}

func (w *Widget) Render() {

	if !w.visible {
		return
	}

	imgui.PushID(w.id)

	// Render pre widgets
	for _, pre := range w.preList {
		pre.Render()
	}

	// Render callbacks
	if w.preRender != nil {
		w.preRender(w.super)
	}
	if w.render != nil {
		w.render()
	}
	if w.postRender != nil {
		w.postRender(w.super)
	}

	// Render post widgets
	for i := len(w.postList) - 1; i >= 0; i-- {
		w.postList[i].Render()
	}

	imgui.PopID()
}

func (w *Widget) RenderChildren() {

	for i := 0; i < len(w.children); i++ {
		w.children[i].Render()
	}
}

func (w *Widget) delPreListItem(idx int) {

	w.preList = append(w.preList[:idx], w.preList[idx+1:]...)
}

//func (w *Widget) delPreListType(v interface{}) bool {
//
//	for idx, item := range w.preList {
//		if reflect.TypeOf(v) == reflect.TypeOf(item) {
//			w.delPreListItem(idx)
//			return true
//		}
//	}
//	return false
//}

func (w *Widget) appendPreList(v IWidget) {

	w.preList = append(w.preList, v.GetWidget())
}

func (w *Widget) appendPostList(v IWidget) {

	w.postList = append(w.postList, v.GetWidget())
}

type PushID struct {
	Widget
	id int
}

func NewPushID(id int) *PushID {

	w := new(PushID)
	w.Init(w)
	w.id = id
	w.SetRender(func() {
		imgui.PushID(w.id)
	})
	return w
}

func (w *PushID) SetID(id int) *PushID {

	w.id = id
	return w
}

type PopID struct {
	Widget
}

func NewPopID() *PopID {

	w := new(PopID)
	w.Init(w)
	w.SetRender(func() {
		imgui.PopID()
	})
	return w
}

type PushFont struct {
	Widget
	font *imgui.Font
}

func NewPushFont(font *imgui.Font) *PushFont {

	p := new(PushFont)
	p.Init(p)
	p.font = font
	p.SetRender(func() {
		imgui.PushFont(p.font)
	})
	return p
}

type PopFont struct {
	Widget
}

func NewPopFont() *PopFont {

	p := new(PopFont)
	p.Init(p)
	p.SetRender(func() {
		imgui.PopFont()
	})
	return p
}

type SetNextItemWidth struct {
	Widget
	width float32
}

func NewSetNextItemWidth(width float32) *SetNextItemWidth {

	w := new(SetNextItemWidth)
	w.Init(w)
	w.width = width
	w.SetRender(func() {
		imgui.SetNextItemWidth(w.width)
	})
	return w
}

func (w *SetNextItemWidth) SetWidth(width float32) *SetNextItemWidth {

	w.width = width
	return w
}

type BeginDisabled struct {
	Widget
	disabled bool
}

func NewBeginDisabled(disabled bool) *BeginDisabled {

	w := new(BeginDisabled)
	w.Init(w)
	w.disabled = disabled
	w.SetRender(func() {
		imgui.BeginDisabled(w.disabled)
	})
	return w
}

func (w *BeginDisabled) SetDisabled(disabled bool) *BeginDisabled {

	w.disabled = disabled
	return w
}

type EndDisabled struct {
	Widget
}

func NewEndDisabled() *EndDisabled {

	w := new(EndDisabled)
	w.Init(w)
	w.SetRender(func() {
		imgui.EndDisabled()
	})
	return w
}

type PushStyleColor struct {
	Widget
	idx   imgui.Col
	color imgui.Vec4
}

func NewPushStyleColor(idx imgui.Col, color imgui.Vec4) *PushStyleColor {

	w := new(PushStyleColor)
	w.Init(w)
	w.idx = idx
	w.color = color
	w.SetRender(func() {
		imgui.PushStyleColor(w.idx, w.color)
	})
	return w
}

func (w *PushStyleColor) Set(idx imgui.Col, color imgui.Vec4) *PushStyleColor {

	w.idx = idx
	w.color = color
	return w
}

type PopStyleColor struct {
	Widget
	count int
}

func NewPopStyleColor(count int) *PopStyleColor {

	w := new(PopStyleColor)
	w.Init(w)
	w.count = count
	w.SetRender(func() {
		imgui.PopStyleColor(w.count)
	})
	return w
}

type PushStyleVar struct {
	Widget
	idx   imgui.StyleVar
	value float32
}

func NewPushStyleVar(idx imgui.StyleVar, value float32) *PushStyleVar {

	w := new(PushStyleVar)
	w.Init(w)
	w.idx = idx
	w.value = value
	w.SetRender(func() {
		imgui.PushStyleVar(w.idx, w.value)
	})
	return w
}

func (w *PushStyleVar) Set(idx imgui.StyleVar, value float32) *PushStyleVar {

	w.idx = idx
	w.value = value
	return w
}

type PushStyleVarVec2 struct {
	Widget
	idx   imgui.StyleVar
	value imgui.Vec2
}

func NewPushStyleVarVec2(idx imgui.StyleVar, value imgui.Vec2) *PushStyleVarVec2 {

	w := new(PushStyleVarVec2)
	w.Init(w)
	w.idx = idx
	w.value = value
	w.SetRender(func() {
		imgui.PushStyleVarVec2(w.idx, w.value)
	})
	return w
}

func (w *PushStyleVarVec2) Set(idx imgui.StyleVar, value imgui.Vec2) *PushStyleVarVec2 {

	w.idx = idx
	w.value = value
	return w
}

type PopStyleVar struct {
	Widget
	count int
}

func NewPopStyleVar(count int) *PopStyleVar {

	w := new(PopStyleVar)
	w.Init(w)
	w.count = count
	w.SetRender(func() {
		imgui.PopStyleVar(w.count)
	})
	return w
}

type IsItemHovered struct {
	Widget
	flags imgui.HoveredFlags
	fn    func(*Widget)
}

func NewIsItemHovered(flags imgui.HoveredFlags, fn func(w *Widget)) *IsItemHovered {

	w := new(IsItemHovered)
	w.Init(w)
	w.flags = flags
	w.fn = fn
	w.SetRender(func() {
		if imgui.IsItemHovered(w.flags) {
			if w.fn != nil {
				w.fn(w.GetWidget())
			}
		}
	})
	return w
}

type IsItemActivated struct {
	Widget
	fn func(*Widget)
}

func NewIsItemActivated(fn func(w *Widget)) *IsItemActivated {

	w := new(IsItemActivated)
	w.Init(w)
	w.fn = fn
	w.SetRender(func() {
		if imgui.IsItemActivated() {
			if w.fn != nil {
				w.fn(w.GetWidget())
			}
		}
	})
	return w
}

type IsItemFocused struct {
	Widget
	fn func(*Widget)
}

func NewIsItemFocused(fn func(w *Widget)) *IsItemFocused {

	w := new(IsItemFocused)
	w.Init(w)
	w.fn = fn
	w.SetRender(func() {
		if imgui.IsItemFocused() {
			if w.fn != nil {
				w.fn(w.GetWidget())
			}
		}
	})
	return w
}
