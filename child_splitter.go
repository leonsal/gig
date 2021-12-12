package gig

import (
	"strconv"

	"github.com/leonsal/gig/imgui"
)

type IChildSplitter interface {
	SetSize(imgui.Vec2)
	Render()
}

type ChildHSplitter struct {
	Widget
	left        IChildSplitter
	right       IChildSplitter
	size        imgui.Vec2
	ratio       float32
	margin      float32
	drag        bool
	sepID       imgui.CString
	sepSize     imgui.Vec2
	sepColor    imgui.Vec4
	sepColorOff imgui.Vec4
	sepColorOn  imgui.Vec4
}

func NewChildHSplitter(left IChildSplitter, right IChildSplitter) *ChildHSplitter {

	s := new(ChildHSplitter)
	s.Init(s)
	s.left = left
	s.right = right
	s.size = imgui.Vec2{-1, -1}
	s.ratio = 0.5
	s.margin = 0.1
	s.sepID.SetString("##" + strconv.Itoa(s.ID()))
	s.sepSize = imgui.Vec2{4, -1}
	s.sepColorOff = imgui.ColorName("Gray")
	s.sepColorOn = imgui.ColorName("White")
	s.sepColor = s.sepColorOff
	s.SetRender(s.Render)
	return s
}

func (s *ChildHSplitter) SetSize(size imgui.Vec2) {

	s.size = size
}

func (s *ChildHSplitter) SetRatio(ratio float32) *ChildHSplitter {

	s.ratio = ClampFloat32(ratio, 0, 1)
	return s
}

func (s *ChildHSplitter) Ratio() float32 {

	return s.ratio
}

func (s *ChildHSplitter) SetMargin(margin float32) *ChildHSplitter {

	s.margin = ClampFloat32(margin, 0, 0.3)
	return s
}

func (s *ChildHSplitter) SetLeft(child IChildSplitter) *ChildHSplitter {

	s.left = child
	return s
}

func (s *ChildHSplitter) SetRight(child IChildSplitter) *ChildHSplitter {

	s.right = child
	return s
}

func (s *ChildHSplitter) Render() {

	if s.size.X == 0 {
		return
	}

	if s.left != nil && s.right == nil {
		s.left.SetSize(imgui.Vec2{-1, -1})
		s.left.Render()
		return
	}

	if s.left == nil && s.right != nil {
		s.right.SetSize(imgui.Vec2{-1, -1})
		s.right.Render()
		return
	}

	totalWidth := s.size.X
	totalHeight := s.size.Y
	avail := imgui.GetContentRegionAvail()
	if totalWidth == -1 {
		totalWidth = avail.X
	}
	if totalHeight == -1 {
		totalHeight = avail.Y
	}

	itemSpacing := imgui.GetStyle().GetItemSpacing().X
	sepWidth := s.sepSize.X
	if s.ratio == 0 || s.ratio == 1 {
		sepWidth += itemSpacing
	} else {
		sepWidth += 2 * itemSpacing
	}

	// Render the left child/splitter
	leftSize := imgui.Vec2{s.ratio * (totalWidth - sepWidth), totalHeight}
	pos := imgui.GetCursorPos()
	if leftSize.X > 0 {
		s.left.SetSize(leftSize)
		s.left.Render()
		pos.X = pos.X + leftSize.X + itemSpacing
		imgui.SetCursorPos(pos)
	}

	// Render the separator
	imgui.PushStyleColor(imgui.Col_Button, s.sepColor)
	imgui.PushStyleColor(imgui.Col_ButtonHovered, s.sepColor)
	imgui.PushStyleColor(imgui.Col_ButtonActive, s.sepColor)
	imgui.ButtonCS(s.sepID, s.sepSize)
	imgui.PopStyleColor(3)
	if imgui.IsItemHovered(imgui.HoveredFlags_None) || s.drag {
		imgui.SetMouseCursor(imgui.MouseCursor_ResizeEW)
		if imgui.IsMouseClicked(imgui.MouseButton_Left, false) {
			s.drag = true
		}
	}

	// Render the right child/splitter
	rightSize := imgui.Vec2{totalWidth - leftSize.X - sepWidth, totalHeight}
	if rightSize.X > 0 {
		pos.X = pos.X + s.sepSize.X + itemSpacing
		imgui.SetCursorPos(pos)
		s.right.SetSize(rightSize)
		s.right.Render()
	}

	// If dragging, get mouse delta
	delta := imgui.Vec2{}
	if s.drag {
		delta = imgui.GetMouseDragDelta(imgui.MouseButton_Left, -1.0)
		s.sepColor = s.sepColorOn
		// If mouse button released, terminates drag and
		// checks if separator is within margins.
		if !imgui.IsMouseDown(imgui.MouseButton_Left) {
			s.drag = false
			s.sepColor = s.sepColorOff
			if s.ratio < s.margin {
				s.ratio = 0
			} else if s.ratio > 1.0-s.margin {
				s.ratio = 1
			}
		}
	}

	// If splitter was dragged recalculates its relative position
	if delta.X != 0 {
		s.ratio = (leftSize.X + delta.X) / (totalWidth - sepWidth)
		if delta.X > 0 && s.ratio > 1.0-s.margin {
			s.ratio = 1.0
			s.drag = false
		}
		if delta.X < 0 && s.ratio < s.margin {
			s.ratio = 0
			s.drag = false
		}
		imgui.ResetMouseDragDelta(imgui.MouseButton_Left)
		if !s.drag {
			s.sepColor = s.sepColorOff
		}
	}
}

