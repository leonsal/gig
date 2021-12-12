package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
// static ImVec2 xGetContentRegionAvail(void) {
// 		ImVec2 pos;
//		igGetContentRegionAvail(&pos);
//		return pos;
// }
// static ImVec2 xGetContentRegionMax(void) {
// 		ImVec2 out;
//		igGetContentRegionMax(&out);
//		return out;
// }
// static ImVec2 xGetWindowContentRegionMin(void) {
// 		ImVec2 out;
//		igGetWindowContentRegionMin(&out);
//		return out;
// }
// static ImVec2 xGetWindowContentRegionMax(void) {
// 		ImVec2 out;
//		igGetWindowContentRegionMax(&out);
//		return out;
// }
import "C"
import "unsafe"

func Begin(name string, popen *bool, flags WindowFlags) bool {

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return bool(C.igBegin(cname, (*C.bool)(popen), C.ImGuiWindowFlags(flags)))
}

func BeginCS(name CString, popen *bool, flags WindowFlags) bool {

	return bool(C.igBegin(name.CString(), (*C.bool)(popen), C.ImGuiWindowFlags(flags)))
}

func End() {

	C.igEnd()
}

func BeginChild(str_id string, size Vec2, border bool, flags WindowFlags) bool {

	cstr_id := C.CString(str_id)
	defer C.free(unsafe.Pointer(cstr_id))
	return bool(C.igBeginChild_Str(cstr_id, size.ImVec2(), C.bool(border), C.ImGuiWindowFlags(flags)))
}

func BeginChildID(id ID, size Vec2, border bool, flags WindowFlags) bool {

	return bool(C.igBeginChild_ID(C.ImGuiID(id), size.ImVec2(), C.bool(border), C.ImGuiWindowFlags(flags)))
}

func EndChild() {

	C.igEndChild()
}

func IsWindowAppearing() bool {

	return bool(C.igIsWindowAppearing())
}

func IsWindowCollapsed() bool {

	return bool(C.igIsWindowCollapsed())
}

func IsWindowFocused(flags FocusedFlags) bool {

	return bool(C.igIsWindowFocused(C.ImGuiFocusedFlags(flags)))
}

func IsWindowHovered(flags HoveredFlags) bool {

	return bool(C.igIsWindowHovered(C.ImGuiHoveredFlags(flags)))
}

func GetWindowDrawList() *DrawList {

	return (*DrawList)(C.igGetWindowDrawList())
}

func GetWindowPos() Vec2 {

	var out C.ImVec2
	C.igGetWindowPos(&out)
	return out.Vec2()
}

func GetWindowSize() Vec2 {

	var out C.ImVec2
	C.igGetWindowSize(&out)
	return out.Vec2()
}

func GetWindowWidth() float32 {

	return float32(C.igGetWindowWidth())
}

func GetWindowHeight() float32 {

	return float32(C.igGetWindowHeight())
}

func SetNextWindowPos(pos Vec2, cond Cond, pivot Vec2) {

	C.igSetNextWindowPos(pos.ImVec2(), C.ImGuiCond(cond), pivot.ImVec2())
}

func SetNextWindowSize(size Vec2, cond Cond) {

	C.igSetNextWindowSize(size.ImVec2(), C.ImGuiCond(cond))
}

// TODO
//CIMGUI_API void igSetNextWindowSizeConstraints(const ImVec2 size_min,const ImVec2 size_max,ImGuiSizeCallback custom_callback,void* custom_callback_data);

func SetNextWindowContentSize(size Vec2) {

	C.igSetNextWindowContentSize(size.ImVec2())
}

func SetNextWindowCollapsed(collapsed bool, cond Cond) {

	C.igSetNextWindowCollapsed(C.bool(collapsed), C.ImGuiCond(cond))
}

func SetNextWindowFocus() {

	C.igSetNextWindowFocus()
}

func SetNextWindowBgAlpha(alpha float32) {

	C.igSetNextWindowBgAlpha(C.float(alpha))
}

func SetWindowPos(pos Vec2, cond Cond) {

	C.igSetWindowPos_Vec2(pos.ImVec2(), C.ImGuiCond(cond))
}

func SetWindowSize(size Vec2, cond Cond) {

	C.igSetWindowSize_Vec2(size.ImVec2(), C.ImGuiCond(cond))
}

func SetWindowCollapsed(collapsed bool, cond Cond) {

	C.igSetWindowCollapsed_Bool(C.bool(collapsed), C.ImGuiCond(cond))
}

func SetWindowFocus() {

	C.igSetWindowFocus_Nil()
}

func SetWindowFontScale(scale float32) {

	C.igSetWindowFontScale(C.float(scale))
}

func SetWindowPosNamed(name string, pos Vec2, cond Cond) {

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.igSetWindowPos_Str(cname, pos.ImVec2(), C.ImGuiCond(cond))
}

func SetWindowSizeNamed(name string, size Vec2, cond Cond) {

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.igSetWindowSize_Str(cname, size.ImVec2(), C.ImGuiCond(cond))
}

func SetWindowCollapsedNamed(name string, collapsed bool, cond Cond) {

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.igSetWindowCollapsed_Str(cname, C.bool(collapsed), C.ImGuiCond(cond))
}

func SetWindowFocusNamed(name string) {

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.igSetWindowFocus_Str(cname)
}

func GetContentRegionAvail() Vec2 {

	return C.xGetContentRegionAvail().Vec2()
}

func GetContentRegionMax() Vec2 {

	return C.xGetContentRegionMax().Vec2()
}

func GetWindowContentRegionMin() Vec2 {

	return C.xGetWindowContentRegionMin().Vec2()
}

func GetWindowContentRegionMax() Vec2 {

	return C.xGetWindowContentRegionMax().Vec2()
}

func GetScrollX() float32 {

	return float32(C.igGetScrollX())
}

func GetScrollY() float32 {

	return float32(C.igGetScrollY())
}

func SetScrollX(scrollX float32) {

	C.igSetScrollX_Float(C.float(scrollX))
}

func SetScrollY(scrollY float32) {

	C.igSetScrollY_Float(C.float(scrollY))
}

func GetScrollMaxX() float32 {

	return float32(C.igGetScrollMaxX())
}

func GetScrollMaxY() float32 {

	return float32(C.igGetScrollMaxY())
}

func SetScrollHereX(centerXratio float32) {

	C.igSetScrollHereX(C.float(centerXratio))
}

func SetScrollHereY(centerYratio float32) {

	C.igSetScrollHereY(C.float(centerYratio))
}

func SetScrollFromPosX(localX, centerXratio float32) {

	C.igSetScrollFromPosX_Float(C.float(localX), C.float(centerXratio))
}

func SetScrollFromPosY(localY, centerYratio float32) {

	C.igSetScrollFromPosY_Float(C.float(localY), C.float(centerYratio))
}
