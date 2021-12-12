package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
// static void xText(const char* s) {
//		igText("%s", s);
// }
// static void xTextColored(ImVec4 c, const char* s) {
//		igTextColored(c, "%s", s);
// }
// static void xTextDisabled(const char* s) {
//		igTextDisabled("%s", s);
// }
// static void xTextWrapped(const char* s) {
//		igTextWrapped("%s", s);
// }
// static void xLabelText(const char *label, const char* s) {
//		igLabelText(label, "%s", s);
// }
// static void xBulletText(const char* s) {
//		igBulletText("%s", s);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

func Text(format string, args ...interface{}) {

	cs := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(cs))
	C.xText(cs)
}

func TextCS(text CString) {

	C.xText(text.CString())
}

func TextColored(c Vec4, format string, args ...interface{}) {

	cs := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(cs))
	C.xTextColored(c.ImVec4(), cs)
}

func TextColoredCS(c Vec4, text CString) {

	C.xTextColored(c.ImVec4(), text.CString())
}

func TextDisabled(format string, args ...interface{}) {

	cs := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(cs))
	C.xTextDisabled(cs)
}

func TextDisabledCS(text CString) {

	C.xTextDisabled(text.CString())
}

func TextWrapped(format string, args ...interface{}) {

	cs := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(cs))
	C.xTextWrapped(cs)
}

func TextWrappedCS(text CString) {

	C.xTextWrapped(text.CString())
}

func LabelText(label string, format string, args ...interface{}) {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	ctext := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(ctext))
	C.xLabelText(clabel, ctext)
}

func LabelTextCS(label CString, text CString) {

	C.xLabelText(label.CString(), text.CString())
}

func BulletText(format string, args ...interface{}) {

	cs := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(cs))
	C.xBulletText(cs)
}

func BulletTextCS(text CString) {

	C.xBulletText(text.CString())
}
