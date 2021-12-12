package gig

import (
	"strconv"

	"github.com/leonsal/gig/imgui"
)

type Button struct {
	Widget
	label   imgui.CString
	size    imgui.Vec2
	repeat  bool
	onclick func(*Button)
}

func NewButton(label string) *Button {

	b := new(Button)
	b.Init(b)
	b.label.SetString(label)
	b.SetRender(func() {
		if b.repeat {
			imgui.PushButtonRepeat(true)
		}
		if imgui.ButtonCS(b.label, b.size) {
			if b.onclick != nil {
				b.onclick(b)
			}
		}
		if b.repeat {
			imgui.PopButtonRepeat()
		}
	})
	return b
}

func (b *Button) Save(pb **Button) *Button {

	*pb = b
	return b
}

func (b *Button) Label() string {

	return b.label.String()
}

func (b *Button) SetLabel(label string) *Button {

	b.label.SetString(label)
	return b
}

func (b *Button) Size() imgui.Vec2 {

	return b.size
}

func (b *Button) SetSize(size imgui.Vec2) *Button {

	b.size = size
	return b
}

func (b *Button) SetRepeat(repeat bool) *Button {

	b.repeat = repeat
	return b
}

func (b *Button) SetOnClick(f func(*Button)) *Button {

	b.onclick = f
	return b
}

type SmallButton struct {
	Widget
	label   imgui.CString
	onclick func(*SmallButton)
}

func NewSmallButton(label string) *SmallButton {

	b := new(SmallButton)
	b.Init(b)
	b.label.SetString(label)
	b.SetRender(func() {
		if imgui.SmallButtonCS(b.label) {
			if b.onclick != nil {
				b.onclick(b)
			}
		}
	})
	return b
}

func (b *SmallButton) SetLabel(label string) *SmallButton {

	b.label.SetString(label)
	return b
}

func (b *SmallButton) SetOnClick(f func(*SmallButton)) *SmallButton {

	b.onclick = f
	return b
}

type InvisibleButton struct {
	Widget
	strId   imgui.CString
	size    imgui.Vec2
	flags   imgui.ButtonFlags
	onclick func(*InvisibleButton)
}

func NewInvisibleButton(size imgui.Vec2) *InvisibleButton {

	b := new(InvisibleButton)
	b.Init(b)
	b.strId.SetString(strconv.Itoa(b.ID()))
	b.SetRender(func() {
		if imgui.InvisibleButtonCS(b.strId, b.size, b.flags) {
			if b.onclick != nil {
				b.onclick(b)
			}
		}
	})
	return b
}

func (b *InvisibleButton) SetSize(size imgui.Vec2) *InvisibleButton {

	b.size = size
	return b
}

func (b *InvisibleButton) SetFlags(flags imgui.ButtonFlags) *InvisibleButton {

	b.flags = flags
	return b
}

func (b *InvisibleButton) SetOnClick(f func(*InvisibleButton)) *InvisibleButton {

	b.onclick = f
	return b
}

type ArrowButton struct {
	Widget
	dir     imgui.Dir
	str_id  imgui.CString
	repeat  bool
	onclick func(*ArrowButton)
}

func NewArrowButton(dir imgui.Dir) *ArrowButton {

	b := new(ArrowButton)
	b.Init(b)
	b.dir = dir
	b.str_id.SetString(strconv.Itoa(b.ID()))
	b.SetRender(func() {
		if b.repeat {
			imgui.PushButtonRepeat(true)
		}
		if imgui.ArrowButtonCS(b.str_id, b.dir) {
			if b.onclick != nil {
				b.onclick(b)
			}
		}
		if b.repeat {
			imgui.PopButtonRepeat()
		}
	})
	return b
}

func (b *ArrowButton) SetOnClick(f func(*ArrowButton)) *ArrowButton {

	b.onclick = f
	return b
}

type Checkbox struct {
	Widget
	label   imgui.CString
	value   bool
	onclick func(*Checkbox)
}

func NewCheckbox(label string) *Checkbox {

	c := new(Checkbox)
	c.Init(c)
	c.label.SetString(label)
	c.SetRender(func() {
		if imgui.CheckboxCS(c.label, &c.value) {
			if c.onclick != nil {
				c.onclick(c)
			}
		}
	})
	return c
}

func (c *Checkbox) Save(pc **Checkbox) *Checkbox {

	*pc = c
	return c
}

