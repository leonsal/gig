package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"
import "unsafe"

func LogToTTY(autoOpenDepth int) {

	C.igLogToTTY(C.int(autoOpenDepth))
}

func LogToFile(autoOpenDepth int, filename string) {

	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	C.igLogToFile(C.int(autoOpenDepth), cfilename)
}

func LogToClipboard(autoOpenDepth int) {

	C.igLogToClipboard(C.int(autoOpenDepth))
}

func LogFinish() {

	C.igLogFinish()
}

func LogButtons() {

	C.igLogButtons()
}

// Not implemented
//CIMGUI_API void igLogTextV(const char* fmt,va_list args);