type ChildVSplitter struct {
	Widget
	top         IChildSplitter
	bottom      IChildSplitter
	size        imgui.Vec2
	ratio       float32
	margin      float32
	drag        bool
	sepID       imgui.CString
	sepSize     imgui.Vec2
	sepColor    imgui.Vec4
	sepColorOff imgui.Vec4
	sepColorOn  imgui.Vec4
}

func NewChildVSplitter(top IChildSplitter, bottom IChildSplitter) *ChildVSplitter {

	s := new(ChildVSplitter)
	s.Init(s)
	s.top = top
	s.bottom = bottom
	s.size = imgui.Vec2{-1, -1}
	s.ratio = 0.5
	s.margin = 0.1
	s.sepID.SetString("##" + strconv.Itoa(s.ID()))
	s.sepSize = imgui.Vec2{-1, 4}
	s.sepColorOff = imgui.ColorName("Gray")
	s.sepColorOn = imgui.ColorName("White")
	s.sepColor = s.sepColorOff
	s.SetRender(s.Render)
	return s
}

func (s *ChildVSplitter) AddChildren(c ...IWidget) *ChildVSplitter {

	s.Widget.AddChildren(c...)
	return s
}

func (s *ChildVSplitter) SetSize(size imgui.Vec2) {

	s.size = size
}

func (s *ChildVSplitter) SetRatio(ratio float32) *ChildVSplitter {

	s.ratio = ClampFloat32(ratio, 0, 1)
	return s
}

func (s *ChildVSplitter) Ratio() float32 {

	return s.ratio
}

func (s *ChildVSplitter) SetMargin(margin float32) *ChildVSplitter {

	s.margin = ClampFloat32(margin, 0, 0.3)
	return s
}

func (s *ChildVSplitter) SetTop(child IChildSplitter) *ChildVSplitter {

	s.top = child
	return s
}

func (s *ChildVSplitter) SetBottom(child IChildSplitter) *ChildVSplitter {

	s.bottom = child
	return s
}

func (s *ChildVSplitter) Render() {

	if s.size.Y == 0 {
		return
	}

	if s.top != nil && s.bottom == nil {
		s.top.SetSize(imgui.Vec2{-1, -1})
		s.top.Render()
		return
	}

	if s.top == nil && s.bottom != nil {
		s.bottom.SetSize(imgui.Vec2{-1, -1})
		s.bottom.Render()
		return
	}

	totalWidth := s.size.X
	totalHeight := s.size.Y
	avail := imgui.GetContentRegionAvail()
	if totalWidth == -1 {
		totalWidth = avail.X
	}
	if totalHeight == -1 {
		totalHeight = avail.Y
	}

	itemSpacing := imgui.GetStyle().GetItemSpacing().Y
	sepHeight := s.sepSize.Y
	if s.ratio == 0 || s.ratio == 1 {
		sepHeight += itemSpacing
	} else {
		sepHeight += 2 * itemSpacing
	}

	// Render the top child/splitter
	topSize := imgui.Vec2{totalWidth, s.ratio * (totalHeight - sepHeight)}
	pos := imgui.GetCursorPos()
	if topSize.Y > 0 {
		s.top.SetSize(topSize)
		s.top.Render()
		pos.Y = pos.Y + topSize.Y + itemSpacing
		imgui.SetCursorPos(pos)
	}

	// Render the separator
	imgui.PushStyleColor(imgui.Col_Button, s.sepColor)
	imgui.PushStyleColor(imgui.Col_ButtonHovered, s.sepColor)
	imgui.PushStyleColor(imgui.Col_ButtonActive, s.sepColor)
	imgui.ButtonCS(s.sepID, s.sepSize)
	imgui.PopStyleColor(3)
	if imgui.IsItemHovered(imgui.HoveredFlags_None) || s.drag {
		imgui.SetMouseCursor(imgui.MouseCursor_ResizeNS)
		if imgui.IsMouseClicked(imgui.MouseButton_Left, false) {
			s.drag = true
		}
	}

	// Render the bottom child/splitter
	bottomSize := imgui.Vec2{totalWidth, totalHeight - topSize.Y - sepHeight}
	if bottomSize.Y > 0 {
		pos.Y = pos.Y + s.sepSize.Y + itemSpacing
		imgui.SetCursorPos(pos)
		s.bottom.SetSize(bottomSize)
		s.bottom.Render()
	}

	// If dragging, get mouse delta
	delta := imgui.Vec2{}
	if s.drag {
		delta = imgui.GetMouseDragDelta(imgui.MouseButton_Left, -1.0)
		s.sepColor = s.sepColorOn
		// If mouse button released, terminates drag and
		// checks if separator is within margins.
		if !imgui.IsMouseDown(imgui.MouseButton_Left) {
			s.drag = false
			s.sepColor = s.sepColorOff
			if s.ratio < s.margin {
				s.ratio = 0
			} else if s.ratio > 1.0-s.margin {
				s.ratio = 1
			}
		}
	}

	// If splitter was dragged recalculates its relative position
	if delta.Y != 0 {
		s.ratio = (topSize.Y + delta.Y) / (totalHeight - sepHeight)
		if delta.Y > 0 && s.ratio > 1.0-s.margin {
			s.ratio = 1.0
			s.drag = false
		}
		if delta.Y < 0 && s.ratio < s.margin {
			s.ratio = 0
			s.drag = false
		}
		imgui.ResetMouseDragDelta(imgui.MouseButton_Left)
		if !s.drag {
			s.sepColor = s.sepColorOff
		}
	}
}
