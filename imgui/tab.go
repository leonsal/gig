package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"
import "unsafe"

func BeginTabBar(strId string, flags TabBarFlags) bool {

	str_id := C.CString(strId)
	defer C.free(unsafe.Pointer(str_id))
	return bool(C.igBeginTabBar(str_id, C.ImGuiTabBarFlags(flags)))
}

func BeginTabBarCS(strId CString, flags TabBarFlags) bool {

	return bool(C.igBeginTabBar(strId.CString(), C.ImGuiTabBarFlags(flags)))
}

func EndTabBar() {

	C.igEndTabBar()
}

func BeginTabItem(label string, pOpen *bool, flags TabItemFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igBeginTabItem(clabel, (*C.bool)(pOpen), C.ImGuiTabItemFlags(flags)))
}

func BeginTabItemCS(label CString, pOpen *bool, flags TabItemFlags) bool {

	return bool(C.igBeginTabItem(label.CString(), (*C.bool)(pOpen), C.ImGuiTabItemFlags(flags)))
}

func EndTabItem() {

	C.igEndTabItem()
}

func TabItemButton(label string, flags TabItemFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igTabItemButton(clabel, C.ImGuiTabItemFlags(flags)))
}

func TabItemButtonCS(label CString, flags TabItemFlags) bool {

	return bool(C.igTabItemButton(label.CString(), C.ImGuiTabItemFlags(flags)))
}

func SetTabItemClosed(tabOrDockedWindowLabel string) {

	tab_or_docked_window_label := C.CString(tabOrDockedWindowLabel)
	defer C.free(unsafe.Pointer(tab_or_docked_window_label))
	C.igSetTabItemClosed(tab_or_docked_window_label)
}
