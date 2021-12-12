package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"
import "unsafe"

type Font C.ImFont

type GlyphRanges []C.ImWchar

type FontConfig C.ImFontConfig

type FontAtlas C.ImFontAtlas

func NewFontConfig() *FontConfig {

	return (*FontConfig)(C.ImFontConfig_ImFontConfig())
}

func (fc *FontConfig) Destroy() {

	C.ImFontConfig_destroy((*C.ImFontConfig)(fc))
}

func (fc *FontConfig) SetGlyphOffset(off Vec2) {

	fc.GlyphOffset.x = C.float(off.X)
	fc.GlyphOffset.y = C.float(off.Y)
}

func (fa *FontAtlas) GetGlyphRangesDefault() GlyphRanges {

	cgr := C.ImFontAtlas_GetGlyphRangesDefault((*C.ImFontAtlas)(fa))
	return cToGlyphRanges(cgr)
}

func (fa *FontAtlas) AddFontDefault(font_cfg *FontConfig) *Font {

	font := C.ImFontAtlas_AddFontDefault((*C.ImFontAtlas)(fa), (*C.ImFontConfig)(font_cfg))
	return (*Font)(font)
}

func (fa *FontAtlas) AddFontFromMemoryTTF(font_data []byte, size_pixels float32, font_cfg *FontConfig, ranges GlyphRanges) *Font {

	// AddFontFromMemory assumes the ownership of the font data passed and destroy it when needed
	fdata := C.CBytes(font_data)
	font := C.ImFontAtlas_AddFontFromMemoryTTF((*C.ImFontAtlas)(fa), fdata, C.int(len(font_data)),
		C.float(size_pixels), (*C.ImFontConfig)(font_cfg), glyphRangesToC(ranges))
	return (*Font)(font)
}

func (fa *FontAtlas) AddFontFromFileTTF(filename string, size_pixels float32, font_cfg *FontConfig, ranges GlyphRanges) *Font {

	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	font := C.ImFontAtlas_AddFontFromFileTTF((*C.ImFontAtlas)(fa), cs, C.float(size_pixels),
		(*C.ImFontConfig)(font_cfg), glyphRangesToC(ranges))
	return (*Font)(font)
}

func (fa *FontAtlas) Build() bool {

	res := C.ImFontAtlas_Build((*C.ImFontAtlas)(fa))
	return bool(res)
}

func (fa *FontAtlas) Clear() {

	C.ImFontAtlas_Clear((*C.ImFontAtlas)(fa))
}

func glyphRangesToC(gr GlyphRanges) *C.ImWchar {

	if len(gr) == 0 {
		return nil
	}
	return (*C.ImWchar)(unsafe.Pointer(&gr[0]))
}

func cToGlyphRanges(cgr *C.ImWchar) GlyphRanges {

	slice := (*[1 << 28]C.ImWchar)(unsafe.Pointer(cgr))[:10]
	var gr GlyphRanges
	for _, v := range slice {
		gr = append(gr, v)
		if v == 0 {
			break
		}
	}
	return gr
}
