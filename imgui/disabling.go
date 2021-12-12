package imgui

// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"

func BeginDisabled(disabled bool) {

	C.igBeginDisabled(C.bool(disabled))
}

func EndDisabled() {

	C.igEndDisabled()
}
