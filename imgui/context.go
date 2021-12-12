package imgui

// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"

func CreateContext(sharedFontAtlas *FontAtlas) *Context {

	return (*Context)(C.igCreateContext((*C.ImFontAtlas)(sharedFontAtlas)))
}

func DestroyContext(ctx *Context) {

	C.igDestroyContext((*C.ImGuiContext)(ctx))
}

func GetCurrentContext() *Context {

	return (*Context)(C.igGetCurrentContext())
}

func SetCurrentContext(ctx *Context) {

	C.igSetCurrentContext((*C.ImGuiContext)(ctx))
}
