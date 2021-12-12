package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"
import "unsafe"

func CalcTextSize(text string, hideTextAfterDoubleHash bool, wrapWidth float32) Vec2 {

	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	var out C.ImVec2
	C.igCalcTextSize(&out, ctext, nil, C.bool(hideTextAfterDoubleHash), C.float(wrapWidth))
	return out.Vec2()
}
