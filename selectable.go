package gig

import (
	"github.com/leonsal/gig/imgui"
)

type Selectable struct {
	Widget
	label    imgui.CString
	selected bool
	flags    imgui.SelectableFlags
	size     imgui.Vec2
	onclick  func(*Selectable)
	on2click func(*Selectable)
}

func NewSelectable(label string) *Selectable {

	s := new(Selectable)
	s.Init(s)
	s.label.SetString(label)
	s.SetRender(func() {
		if imgui.SelectableCS(s.label, s.selected, s.flags, s.size) {
			if s.on2click != nil {
				if imgui.IsMouseDoubleClicked(imgui.MouseButton_Left) {
					s.on2click(s)
					return
				}
			}
			if s.onclick != nil {
				s.onclick(s)
			}
		}
	})
	return s
}

func (s *Selectable) Label() string {

	return s.label.String()
}

func (s *Selectable) SetLabel(label string) *Selectable {

	s.label.SetString(label)
	return s
}

func (s *Selectable) Selected() bool {

	return s.selected
}

func (s *Selectable) SetSelected(selected bool) *Selectable {

	s.selected = selected
	return s
}

func (s *Selectable) SetFlags(flags imgui.SelectableFlags) *Selectable {

	s.flags = flags
	return s
}

func (s *Selectable) SetSize(size imgui.Vec2) *Selectable {

	s.size = size
	return s
}

func (s *Selectable) SetOnClick(f func(*Selectable)) *Selectable {

	s.onclick = f
	return s
}

func (s *Selectable) SetOnDoubleclick(f func(*Selectable)) *Selectable {

	s.on2click = f
	return s
}
