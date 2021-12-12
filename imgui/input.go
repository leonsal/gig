package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
// static bool xInputInt32N(const char* label, int32_t* pData, int components, int32_t* pStep, int32_t* pStepFast, const char* format, ImGuiInputTextFlags flags) {
//  	return igInputScalarN(label, ImGuiDataType_S32, pData, components, pStep, pStepFast, format, flags);
// }
// static bool xInputInt64N(const char* label, int64_t* pData, int components, int64_t* pStep, int64_t* pStepFast, const char* format, ImGuiInputTextFlags flags) {
//  	return igInputScalarN(label, ImGuiDataType_S64, pData, components, pStep, pStepFast, format, flags);
// }
// static bool xInputFloatN(const char* label, float* pData, int components, float* pStep, float* pStepFast, const char* format, ImGuiInputTextFlags flags) {
//  	return igInputScalarN(label, ImGuiDataType_Float, pData, components, pStep, pStepFast, format, flags);
// }
// static bool xInputDoubleN(const char* label, double* pData, int components, double* pStep, double* pStepFast, const char* format, ImGuiInputTextFlags flags) {
//  	return igInputScalarN(label, ImGuiDataType_Double, pData, components, pStep, pStepFast, format, flags);
// }
// extern int goInputTextCallback(ImGuiInputTextCallbackData* t);
// extern int goInputTextCallback2(ImGuiInputTextCallbackData* t);
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

func InputText(label string, buf []byte, flags InputTextFlags, cb InputTextCallback, userData interface{}) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	var handle cgo.Handle
	if cb != nil {
		info := InputTextCallbackInfo{
			Cb: cb,
			Ud: userData,
		}
		handle = cgo.NewHandle(info)
	}
	return bool(C.igInputText(clabel, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)), C.ImGuiInputTextFlags(flags),
		C.ImGuiInputTextCallback(C.goInputTextCallback), unsafe.Pointer(handle)))
}

func InputTextCS(label CString, buf []byte, flags InputTextFlags, cbh cgo.Handle) bool {

	return bool(C.igInputText(label.CString(), (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)), C.ImGuiInputTextFlags(flags),
		C.ImGuiInputTextCallback(C.goInputTextCallback2), unsafe.Pointer(cbh)))
}

func InputTextMultiline(label string, buf []byte, size Vec2, flags InputTextFlags, cb InputTextCallback, userData interface{}) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	var handle cgo.Handle
	if cb != nil {
		info := InputTextCallbackInfo{
			Cb: cb,
			Ud: userData,
		}
		handle = cgo.NewHandle(info)
	}
	return bool(C.igInputTextMultiline(clabel, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)), size.ImVec2(),
		C.ImGuiInputTextFlags(flags), C.ImGuiInputTextCallback(C.goInputTextCallback), unsafe.Pointer(handle)))
}

func InputTextMultilineCS(label CString, buf []byte, size Vec2, flags InputTextFlags, cbh cgo.Handle) bool {

	return bool(C.igInputTextMultiline(label.CString(), (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)), size.ImVec2(),
		C.ImGuiInputTextFlags(flags), C.ImGuiInputTextCallback(C.goInputTextCallback2), unsafe.Pointer(cbh)))
}

func InputTextWithHint(label string, hint string, buf []byte, flags InputTextFlags, cb InputTextCallback, userData interface{}) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	chint := C.CString(hint)
	defer C.free(unsafe.Pointer(chint))

	var handle cgo.Handle
	if cb != nil {
		info := InputTextCallbackInfo{
			Cb: cb,
			Ud: userData,
		}
		handle = cgo.NewHandle(info)
	}
	return bool(C.igInputTextWithHint(clabel, chint, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)),
		C.ImGuiInputTextFlags(flags), C.ImGuiInputTextCallback(C.goInputTextCallback), unsafe.Pointer(handle)))
}

func InputTextWithHintCS(label CString, hint CString, buf []byte, flags InputTextFlags, cbh cgo.Handle) bool {

	return bool(C.igInputTextWithHint(label.CString(), hint.CString(), (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)),
		C.ImGuiInputTextFlags(flags), C.ImGuiInputTextCallback(C.goInputTextCallback2), unsafe.Pointer(cbh)))
}

