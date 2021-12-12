package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"
import "unsafe"

func GetClipboardText() string {

	return C.GoString(C.igGetClipboardText())
}

func SetClipboardText(text string) {

	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.igSetClipboardText(ctext)
}
