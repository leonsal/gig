package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
// static bool xDragInt32N(const char* label, int32_t* pData, int components, float speed, int32_t* pMin, int32_t* pMax, const char* format, ImGuiSliderFlags flags) {
//		return igDragScalarN(label, ImGuiDataType_S32, pData, components, speed, pMin, pMax, format, flags);
// }
// static bool xDragInt64N(const char* label, int64_t* pData, int components, float speed, int64_t* pMin, int64_t* pMax, const char* format, ImGuiSliderFlags flags) {
//		return igDragScalarN(label, ImGuiDataType_S64, pData, components, speed, pMin, pMax, format, flags);
// }
// static bool xDragFloatN(const char* label, float* pData, int components, float speed, float* pMin, float* pMax, const char* format, ImGuiSliderFlags flags) {
//		return igDragScalarN(label, ImGuiDataType_Float, pData, components, speed, pMin, pMax, format, flags);
// }
// static bool xDragDoubleN(const char* label, double* pData, int components, float speed, double* pMin, double* pMax, const char* format, ImGuiSliderFlags flags) {
//		return igDragScalarN(label, ImGuiDataType_Double, pData, components, speed, pMin, pMax, format, flags);
// }
import "C"
import "unsafe"

func DragFloat(label string, v *float32, vSpeed, vMin, vMax float32, format string, flags SliderFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igDragFloat(clabel, (*C.float)(v), C.float(vSpeed), C.float(vMin), C.float(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

func DragFloatCS(label CString, v *float32, vSpeed, vMin, vMax float32, format CString, flags SliderFlags) bool {

	return bool(C.igDragFloat(label.CString(), (*C.float)(v), C.float(vSpeed), C.float(vMin), C.float(vMax),
		format.CString(), C.ImGuiSliderFlags(flags)))
}

func DragFloat2(label string, v []float32, vSpeed, vMin, vMax float32, format string, flags SliderFlags) bool {

	if len(v) < 2 {
		panic("DragFloat2 slice length must be >= 2")
	}
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igDragFloat2(clabel, (*C.float)(&v[0]), C.float(vSpeed), C.float(vMin), C.float(vMax),
		cformat, C.ImGuiSliderFlags(flags)))
}

func DragFloat2CS(label CString, v []float32, vSpeed, vMin, vMax float32, format CString, flags SliderFlags) bool {

	return bool(C.igDragFloat2(label.CString(), (*C.float)(&v[0]), C.float(vSpeed), C.float(vMin), C.float(vMax),
		format.CString(), C.ImGuiSliderFlags(flags)))
}

func DragFloat3(label string, v []float32, vSpeed, vMin, vMax float32, format string, flags SliderFlags) bool {

	if len(v) < 3 {
		panic("DragFloat3 slice length must be >= 3")
	}
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igDragFloat3(clabel, (*C.float)(&v[0]), C.float(vSpeed), C.float(vMin), C.float(vMax),
		cformat, C.ImGuiSliderFlags(flags)))
}

func DragFloat3CS(label CString, v []float32, vSpeed, vMin, vMax float32, format CString, flags SliderFlags) bool {

	return bool(C.igDragFloat3(label.CString(), (*C.float)(&v[0]), C.float(vSpeed), C.float(vMin), C.float(vMax),
		format.CString(), C.ImGuiSliderFlags(flags)))
}

func DragFloat4(label string, v []float32, vSpeed, vMin, vMax float32, format string, flags SliderFlags) bool {

	if len(v) < 4 {
		panic("DragFloat4 slice length must be >= 4")
	}
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igDragFloat4(clabel, (*C.float)(&v[0]), C.float(vSpeed), C.float(vMin), C.float(vMax),
		cformat, C.ImGuiSliderFlags(flags)))
}

func DragFloat4CS(label CString, v []float32, vSpeed, vMin, vMax float32, format CString, flags SliderFlags) bool {

	return bool(C.igDragFloat4(label.CString(), (*C.float)(&v[0]), C.float(vSpeed), C.float(vMin), C.float(vMax),
		format.CString(), C.ImGuiSliderFlags(flags)))
}

