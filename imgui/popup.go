package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"
import "unsafe"

func BeginPopup(strId string, flags WindowFlags) bool {

	str_id := C.CString(strId)
	defer C.free(unsafe.Pointer(str_id))
	return bool(C.igBeginPopup(str_id, C.ImGuiWindowFlags(flags)))
}

func BeginPopupCS(strId CString, flags WindowFlags) bool {

	return bool(C.igBeginPopup(strId.CString(), C.ImGuiWindowFlags(flags)))
}

func BeginPopupModal(name string, pOpen *bool, flags WindowFlags) bool {

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return bool(C.igBeginPopupModal(cname, (*C.bool)(pOpen), C.ImGuiWindowFlags(flags)))
}

func BeginPopupModalCS(name CString, pOpen *bool, flags WindowFlags) bool {

	return bool(C.igBeginPopupModal(name.CString(), (*C.bool)(pOpen), C.ImGuiWindowFlags(flags)))
}

func EndPopup() {

	C.igEndPopup()
}

func OpenPopup(strId string, popupFlags PopupFlags) {

	str_id := C.CString(strId)
	defer C.free(unsafe.Pointer(str_id))
	C.igOpenPopup_Str(str_id, C.ImGuiPopupFlags(popupFlags))
}

func OpenPopupCS(strId CString, popupFlags PopupFlags) {

	C.igOpenPopup_Str(strId.CString(), C.ImGuiPopupFlags(popupFlags))
}

func OpenPopupId(id ID, popupFlags PopupFlags) {

	C.igOpenPopup_ID(C.ImGuiID(id), C.ImGuiPopupFlags(popupFlags))
}

func OpenPopupOnItemClick(strId string, popupFlags PopupFlags) {

	str_id := C.CString(strId)
	defer C.free(unsafe.Pointer(str_id))
	C.igOpenPopupOnItemClick(str_id, C.ImGuiPopupFlags(popupFlags))
}

func CloseCurrentPopup() {

	C.igCloseCurrentPopup()
}

func BeginPopupContextItem(strId string, popupFlags PopupFlags) bool {

	str_id := C.CString(strId)
	defer C.free(unsafe.Pointer(str_id))
	return bool(C.igBeginPopupContextItem(str_id, C.ImGuiPopupFlags(popupFlags)))
}

func BeginPopupContextWindow(strId string, popupFlags PopupFlags) bool {

	str_id := C.CString(strId)
	defer C.free(unsafe.Pointer(str_id))
	return bool(C.igBeginPopupContextWindow(str_id, C.ImGuiPopupFlags(popupFlags)))
}

func BeginPopupContextVoid(strId string, popupFlags PopupFlags) bool {

	str_id := C.CString(strId)
	defer C.free(unsafe.Pointer(str_id))
	return bool(C.igBeginPopupContextVoid(str_id, C.ImGuiPopupFlags(popupFlags)))
}

func IsPopupOpen(strId string, flags PopupFlags) bool {

	str_id := C.CString(strId)
	defer C.free(unsafe.Pointer(str_id))
	return bool(C.igIsPopupOpen_Str(str_id, C.ImGuiPopupFlags(flags)))
}
