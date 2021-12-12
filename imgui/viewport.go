package imgui

// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"

type Viewport struct {
	cp *C.ImGuiViewport
}

func GetMainViewport() Viewport {

	return Viewport{C.igGetMainViewport()}
}

func (v Viewport) Flags() ViewportFlags {

	return ViewportFlags(v.cp.Flags)
}

func (v Viewport) Pos() Vec2 {

	return v.cp.Pos.Vec2()
}

func (v Viewport) Size() Vec2 {

	return v.cp.Size.Vec2()
}

func (v Viewport) WorkPos() Vec2 {

	return v.cp.WorkPos.Vec2()
}

func (v Viewport) WorkSize() Vec2 {

	return v.cp.WorkSize.Vec2()
}

func (v Viewport) GetCenter() Vec2 {

	var center C.ImVec2
	C.ImGuiViewport_GetCenter(&center, v.cp)
	return center.Vec2()
}

func (v Viewport) GetWorkCenter() Vec2 {

	var center C.ImVec2
	C.ImGuiViewport_GetWorkCenter(&center, v.cp)
	return center.Vec2()
}
