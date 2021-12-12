package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
// static bool xSliderInt32N(const char* label, int *p_data, int components, int* p_min, int* p_max, const char* format, ImGuiSliderFlags flags) {
// 		return igSliderScalarN(label, ImGuiDataType_S32, p_data, components,  p_min,  p_max, format, flags);
// }
// static bool xSliderInt64N(const char* label, int64_t *p_data, int components, int64_t* p_min, int64_t* p_max, const char* format, ImGuiSliderFlags flags) {
// 		return igSliderScalarN(label, ImGuiDataType_S64, p_data, components,  p_min,  p_max, format, flags);
// }
// static bool xSliderFloatN(const char* label, float* p_data, int components, float* p_min, float* p_max, const char* format, ImGuiSliderFlags flags) {
// 		return igSliderScalarN(label, ImGuiDataType_Float, p_data, components,  p_min,  p_max, format, flags);
// }
// static bool xSliderDoubleN(const char* label, double* p_data, int components, double* p_min, double* p_max, const char* format, ImGuiSliderFlags flags) {
// 		return igSliderScalarN(label, ImGuiDataType_Double, p_data, components,  p_min,  p_max, format, flags);
// }
// static bool xVSliderInt32(const char* label, const ImVec2 size, int* p_data, int* p_min, int* p_max, const char* format, ImGuiSliderFlags flags) {
//		return igVSliderScalar(label, size, ImGuiDataType_S32, p_data, p_min, p_max, format, flags);
// }
// static bool xVSliderInt64(const char* label, const ImVec2 size, int64_t* p_data, int64_t* p_min, int64_t* p_max, const char* format, ImGuiSliderFlags flags) {
//		return igVSliderScalar(label, size, ImGuiDataType_S64, p_data, p_min, p_max, format, flags);
// }
// static bool xVSliderFloat(const char* label, const ImVec2 size, float* p_data, float* p_min, float* p_max, const char* format, ImGuiSliderFlags flags) {
//		return igVSliderScalar(label, size, ImGuiDataType_Float, p_data, p_min, p_max, format, flags);
// }
// static bool xVSliderDouble(const char* label, const ImVec2 size, double* p_data, double* p_min, double* p_max, const char* format, ImGuiSliderFlags flags) {
//		return igVSliderScalar(label, size, ImGuiDataType_Double, p_data, p_min, p_max, format, flags);
// }
import "C"
import "unsafe"