func (c *Checkbox) SetLabel(label string) *Checkbox {

	c.label.SetString(label)
	return c
}

func (c *Checkbox) Value() bool {

	return c.value
}

func (c *Checkbox) SetValue(v bool) *Checkbox {

	c.value = v
	return c
}

func (c *Checkbox) SetOnClick(f func(*Checkbox)) *Checkbox {

	c.onclick = f
	return c
}

type RadioButton struct {
	Widget
	label   imgui.CString
	active  bool
	group   *RadioGroup
	onclick func(*RadioButton)
}

func NewRadioButton(label string) *RadioButton {

	r := new(RadioButton)
	r.Init(r)
	r.label.SetString(label)
	r.SetRender(func() {
		if imgui.RadioButtonCS(r.label, r.active) {
			if r.group != nil {
				r.group.RadioClicked(r)
			}
			if r.onclick != nil {
				r.onclick(r)
			}
		}
	})
	return r
}

func (r *RadioButton) Label() string {

	return r.label.String()
}

func (r *RadioButton) SetLabel(label string) *RadioButton {

	r.label.SetString(label)
	return r
}

func (r *RadioButton) Value() bool {

	return r.active
}

func (r *RadioButton) SetValue(active bool) *RadioButton {

	r.active = active
	return r
}

func (r *RadioButton) SetGroup(g *RadioGroup) *RadioButton {

	r.group = g
	return r
}

func (r *RadioButton) SetOnClick(f func(*RadioButton)) *RadioButton {

	r.onclick = f
	return r
}

type RadioGroup struct {
	Widget
	radios   []*RadioButton
	active   int
	onchange func(*RadioGroup)
}

func NewRadioGroup(c ...IWidget) *RadioGroup {

	g := new(RadioGroup)
	g.Init(g)
	g.AddChildren(c...)
	g.radios = make([]*RadioButton, 0)
	g.active = -1
	for _, w := range c {
		c := w.Super()
		rb, ok := c.(*RadioButton)
		if ok {
			rb.SetGroup(g)
			rb.SetValue(false)
			g.radios = append(g.radios, rb)
		}
	}
	g.SetRender(func() {
		g.RenderChildren()
	})
	return g
}

func (g *RadioGroup) Len() int {

	return len(g.radios)
}

func (g *RadioGroup) Active() (int, *RadioButton) {

	for idx, rb := range g.radios {
		if rb.Value() {
			return idx, rb
		}
	}
	return -1, nil
}

func (g *RadioGroup) SetActive(idx int) *RadioGroup {

	if idx < 0 || idx >= len(g.radios) {
		return g
	}
	for i, rb := range g.radios {
		if i == idx {
			rb.SetValue(true)
		} else {
			rb.SetValue(false)
		}
	}
	g.active = idx
	return g
}

func (g *RadioGroup) RadioClicked(rb *RadioButton) {

	found := -1
	for idx, r := range g.radios {
		if r == rb {
			r.SetValue(true)
			found = idx
		} else {
			r.SetValue(false)
		}
	}
	if found != g.active {
		g.active = found
		if g.onchange != nil {
			g.onchange(g)
		}
	}
}

func (g *RadioGroup) SetOnChange(f func(*RadioGroup)) *RadioGroup {

	g.onchange = f
	return g
}

type ProgressBar struct {
	Widget
	fraction float32
	size     imgui.Vec2
	overlay  imgui.CString
}

func NewProgressBar(fraction float32) *ProgressBar {

	p := new(ProgressBar)
	p.Init(p)
	p.fraction = fraction
	p.size = imgui.Vec2{-1, 0}
	p.SetRender(func() {
		imgui.ProgressBarCS(p.fraction, p.size, p.overlay)
	})
	return p
}

func (p *ProgressBar) SetFraction(fraction float32) *ProgressBar {

	p.fraction = fraction
	return p
}

func (p *ProgressBar) SetSize(size imgui.Vec2) *ProgressBar {

	p.size = size
	return p
}

func (p *ProgressBar) SetOverlay(overlay string) *ProgressBar {

	p.overlay.SetString(overlay)
	return p
}

type Bullet struct {
	Widget
}

func NewBullet() *Bullet {

	b := new(Bullet)
	b.Init(b)
	b.SetRender(func() {
		imgui.Bullet()
	})
	return b
}
