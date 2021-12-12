package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
// static bool xTreeNode_StrStr(const char* str_id, const char* text) {
// 		return  igTreeNode_StrStr(str_id, "%s", text);
//}
// static bool xTreeNode_Ptr(const void* ptr_id, const char* text) {
// 		return  igTreeNode_Ptr(ptr_id, "%s", text);
//}
// static bool xTreeNodeEx_StrStr(const char* str_id, ImGuiTreeNodeFlags flags, const char* text) {
// 		return  igTreeNodeEx_StrStr(str_id, flags, "%s", text);
//}
// static bool xTreeNodeEx_Ptr(const void* ptr_id, ImGuiTreeNodeFlags flags, const char* text) {
//		return igTreeNodeEx_Ptr( ptr_id, flags, "%s", text);
//}
import "C"
import (
	"fmt"
	"unsafe"
)

func TreeNode(label string) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igTreeNode_Str(clabel))
}

func TreeNodeStrId(strId string, format string, args ...interface{}) bool {

	str_id := C.CString(strId)
	defer C.free(unsafe.Pointer(str_id))
	ctext := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(ctext))
	return bool(C.xTreeNode_StrStr(str_id, ctext))
}

func TreeNodePtrId(ptrId unsafe.Pointer, format string, args ...interface{}) bool {

	ctext := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(ctext))
	return bool(C.xTreeNode_Ptr(ptrId, ctext))
}

// Not implemented
//CIMGUI_API bool igTreeNodeV_Str(const char* str_id,const char* fmt,va_list args);
//CIMGUI_API bool igTreeNodeV_Ptr(const void* ptr_id,const char* fmt,va_list args);

func TreeNodeEx(label string, flags TreeNodeFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igTreeNodeEx_Str(clabel, C.ImGuiTreeNodeFlags(flags)))
}

func TreeNodeExCS(label CString, flags TreeNodeFlags) bool {

	return bool(C.igTreeNodeEx_Str(label.CString(), C.ImGuiTreeNodeFlags(flags)))
}

func TreeNodeExStrId(strId string, flags TreeNodeFlags, format string, args ...interface{}) bool {

	str_id := C.CString(strId)
	defer C.free(unsafe.Pointer(str_id))
	ctext := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(ctext))
	return bool(C.xTreeNodeEx_StrStr(str_id, C.ImGuiTreeNodeFlags(flags), ctext))
}

func TreeNodeExStrIdCS(strId CString, flags TreeNodeFlags, text CString) bool {

	return bool(C.xTreeNodeEx_StrStr(strId.CString(), C.ImGuiTreeNodeFlags(flags), text.CString()))
}

func TreeNodeExPtrId(ptrId unsafe.Pointer, flags TreeNodeFlags, format string, args ...interface{}) bool {

	ctext := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(ctext))
	return bool(C.xTreeNodeEx_Ptr(ptrId, C.ImGuiTreeNodeFlags(flags), ctext))
}

// Not implemented
//CIMGUI_API bool igTreeNodeExV_Str(const char* str_id,ImGuiTreeNodeFlags flags,const char* fmt,va_list args);
//CIMGUI_API bool igTreeNodeExV_Ptr(const void* ptr_id,ImGuiTreeNodeFlags flags,const char* fmt,va_list args);

func TreePush(strId string) {

	str_id := C.CString(strId)
	defer C.free(unsafe.Pointer(str_id))
	C.igTreePush_Str(str_id)
}

func TreePushPtr(ptrId unsafe.Pointer) {

	C.igTreePush_Ptr(ptrId)
}

func TreePop() {

	C.igTreePop()
}

func TreeNodeToLabelSpacing() float32 {

	return float32(C.igGetTreeNodeToLabelSpacing())
}

func CollapsingHeader(label string, flags TreeNodeFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igCollapsingHeader_TreeNodeFlags(clabel, C.ImGuiTreeNodeFlags(flags)))
}

func CollapsingHeaderPtr(label string, pVisible *bool, flags TreeNodeFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igCollapsingHeader_BoolPtr(clabel, (*C.bool)(pVisible), C.ImGuiTreeNodeFlags(flags)))
}

func CollapsingHeaderPtrCS(label CString, pVisible *bool, flags TreeNodeFlags) bool {

	return bool(C.igCollapsingHeader_BoolPtr(label.CString(), (*C.bool)(pVisible), C.ImGuiTreeNodeFlags(flags)))
}

func SetNextItemOpen(isOpen bool, cond Cond) {

	C.igSetNextItemOpen(C.bool(isOpen), C.ImGuiCond(cond))
}
