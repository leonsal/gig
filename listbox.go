package gig

import "github.com/leonsal/gig/imgui"

type ListBox struct {
	Widget
	label    imgui.CString
	items    []*Selectable
	size     imgui.Vec2
	multi    bool
	onchange func(*ListBox)
}

func NewListBox(label string) *ListBox {

	l := new(ListBox)
	l.Init(l)
	l.label.SetString(label)
	l.SetRender(func() {
		if imgui.BeginListBoxCS(l.label, l.size) {
			for i := 0; i < len(l.items); i++ {
				l.items[i].Render()
			}
			imgui.EndListBox()
		}
	})
	return l
}

func (l *ListBox) Save(pl **ListBox) *ListBox {

	*pl = l
	return l
}

func (l *ListBox) SetLabel(label string) *ListBox {

	l.label.SetString(label)
	return l
}

func (l *ListBox) AddItems(item ...*Selectable) *ListBox {

	for _, it := range item {
		it.SetOnClick(l.itemSelected)
		l.items = append(l.items, it)
	}
	return l
}

func (l *ListBox) ClearItems() *ListBox {

	l.items = l.items[:0]
	return l
}

func (l *ListBox) AddStrings(item ...string) *ListBox {

	for _, it := range item {
		sel := NewSelectable(it)
		sel.SetOnClick(l.itemSelected)
		l.items = append(l.items, sel)
	}
	return l
}

func (l *ListBox) Selected() []*Selectable {

	sel := make([]*Selectable, 0)
	for _, it := range l.items {
		if it.Selected() {
			sel = append(sel, it)
		}
	}
	return sel
}

func (l *ListBox) SelectedIndex() []int {

	sel := make([]int, 0)
	for idx, it := range l.items {
		if it.Selected() {
			sel = append(sel, idx)
		}
	}
	return sel
}

func (l *ListBox) UnselectAll() *ListBox {

	for _, it := range l.items {
		it.SetSelected(false)
	}
	return l
}

func (l *ListBox) SetSize(size imgui.Vec2) *ListBox {

	l.size = size
	return l
}

func (l *ListBox) SetMulti(multi bool) *ListBox {

	l.multi = multi
	return l
}

func (l *ListBox) SetOnChange(f func(*ListBox)) *ListBox {

	l.onchange = f
	return l
}

func (l *ListBox) itemSelected(it *Selectable) {

	if !l.multi {
		it.SetSelected(true)
		for i := 0; i < len(l.items); i++ {
			if l.items[i] != it {
				l.items[i].SetSelected(false)
			}
		}
	} else {
		it.SetSelected(!it.Selected())
	}
	if l.onchange != nil {
		l.onchange(l)
	}
}