func InputFloat(label string, v *float32, step, stepFast float32, format string, flags InputTextFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igInputFloat(clabel, (*C.float)(v), C.float(step), C.float(stepFast), cformat, C.ImGuiInputTextFlags(flags)))
}

func InputFloatCS(label CString, v *float32, step, stepFast float32, format CString, flags InputTextFlags) bool {

	return bool(C.igInputFloat(label.CString(), (*C.float)(v), C.float(step), C.float(stepFast), format.CString(), C.ImGuiInputTextFlags(flags)))
}

func InputFloat2(label string, v []float32, format string, flags InputTextFlags) bool {

	if len(v) < 2 {
		panic("InputFloat2 slice length must be >= 2")
	}
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igInputFloat2(clabel, (*C.float)(&v[0]), cformat, C.ImGuiInputTextFlags(flags)))
}

func InputFloat2CS(label CString, v []float32, format CString, flags InputTextFlags) bool {

	return bool(C.igInputFloat2(label.CString(), (*C.float)(&v[0]), format.CString(), C.ImGuiInputTextFlags(flags)))
}

func InputFloat3(label string, v []float32, format string, flags InputTextFlags) bool {

	if len(v) < 3 {
		panic("InputFloat3 slice length must be >= 3")
	}
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igInputFloat3(clabel, (*C.float)(&v[0]), cformat, C.ImGuiInputTextFlags(flags)))
}

func InputFloat3CS(label CString, v []float32, format CString, flags InputTextFlags) bool {

	return bool(C.igInputFloat3(label.CString(), (*C.float)(&v[0]), format.CString(), C.ImGuiInputTextFlags(flags)))
}

func InputFloat4(label string, v *[4]float32, format string, flags InputTextFlags) bool {

	if len(v) < 4 {
		panic("InputFloat4 slice length must be >= 4")
	}
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igInputFloat4(clabel, (*C.float)(&v[0]), cformat, C.ImGuiInputTextFlags(flags)))
}

func InputFloat4CS(label CString, v []float32, format CString, flags InputTextFlags) bool {

	return bool(C.igInputFloat4(label.CString(), (*C.float)(&v[0]), format.CString(), C.ImGuiInputTextFlags(flags)))
}

func InputInt(label string, v *int32, step, stepFast int, flags InputTextFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igInputInt(clabel, (*C.int)(v), C.int(step), C.int(stepFast), C.ImGuiInputTextFlags(flags)))
}

func InputIntCS(label CString, v *int32, step, stepFast int, flags InputTextFlags) bool {

	return bool(C.igInputInt(label.CString(), (*C.int)(v), C.int(step), C.int(stepFast), C.ImGuiInputTextFlags(flags)))
}

func InputInt2(label string, v []int32, flags InputTextFlags) bool {

	if len(v) < 4 {
		panic("InputInt2 slice length must be >= 2")
	}
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igInputInt2(clabel, (*C.int)(&v[0]), C.ImGuiInputTextFlags(flags)))
}

func InputInt2CS(label CString, v []int32, flags InputTextFlags) bool {

	return bool(C.igInputInt2(label.CString(), (*C.int)(&v[0]), C.ImGuiInputTextFlags(flags)))
}

func InputInt3(label string, v []int32, flags InputTextFlags) bool {

	if len(v) < 3 {
		panic("InputInt3 slice length must be >= 3")
	}
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igInputInt3(clabel, (*C.int)(&v[0]), C.ImGuiInputTextFlags(flags)))
}

func InputInt3CS(label CString, v []int32, flags InputTextFlags) bool {

	return bool(C.igInputInt3(label.CString(), (*C.int)(&v[0]), C.ImGuiInputTextFlags(flags)))
}

func InputInt4(label string, v []int32, flags InputTextFlags) bool {

	if len(v) < 4 {
		panic("InputInt4 slice length must be >= 4")
	}
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igInputInt4(clabel, (*C.int)(&v[0]), C.ImGuiInputTextFlags(flags)))
}

func InputInt4CS(label CString, v []int32, flags InputTextFlags) bool {

	return bool(C.igInputInt4(label.CString(), (*C.int)(&v[0]), C.ImGuiInputTextFlags(flags)))
}

func InputDouble(label string, v *float64, step, stepFast float64, format string, flags InputTextFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igInputDouble(clabel, (*C.double)(v), C.double(step), C.double(stepFast), cformat, C.ImGuiInputTextFlags(flags)))
}

