package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"
import "unsafe"

func ShowDemoWindow(pOpen *bool) {

	C.igShowDemoWindow((*C.bool)(pOpen))
}

func ShowMetricsWindow(pOpen *bool) {

	C.igShowMetricsWindow((*C.bool)(pOpen))
}

func ShowStackToolWindow(pOpen *bool) {

	C.igShowStackToolWindow((*C.bool)(pOpen))
}

func ShowAboutWindow(pOpen *bool) {

	C.igShowAboutWindow((*C.bool)(pOpen))
}

func ShowStyleEditor(ref *Style) {

	C.igShowStyleEditor((*C.ImGuiStyle)(ref))
}

func ShowStyleSelector(label string) {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	C.igShowStyleSelector(clabel)
}

func ShowFontSelector(label string) {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	C.igShowFontSelector(clabel)
}

func ShowUserGuide() {

	C.igShowUserGuide()
}

func GetVersion() string {

	return C.GoString(C.igGetVersion())
}
