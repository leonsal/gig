package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"
import "unsafe"

func ColorEdit3(label string, col *float32, flags ColorEditFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igColorEdit3(clabel, (*C.float)(col), C.ImGuiColorEditFlags(flags)))
}

func ColorEdit3CS(label CString, col *float32, flags ColorEditFlags) bool {

	return bool(C.igColorEdit3(label.CString(), (*C.float)(col), C.ImGuiColorEditFlags(flags)))
}

func ColorEdit4(label string, col *float32, flags ColorEditFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igColorEdit4(clabel, (*C.float)(col), C.ImGuiColorEditFlags(flags)))
}

func ColorEdit4CS(label CString, col *float32, flags ColorEditFlags) bool {

	return bool(C.igColorEdit4(label.CString(), (*C.float)(col), C.ImGuiColorEditFlags(flags)))
}

func ColorPicker3(label string, col *float32, flags ColorEditFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igColorPicker3(clabel, (*C.float)(col), C.ImGuiColorEditFlags(flags)))
}

func ColorPicker3CS(label CString, col *float32, flags ColorEditFlags) bool {

	return bool(C.igColorPicker3(label.CString(), (*C.float)(col), C.ImGuiColorEditFlags(flags)))
}

func ColorPicker4(label string, col *float32, flags ColorEditFlags, refCol *float32) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igColorPicker4(clabel, (*C.float)(col), C.ImGuiColorEditFlags(flags), (*C.float)(refCol)))
}

func ColorPicker4CS(label CString, col *float32, flags ColorEditFlags, refCol *float32) bool {

	return bool(C.igColorPicker4(label.CString(), (*C.float)(col), C.ImGuiColorEditFlags(flags), (*C.float)(refCol)))
}

func ColorButton(descId string, col Vec4, flags ColorEditFlags, size Vec2) bool {

	desc_id := C.CString(descId)
	defer C.free(unsafe.Pointer(desc_id))
	return bool(C.igColorButton(desc_id, col.ImVec4(), C.ImGuiColorEditFlags(flags), size.ImVec2()))
}

func ColorButtonCS(descId CString, col Vec4, flags ColorEditFlags, size Vec2) bool {

	return bool(C.igColorButton(descId.CString(), col.ImVec4(), C.ImGuiColorEditFlags(flags), size.ImVec2()))
}

func SetColorEditOptions(flags ColorEditFlags) {

	C.igSetColorEditOptions(C.ImGuiColorEditFlags(flags))
}
