package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"
import "unsafe"

func BeginCombo(label, previewValue string, flags ComboFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	preview_value := C.CString(previewValue)
	defer C.free(unsafe.Pointer(preview_value))
	return bool(C.igBeginCombo(clabel, preview_value, C.ImGuiComboFlags(flags)))
}

func BeginComboCS(label, previewValue CString, flags ComboFlags) bool {

	return bool(C.igBeginCombo(label.CString(), previewValue.CString(), C.ImGuiComboFlags(flags)))
}

func EndCombo() {

	C.igEndCombo()
}

//func Combo(label string, currentItem *int, items []string, popupMaxHeightInItems int) bool {
//
//	clabel := C.CString(label)
//	defer C.free(unsafe.Pointer(clabel))
//	citems := make([]*C.char, len(items))
//	for idx, item := range items {
//		citem := C.CString(item)
//		citems[idx] = citem
//		defer C.free(unsafe.Pointer(citem))
//	}
//	current_item := C.int(*currentItem)
//	res := C.igCombo_Str_arr(clabel, &current_item, &citems[0], C.int(len(items)), C.int(popupMaxHeightInItems))
//	*currentItem = int(current_item)
//	return bool(res)
//}

func Combo(label string, currentItem *int, items []string, popupMaxHeightInItems int) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	buf := make([]byte, 0)
	for _, item := range items {
		buf = append(buf, []byte(item)...)
		buf = append(buf, 0)
	}
	current_item := C.int(*currentItem)
	res := C.igCombo_Str(clabel, &current_item, (*C.char)(unsafe.Pointer(&buf[0])), C.int(popupMaxHeightInItems))
	*currentItem = int(current_item)
	return bool(res)
}

// TODO?
//CIMGUI_API bool igCombo_FnBoolPtr(const char* label,int* current_item,bool(*items_getter)(void* data,int idx,const char** out_text),void* data,int items_count,int popup_max_height_in_items);
