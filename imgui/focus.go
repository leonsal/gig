package imgui

// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"

func SetItemDefaultFocus() {

	C.igSetItemDefaultFocus()
}

func SetKeyboardFocusHere(offset int) {

	C.igSetKeyboardFocusHere(C.int(offset))
}
