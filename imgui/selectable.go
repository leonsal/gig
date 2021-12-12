package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"
import "unsafe"

func Selectable(label string, selected bool, flags SelectableFlags, size Vec2) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igSelectable_Bool(clabel, C.bool(selected), C.ImGuiSelectableFlags(flags), size.ImVec2()))
}

func SelectableCS(label CString, selected bool, flags SelectableFlags, size Vec2) bool {

	return bool(C.igSelectable_Bool(label.CString(), C.bool(selected), C.ImGuiSelectableFlags(flags), size.ImVec2()))
}

func SelectablePtr(label string, pSelected *bool, flags SelectableFlags, size Vec2) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igSelectable_BoolPtr(clabel, (*C.bool)(pSelected), C.ImGuiSelectableFlags(flags), size.ImVec2()))
}

func SelectablePtrCS(label CString, pSelected *bool, flags SelectableFlags, size Vec2) bool {

	return bool(C.igSelectable_BoolPtr(label.CString(), (*C.bool)(pSelected), C.ImGuiSelectableFlags(flags), size.ImVec2()))
}
