package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
// static void xSetTooltip(const char* text) {
//    igSetTooltip("%s", text);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

func BeginTooltip() {

	C.igBeginTooltip()
}

func EndTooltip() {

	C.igEndTooltip()
}

func SetTooltip(format string, args ...interface{}) {

	text := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(text))
	C.xSetTooltip(text)
}

func SetTooltipCS(text CString) {

	C.xSetTooltip(text.CString())
}
