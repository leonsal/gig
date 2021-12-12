package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"
import "unsafe"

func BeginListBox(label string, size Vec2) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igBeginListBox(clabel, size.ImVec2()))
}

func BeginListBoxCS(label CString, size Vec2) bool {

	return bool(C.igBeginListBox(label.CString(), size.ImVec2()))
}

func EndListBox() {

	C.igEndListBox()
}

func ListBox(label string, currentItem *int, items []string, heightInItems int) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	citems := make([]*C.char, len(items))
	current_item := C.int(*currentItem)
	for i := 0; i < len(items); i++ {
		citems[i] = C.CString(items[i])
		defer C.free(unsafe.Pointer(citems[i]))
	}
	res := C.igListBox_Str_arr(clabel, &current_item, &citems[0], C.int(len(citems)), C.int(heightInItems))
	*currentItem = int(current_item)
	return bool(res)
}

// TODO ?
//CIMGUI_API bool igListBox_FnBoolPtr(const char* label,int* current_item,bool(*items_getter)(void* data,int idx,const char** out_text),void* data,int items_count,int height_in_items);
