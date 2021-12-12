package imgui

// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"

func PushClipRect(clipRectMin, clipRectMax Vec2, intersectWithCurrentClipRect bool) {

	C.igPushClipRect(clipRectMin.ImVec2(), clipRectMax.ImVec2(), C.bool(intersectWithCurrentClipRect))
}

func PopClipRect() {

	C.igPopClipRect()
}
