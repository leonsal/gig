package gig

import (
	"strconv"

	"github.com/leonsal/gig/imgui"
)

type ISplitted interface {
	SetPos(imgui.Vec2, imgui.Cond)
	SetSize(imgui.Vec2, imgui.Cond)
	Render()
}

type WindowHSplitter struct {
	Widget
	left        ISplitted
	right       ISplitted
	pos         imgui.Vec2
	size        imgui.Vec2
	sepWindow   *Window
	ratio       float32
	sepWidth    float32
	sepColor    imgui.Vec4
	sepColorSel imgui.Vec4
	drag        bool
	margin      float32
	maximize    bool
}

func NewWindowHSplitter(left ISplitted, right ISplitted) *WindowHSplitter {

	s := new(WindowHSplitter)
	s.Init(s)
	s.SetLeft(left)
	s.SetRight(right)
	s.ratio = 0.5
	s.margin = 0.1
	s.sepWidth = 4
	s.sepColor = imgui.ColorName("Gray")
	s.sepColorSel = imgui.ColorName("White")

	// Creates separator window
	s.sepWindow = NewWindow(strconv.Itoa(s.ID()))
	s.sepWindow.SetFlags(imgui.WindowFlags_NoDecoration | imgui.WindowFlags_NoMove)
	s.sepWindow.SetStyleColor(imgui.Col_WindowBg, s.sepColor)
	s.sepWindow.SetStyleVarVec2(imgui.StyleVar_WindowMinSize, imgui.Vec2{s.sepWidth, -1})
	s.sepWindow.SetRendering(func(sw *Window) {
		if imgui.IsWindowHovered(imgui.HoveredFlags_None) || s.drag {
			imgui.SetMouseCursor(imgui.MouseCursor_ResizeEW)
			if imgui.IsMouseClicked(imgui.MouseButton_Left, false) {
				s.drag = true
			}
		}
	})
	s.SetRender(s.Render)
	return s
}

func (s *WindowHSplitter) SetLeft(left ISplitted) *WindowHSplitter {

	s.left = left
	if s.left == nil {
		return s
	}
	wflags := imgui.WindowFlags_NoCollapse | imgui.WindowFlags_NoMove | imgui.WindowFlags_NoResize
	wleft, ok := s.left.(*Window)
	if ok {
		wleft.SetFlags(wleft.Flags() | wflags)
		wleft.SetStyleVar(imgui.StyleVar_WindowRounding, 0)
		wleft.SetStyleVarVec2(imgui.StyleVar_WindowMinSize, imgui.Vec2{0, -1})
	}
	return s
}

func (s *WindowHSplitter) SetRight(right ISplitted) *WindowHSplitter {

	s.right = right
	if s.right == nil {
		return s
	}
	wflags := imgui.WindowFlags_NoCollapse | imgui.WindowFlags_NoMove | imgui.WindowFlags_NoResize
	wright, ok := s.right.(*Window)
	if ok {
		wright.SetFlags(wright.Flags() | wflags)
		wright.SetStyleVar(imgui.StyleVar_WindowRounding, 0)
		wright.SetStyleVarVec2(imgui.StyleVar_WindowMinSize, imgui.Vec2{0, -1})
	}
	return s
}

func (s *WindowHSplitter) SetPos(pos imgui.Vec2, cond imgui.Cond) {

	s.pos = pos
}

func (s *WindowHSplitter) SetSize(size imgui.Vec2, cond imgui.Cond) {

	s.size = size
}

func (s *WindowHSplitter) SetRatio(ratio float32) *WindowHSplitter {

	s.ratio = ClampFloat32(ratio, 0, 1)
	return s
}

func (s *WindowHSplitter) SetMargin(margin float32) *WindowHSplitter {

	s.margin = ClampFloat32(margin, 0, 0.4)
	return s
}

func (s *WindowHSplitter) SetMaximize(maximize bool) *WindowHSplitter {

	s.maximize = maximize
	return s
}

