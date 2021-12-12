package gig

import (
	"fmt"
	"strconv"

	"github.com/leonsal/gig/imgui"
)

type TreeNode struct {
	Widget
	strId imgui.CString
	flags imgui.TreeNodeFlags
	text  imgui.CString
}

func NewTreeNode(format string, args ...interface{}) *TreeNode {

	t := new(TreeNode)
	t.Init(t, format, args...)
	return t
}

func (t *TreeNode) Init(wi IWidget, format string, args ...interface{}) {

	t.Widget.Init(wi)
	t.strId.SetString(strconv.Itoa(t.ID()))
	t.text.SetString(fmt.Sprintf(format, args...))
	t.SetRender(func() {
		if imgui.TreeNodeExStrIdCS(t.strId, t.flags, t.text) {
			t.RenderChildren()
			imgui.TreePop()
		}
	})
}

func (t *TreeNode) Flags() imgui.TreeNodeFlags {

	return t.flags
}

func (t *TreeNode) SetFlags(flags imgui.TreeNodeFlags) *TreeNode {

	t.flags = flags
	return t
}

func (t *TreeNode) SetText(format string, args ...interface{}) *TreeNode {

	t.text.SetString(fmt.Sprintf(format, args...))
	return t
}

type CollapsingHeader struct {
	Widget
	label imgui.CString
	flags imgui.TreeNodeFlags
	open  bool
	popen *bool
}

func NewCollapsingHeader(label string) *CollapsingHeader {

	h := new(CollapsingHeader)
	h.Init(h)
	h.label.SetString(label)
	h.SetRender(func() {
		if imgui.CollapsingHeaderPtrCS(h.label, h.popen, h.flags) {
			h.RenderChildren()
		}
		if h.popen != nil && h.open == false {
			h.visible = false
			h.open = true
		}
	})
	return h
}

func (h *CollapsingHeader) Save(ph **CollapsingHeader) *CollapsingHeader {

	*ph = h
	return h
}

func (h *CollapsingHeader) SetLabel(label string) *CollapsingHeader {

	h.label.SetString(label)
	return h
}

func (h *CollapsingHeader) Flags() imgui.TreeNodeFlags {

	return h.flags
}

func (h *CollapsingHeader) SetFlags(flags imgui.TreeNodeFlags) *CollapsingHeader {

	h.flags = flags
	return h
}

func (h *CollapsingHeader) SetCanClose(canclose bool) *CollapsingHeader {

	if canclose {
		h.open = true
		h.popen = &h.open
	} else {
		h.popen = nil
	}
	return h
}
