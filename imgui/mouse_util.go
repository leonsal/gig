package imgui

// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
// static ImVec2 xGetMouseDragDelta(ImGuiMouseButton button, float lock_threshold) {
// 		ImVec2 delta;
//		igGetMouseDragDelta(&delta, button, lock_threshold);
//		return delta;
// }
// static ImVec2 xGetMousePos() {
//		ImVec2 out;
//		igGetMousePos(&out);
//		return out;
// }
import "C"

func IsMouseDown(button MouseButton) bool {

	return bool(C.igIsMouseDown(C.ImGuiMouseButton(button)))
}

func IsMouseClicked(button MouseButton, repeat bool) bool {

	return bool(C.igIsMouseClicked(C.ImGuiMouseButton(button), C.bool(repeat)))
}

func IsMouseReleased(button MouseButton) bool {

	return bool(C.igIsMouseReleased(C.ImGuiMouseButton(button)))
}

func IsMouseDoubleClicked(button MouseButton) bool {

	return bool(C.igIsMouseDoubleClicked(C.ImGuiMouseButton(button)))
}

func IsMouseHoveringRect(rMin, rMax Vec2, clip bool) bool {

	return bool(C.igIsMouseHoveringRect(rMin.ImVec2(), rMax.ImVec2(), C.bool(clip)))
}

func IsMousePosValid(mousePos *Vec2) bool {

	mouse_pos := mousePos.ImVec2()
	res := bool(C.igIsMousePosValid(&mouse_pos))
	*mousePos = mouse_pos.Vec2()
	return bool(res)
}

func IsAnyMouseDown() bool {

	return bool(C.igIsAnyMouseDown())
}

func GetMousePos() Vec2 {

	return C.xGetMousePos().Vec2()
}

func GetMousePosOnOpeningCurrentPopup() Vec2 {

	var out C.ImVec2
	C.igGetMousePosOnOpeningCurrentPopup(&out)
	return out.Vec2()
}

func IsMouseDraggin(button MouseButton, lockThreshold float32) bool {

	return bool(C.igIsMouseDragging(C.ImGuiMouseButton(button), C.float(lockThreshold)))
}

func GetMouseDragDelta(button MouseButton, lockThreshold float32) Vec2 {

	return C.xGetMouseDragDelta(C.ImGuiMouseButton(button), C.float(lockThreshold)).Vec2()
}

func ResetMouseDragDelta(button MouseButton) {

	C.igResetMouseDragDelta(C.ImGuiMouseButton(button))
}

func GetMouseCursor() MouseCursor {

	return MouseCursor(C.igGetMouseCursor())
}

func SetMouseCursor(cursorType MouseCursor) {

	C.igSetMouseCursor(C.ImGuiMouseCursor(cursorType))
}

func CaptureMouseFromApp(wantCaptureMouseValue bool) {

	C.igCaptureMouseFromApp(C.bool(wantCaptureMouseValue))
}
