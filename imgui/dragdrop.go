package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"
import "unsafe"

func BeginDragDropSource(flags DragDropFlags) bool {

	return bool(C.igBeginDragDropSource(C.ImGuiDragDropFlags(flags)))
}

func SetDragDropPayload(typ string, data []byte, cond Cond) bool {

	ctyp := C.CString(typ)
	defer C.free(unsafe.Pointer(ctyp))
	return bool(C.igSetDragDropPayload(ctyp, unsafe.Pointer(&data[0]), C.size_t(len(data)), C.ImGuiCond(cond)))
}

func EndDragDropSource() {

	C.igEndDragDropSource()
}

func BeginDragDropTarget() bool {

	return bool(C.igBeginDragDropTarget())
}

func AcceptDragDropPayload(typ string, flags DragDropFlags) *Payload {

	ctyp := C.CString(typ)
	defer C.free(unsafe.Pointer(ctyp))
	return (*Payload)(C.igAcceptDragDropPayload(ctyp, C.ImGuiDragDropFlags(flags)))
}

func EndDragDropTarget() {

	C.igEndDragDropTarget()
}

func GetDragDropPayload() *Payload {

	return (*Payload)(C.igGetDragDropPayload())
}