func DragFloat32Range2(label string, vCurrentMin, vCurrentMax *float32, vSpeed, vMin, vMax float32, format, formatMax string, flags SliderFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	format_max := C.CString(formatMax)
	defer C.free(unsafe.Pointer(format_max))
	return bool(C.igDragFloatRange2(clabel, (*C.float)(vCurrentMin), (*C.float)(vCurrentMax), C.float(vSpeed),
		C.float(vMin), C.float(vMax), cformat, format_max, C.ImGuiSliderFlags(flags)))
}

func DragFloatRange2CS(label CString, vCurrentMin, vCurrentMax *float32, vSpeed, vMin, vMax float32, format, formatMax CString, flags SliderFlags) bool {

	return bool(C.igDragFloatRange2(label.CString(), (*C.float)(vCurrentMin), (*C.float)(vCurrentMax), C.float(vSpeed),
		C.float(vMin), C.float(vMax), format.CString(), formatMax.CString(), C.ImGuiSliderFlags(flags)))
}

func DragInt(label string, v *int32, vSpeed float32, vMin, vMax int32, format string, flags SliderFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igDragInt(clabel, (*C.int)(v), C.float(vSpeed), C.int(vMin), C.int(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

func DragIntCS(label CString, v *int32, vSpeed float32, vMin, vMax int32, format CString, flags SliderFlags) bool {

	return bool(C.igDragInt(label.CString(), (*C.int)(v), C.float(vSpeed), C.int(vMin), C.int(vMax), format.CString(), C.ImGuiSliderFlags(flags)))
}

func DragInt2(label string, v []int32, vSpeed float32, vMin, vMax int32, format string, flags SliderFlags) bool {

	if len(v) < 2 {
		panic("DragInt2 slice length must be >= 2")
	}
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igDragInt2(clabel, (*C.int)(&v[0]), C.float(vSpeed), C.int(vMin), C.int(vMax),
		cformat, C.ImGuiSliderFlags(flags)))
}

func DragInt2CS(label CString, v []int32, vSpeed float32, vMin, vMax int32, format CString, flags SliderFlags) bool {

	return bool(C.igDragInt2(label.CString(), (*C.int)(&v[0]), C.float(vSpeed), C.int(vMin), C.int(vMax),
		format.CString(), C.ImGuiSliderFlags(flags)))
}

func DragInt3(label string, v []int32, vSpeed float32, vMin, vMax int32, format string, flags SliderFlags) bool {

	if len(v) < 3 {
		panic("DragInt3 slice length must be >= 3")
	}
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igDragInt3(clabel, (*C.int)(&v[0]), C.float(vSpeed), C.int(vMin), C.int(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

func DragInt3CS(label CString, v []int32, vSpeed float32, vMin, vMax int32, format CString, flags SliderFlags) bool {

	return bool(C.igDragInt3(label.CString(), (*C.int)(&v[0]), C.float(vSpeed), C.int(vMin), C.int(vMax), format.CString(), C.ImGuiSliderFlags(flags)))
}

func DragInt4(label string, v []int32, vSpeed float32, vMin, vMax int32, format string, flags SliderFlags) bool {

	if len(v) < 4 {
		panic("DragInt3 slice length must be >= 4")
	}
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igDragInt4(clabel, (*C.int)(&v[0]), C.float(vSpeed), C.int(vMin), C.int(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

func DragInt4CS(label CString, v []int32, vSpeed float32, vMin, vMax int32, format CString, flags SliderFlags) bool {

	return bool(C.igDragInt4(label.CString(), (*C.int)(&v[0]), C.float(vSpeed), C.int(vMin), C.int(vMax),
		format.CString(), C.ImGuiSliderFlags(flags)))
}

func DragIntRange2(label string, vCurrentMin, vCurrentMax *int32, vSpeed float32, vMin, vMax int32, format, formatMax string, flags SliderFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	format_max := C.CString(formatMax)
	defer C.free(unsafe.Pointer(format_max))
	return bool(C.igDragIntRange2(clabel, (*C.int)(vCurrentMin), (*C.int)(vCurrentMax), C.float(vSpeed),
		C.int(vMin), C.int(vMax), cformat, format_max, C.ImGuiSliderFlags(flags)))
}

func DragIntRange2CS(label CString, vCurrentMin, vCurrentMax *int32, vSpeed float32, vMin, vMax int32, format, formatMax CString, flags SliderFlags) bool {

	return bool(C.igDragIntRange2(label.CString(), (*C.int)(vCurrentMin), (*C.int)(vCurrentMax), C.float(vSpeed),
		C.int(vMin), C.int(vMax), format.CString(), formatMax.CString(), C.ImGuiSliderFlags(flags)))
}

func DragScalar(label string, dataType DataType, pData unsafe.Pointer, vSpeed float32, pMin, pMax unsafe.Pointer, format string, flags SliderFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igDragScalar(clabel, C.ImGuiDataType(dataType), pData, C.float(vSpeed), pMin, pMax, cformat, C.ImGuiSliderFlags(flags)))
}

func DragScalarCS(label CString, dataType DataType, pData unsafe.Pointer, vSpeed float32, pMin, pMax unsafe.Pointer, format CString, flags SliderFlags) bool {

	return bool(C.igDragScalar(label.CString(), C.ImGuiDataType(dataType), pData, C.float(vSpeed), pMin, pMax, format.CString(), C.ImGuiSliderFlags(flags)))
}

func DragScalarN(label string, dataType DataType, pData unsafe.Pointer, components int, vSpeed float32, pMin, pMax unsafe.Pointer, format string, flags SliderFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igDragScalarN(clabel, C.ImGuiDataType(dataType), pData, C.int(components), C.float(vSpeed),
		pMin, pMax, cformat, C.ImGuiSliderFlags(flags)))
}

func DragScalarNCS(label CString, dataType DataType, pData unsafe.Pointer, components int, vSpeed float32, pMin, pMax unsafe.Pointer, format CString, flags SliderFlags) bool {

	return bool(C.igDragScalarN(label.CString(), C.ImGuiDataType(dataType), pData, C.int(components), C.float(vSpeed),
		pMin, pMax, format.CString(), C.ImGuiSliderFlags(flags)))
}

func DragInt32N(label CString, pData *int32, components int, speed float32, pMin, pMax *int32, format CString, flags SliderFlags) bool {

	return bool(C.xDragInt32N(label.CString(), (*C.int32_t)(pData), C.int(components), C.float(speed),
		(*C.int32_t)(pMin), (*C.int32_t)(pMax), format.CString(), C.ImGuiInputTextFlags(flags)))
}

func DragInt64N(label CString, pData *int64, components int, speed float32, pMin, pMax *int64, format CString, flags SliderFlags) bool {

	return bool(C.xDragInt64N(label.CString(), (*C.int64_t)(pData), C.int(components), C.float(speed),
		(*C.int64_t)(pMin), (*C.int64_t)(pMax), format.CString(), C.ImGuiInputTextFlags(flags)))
}

func DragFloat32N(label CString, pData *float32, components int, speed float32, pMin, pMax *float32, format CString, flags SliderFlags) bool {

	return bool(C.xDragFloatN(label.CString(), (*C.float)(pData), C.int(components), C.float(speed),
		(*C.float)(pMin), (*C.float)(pMax), format.CString(), C.ImGuiInputTextFlags(flags)))
}

func DragFloat64N(label CString, pData *float64, components int, speed float32, pMin, pMax *float64, format CString, flags SliderFlags) bool {

	return bool(C.xDragDoubleN(label.CString(), (*C.double)(pData), C.int(components), C.float(speed),
		(*C.double)(pMin), (*C.double)(pMax), format.CString(), C.ImGuiInputTextFlags(flags)))
}
