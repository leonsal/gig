package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"
import "unsafe"

type IO struct {
	c           *C.ImGuiIO
	iniFilename *C.char
}

func GetIO() *IO {

	return &IO{C.igGetIO(), nil}
}

func (io *IO) ConfigFlags() ConfigFlags {

	return ConfigFlags(io.c.ConfigFlags)
}

func (io IO) SetConfigFlags(flags ConfigFlags) IO {

	io.c.ConfigFlags = C.ImGuiConfigFlags(flags)
	return io
}

func (io IO) DeltaTime() float32 {

	return float32(io.c.DeltaTime)
}

func (io IO) Fonts() *FontAtlas {

	return (*FontAtlas)(io.c.Fonts)
}

func (io IO) SetIniFilename(name string) {

	if io.iniFilename != nil {
		C.free(unsafe.Pointer(io.iniFilename))
		io.iniFilename = nil
	}
	if len(name) == 0 {
		io.c.IniFilename = nil
		return
	}
	io.iniFilename = C.CString(name)
	io.c.IniFilename = io.iniFilename
}