func InputDoubleCS(label CString, v *float64, step, stepFast float64, format CString, flags InputTextFlags) bool {

	return bool(C.igInputDouble(label.CString(), (*C.double)(v), C.double(step), C.double(stepFast),
		format.CString(), C.ImGuiInputTextFlags(flags)))
}

func InputScalar(label string, dataType DataType, pData unsafe.Pointer, pStep, pStepFast unsafe.Pointer, format string, flags InputTextFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igInputScalar(clabel, C.ImGuiDataType(dataType), pData, pStep, pStepFast, cformat, C.ImGuiInputTextFlags(flags)))
}

func InputScalarCS(label CString, dataType DataType, pData unsafe.Pointer, pStep, pStepFast unsafe.Pointer, format CString, flags InputTextFlags) bool {

	return bool(C.igInputScalar(label.CString(), C.ImGuiDataType(dataType), pData, pStep, pStepFast, format.CString(), C.ImGuiInputTextFlags(flags)))
}

func InputScalarN(label string, dataType DataType, pData unsafe.Pointer, components int, pStep, pStepFast unsafe.Pointer,
	format string, flags InputTextFlags) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igInputScalarN(clabel, C.ImGuiDataType(dataType), pData, C.int(components), pStep, pStepFast,
		cformat, C.ImGuiInputTextFlags(flags)))
}

func InputScalarNCS(label CString, dataType DataType, pData unsafe.Pointer, components int, pStep, pStepFast unsafe.Pointer,
	format CString, flags InputTextFlags) bool {

	return bool(C.igInputScalarN(label.CString(), C.ImGuiDataType(dataType), pData, C.int(components), pStep, pStepFast,
		format.CString(), C.ImGuiInputTextFlags(flags)))
}

func InputInt32N(label CString, pData *int32, components int, pStep, pStepFast *int32, format CString, flags InputTextFlags) bool {

	return bool(C.xInputInt32N(label.CString(), (*C.int32_t)(pData), C.int(components), (*C.int32_t)(pStep), (*C.int32_t)(pStepFast),
		format.CString(), C.ImGuiInputTextFlags(flags)))
}

func InputInt64N(label CString, pData *int64, components int, pStep, pStepFast *int64, format CString, flags InputTextFlags) bool {

	return bool(C.xInputInt64N(label.CString(), (*C.int64_t)(pData), C.int(components), (*C.int64_t)(pStep), (*C.int64_t)(pStepFast),
		format.CString(), C.ImGuiInputTextFlags(flags)))
}

func InputFloat32N(label CString, pData *float32, components int, pStep, pStepFast *float32, format CString, flags InputTextFlags) bool {

	return bool(C.xInputFloatN(label.CString(), (*C.float)(pData), C.int(components), (*C.float)(pStep), (*C.float)(pStepFast),
		format.CString(), C.ImGuiInputTextFlags(flags)))
}

func InputFloat64N(label CString, pData *float64, components int, pStep, pStepFast *float64, format CString, flags InputTextFlags) bool {

	return bool(C.xInputDoubleN(label.CString(), (*C.double)(pData), C.int(components), (*C.double)(pStep), (*C.double)(pStepFast),
		format.CString(), C.ImGuiInputTextFlags(flags)))
}

// Auxiliary type to allow to pass the final Go callback and user data.
type InputTextCallbackInfo struct {
	Cb InputTextCallback
	Ud interface{}
}

// goInputTextCallback is called from C.igInputText when a callback was specified
//export goInputTextCallback
func goInputTextCallback(data *C.ImGuiInputTextCallbackData) C.int {

	// Get the Go function callback from the UserData field
	h := cgo.Handle(data.UserData)
	info := h.Value().(InputTextCallbackInfo)
	h.Delete()

	// Call user callback
	return C.int(info.Cb(InputTextCallbackData{data}))
}

// goInputTextCallback2 is called from C.igInputTextCS when a callback was specified
//export goInputTextCallback2
func goInputTextCallback2(data *C.ImGuiInputTextCallbackData) C.int {

	// Get the Go function callback from the UserData field
	// Do not delete the handle.
	h := cgo.Handle(data.UserData)
	cb := h.Value().(InputTextCallback)

	// Call user callback
	return C.int(cb(InputTextCallbackData{data}))
}
