package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"
import "unsafe"

func ValueBool(prefix string, b bool) {

	cprefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(cprefix))
	C.igValue_Bool(cprefix, C.bool(b))
}

func ValueInt(prefix string, v int) {

	cprefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(cprefix))
	C.igValue_Int(cprefix, C.int(v))
}

func ValueUint(prefix string, v uint) {

	cprefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(cprefix))
	C.igValue_Uint(cprefix, C.uint(v))
}

func ValueFloat(prefix string, v float32, floatFormat string) {

	cprefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(cprefix))
	float_format := C.CString(prefix)
	defer C.free(unsafe.Pointer(float_format))
	C.igValue_Float(cprefix, C.float(v), float_format)
}
