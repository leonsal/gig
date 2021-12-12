package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"
import "unsafe"

func BeginMenuBar() bool {

	return bool(C.igBeginMenuBar())
}

func EndMenuBar() {

	C.igEndMenuBar()
}

func BeginMainMenuBar() bool {

	return bool(C.igBeginMainMenuBar())
}

func EndMainMenuBar() {

	C.igEndMainMenuBar()
}

func BeginMenu(label string, enabled bool) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igBeginMenu(clabel, C.bool(enabled)))
}

func BeginMenuCS(label CString, enabled bool) bool {

	return bool(C.igBeginMenu(label.CString(), C.bool(enabled)))
}

func EndMenu() {

	C.igEndMenu()
}

func MenuItem(label, shortcut string, selected, enabled bool) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cshortcut := C.CString(shortcut)
	defer C.free(unsafe.Pointer(cshortcut))
	return bool(C.igMenuItem_Bool(clabel, cshortcut, C.bool(selected), C.bool(enabled)))
}

func MenuItemCS(label, shortcut CString, selected, enabled bool) bool {

	return bool(C.igMenuItem_Bool(label.CString(), shortcut.CString(), C.bool(selected), C.bool(enabled)))
}

func MenuItemPtr(label, shortcut string, pSelected *bool, enabled bool) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cshortcut := C.CString(shortcut)
	defer C.free(unsafe.Pointer(cshortcut))
	return bool(C.igMenuItem_BoolPtr(clabel, cshortcut, (*C.bool)(pSelected), C.bool(enabled)))
}

func MenuItemPtrCS(label, shortcut CString, pSelected *bool, enabled bool) bool {

	return bool(C.igMenuItem_BoolPtr(label.CString(), shortcut.CString(), (*C.bool)(pSelected), C.bool(enabled)))
}