func (s *WindowHSplitter) Render() {

	if s.maximize {
		vp := imgui.GetMainViewport()
		s.SetPos(vp.WorkPos(), 0)
		s.SetSize(vp.WorkSize(), 0)
	}

	// Checks if only left window/splitter is set
	if s.left != nil && s.right == nil {
		s.left.SetPos(s.pos, imgui.Cond_None)
		s.left.SetSize(s.size, imgui.Cond_None)
		s.left.Render()
		return
	}

	// Checks if only right window/splitter is set
	if s.left == nil && s.right != nil {
		s.right.SetPos(s.pos, imgui.Cond_None)
		s.right.SetSize(s.size, imgui.Cond_None)
		s.right.Render()
		return
	}

	// Render left window/splitter
	leftWidth := s.ratio * (s.size.X - s.sepWidth)
	if leftWidth > 0 {
		s.left.SetPos(s.pos, imgui.Cond_None)
		s.left.SetSize(imgui.Vec2{leftWidth, s.size.Y}, imgui.Cond_None)
		s.left.Render()
	}
	cx := s.pos.X + leftWidth

	// Render splitter separator window
	s.sepWindow.SetPos(imgui.Vec2{cx, s.pos.Y}, imgui.Cond_None)
	s.sepWindow.SetSize(imgui.Vec2{s.sepWidth, s.size.Y}, imgui.Cond_None)
	s.sepWindow.Render()
	cx += s.sepWidth

	// Render right window
	rightWidth := (1 - s.ratio) * (s.size.X - s.sepWidth)
	if rightWidth > 0 {
		s.right.SetPos(imgui.Vec2{cx, s.pos.Y}, imgui.Cond_None)
		s.right.SetSize(imgui.Vec2{rightWidth, s.size.Y}, imgui.Cond_None)
		s.right.Render()
	}

	// If dragging, get mouse delta
	delta := imgui.Vec2{}
	if s.drag {
		delta = imgui.GetMouseDragDelta(imgui.MouseButton_Left, -1.0)
		s.sepWindow.SetStyleColor(imgui.Col_WindowBg, s.sepColorSel)
		// If mouse button released, terminates drag and
		// checks if separator is within margins.
		if !imgui.IsMouseDown(imgui.MouseButton_Left) {
			s.drag = false
			s.sepWindow.SetStyleColor(imgui.Col_WindowBg, s.sepColor)
			if s.ratio < s.margin {
				s.ratio = 0
			} else if s.ratio > 1.0-s.margin {
				s.ratio = 1
			}
		}
	}

	// If splitter was dragged recalculates its relative position
	if delta.X != 0 {
		s.ratio = (leftWidth + delta.X) / (s.size.X - s.sepWidth)
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
			s.sepWindow.SetStyleColor(imgui.Col_WindowBg, s.sepColor)
		}
	}
}

type WindowVSplitter struct {
	Widget
	top         ISplitted
	bottom      ISplitted
	pos         imgui.Vec2
	size        imgui.Vec2
	sepWindow   *Window
	ratio       float32
	sepHeight   float32
	sepColor    imgui.Vec4
	sepColorSel imgui.Vec4
	drag        bool
	margin      float32
	maximize    bool
}

func NewWindowVSplitter(top ISplitted, bottom ISplitted) *WindowVSplitter {

	s := new(WindowVSplitter)
	s.Init(s)
	s.SetTop(top)
	s.SetBottom(bottom)
	s.ratio = 0.5
	s.margin = 0.1
	s.sepHeight = 4
	s.sepColor = imgui.ColorName("Gray")
	s.sepColorSel = imgui.ColorName("White")

	// Creates separator window
	s.sepWindow = NewWindow(strconv.Itoa(s.ID()))
	s.sepWindow.SetFlags(imgui.WindowFlags_NoDecoration | imgui.WindowFlags_NoMove)
	s.sepWindow.SetStyleColor(imgui.Col_WindowBg, s.sepColor)
	s.sepWindow.SetStyleVarVec2(imgui.StyleVar_WindowMinSize, imgui.Vec2{-1, s.sepHeight})
	s.sepWindow.SetRendering(func(sw *Window) {
		if imgui.IsWindowHovered(imgui.HoveredFlags_None) || s.drag {
			imgui.SetMouseCursor(imgui.MouseCursor_ResizeNS)
			if imgui.IsMouseClicked(imgui.MouseButton_Left, false) {
				s.drag = true
			}
		}
	})
	s.SetRender(s.Render)
	return s
}

func (s *WindowVSplitter) SetTop(top ISplitted) *WindowVSplitter {

	s.top = top
	if s.top == nil {
		return s
	}
	wtop, ok := top.(*Window)
	if ok {
		wflags := imgui.WindowFlags_NoCollapse | imgui.WindowFlags_NoMove | imgui.WindowFlags_NoResize
		wtop.SetFlags(wtop.Flags() | wflags)
		wtop.SetStyleVar(imgui.StyleVar_WindowRounding, 0)
		wtop.SetStyleVarVec2(imgui.StyleVar_WindowMinSize, imgui.Vec2{-1, 0})
	}
	return s
}