func SliderFloat(label string, v *float32, vMin, vMax float32, format string, flags SliderFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igSliderFloat(clabel, (*C.float)(v), C.float(vMin), C.float(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

func SliderFloat2(label string, v *[2]float32, vMin, vMax float32, format string, flags SliderFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igSliderFloat2(clabel, (*C.float)(&v[0]), C.float(vMin), C.float(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

func SliderFloat3(label string, v *[3]float32, vMin, vMax float32, format string, flags SliderFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igSliderFloat3(clabel, (*C.float)(&v[0]), C.float(vMin), C.float(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

func SliderFloat4(label string, v *[4]float32, vMin, vMax float32, format string, flags SliderFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igSliderFloat4(clabel, (*C.float)(&v[0]), C.float(vMin), C.float(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

func SliderAngle(label string, vRad *float32, vDegreesMin, vDegreesMax float32, format string, flags SliderFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igSliderAngle(clabel, (*C.float)(vRad), C.float(vDegreesMin), C.float(vDegreesMax), cformat, C.ImGuiSliderFlags(flags)))
}

func SliderAngleCS(label CString, vRad *float32, vDegreesMin, vDegreesMax float32, format CString, flags SliderFlags) bool {

	return bool(C.igSliderAngle(label.CString(), (*C.float)(vRad), C.float(vDegreesMin), C.float(vDegreesMax),
		format.CString(), C.ImGuiSliderFlags(flags)))
}

func SliderInt(label string, v *int32, vMin, vMax int, format string, flags SliderFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igSliderInt(clabel, (*C.int)(v), C.int(vMin), C.int(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

func SliderIntCS(label CString, v *int32, vMin, vMax int, format CString, flags SliderFlags) bool {

	return bool(C.igSliderInt(label.CString(), (*C.int)(v), C.int(vMin), C.int(vMax), format.CString(), C.ImGuiSliderFlags(flags)))
}

func SliderInt2(label string, v []int32, vMin, vMax int, format string, flags SliderFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igSliderInt2(clabel, (*C.int)(&v[0]), C.int(vMin), C.int(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

func SliderInt3(label string, v []int32, vMin, vMax int, format string, flags SliderFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igSliderInt3(clabel, (*C.int)(&v[0]), C.int(vMin), C.int(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

func SliderInt4(label string, v []int32, vMin, vMax int, format string, flags SliderFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igSliderInt4(clabel, (*C.int)(&v[0]), C.int(vMin), C.int(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

func SliderInt32N(label string, data []int32, pMin, pMax *int32, format string, flags SliderFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.xSliderInt32N(clabel, (*C.int)(&data[0]), C.int(len(data)), (*C.int)(pMin), (*C.int)(pMax),
		cformat, C.ImGuiSliderFlags(flags)))
}

func SliderInt32NCS(label CString, data []int32, pMin, pMax *int32, format CString, flags SliderFlags) bool {

	return bool(C.xSliderInt32N(label.CString(), (*C.int)(&data[0]), C.int(len(data)), (*C.int)(pMin), (*C.int)(pMax),
		format.CString(), C.ImGuiSliderFlags(flags)))
}

func SliderInt64N(label string, data []int64, pMin, pMax *int64, format string, flags SliderFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.xSliderInt64N(clabel, (*C.int64_t)(&data[0]), C.int(len(data)), (*C.int64_t)(pMin), (*C.int64_t)(pMax),
		cformat, C.ImGuiSliderFlags(flags)))
}

func SliderInt64NCS(label CString, data []int64, pMin, pMax *int64, format CString, flags SliderFlags) bool {

	return bool(C.xSliderInt64N(label.CString(), (*C.int64_t)(&data[0]), C.int(len(data)), (*C.int64_t)(pMin), (*C.int64_t)(pMax),
		format.CString(), C.ImGuiSliderFlags(flags)))
}

func SliderFloat32N(label string, data []float32, pMin, pMax *float32, format string, flags SliderFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.xSliderFloatN(clabel, (*C.float)(&data[0]), C.int(len(data)), (*C.float)(pMin), (*C.float)(pMax),
		cformat, C.ImGuiSliderFlags(flags)))
}

func SliderFloat32NCS(label CString, data []float32, pMin, pMax *float32, format CString, flags SliderFlags) bool {

	return bool(C.xSliderFloatN(label.CString(), (*C.float)(&data[0]), C.int(len(data)), (*C.float)(pMin), (*C.float)(pMax),
		format.CString(), C.ImGuiSliderFlags(flags)))
}

func SliderFloat64N(label string, data []float64, pMin, pMax *float64, format string, flags SliderFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.xSliderDoubleN(clabel, (*C.double)(&data[0]), C.int(len(data)), (*C.double)(pMin), (*C.double)(pMax), cformat, C.ImGuiSliderFlags(flags)))
}

func SliderFloat64NCS(label CString, data []float64, pMin, pMax *float64, format CString, flags SliderFlags) bool {

	return bool(C.xSliderDoubleN(label.CString(), (*C.double)(&data[0]), C.int(len(data)), (*C.double)(pMin), (*C.double)(pMax),
		format.CString(), C.ImGuiSliderFlags(flags)))
}

func VSliderFloat(label string, size Vec2, v *float32, vMin, vMax float32, format string, flags SliderFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igVSliderFloat(clabel, size.ImVec2(), (*C.float)(v), C.float(vMin), C.float(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

func VSliderInt(label string, size Vec2, v *int32, vMin, vMax int, format string, flags SliderFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igVSliderInt(clabel, size.ImVec2(), (*C.int)(v), C.int(vMin), C.int(vMax), cformat, C.ImGuiSliderFlags(flags)))
}
func VSliderInt32CS(label CString, size Vec2, pv, pMin, pMax *int32, format CString, flags SliderFlags) bool {

	return bool(C.xVSliderInt32(label.CString(), size.ImVec2(), (*C.int)(pv), (*C.int)(pMin), (*C.int)(pMax),
		format.CString(), C.ImGuiSliderFlags(flags)))
}

func VSliderInt64CS(label CString, size Vec2, pv, pMin, pMax *int64, format CString, flags SliderFlags) bool {

	return bool(C.xVSliderInt64(label.CString(), size.ImVec2(), (*C.int64_t)(pv), (*C.int64_t)(pMin), (*C.int64_t)(pMax),
		format.CString(), C.ImGuiSliderFlags(flags)))
}

func VSliderFloat32CS(label CString, size Vec2, pv, pMin, pMax *float32, format CString, flags SliderFlags) bool {

	return bool(C.xVSliderFloat(label.CString(), size.ImVec2(), (*C.float)(pv), (*C.float)(pMin), (*C.float)(pMax),
		format.CString(), C.ImGuiSliderFlags(flags)))
}

func VSliderFloat64CS(label CString, size Vec2, pv, pMin, pMax *float64, format CString, flags SliderFlags) bool {

	return bool(C.xVSliderDouble(label.CString(), size.ImVec2(), (*C.double)(pv), (*C.double)(pMin), (*C.double)(pMax),
		format.CString(), C.ImGuiSliderFlags(flags)))
}
