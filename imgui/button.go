package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"
import "unsafe"

func Button(label string, size Vec2) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igButton(clabel, size.ImVec2()))
}

func ButtonCS(label CString, size Vec2) bool {

	return bool(C.igButton(label.CString(), size.ImVec2()))
}

func SmallButton(label string) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igSmallButton(clabel))
}

func SmallButtonCS(label CString) bool {

	return bool(C.igSmallButton(label.CString()))
}

func InvisibleButton(strId string, size Vec2, flags ButtonFlags) bool {

	str_id := C.CString(strId)
	defer C.free(unsafe.Pointer(str_id))
	return bool(C.igInvisibleButton(str_id, size.ImVec2(), C.ImGuiButtonFlags(flags)))
}

func InvisibleButtonCS(strId CString, size Vec2, flags ButtonFlags) bool {

	return bool(C.igInvisibleButton(strId.CString(), size.ImVec2(), C.ImGuiButtonFlags(flags)))
}

func ArrowButton(strId string, dir Dir) bool {

	str_id := C.CString(strId)
	defer C.free(unsafe.Pointer(str_id))
	return bool(C.igArrowButton(str_id, C.ImGuiDir(dir)))
}

func ArrowButtonCS(strId CString, dir Dir) bool {

	return bool(C.igArrowButton(strId.CString(), C.ImGuiDir(dir)))
}

func Image(userTextureId TextureID, size, uv0, uv1 Vec2, tintCol Vec4, borderCol Vec4) {

	C.igImage(C.ImTextureID(userTextureId), size.ImVec2(), uv0.ImVec2(), uv1.ImVec2(), tintCol.ImVec4(), borderCol.ImVec4())
}

func ImageButton(userTextureId TextureID, size, uv0, uv1 Vec2, framePadding int, bgCol Vec4, tintCol Vec4) bool {

	return bool(C.igImageButton(C.ImTextureID(userTextureId), size.ImVec2(), uv0.ImVec2(), uv1.ImVec2(),
		C.int(framePadding), bgCol.ImVec4(), tintCol.ImVec4()))
}

func Checkbox(label string, v *bool) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igCheckbox(clabel, (*C.bool)(v)))
}

func CheckboxCS(label CString, v *bool) bool {

	return bool(C.igCheckbox(label.CString(), (*C.bool)(v)))
}

func CheckboxFlags(label string, flags *int32, flagsValue int) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igCheckboxFlags_IntPtr(clabel, (*C.int)(flags), C.int(flagsValue)))
}

func CheckboxFlagsCS(label CString, flags *int32, flagsValue int) bool {

	return bool(C.igCheckboxFlags_IntPtr(label.CString(), (*C.int)(flags), C.int(flagsValue)))
}

func CheckboxFlagsUint(label string, flags *uint32, flagsValue uint) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igCheckboxFlags_UintPtr(clabel, (*C.uint)(flags), C.uint(flagsValue)))
}

func CheckboxFlagsUintCS(label CString, flags *uint32, flagsValue uint) bool {

	return bool(C.igCheckboxFlags_UintPtr(label.CString(), (*C.uint)(flags), C.uint(flagsValue)))
}

func RadioButton(label string, active bool) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igRadioButton_Bool(clabel, C.bool(active)))
}

func RadioButtonCS(label CString, active bool) bool {

	return bool(C.igRadioButton_Bool(label.CString(), C.bool(active)))
}

func RadioButtonPtr(label string, v *int32, vButton int) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igRadioButton_IntPtr(clabel, (*C.int)(v), C.int(vButton)))
}

func ProgressBar(fraction float32, sizeArg Vec2, overlay string) {

	coverlay := (*C.char)(nil)
	if len(overlay) > 0 {
		coverlay = C.CString(overlay)
		defer C.free(unsafe.Pointer(coverlay))
	}
	C.igProgressBar(C.float(fraction), sizeArg.ImVec2(), coverlay)
}

func ProgressBarCS(fraction float32, sizeArg Vec2, overlay CString) {

	C.igProgressBar(C.float(fraction), sizeArg.ImVec2(), overlay.CString())
}

func Bullet() {

	C.igBullet()
}