func (s *WindowVSplitter) SetBottom(bottom ISplitted) *WindowVSplitter {

	s.bottom = bottom
	if s.bottom == nil {
		return s
	}
	wbottom, ok := bottom.(*Window)
	if ok {
		wflags := imgui.WindowFlags_NoCollapse | imgui.WindowFlags_NoMove | imgui.WindowFlags_NoResize
		wbottom.SetFlags(wbottom.Flags() | wflags)
		wbottom.SetStyleVar(imgui.StyleVar_WindowRounding, 0)
		wbottom.SetStyleVarVec2(imgui.StyleVar_WindowMinSize, imgui.Vec2{-1, 0})
	}
	return s
}

func (s *WindowVSplitter) SetPos(pos imgui.Vec2, cond imgui.Cond) {

	s.pos = pos
}

func (s *WindowVSplitter) SetSize(size imgui.Vec2, cond imgui.Cond) {

	s.size = size
}

func (s *WindowVSplitter) SetRatio(ratio float32) *WindowVSplitter {

	s.ratio = ClampFloat32(ratio, 0, 1)
	return s
}

func (s *WindowVSplitter) SetMargin(margin float32) *WindowVSplitter {

	s.margin = ClampFloat32(margin, 0, 0.4)
	return s
}

func (s *WindowVSplitter) SetMaximize(maximize bool) *WindowVSplitter {

	s.maximize = maximize
	return s
}

func (s *WindowVSplitter) Render() {

	if s.maximize {
		vp := imgui.GetMainViewport()
		s.SetPos(vp.WorkPos(), imgui.Cond_None)
		s.SetSize(vp.WorkSize(), imgui.Cond_None)
	}

	// Checks if only top window/splitter is set
	if s.top != nil && s.bottom == nil {
		s.top.SetPos(s.pos, imgui.Cond_None)
		s.top.SetSize(s.size, imgui.Cond_None)
		s.top.Render()
		return
	}

	// Checks if only bottom window/splitter is set
	if s.top == nil && s.bottom != nil {
		s.bottom.SetPos(s.pos, imgui.Cond_None)
		s.bottom.SetSize(s.size, imgui.Cond_None)
		s.bottom.Render()
		return
	}

	// Render top window/splitter
	topHeight := s.ratio * (s.size.Y - s.sepHeight)
	if topHeight > 0 {
		s.top.SetPos(s.pos, imgui.Cond_None)
		s.top.SetSize(imgui.Vec2{s.size.X, topHeight}, imgui.Cond_None)
		s.top.Render()
	}
	cy := s.pos.Y + topHeight

	// Render splitter separator window
	s.sepWindow.SetPos(imgui.Vec2{s.pos.X, cy}, imgui.Cond_None)
	s.sepWindow.SetSize(imgui.Vec2{s.size.X, s.sepHeight}, imgui.Cond_None)
	s.sepWindow.Render()
	cy += s.sepHeight

	// Render bottom window/splitter
	bottomHeight := (1 - s.ratio) * (s.size.Y - s.sepHeight)
	if bottomHeight > 0 {
		s.bottom.SetPos(imgui.Vec2{s.pos.X, cy}, imgui.Cond_None)
		s.bottom.SetSize(imgui.Vec2{s.size.X, bottomHeight}, imgui.Cond_None)
		s.bottom.Render()
	}

	// If dragging, get mouse delta
	delta := imgui.Vec2{}
	if s.drag {
		delta = imgui.GetMouseDragDelta(imgui.MouseButton_Left, -1.0)
		s.sepWindow.SetStyleColor(imgui.Col_WindowBg, s.sepColorSel)
		// If mouse button released, terminates drag and
		// checks if separator is within margins.
		if !imgui.IsMouseDown(imgui.MouseButton_Left) {
			s.drag = false
			s.sepWindow.SetStyleColor(imgui.Col_WindowBg, s.sepColor)
			if s.ratio < s.margin {
				s.ratio = 0
			} else if s.ratio > 1.0-s.margin {
				s.ratio = 1
			}
		}
	}

	// If splitter was dragged recalculates its relative position
	if delta.Y != 0 {
		s.ratio = (topHeight + delta.Y) / (s.size.Y - s.sepHeight)
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
			s.sepWindow.SetStyleColor(imgui.Col_WindowBg, s.sepColor)
		}
	}
}
