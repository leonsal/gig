package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"
import "unsafe"

func PushFont(font *Font) {

	C.igPushFont((*C.ImFont)(font))
}

func PopFont() {

	C.igPopFont()
}

func PushStyleColorU32(idx Col, col U32) {

	C.igPushStyleColor_U32(C.ImGuiCol(idx), C.ImU32(col))
}

func PushStyleColor(idx Col, col Vec4) {

	C.igPushStyleColor_Vec4(C.ImGuiCol(idx), col.ImVec4())
}

func PopStyleColor(count int) {

	C.igPopStyleColor(C.int(count))
}

func PushStyleVar(idx StyleVar, val float32) {

	C.igPushStyleVar_Float(C.ImGuiStyleVar(idx), C.float(val))
}

func PushStyleVarVec2(idx StyleVar, val Vec2) {

	C.igPushStyleVar_Vec2(C.ImGuiStyleVar(idx), val.ImVec2())
}

func PopStyleVar(count int) {

	C.igPopStyleVar(C.int(count))
}

func PushAllowKeyboardFocus(allowKeyboardFocus bool) {

	C.igPushAllowKeyboardFocus(C.bool(allowKeyboardFocus))
}

func PopAllowKeyboardFocus() {

	C.igPopAllowKeyboardFocus()
}

func PushButtonRepeat(repeat bool) {

	C.igPushButtonRepeat(C.bool(repeat))
}

func PopButtonRepeat() {

	C.igPopButtonRepeat()
}

func PushItemWidth(itemWidth float32) {

	C.igPushItemWidth(C.float(itemWidth))
}

func PopItemWidth() {

	C.igPopItemWidth()
}

func SetNextItemWidth(itemWidth float32) {

	C.igSetNextItemWidth(C.float(itemWidth))
}

func CalcItemWidth() float32 {

	return float32(C.igCalcItemWidth())
}

func PushTextWrapPos(wrapLocalPosX float32) {

	C.igPushTextWrapPos(C.float(wrapLocalPosX))
}

func PopTextWrapPos() {

	C.igPopTextWrapPos()
}

func PushIDStr(strId string) {

	cstr_id := C.CString(strId)
	defer C.free(unsafe.Pointer(cstr_id))
	C.igPushID_Str(cstr_id)
}

// TODO ?
//CIMGUI_API void igPushID_StrStr(const char* str_id_begin,const char* str_id_end);

// TODO ?
//CIMGUI_API void igPushID_Ptr(const void* ptr_id);

func PushID(intId int) {

	C.igPushID_Int(C.int(intId))
}

func PopID() {

	C.igPopID()
}

func GetIDStr(strId string) ID {

	cstr_id := C.CString(strId)
	defer C.free(unsafe.Pointer(cstr_id))
	return ID(C.igGetID_Str(cstr_id))
}

// TODO ?
//CIMGUI_API ImGuiID igGetID_Ptr(const void* ptr_id);
