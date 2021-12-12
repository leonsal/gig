package gig

import (
	"fmt"

	"github.com/leonsal/gig/imgui"
)

type Text struct {
	Widget
	text imgui.CString
}

func NewText(format string, args ...interface{}) *Text {

	t := new(Text)
	t.Init(t)
	t.text.SetString(fmt.Sprintf(format, args...))
	t.SetRender(func() {
		imgui.TextCS(t.text)
	})
	return t
}

func (t *Text) Save(pt **Text) *Text {

	*pt = t
	return t
}

func (t *Text) Text() string {

	return t.text.String()
}

func (t *Text) SetText(format string, args ...interface{}) *Text {

	t.text.SetString(fmt.Sprintf(format, args...))
	return t
}

type TextColored struct {
	Widget
	color imgui.Vec4
	text  imgui.CString
}

func NewTextColored(color imgui.Vec4, format string, args ...interface{}) *TextColored {

	tc := new(TextColored)
	tc.Init(tc)
	tc.color = color
	tc.text.SetString(fmt.Sprintf(format, args...))
	tc.SetRender(func() {
		imgui.TextColoredCS(tc.color, tc.text)
	})
	return tc
}

func (t *TextColored) Save(pt **TextColored) *TextColored {

	*pt = t
	return t
}

func (tc *TextColored) SetText(format string, args ...interface{}) *TextColored {

	tc.text.SetString(fmt.Sprintf(format, args...))
	return tc
}

func (tc *TextColored) SetColor(color imgui.Vec4) *TextColored {

	tc.color = color
	return tc
}

type TextDisabled struct {
	Widget
	text imgui.CString
}

func NewTextDisabled(format string, args ...interface{}) *TextDisabled {

	t := new(TextDisabled)
	t.Init(t)
	t.text.SetString(fmt.Sprintf(format, args...))
	t.SetRender(func() {
		imgui.TextDisabledCS(t.text)
	})
	return t
}

func (t *TextDisabled) Text() string {

	return t.text.String()
}

func (t *TextDisabled) SetText(format string, args ...interface{}) *TextDisabled {

	t.text.SetString(fmt.Sprintf(format, args...))
	return t
}

type TextWrapped struct {
	Widget
	text imgui.CString
}

func NewTextWrapped(format string, args ...interface{}) *TextWrapped {

	t := new(TextWrapped)
	t.Init(t)
	t.text.SetString(fmt.Sprintf(format, args...))
	t.SetRender(func() {
		imgui.TextWrappedCS(t.text)
	})
	return t
}

func (t *TextWrapped) SetText(format string, args ...interface{}) *TextWrapped {

	t.text.SetString(fmt.Sprintf(format, args...))
	return t
}

type LabelText struct {
	Widget
	label imgui.CString
	text  imgui.CString
}

func NewLabelText(label string, format string, args ...interface{}) *LabelText {

	t := new(LabelText)
	t.Init(t)
	t.label.SetString(label)
	t.text.SetString(fmt.Sprintf(format, args...))
	t.SetRender(func() {
		imgui.LabelTextCS(t.label, t.text)
	})
	return t
}

func (t *LabelText) SetLabel(label string) *LabelText {

	t.label.SetString(label)
	return t
}

func (t *LabelText) SetText(format string, args ...interface{}) *LabelText {

	t.text.SetString(fmt.Sprintf(format, args...))
	return t
}

type BulletText struct {
	Widget
	text imgui.CString
}

func NewBulletText(format string, args ...interface{}) *BulletText {

	bt := new(BulletText)
	bt.Init(bt)
	bt.text.SetString(fmt.Sprintf(format, args...))
	bt.SetRender(func() {
		imgui.BulletTextCS(bt.text)
	})
	return bt
}

func (bt *BulletText) SetText(format string, args ...interface{}) *BulletText {

	bt.text.SetString(fmt.Sprintf(format, args...))
	return bt
}
