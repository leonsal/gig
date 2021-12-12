package gig

import (
	"github.com/leonsal/gig/imgui"
)

type SameLine struct {
	Widget
	offset  float32
	spacing float32
}

func NewSameLine() *SameLine {

	s := new(SameLine)
	s.Init(s)
	s.spacing = -1.0
	s.SetRender(func() {
		imgui.SameLine(s.offset, s.spacing)
	})
	return s
}

func (s *SameLine) SetOffset(offset float32) *SameLine {

	s.offset = offset
	return s
}

func (s *SameLine) SetSpacing(spacing float32) *SameLine {

	s.spacing = spacing
	return s
}

type NewLine struct {
	Widget
}

func NewNewLine() *NewLine {

	n := new(NewLine)
	n.Init(n)
	n.SetRender(func() {
		imgui.NewLine()
	})
	return n
}

type Separator struct {
	Widget
}

func NewSeparator() *Separator {

	s := new(Separator)
	s.Init(s)
	s.SetRender(func() {
		imgui.Separator()
	})
	return s
}

type Spacing struct {
	Widget
}

func NewSpacing() *Spacing {

	s := new(Spacing)
	s.Init(s)
	s.SetRender(func() {
		imgui.Spacing()
	})
	return s
}

type Dummy struct {
	Widget
	size imgui.Vec2
}

func NewDummy(size imgui.Vec2) *Dummy {

	d := new(Dummy)
	d.Init(d)
	d.size = size
	d.SetRender(func() {
		imgui.Dummy(d.size)
	})
	return d
}

type Indent struct {
	Widget
	indent float32
}

func NewIndent(indent float32) *Indent {

	i := new(Indent)
	i.Init(i)
	i.indent = indent
	i.SetRender(func() {
		imgui.Indent(i.indent)
	})
	return i
}

type Unindent struct {
	Widget
	unindent float32
}

func NewUnindent(unindent float32) *Unindent {

	u := new(Unindent)
	u.Init(u)
	u.unindent = unindent
	u.SetRender(func() {
		imgui.Unindent(u.unindent)
	})
	return u
}
