package imgui

// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"

func ColorConvertU32ToFloat4(in U32) Vec4 {

	var out C.ImVec4
	C.igColorConvertU32ToFloat4(&out, C.ImU32(in))
	return out.Vec4()
}

func ColorConvertFloat4ToU32(in Vec4) U32 {

	return U32(C.igColorConvertFloat4ToU32(in.ImVec4()))
}

func ColorConvertRGBtoHSV(r, g, b float32) (h, s, v float32) {

	var out_h, out_s, out_v C.float
	C.igColorConvertRGBtoHSV(C.float(r), C.float(g), C.float(b), &out_h, &out_s, &out_v)
	return float32(out_h), float32(out_s), float32(out_v)
}

func ColorConvertHSVtoRGB(h, s, v float32) (r, g, b float32) {

	var out_r, out_b, out_g C.float
	C.igColorConvertHSVtoRGB(C.float(h), C.float(s), C.float(v), &out_r, &out_g, &out_b)
	return float32(out_r), float32(out_g), float32(out_b)
}
