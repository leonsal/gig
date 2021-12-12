package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"
import "unsafe"

func LoadIniSettingsFromDisk(iniFilename string) {

	ini_filename := C.CString(iniFilename)
	defer C.free(unsafe.Pointer(ini_filename))
	C.igLoadIniSettingsFromDisk(ini_filename)
}

func LoadSettingsFromMemory(iniData []byte) {

	C.igLoadIniSettingsFromMemory((*C.char)(unsafe.Pointer(&iniData[0])), C.size_t(len(iniData)))
}

func SaveIniSettingsToDisk(iniFilename string) {

	ini_filename := C.CString(iniFilename)
	defer C.free(unsafe.Pointer(ini_filename))
	C.igSaveIniSettingsToDisk(ini_filename)
}

func SaveIniSettingsToMemory() []byte {

	var out_ini_size C.size_t
	pdata := C.igSaveIniSettingsToMemory(&out_ini_size)
	return C.GoBytes(unsafe.Pointer(pdata), C.int(out_ini_size))
}
