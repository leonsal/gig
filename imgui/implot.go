package imgui

// #include <stdlib.h>
// #include <time.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
// #include "cimplot.h"
// static ImVec2 xGetPlotPos() {
//		ImVec2 out;
//		ImPlot_GetPlotPos(&out);
//		return out;
//	}
// static ImVec2 xGetPlotSize() {
//		ImVec2 out;
//		ImPlot_GetPlotSize(&out);
//		return out;
//	}
// static ImPlotPoint xGetPlotMousePos(ImPlotYAxis y_axis) {
//		ImPlotPoint out;
//		ImPlot_GetPlotMousePos(&out, y_axis);
//		return out;
// }
// static ImPlotLimits xGetPlotLimits(ImPlotYAxis y_axis) {
//		ImPlotLimits out;
//		ImPlot_GetPlotLimits(&out, y_axis);
//		return out;
//	}
// static ImVec4 xNextColormapColor() {
//		ImVec4 out;
//		ImPlot_NextColormapColor(&out);
//		return out;
// }
// static ImVec4 xGetLastItemColor() {
//		ImVec4 out;
//		ImPlot_GetLastItemColor(&out);
//		return out;
// }
// static ImVec4 xGetColormapColor(int idx, ImPlotColormap cmap) {
//		ImVec4 out;
//		ImPlot_GetColormapColor(&out, idx, cmap);
//		return out;
// }
// static ImVec4 xSampleColormap(float t, ImPlotColormap cmap){
//		ImVec4 out;
//		ImPlot_SampleColormap(&out, t, cmap);
//		return out;
// }
import "C"
import (
	"unsafe"
)

const (
	sizeofI32 = int(unsafe.Sizeof(int32(0)))
	sizeofF32 = int(unsafe.Sizeof(float32(0)))
	sizeofF64 = int(unsafe.Sizeof(float64(0)))
)

func BeginPlot(titleId, xLabel, yLabel string, size Vec2, flags PlotFlags, xFlags, yFlags, y2Flags, y3Flags PlotAxisFlags,
	y2Label, y3Label string) bool {

	title_id := C.CString(titleId)
	defer C.free(unsafe.Pointer(title_id))
	x_label := C.CString(xLabel)
	defer C.free(unsafe.Pointer(x_label))
	y_label := C.CString(yLabel)
	defer C.free(unsafe.Pointer(y_label))
	y2_label := C.CString(y2Label)
	defer C.free(unsafe.Pointer(y2_label))
	y3_label := C.CString(y2Label)
	defer C.free(unsafe.Pointer(y3_label))
	return bool(C.ImPlot_BeginPlot(title_id, x_label, y_label, size.ImVec2(), C.ImPlotFlags(flags), C.ImPlotAxisFlags(xFlags),
		C.ImPlotAxisFlags(yFlags), C.ImPlotAxisFlags(y2Flags), C.ImPlotAxisFlags(y3Flags), y2_label, y3_label))
}

func BeginPlotCS(titleId, xLabel, yLabel CString, size Vec2, flags PlotFlags, xFlags, yFlags, y2Flags, y3Flags PlotAxisFlags, y2Label, y3Label CString) bool {

	return bool(C.ImPlot_BeginPlot(titleId.CString(), xLabel.CString(), yLabel.CString(), size.ImVec2(), C.ImPlotFlags(flags),
		C.ImPlotAxisFlags(xFlags), C.ImPlotAxisFlags(yFlags), C.ImPlotAxisFlags(y2Flags), C.ImPlotAxisFlags(y3Flags),
		y2Label.CString(), y3Label.CString()))
}

func EndPlot() {

	C.ImPlot_EndPlot()
}

func BeginSubplots(titleId string, rows, cols int, size Vec2, flags PlotSubplotFlags, rowRatios, colRatios *float32) bool {

	title_id := C.CString(titleId)
	defer C.free(unsafe.Pointer(title_id))
	return bool(C.ImPlot_BeginSubplots(title_id, C.int(rows), C.int(cols), size.ImVec2(), C.ImPlotSubplotFlags(flags),
		(*C.float)(rowRatios), (*C.float)(colRatios)))
}

func BeginSubplotsCS(titleId CString, rows, cols int, size Vec2, flags PlotSubplotFlags, rowRatios, colRatios *float32) bool {

	return bool(C.ImPlot_BeginSubplots(titleId.CString(), C.int(rows), C.int(cols), size.ImVec2(), C.ImPlotSubplotFlags(flags),
		(*C.float)(rowRatios), (*C.float)(colRatios)))
}

func EndSubplots() {

	C.ImPlot_EndSubplots()
}

func PlotLineInt32(labelId string, values *int32, count int, xscale, x0 float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotLine_S32PtrInt(label_id, (*C.int)(values), C.int(count), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotLineInt32CS(labelId CString, values *int32, count int, xscale, x0 float64, offset, stride int) {

	C.ImPlot_PlotLine_S32PtrInt(labelId.CString(), (*C.int)(values), C.int(count), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotLineFloat32(labelId string, values *float32, count int, xscale, x0 float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotLine_FloatPtrInt(label_id, (*C.float)(values), C.int(count), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotLineFloat32CS(labelId CString, values *float32, count int, xscale, x0 float64, offset, stride int) {

	C.ImPlot_PlotLine_FloatPtrInt(labelId.CString(), (*C.float)(values), C.int(count), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotLineFloat64(labelId string, values *float64, count int, xscale, x0 float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotLine_doublePtrInt(label_id, (*C.double)(values), C.int(count), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotLineFloat64CS(labelId CString, values *float64, count int, xscale, x0 float64, offset, stride int) {

	C.ImPlot_PlotLine_doublePtrInt(labelId.CString(), (*C.double)(values), C.int(count), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotLineXYInt32(labelId string, xs, ys *int32, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotLine_S32PtrS32Ptr(label_id, (*C.ImS32)(xs), (*C.ImS32)(ys), C.int(count),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotLineXYInt32CS(labelId CString, xs, ys *int32, count, offset, stride int) {

	C.ImPlot_PlotLine_S32PtrS32Ptr(labelId.CString(), (*C.ImS32)(xs), (*C.ImS32)(ys), C.int(count),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotLineXYFloat32(labelId string, xs, ys *float32, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotLine_FloatPtrFloatPtr(label_id, (*C.float)(xs), (*C.float)(ys), C.int(count),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotLineXYFloat32CS(labelId CString, xs, ys *float32, count, offset, stride int) {

	C.ImPlot_PlotLine_FloatPtrFloatPtr(labelId.CString(), (*C.float)(xs), (*C.float)(ys), C.int(count),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotLineXYFloat64(labelId string, xs, ys *float64, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotLine_doublePtrdoublePtr(label_id, (*C.double)(xs), (*C.double)(ys), C.int(count),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotLineXYFloat64CS(labelId CString, xs, ys *float64, count, offset, stride int) {

	C.ImPlot_PlotLine_doublePtrdoublePtr(labelId.CString(), (*C.double)(xs), (*C.double)(ys), C.int(count),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotScatterInt32(labelId string, values *int32, count int, xscale, x0 float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotScatter_S32PtrInt(label_id, (*C.ImS32)(values), C.int(count), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotScatterInt32CS(labelId CString, values *int32, count int, xscale, x0 float64, offset, stride int) {

	C.ImPlot_PlotScatter_S32PtrInt(labelId.CString(), (*C.ImS32)(values), C.int(count), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotScatterFloat32(labelId string, values *float32, count int, xscale, x0 float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotScatter_FloatPtrInt(label_id, (*C.float)(values), C.int(count), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotScatterFloat32CS(labelId CString, values *float32, count int, xscale, x0 float64, offset, stride int) {

	C.ImPlot_PlotScatter_FloatPtrInt(labelId.CString(), (*C.float)(values), C.int(count), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotScatterFloat64(labelId string, values *float64, count int, xscale, x0 float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotScatter_doublePtrInt(label_id, (*C.double)(values), C.int(count), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotScatterFloat64CS(labelId CString, values *float64, count int, xscale, x0 float64, offset, stride int) {

	C.ImPlot_PlotScatter_doublePtrInt(labelId.CString(), (*C.double)(values), C.int(count), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotScatterXYInt32(labelId string, xs, ys *int32, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotScatter_S32PtrS32Ptr(label_id, (*C.ImS32)(xs), (*C.ImS32)(ys), C.int(count),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotScatterXYInt32CS(labelId CString, xs, ys *int32, count, offset, stride int) {

	C.ImPlot_PlotScatter_S32PtrS32Ptr(labelId.CString(), (*C.ImS32)(xs), (*C.ImS32)(ys), C.int(count),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotScatterXYFloat32(labelId string, xs, ys *float32, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotScatter_FloatPtrFloatPtr(label_id, (*C.float)(xs), (*C.float)(ys), C.int(count),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotScatterXYFloat32CS(labelId CString, xs, ys *float32, count, offset, stride int) {

	C.ImPlot_PlotScatter_FloatPtrFloatPtr(labelId.CString(), (*C.float)(xs), (*C.float)(ys), C.int(count),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotScatterXYFloat64(labelId string, xs, ys *float64, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotScatter_doublePtrdoublePtr(label_id, (*C.double)(xs), (*C.double)(ys), C.int(count),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotScatterXYFloat64CS(labelId CString, xs, ys *float64, count, offset, stride int) {

	C.ImPlot_PlotScatter_doublePtrdoublePtr(labelId.CString(), (*C.double)(xs), (*C.double)(ys), C.int(count),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotStairsInt32(labelId string, values *int32, count int, xscale, x0 float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotStairs_S32PtrInt(label_id, (*C.ImS32)(values), C.int(count), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotStairsInt32CS(labelId CString, values *int32, count int, xscale, x0 float64, offset, stride int) {

	C.ImPlot_PlotStairs_S32PtrInt(labelId.CString(), (*C.ImS32)(values), C.int(count), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotStairsFloat32(labelId string, values *float32, count int, xscale, x0 float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotStairs_FloatPtrInt(label_id, (*C.float)(values), C.int(count), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotStairsFloat32CS(labelId CString, values *float32, count int, xscale, x0 float64, offset, stride int) {

	C.ImPlot_PlotStairs_FloatPtrInt(labelId.CString(), (*C.float)(values), C.int(count), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotStairsFloat64(labelId string, values *float64, count int, xscale, x0 float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotStairs_doublePtrInt(label_id, (*C.double)(values), C.int(count), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotStairsFloat64CS(labelId CString, values *float64, count int, xscale, x0 float64, offset, stride int) {

	C.ImPlot_PlotStairs_doublePtrInt(labelId.CString(), (*C.double)(values), C.int(count), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotStairsXYInt32(labelId string, xs, ys *int32, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotStairs_S32PtrS32Ptr(label_id, (*C.ImS32)(xs), (*C.ImS32)(ys), C.int(count),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotStairsXYInt32CS(labelId CString, xs, ys *int32, count, offset, stride int) {

	C.ImPlot_PlotStairs_S32PtrS32Ptr(labelId.CString(), (*C.ImS32)(xs), (*C.ImS32)(ys), C.int(count), C.int(offset),
		C.int(stride*sizeofI32))
}

func PlotStairsXYFloat32(labelId string, xs, ys *float32, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotStairs_FloatPtrFloatPtr(label_id, (*C.float)(xs), (*C.float)(ys), C.int(count),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotStairsXYFloat32CS(labelId CString, xs, ys *float32, count, offset, stride int) {

	C.ImPlot_PlotStairs_FloatPtrFloatPtr(labelId.CString(), (*C.float)(xs), (*C.float)(ys), C.int(count),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotStairsXYFloat64(labelId string, xs, ys *float64, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotStairs_doublePtrdoublePtr(label_id, (*C.double)(xs), (*C.double)(ys), C.int(count),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotStairsXYFloat64CS(labelId CString, xs, ys *float64, count, offset, stride int) {

	C.ImPlot_PlotStairs_doublePtrdoublePtr(labelId.CString(), (*C.double)(xs), (*C.double)(ys), C.int(count),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotShadedInt32(labelId string, values *int32, count int, yRef, xscale, x0 float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotShaded_S32PtrInt(label_id, (*C.ImS32)(values), C.int(count), C.double(yRef), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotShadedInt32CS(labelId CString, values *int32, count int, yRef, xscale, x0 float64, offset, stride int) {

	C.ImPlot_PlotShaded_S32PtrInt(labelId.CString(), (*C.ImS32)(values), C.int(count), C.double(yRef), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotShadedFloat32(labelId string, values *float32, count int, yRef, xscale, x0 float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotShaded_FloatPtrInt(label_id, (*C.float)(values), C.int(count), C.double(yRef), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotShadedFloat32CS(labelId CString, values *float32, count int, yRef, xscale, x0 float64, offset, stride int) {

	C.ImPlot_PlotShaded_FloatPtrInt(labelId.CString(), (*C.float)(values), C.int(count), C.double(yRef), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotShadedFloat64(labelId string, values *float64, count int, yRef, xscale, x0 float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotShaded_doublePtrInt(label_id, (*C.double)(values), C.int(count), C.double(yRef), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotShadedFloat64CS(labelId CString, values *float64, count int, yRef, xscale, x0 float64, offset, stride int) {

	C.ImPlot_PlotShaded_doublePtrInt(labelId.CString(), (*C.double)(values), C.int(count), C.double(yRef), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotShadedXYInt32(labelId string, xs, ys *int32, count int, yRef float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotShaded_S32PtrS32PtrInt(label_id, (*C.ImS32)(xs), (*C.ImS32)(ys), C.int(count), C.double(yRef),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotShadedXYInt32CS(labelId CString, xs, ys *int32, count int, yRef float64, offset, stride int) {

	C.ImPlot_PlotShaded_S32PtrS32PtrInt(labelId.CString(), (*C.ImS32)(xs), (*C.ImS32)(ys), C.int(count), C.double(yRef),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotShadedXYFloat32(labelId string, xs, ys *float32, count int, yRef float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotShaded_FloatPtrFloatPtrInt(label_id, (*C.float)(xs), (*C.float)(ys), C.int(count), C.double(yRef),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotShadedXYFloat32CS(labelId CString, xs, ys *float32, count int, yRef float64, offset, stride int) {

	C.ImPlot_PlotShaded_FloatPtrFloatPtrInt(labelId.CString(), (*C.float)(xs), (*C.float)(ys), C.int(count), C.double(yRef),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotShadedXYFloat64(labelId string, xs, ys *float64, count int, yRef float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotShaded_doublePtrdoublePtrInt(label_id, (*C.double)(xs), (*C.double)(ys), C.int(count), C.double(yRef),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotShadedXYFloat64CS(labelId CString, xs, ys *float64, count int, yRef float64, offset, stride int) {

	C.ImPlot_PlotShaded_doublePtrdoublePtrInt(labelId.CString(), (*C.double)(xs), (*C.double)(ys), C.int(count), C.double(yRef),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotShadedXYYInt32(labelId string, xs, ys1, ys2 *int32, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotShaded_S32PtrS32PtrS32Ptr(label_id, (*C.ImS32)(xs), (*C.ImS32)(ys1), (*C.ImS32)(ys2),
		C.int(count), C.int(offset), C.int(stride*sizeofI32))
}

func PlotShadedXYYInt32CS(labelId CString, xs, ys1, ys2 *int32, count, offset, stride int) {

	C.ImPlot_PlotShaded_S32PtrS32PtrS32Ptr(labelId.CString(), (*C.ImS32)(xs), (*C.ImS32)(ys1), (*C.ImS32)(ys2),
		C.int(count), C.int(offset), C.int(stride*sizeofI32))
}

func PlotShadedXYYFloat32(labelId string, xs, ys1, ys2 *float32, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotShaded_FloatPtrFloatPtrFloatPtr(label_id, (*C.float)(xs), (*C.float)(ys1), (*C.float)(ys2),
		C.int(count), C.int(offset), C.int(stride*sizeofF32))
}

func PlotShadedXYYFloat32CS(labelId CString, xs, ys1, ys2 *float32, count, offset, stride int) {

	C.ImPlot_PlotShaded_FloatPtrFloatPtrFloatPtr(labelId.CString(), (*C.float)(xs), (*C.float)(ys1), (*C.float)(ys2),
		C.int(count), C.int(offset), C.int(stride*sizeofF32))
}

func PlotShadedXYYFloat64(labelId string, xs, ys1, ys2 *float64, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotShaded_doublePtrdoublePtrdoublePtr(label_id, (*C.double)(xs), (*C.double)(ys1), (*C.double)(ys2),
		C.int(count), C.int(offset), C.int(stride*sizeofF64))
}

func PlotShadedXYYFloat64CS(labelId CString, xs, ys1, ys2 *float64, count, offset, stride int) {

	C.ImPlot_PlotShaded_doublePtrdoublePtrdoublePtr(labelId.CString(), (*C.double)(xs), (*C.double)(ys1), (*C.double)(ys2),
		C.int(count), C.int(offset), C.int(stride*sizeofF64))
}

func PlotBarsInt32(labelId string, values *int32, count int, width, shift float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotBars_S32PtrInt(label_id, (*C.ImS32)(values), C.int(count), C.double(width), C.double(shift),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotBarsInt32CS(labelId CString, values *int32, count int, width, shift float64, offset, stride int) {

	C.ImPlot_PlotBars_S32PtrInt(labelId.CString(), (*C.ImS32)(values), C.int(count), C.double(width), C.double(shift),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotBarsFloat32(labelId string, values *float32, count int, width, shift float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotBars_FloatPtrInt(label_id, (*C.float)(values), C.int(count), C.double(width), C.double(shift),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotBarsFloat32CS(labelId CString, values *float32, count int, width, shift float64, offset, stride int) {

	C.ImPlot_PlotBars_FloatPtrInt(labelId.CString(), (*C.float)(values), C.int(count), C.double(width), C.double(shift),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotBarsFloat64(labelId string, values *float64, count int, width, shift float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotBars_doublePtrInt(label_id, (*C.double)(values), C.int(count), C.double(width), C.double(shift),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotBarsFloat64CS(labelId CString, values *float64, count int, width, shift float64, offset, stride int) {

	C.ImPlot_PlotBars_doublePtrInt(labelId.CString(), (*C.double)(values), C.int(count), C.double(width), C.double(shift),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotBarsXYInt32(labelId string, xs, ys *int32, count int, width float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotBars_S32PtrS32Ptr(label_id, (*C.ImS32)(xs), (*C.ImS32)(ys), C.int(count), C.double(width),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotBarsXYInt32CS(labelId CString, xs, ys *int32, count int, width float64, offset, stride int) {

	C.ImPlot_PlotBars_S32PtrS32Ptr(labelId.CString(), (*C.ImS32)(xs), (*C.ImS32)(ys), C.int(count), C.double(width),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotBarsXYFloat32(labelId string, xs, ys *float32, count int, width float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotBars_FloatPtrFloatPtr(label_id, (*C.float)(xs), (*C.float)(ys), C.int(count), C.double(width),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotBarsXYFloat32CS(labelId CString, xs, ys *float32, count int, width float64, offset, stride int) {

	C.ImPlot_PlotBars_FloatPtrFloatPtr(labelId.CString(), (*C.float)(xs), (*C.float)(ys), C.int(count), C.double(width),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotBarsXYFloat64(labelId string, xs, ys *float64, count int, width float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotBars_doublePtrdoublePtr(label_id, (*C.double)(xs), (*C.double)(ys), C.int(count), C.double(width),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotBarsXYFloat64CS(labelId CString, xs, ys *float64, count int, width float64, offset, stride int) {

	C.ImPlot_PlotBars_doublePtrdoublePtr(labelId.CString(), (*C.double)(xs), (*C.double)(ys), C.int(count), C.double(width),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotBarsHInt32(labelId string, values *int32, count int, height, shift float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotBarsH_S32PtrInt(label_id, (*C.ImS32)(values), C.int(count), C.double(height), C.double(shift),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotBarsHInt32CS(labelId CString, values *int32, count int, height, shift float64, offset, stride int) {

	C.ImPlot_PlotBarsH_S32PtrInt(labelId.CString(), (*C.ImS32)(values), C.int(count), C.double(height), C.double(shift),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotBarsHFloat32(labelId string, values []float32, height, shift float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotBarsH_FloatPtrInt(label_id, (*C.float)(&values[0]), C.int(len(values)), C.double(height), C.double(shift),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotBarsHFloat32CS(labelId CString, values *float32, count int, height, shift float64, offset, stride int) {

	C.ImPlot_PlotBarsH_FloatPtrInt(labelId.CString(), (*C.float)(values), C.int(count), C.double(height), C.double(shift),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotBarsHFloat64(labelId string, values *float64, count int, height, shift float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotBarsH_doublePtrInt(label_id, (*C.double)(values), C.int(count), C.double(height), C.double(shift),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotBarsHFloat64CS(labelId CString, values *float64, count int, height, shift float64, offset, stride int) {

	C.ImPlot_PlotBarsH_doublePtrInt(labelId.CString(), (*C.double)(values), C.int(count), C.double(height), C.double(shift),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotBarsHXYInt32(labelId string, xs, ys *int32, count int, height float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotBarsH_S32PtrS32Ptr(label_id, (*C.ImS32)(xs), (*C.ImS32)(ys), C.int(count), C.double(height),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotBarsHXYInt32CS(labelId CString, xs, ys *int32, count int, height float64, offset, stride int) {

	C.ImPlot_PlotBarsH_S32PtrS32Ptr(labelId.CString(), (*C.ImS32)(xs), (*C.ImS32)(ys), C.int(count), C.double(height),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotBarsHXYFloat32(labelId string, xs, ys *float32, count int, height float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotBarsH_FloatPtrFloatPtr(label_id, (*C.float)(xs), (*C.float)(ys), C.int(count), C.double(height),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotBarsHXYFloat32CS(labelId CString, xs, ys *float32, count int, height float64, offset, stride int) {

	C.ImPlot_PlotBarsH_FloatPtrFloatPtr(labelId.CString(), (*C.float)(xs), (*C.float)(ys), C.int(count), C.double(height),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotBarsHXYFloat64(labelId string, xs, ys *float64, count int, height float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotBarsH_doublePtrdoublePtr(label_id, (*C.double)(xs), (*C.double)(ys), C.int(count), C.double(height),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotBarsHXYFloat64CS(labelId CString, xs, ys *float64, count int, height float64, offset, stride int) {

	C.ImPlot_PlotBarsH_doublePtrdoublePtr(labelId.CString(), (*C.double)(xs), (*C.double)(ys), C.int(count), C.double(height),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotErrorBarsInt32(labelId string, xs, ys, err *int32, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotErrorBars_S32PtrS32PtrS32PtrInt(label_id, (*C.ImS32)(xs), (*C.ImS32)(ys), (*C.ImS32)(err), C.int(count),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotErrorBarsInt32CS(labelId CString, xs, ys, err *int32, count, offset, stride int) {

	C.ImPlot_PlotErrorBars_S32PtrS32PtrS32PtrInt(labelId.CString(), (*C.ImS32)(xs), (*C.ImS32)(ys), (*C.ImS32)(err), C.int(count),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotErrorBarsFloat32(labelId string, xs, ys, err *float32, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotErrorBars_FloatPtrFloatPtrFloatPtrInt(label_id, (*C.float)(xs), (*C.float)(ys), (*C.float)(err), C.int(count),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotErrorBarsFloat32CS(labelId CString, xs, ys, err *float32, count, offset, stride int) {

	C.ImPlot_PlotErrorBars_FloatPtrFloatPtrFloatPtrInt(labelId.CString(), (*C.float)(xs), (*C.float)(ys), (*C.float)(err),
		C.int(count), C.int(offset), C.int(stride*sizeofF32))
}

func PlotErrorBarsFloat64(labelId string, xs, ys, err *float64, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotErrorBars_doublePtrdoublePtrdoublePtrInt(label_id, (*C.double)(xs), (*C.double)(ys), (*C.double)(err),
		C.int(count), C.int(offset), C.int(stride*sizeofF64))
}

func PlotErrorBarsFloat64CS(labelId CString, xs, ys, err *float64, count, offset, stride int) {

	C.ImPlot_PlotErrorBars_doublePtrdoublePtrdoublePtrInt(labelId.CString(), (*C.double)(xs), (*C.double)(ys), (*C.double)(err),
		C.int(count), C.int(offset), C.int(stride*sizeofF64))
}

func PlotErrorBars2Int32(labelId string, xs, ys, neg, pos *int32, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotErrorBars_S32PtrS32PtrS32PtrS32Ptr(label_id, (*C.ImS32)(xs), (*C.ImS32)(ys), (*C.ImS32)(neg), (*C.ImS32)(pos),
		C.int(count), C.int(offset), C.int(stride*sizeofI32))
}

func PlotErrorBars2Int32CS(labelId CString, xs, ys, neg, pos *int32, count, offset, stride int) {

	C.ImPlot_PlotErrorBars_S32PtrS32PtrS32PtrS32Ptr(labelId.CString(), (*C.ImS32)(xs), (*C.ImS32)(ys), (*C.ImS32)(neg),
		(*C.ImS32)(pos), C.int(count), C.int(offset), C.int(stride*sizeofI32))
}

func PlotErrorBars2Float32(labelId string, xs, ys, neg, pos *float32, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotErrorBars_FloatPtrFloatPtrFloatPtrFloatPtr(label_id, (*C.float)(xs), (*C.float)(ys), (*C.float)(neg),
		(*C.float)(pos), C.int(count), C.int(offset), C.int(stride*sizeofF32))
}

func PlotErrorBars2Float32CS(labelId CString, xs, ys, neg, pos *float32, count, offset, stride int) {

	C.ImPlot_PlotErrorBars_FloatPtrFloatPtrFloatPtrFloatPtr(labelId.CString(), (*C.float)(xs), (*C.float)(ys), (*C.float)(neg),
		(*C.float)(pos), C.int(count), C.int(offset), C.int(stride*sizeofF32))
}

func PlotErrorBars2Float64(labelId string, xs, ys, neg, pos *float64, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotErrorBars_doublePtrdoublePtrdoublePtrdoublePtr(label_id, (*C.double)(xs), (*C.double)(ys), (*C.double)(neg),
		(*C.double)(pos), C.int(count), C.int(offset), C.int(stride*sizeofF64))
}

func PlotErrorBars2Float64CS(labelId CString, xs, ys, neg, pos *float64, count, offset, stride int) {

	C.ImPlot_PlotErrorBars_doublePtrdoublePtrdoublePtrdoublePtr(labelId.CString(), (*C.double)(xs), (*C.double)(ys), (*C.double)(neg),
		(*C.double)(pos), C.int(count), C.int(offset), C.int(stride*sizeofF64))
}

func PlotErrorBarsHInt32(labelId string, xs, ys, err *int32, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotErrorBarsH_S32PtrS32PtrS32PtrInt(label_id, (*C.ImS32)(xs), (*C.ImS32)(ys), (*C.ImS32)(err), C.int(count),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotErrorBarsHInt32CS(labelId CString, xs, ys, err *int32, count, offset, stride int) {

	C.ImPlot_PlotErrorBarsH_S32PtrS32PtrS32PtrInt(labelId.CString(), (*C.ImS32)(xs), (*C.ImS32)(ys), (*C.ImS32)(err), C.int(count),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotErrorBarsHFloat32(labelId string, xs, ys, err *float32, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotErrorBarsH_FloatPtrFloatPtrFloatPtrInt(label_id, (*C.float)(xs), (*C.float)(ys), (*C.float)(err),
		C.int(count), C.int(offset), C.int(stride*sizeofF32))
}

func PlotErrorBarsHFloat32CS(labelId CString, xs, ys, err *float32, count, offset, stride int) {

	C.ImPlot_PlotErrorBarsH_FloatPtrFloatPtrFloatPtrInt(labelId.CString(), (*C.float)(xs), (*C.float)(ys), (*C.float)(err),
		C.int(count), C.int(offset), C.int(stride*sizeofF32))
}

func PlotErrorBarsHFloat64(labelId string, xs, ys, err *float64, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotErrorBarsH_doublePtrdoublePtrdoublePtrInt(label_id, (*C.double)(xs), (*C.double)(ys), (*C.double)(err),
		C.int(count), C.int(offset), C.int(stride*sizeofF64))
}

func PlotErrorBarsHFloat64CS(labelId CString, xs, ys, err *float64, count, offset, stride int) {

	C.ImPlot_PlotErrorBarsH_doublePtrdoublePtrdoublePtrInt(labelId.CString(), (*C.double)(xs), (*C.double)(ys), (*C.double)(err),
		C.int(count), C.int(offset), C.int(stride*sizeofF64))
}

func PlotErrorBarsH2Int32(labelId string, xs, ys, neg, pos *int32, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotErrorBarsH_S32PtrS32PtrS32PtrS32Ptr(label_id, (*C.ImS32)(xs), (*C.ImS32)(ys), (*C.ImS32)(neg), (*C.ImS32)(pos),
		C.int(count), C.int(offset), C.int(stride*sizeofI32))
}

func PlotErrorBarsH2Int32CS(labelId CString, xs, ys, neg, pos *int32, count, offset, stride int) {

	C.ImPlot_PlotErrorBarsH_S32PtrS32PtrS32PtrS32Ptr(labelId.CString(), (*C.ImS32)(xs), (*C.ImS32)(ys), (*C.ImS32)(neg), (*C.ImS32)(pos),
		C.int(count), C.int(offset), C.int(stride*sizeofI32))
}

func PlotErrorBarsH2Float32(labelId string, xs, ys, neg, pos *float32, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotErrorBarsH_FloatPtrFloatPtrFloatPtrFloatPtr(label_id, (*C.float)(xs), (*C.float)(ys), (*C.float)(neg), (*C.float)(pos),
		C.int(count), C.int(offset), C.int(stride*sizeofF32))
}

func PlotErrorBarsH2Float32CS(labelId CString, xs, ys, neg, pos *float32, count, offset, stride int) {

	C.ImPlot_PlotErrorBarsH_FloatPtrFloatPtrFloatPtrFloatPtr(labelId.CString(), (*C.float)(xs), (*C.float)(ys), (*C.float)(neg), (*C.float)(pos),
		C.int(count), C.int(offset), C.int(stride*sizeofF32))
}

func PlotErrorBarsH2Float64(labelId string, xs, ys, neg, pos *float64, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotErrorBarsH_doublePtrdoublePtrdoublePtrdoublePtr(label_id, (*C.double)(xs), (*C.double)(ys), (*C.double)(neg), (*C.double)(pos),
		C.int(count), C.int(offset), C.int(stride*sizeofF64))
}

func PlotErrorBarsH2Float64CS(labelId CString, xs, ys, neg, pos *float64, count, offset, stride int) {

	C.ImPlot_PlotErrorBarsH_doublePtrdoublePtrdoublePtrdoublePtr(labelId.CString(), (*C.double)(xs), (*C.double)(ys), (*C.double)(neg), (*C.double)(pos),
		C.int(count), C.int(offset), C.int(stride*sizeofF64))
}

func PlotStemsInt32(labelId string, values *int32, count int, yRef, xscale, x0 float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotStems_S32PtrInt(label_id, (*C.ImS32)(values), C.int(count), C.double(yRef), C.double(xscale), C.double(x0),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotStemsInt32CS(labelId CString, values *int32, count int, yRef, xscale, x0 float64, offset, stride int) {

	C.ImPlot_PlotStems_S32PtrInt(labelId.CString(), (*C.ImS32)(values), C.int(count), C.double(yRef), C.double(xscale),
		C.double(x0), C.int(offset), C.int(stride*sizeofI32))
}

func PlotStemsFloat32(labelId string, values *float32, count int, yRef, xscale, x0 float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotStems_FloatPtrInt(label_id, (*C.float)(values), C.int(count), C.double(yRef), C.double(xscale),
		C.double(x0), C.int(offset), C.int(stride*sizeofF32))
}

func PlotStemsFloat32CS(labelId CString, values *float32, count int, yRef, xscale, x0 float64, offset, stride int) {

	C.ImPlot_PlotStems_FloatPtrInt(labelId.CString(), (*C.float)(values), C.int(count), C.double(yRef), C.double(xscale),
		C.double(x0), C.int(offset), C.int(stride*sizeofF32))
}

func PlotStemsFloat64(labelId string, values *float64, count int, yRef, xscale, x0 float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotStems_doublePtrInt(label_id, (*C.double)(values), C.int(count), C.double(yRef), C.double(xscale),
		C.double(x0), C.int(offset), C.int(stride*sizeofF64))
}

func PlotStemsFloat64CS(labelId CString, values *float64, count int, yRef, xscale, x0 float64, offset, stride int) {

	C.ImPlot_PlotStems_doublePtrInt(labelId.CString(), (*C.double)(values), C.int(count), C.double(yRef), C.double(xscale),
		C.double(x0), C.int(offset), C.int(stride*sizeofF64))
}

func PlotStemsXYInt32(labelId string, xs, ys *int32, count int, yRef float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotStems_S32PtrS32Ptr(label_id, (*C.ImS32)(xs), (*C.ImS32)(ys), C.int(count), C.double(yRef),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotStemsXYInt32CS(labelId CString, xs, ys *int32, count int, yRef float64, offset, stride int) {

	C.ImPlot_PlotStems_S32PtrS32Ptr(labelId.CString(), (*C.ImS32)(xs), (*C.ImS32)(ys), C.int(count), C.double(yRef),
		C.int(offset), C.int(stride*sizeofI32))
}

func PlotStemsXYFloat32(labelId string, xs, ys *float32, count int, yRef float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotStems_FloatPtrFloatPtr(label_id, (*C.float)(xs), (*C.float)(ys), C.int(count), C.double(yRef),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotStemsXYFloat32CS(labelId CString, xs, ys *float32, count int, yRef float64, offset, stride int) {

	C.ImPlot_PlotStems_FloatPtrFloatPtr(labelId.CString(), (*C.float)(xs), (*C.float)(ys), C.int(count), C.double(yRef),
		C.int(offset), C.int(stride*sizeofF32))
}

func PlotStemsXYFloat64(labelId string, xs, ys *float64, count int, yRef float64, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotStems_doublePtrdoublePtr(label_id, (*C.double)(xs), (*C.double)(ys), C.int(count), C.double(yRef),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotStemsXYFloat64CS(labelId CString, xs, ys *float64, count int, yRef float64, offset, stride int) {

	C.ImPlot_PlotStems_doublePtrdoublePtr(labelId.CString(), (*C.double)(xs), (*C.double)(ys), C.int(count), C.double(yRef),
		C.int(offset), C.int(stride*sizeofF64))
}

func PlotVLinesInt32(labelId string, xs *int32, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotVLines_S32Ptr(label_id, (*C.ImS32)(xs), C.int(count), C.int(offset), C.int(stride*sizeofI32))
}

func PlotVLinesInt32CS(labelId CString, xs *int32, count, offset, stride int) {

	C.ImPlot_PlotVLines_S32Ptr(labelId.CString(), (*C.ImS32)(xs), C.int(count), C.int(offset), C.int(stride*sizeofI32))
}

func PlotVLinesFloat32(labelId string, xs *float32, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotVLines_FloatPtr(label_id, (*C.float)(xs), C.int(count), C.int(offset), C.int(stride*sizeofF32))
}

func PlotVLinesFloat32CS(labelId CString, xs *float32, count, offset, stride int) {

	C.ImPlot_PlotVLines_FloatPtr(labelId.CString(), (*C.float)(xs), C.int(count), C.int(offset), C.int(stride*sizeofF32))
}

func PlotVLinesFloat64(labelId string, xs *float64, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotVLines_doublePtr(label_id, (*C.double)(xs), C.int(count), C.int(offset), C.int(stride*sizeofF64))
}

func PlotVLinesFloat64CS(labelId CString, xs *float64, count, offset, stride int) {

	C.ImPlot_PlotVLines_doublePtr(labelId.CString(), (*C.double)(xs), C.int(count), C.int(offset), C.int(stride*sizeofF64))
}

func PlotHLinesInt32(labelId string, ys *int32, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotHLines_S32Ptr(label_id, (*C.ImS32)(ys), C.int(count), C.int(offset), C.int(stride*sizeofI32))
}

func PlotHLinesInt32CS(labelId CString, ys *int32, count, offset, stride int) {

	C.ImPlot_PlotHLines_S32Ptr(labelId.CString(), (*C.ImS32)(ys), C.int(count), C.int(offset), C.int(stride*sizeofI32))
}

func PlotHLinesFloat32(labelId string, ys *float32, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotHLines_FloatPtr(label_id, (*C.float)(ys), C.int(count), C.int(offset), C.int(stride*sizeofF32))
}

func PlotHLinesFloat32CS(labelId CString, ys *float32, count, offset, stride int) {

	C.ImPlot_PlotHLines_FloatPtr(labelId.CString(), (*C.float)(ys), C.int(count), C.int(offset), C.int(stride*sizeofF32))
}

func PlotHLinesFloat64(labelId string, ys *float64, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotHLines_doublePtr(label_id, (*C.double)(ys), C.int(count), C.int(offset), C.int(stride*sizeofF64))
}

func PlotHLinesFloat64CS(labelId CString, ys *float64, count, offset, stride int) {

	C.ImPlot_PlotHLines_doublePtr(labelId.CString(), (*C.double)(ys), C.int(count), C.int(offset), C.int(stride*sizeofF64))
}

func PlotPieChartInt32(labelIds []string, values *int32, count int, x, y, radius float64, normalize bool, labelFmt string, angle0 float64) {

	var csa CStrArr
	csa.Append(labelIds...)
	defer csa.Clear()
	label_fmt := C.CString(labelFmt)
	defer C.free(unsafe.Pointer(label_fmt))
	C.ImPlot_PlotPieChart_S32Ptr(csa.Ptr(), (*C.ImS32)(values), C.int(count),
		C.double(x), C.double(y), C.double(radius), C.bool(normalize), label_fmt, C.double(angle0))
}

func PlotPieChartInt32CS(labelIds CStrArr, values *int32, count int, x, y, radius float64, normalize bool, labelFmt CString, angle0 float64) {

	C.ImPlot_PlotPieChart_S32Ptr(labelIds.Ptr(), (*C.ImS32)(values), C.int(count),
		C.double(x), C.double(y), C.double(radius), C.bool(normalize), labelFmt.CString(), C.double(angle0))
}

func PlotPieChartFloat32(labelIds []string, values *float32, count int, x, y, radius float64, normalize bool, labelFmt string, angle0 float64) {

	var csa CStrArr
	csa.Append(labelIds...)
	defer csa.Clear()
	label_fmt := C.CString(labelFmt)
	defer C.free(unsafe.Pointer(label_fmt))
	C.ImPlot_PlotPieChart_FloatPtr(csa.Ptr(), (*C.float)(values), C.int(count),
		C.double(x), C.double(y), C.double(radius), C.bool(normalize), label_fmt, C.double(angle0))
}

func PlotPieChartFloat32CS(labelIds CStrArr, values *float32, count int, x, y, radius float64, normalize bool, labelFmt CString, angle0 float64) {

	C.ImPlot_PlotPieChart_FloatPtr(labelIds.Ptr(), (*C.float)(values), C.int(count),
		C.double(x), C.double(y), C.double(radius), C.bool(normalize), labelFmt.CString(), C.double(angle0))
}

func PlotPieChartFloat64(labelIds []string, values *float64, count int, x, y, radius float64, normalize bool, labelFmt string, angle0 float64) {

	var csa CStrArr
	csa.Append(labelIds...)
	defer csa.Clear()
	label_fmt := C.CString(labelFmt)
	defer C.free(unsafe.Pointer(label_fmt))
	C.ImPlot_PlotPieChart_doublePtr(csa.Ptr(), (*C.double)(values), C.int(count),
		C.double(x), C.double(y), C.double(radius), C.bool(normalize), label_fmt, C.double(angle0))
}

func PlotPieChartFloat64CS(labelIds CStrArr, values *float64, count int, x, y, radius float64, normalize bool, labelFmt CString, angle0 float64) {

	C.ImPlot_PlotPieChart_doublePtr(labelIds.Ptr(), (*C.double)(values), C.int(count),
		C.double(x), C.double(y), C.double(radius), C.bool(normalize), labelFmt.CString(), C.double(angle0))
}

func PlotHeatmapInt32(labelId string, values *int32, rows, cols int, scaleMin, scaleMax float64, labelFmt string, boundsMin, boundsMax PlotPoint) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	label_fmt := C.CString(labelFmt)
	defer C.free(unsafe.Pointer(label_fmt))
	C.ImPlot_PlotHeatmap_S32Ptr(label_id, (*C.ImS32)(values), C.int(rows), C.int(cols), C.double(scaleMin), C.double(scaleMax),
		label_fmt, boundsMin.ImPlotPoint(), boundsMax.ImPlotPoint())
}

func PlotHeatmapInt32CS(labelId CString, values *int32, rows, cols int, scaleMin, scaleMax float64, labelFmt CString, boundsMin, boundsMax PlotPoint) {

	C.ImPlot_PlotHeatmap_S32Ptr(labelId.CString(), (*C.ImS32)(values), C.int(rows), C.int(cols), C.double(scaleMin), C.double(scaleMax),
		labelFmt.CString(), boundsMin.ImPlotPoint(), boundsMax.ImPlotPoint())
}

func PlotHeatmapFloat32(labelId string, values *float32, rows, cols int, scaleMin, scaleMax float64, labelFmt string, boundsMin, boundsMax PlotPoint) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	label_fmt := C.CString(labelFmt)
	defer C.free(unsafe.Pointer(label_fmt))
	C.ImPlot_PlotHeatmap_FloatPtr(label_id, (*C.float)(values), C.int(rows), C.int(cols), C.double(scaleMin), C.double(scaleMax),
		label_fmt, boundsMin.ImPlotPoint(), boundsMax.ImPlotPoint())
}

func PlotHeatmapFloat32CS(labelId CString, values *float32, rows, cols int, scaleMin, scaleMax float64, labelFmt CString,
	boundsMin, boundsMax PlotPoint) {

	C.ImPlot_PlotHeatmap_FloatPtr(labelId.CString(), (*C.float)(values), C.int(rows), C.int(cols), C.double(scaleMin), C.double(scaleMax),
		labelFmt.CString(), boundsMin.ImPlotPoint(), boundsMax.ImPlotPoint())
}

func PlotHeatmapFloat64(labelId string, values *float64, rows, cols int, scaleMin, scaleMax float64, labelFmt string, boundsMin, boundsMax PlotPoint) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	label_fmt := C.CString(labelFmt)
	defer C.free(unsafe.Pointer(label_fmt))
	C.ImPlot_PlotHeatmap_doublePtr(label_id, (*C.double)(values), C.int(rows), C.int(cols), C.double(scaleMin), C.double(scaleMax),
		label_fmt, boundsMin.ImPlotPoint(), boundsMax.ImPlotPoint())
}

func PlotHeatmapFloat64CS(labelId CString, values *float64, rows, cols int, scaleMin, scaleMax float64, labelFmt CString,
	boundsMin, boundsMax PlotPoint) {

	C.ImPlot_PlotHeatmap_doublePtr(labelId.CString(), (*C.double)(values), C.int(rows), C.int(cols), C.double(scaleMin), C.double(scaleMax),
		labelFmt.CString(), boundsMin.ImPlotPoint(), boundsMax.ImPlotPoint())
}

func PlotHistogramInt32(labelId string, values *int32, count, bins int, cumulative, density bool, prange PlotRange, outliers bool, barScale float64) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotHistogram_S32Ptr(label_id, (*C.ImS32)(values), C.int(count), C.int(bins), C.bool(cumulative), C.bool(density),
		prange.ImPlotRange(), C.bool(outliers), C.double(barScale))
}

func PlotHistogramInt32CS(labelId CString, values *int32, count, bins int, cumulative, density bool, prange PlotRange, outliers bool, barScale float64) {

	C.ImPlot_PlotHistogram_S32Ptr(labelId.CString(), (*C.ImS32)(values), C.int(count), C.int(bins), C.bool(cumulative), C.bool(density),
		prange.ImPlotRange(), C.bool(outliers), C.double(barScale))
}

func PlotHistogramFloat32(labelId string, values *float32, count, bins int, cumulative, density bool, prange PlotRange, outliers bool, barScale float64) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotHistogram_FloatPtr(label_id, (*C.float)(values), C.int(count), C.int(bins), C.bool(cumulative), C.bool(density),
		prange.ImPlotRange(), C.bool(outliers), C.double(barScale))
}

func PlotHistogramFloat32CS(labelId CString, values *float32, count, bins int, cumulative, density bool, prange PlotRange, outliers bool, barScale float64) {

	C.ImPlot_PlotHistogram_FloatPtr(labelId.CString(), (*C.float)(values), C.int(count), C.int(bins), C.bool(cumulative), C.bool(density),
		prange.ImPlotRange(), C.bool(outliers), C.double(barScale))
}

func PlotHistogramFloat64(labelId string, values *float64, count, bins int, cumulative, density bool, prange PlotRange, outliers bool, barScale float64) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotHistogram_doublePtr(label_id, (*C.double)(values), C.int(count), C.int(bins), C.bool(cumulative), C.bool(density),
		prange.ImPlotRange(), C.bool(outliers), C.double(barScale))
}

func PlotHistogramFloat64CS(labelId CString, values *float64, count, bins int, cumulative, density bool, prange PlotRange, outliers bool, barScale float64) {

	C.ImPlot_PlotHistogram_doublePtr(labelId.CString(), (*C.double)(values), C.int(count), C.int(bins), C.bool(cumulative), C.bool(density),
		prange.ImPlotRange(), C.bool(outliers), C.double(barScale))
}

func PlotHistogram2DInt32(labelId string, xs, ys *int32, count, xbins, ybins int, density bool, prange PlotLimits, outliers bool) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotHistogram2D_S32Ptr(label_id, (*C.ImS32)(xs), (*C.ImS32)(ys), C.int(count), C.int(xbins), C.int(ybins),
		C.bool(density), prange.ImPlotLimits(), C.bool(outliers))
}

func PlotHistogram2DInt32CS(labelId CString, xs, ys *int32, count int, xbins, ybins int, density bool, prange PlotLimits, outliers bool) {

	C.ImPlot_PlotHistogram2D_S32Ptr(labelId.CString(), (*C.ImS32)(xs), (*C.ImS32)(ys), C.int(count), C.int(xbins), C.int(ybins),
		C.bool(density), prange.ImPlotLimits(), C.bool(outliers))
}

func PlotHistogram2DFloat32(labelId string, xs, ys *float32, count, xbins, ybins int, density bool, prange PlotLimits, outliers bool) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotHistogram2D_FloatPtr(label_id, (*C.float)(xs), (*C.float)(ys), C.int(count), C.int(xbins), C.int(ybins),
		C.bool(density), prange.ImPlotLimits(), C.bool(outliers))
}

func PlotHistogram2DFloat32CS(labelId CString, xs, ys *float32, count, xbins, ybins int, density bool, prange PlotLimits, outliers bool) {

	C.ImPlot_PlotHistogram2D_FloatPtr(labelId.CString(), (*C.float)(xs), (*C.float)(ys), C.int(count), C.int(xbins), C.int(ybins),
		C.bool(density), prange.ImPlotLimits(), C.bool(outliers))
}

func PlotHistogram2DFloat64(labelId string, xs, ys *float64, count, xbins, ybins int, density bool, prange PlotLimits, outliers bool) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotHistogram2D_doublePtr(label_id, (*C.double)(xs), (*C.double)(ys), C.int(count), C.int(xbins), C.int(ybins),
		C.bool(density), prange.ImPlotLimits(), C.bool(outliers))
}

func PlotHistogram2DFloat64CS(labelId CString, xs, ys *float64, count, xbins, ybins int, density bool, prange PlotLimits, outliers bool) {

	C.ImPlot_PlotHistogram2D_doublePtr(labelId.CString(), (*C.double)(xs), (*C.double)(ys), C.int(count), C.int(xbins), C.int(ybins),
		C.bool(density), prange.ImPlotLimits(), C.bool(outliers))
}

func PlotDigitalInt32(labelId string, xs, ys *int32, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotDigital_S32Ptr(label_id, (*C.ImS32)(xs), (*C.ImS32)(ys), C.int(count), C.int(offset), C.int(stride*sizeofI32))
}

func PlotDigitalInt32CS(labelId CString, xs, ys *int32, count, offset, stride int) {

	C.ImPlot_PlotDigital_S32Ptr(labelId.CString(), (*C.ImS32)(xs), (*C.ImS32)(ys), C.int(count), C.int(offset), C.int(stride*sizeofI32))
}

func PlotDigitalFloat32(labelId string, xs, ys *float32, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotDigital_FloatPtr(label_id, (*C.float)(xs), (*C.float)(ys), C.int(count), C.int(offset), C.int(stride*sizeofF32))
}

func PlotDigitalFloat32CS(labelId CString, xs, ys *float32, count, offset, stride int) {

	C.ImPlot_PlotDigital_FloatPtr(labelId.CString(), (*C.float)(xs), (*C.float)(ys), C.int(count), C.int(offset), C.int(stride*sizeofF32))
}

func PlotDigitalFloat64(labelId string, xs, ys *float64, count, offset, stride int) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotDigital_doublePtr(label_id, (*C.double)(xs), (*C.double)(ys), C.int(count), C.int(offset), C.int(stride*sizeofF64))
}

func PlotDigitalFloat64CS(labelId CString, xs, ys *float64, count, offset, stride int) {

	C.ImPlot_PlotDigital_doublePtr(labelId.CString(), (*C.double)(xs), (*C.double)(ys), C.int(count), C.int(offset), C.int(stride*sizeofF64))
}

func PlotImage(labelId string, userTextureID TextureID, boundsMin, boundsMax PlotPoint, uv0, uv1 Vec2, tintCol Vec4) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotImage(label_id, C.ImTextureID(userTextureID), boundsMin.ImPlotPoint(), boundsMax.ImPlotPoint(),
		uv0.ImVec2(), uv1.ImVec2(), tintCol.ImVec4())
}

func PlotImageCS(labelId CString, userTextureID TextureID, boundsMin, boundsMax PlotPoint, uv0, uv1 Vec2, tintCol Vec4) {

	C.ImPlot_PlotImage(labelId.CString(), C.ImTextureID(userTextureID), boundsMin.ImPlotPoint(), boundsMax.ImPlotPoint(),
		uv0.ImVec2(), uv1.ImVec2(), tintCol.ImVec4())
}

func PlotText(text string, x, y float64, vertical bool, pixOffset Vec2) {

	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.ImPlot_PlotText(ctext, C.double(x), C.double(y), C.bool(vertical), pixOffset.ImVec2())
}

func PlotTextCS(text CString, x, y float64, vertical bool, pixOffset Vec2) {

	C.ImPlot_PlotText(text.CString(), C.double(x), C.double(y), C.bool(vertical), pixOffset.ImVec2())
}

func PlotDummy(labelId string) {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	C.ImPlot_PlotDummy(label_id)
}

func PlotDummyCS(labelId CString) {

	C.ImPlot_PlotDummy(labelId.CString())
}

func SetNextPlotLimits(xmin, xmax, ymin, ymax float64, cond Cond) {

	C.ImPlot_SetNextPlotLimits(C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax), C.ImGuiCond(cond))
}

func SetNextPlotLimitsX(xmin, xmax float64, cond Cond) {

	C.ImPlot_SetNextPlotLimitsX(C.double(xmin), C.double(xmax), C.ImGuiCond(cond))
}

func SetNextPlotLimitsY(ymin, ymax float64, cond Cond, yaxis PlotYAxis) {

	C.ImPlot_SetNextPlotLimitsY(C.double(ymin), C.double(ymax), C.ImGuiCond(cond), C.ImPlotYAxis(yaxis))
}

func LinkNextPlotLimits(xmin, xmax, ymin, ymax, ymin2, ymax2, ymin3, ymax3 *float64) {

	C.ImPlot_LinkNextPlotLimits((*C.double)(xmin), (*C.double)(xmax), (*C.double)(ymin), (*C.double)(ymax), (*C.double)(ymin2),
		(*C.double)(ymax2), (*C.double)(ymin3), (*C.double)(ymax3))
}

func FitNextPlotAxes(x, y, y2, y3 bool) {

	C.ImPlot_FitNextPlotAxes(C.bool(x), C.bool(y), C.bool(y2), C.bool(y3))
}

func SetNextPlotTicksX(values []float64, labels []string, keepDefault bool) {

	var csa CStrArr
	csa.Append(labels...)
	defer csa.Clear()
	C.ImPlot_SetNextPlotTicksX_doublePtr((*C.double)(&values[0]), C.int(len(values)), csa.Ptr(), C.bool(keepDefault))
}

func SetNextPlotTicksXCS(values []float64, labels CStrArr, keepDefault bool) {

	C.ImPlot_SetNextPlotTicksX_doublePtr((*C.double)(&values[0]), C.int(len(values)), labels.Ptr(), C.bool(keepDefault))
}

func SetNextPlotTicksXRange(xMin, xMax float64, nticks int, labels []string, keepDefault bool) {

	var csa CStrArr
	csa.Append(labels...)
	defer csa.Clear()
	C.ImPlot_SetNextPlotTicksX_double(C.double(xMin), C.double(xMax), C.int(nticks), csa.Ptr(), C.bool(keepDefault))
}

func SetNextPlotTicksXRangeCS(xMin, xMax float64, nticks int, labels CStrArr, keepDefault bool) {

	C.ImPlot_SetNextPlotTicksX_double(C.double(xMin), C.double(xMax), C.int(nticks), labels.Ptr(), C.bool(keepDefault))
}

func SetNextPlotTicksY(values []float64, labels []string, keepDefault bool, yAxis PlotYAxis) {

	var csa CStrArr
	csa.Append(labels...)
	defer csa.Clear()
	C.ImPlot_SetNextPlotTicksY_doublePtr((*C.double)(&values[0]), C.int(len(values)), csa.Ptr(), C.bool(keepDefault), C.ImPlotYAxis(yAxis))
}

func SetNextPlotTicksYCS(values []float64, labels CStrArr, keepDefault bool, yAxis PlotYAxis) {

	C.ImPlot_SetNextPlotTicksY_doublePtr((*C.double)(&values[0]), C.int(len(values)), labels.Ptr(), C.bool(keepDefault), C.ImPlotYAxis(yAxis))
}

func SetNextPlotTicksYRange(yMin, yMax float64, nticks int, labels []string, keepDefault bool, yAxis PlotYAxis) {

	var csa CStrArr
	csa.Append(labels...)
	defer csa.Clear()
	C.ImPlot_SetNextPlotTicksY_double(C.double(yMin), C.double(yMax), C.int(nticks), csa.Ptr(), C.bool(keepDefault), C.ImPlotYAxis(yAxis))
}

func SetNextPlotTicksYRangeCS(yMin, yMax float64, nticks int, labels CStrArr, keepDefault bool, yAxis PlotYAxis) {

	C.ImPlot_SetNextPlotTicksY_double(C.double(yMin), C.double(yMax), C.int(nticks), labels.Ptr(), C.bool(keepDefault), C.ImPlotYAxis(yAxis))
}

func SetNextPlotFormatX(fmt string) {

	cfmt := C.CString(fmt)
	defer C.free(unsafe.Pointer(cfmt))
	C.ImPlot_SetNextPlotFormatX(cfmt)
}

func SetNextPlotFormatXCS(fmt CString) {

	C.ImPlot_SetNextPlotFormatX(fmt.CString())
}

func SetNextPlotFormatY(fmt string, yAxis PlotYAxis) {

	cfmt := C.CString(fmt)
	defer C.free(unsafe.Pointer(cfmt))
	C.ImPlot_SetNextPlotFormatY(cfmt, C.ImPlotYAxis(yAxis))
}

func SetNextPlotFormatYCS(fmt CString, yAxis PlotYAxis) {

	C.ImPlot_SetNextPlotFormatY(fmt.CString(), C.ImPlotYAxis(yAxis))
}

func SetPlotYAxis(yAxis PlotYAxis) {

	C.ImPlot_SetPlotYAxis(C.ImPlotYAxis(yAxis))
}

func HideNextItem(hidden bool, cond Cond) {

	C.ImPlot_HideNextItem(C.bool(hidden), C.ImGuiCond(cond))
}

//CIMGUI_API void ImPlot_PixelsToPlot_Vec2(ImPlotPoint *pOut,const ImVec2 pix,ImPlotYAxis y_axis);
//CIMGUI_API void ImPlot_PixelsToPlot_Float(ImPlotPoint *pOut,float x,float y,ImPlotYAxis y_axis);
//CIMGUI_API void ImPlot_PlotToPixels_PlotPoInt(ImVec2 *pOut,const ImPlotPoint plt,ImPlotYAxis y_axis);
//CIMGUI_API void ImPlot_PlotToPixels_double(ImVec2 *pOut,double x,double y,ImPlotYAxis y_axis);

func GetPlotPos() Vec2 {

	return C.xGetPlotPos().Vec2()
}

func GetPlotSize() Vec2 {

	return C.xGetPlotSize().Vec2()
}

func IsPlotHovered() bool {

	return bool(C.ImPlot_IsPlotHovered())
}

func IsPlotXAxisHovered() bool {

	return bool(C.ImPlot_IsPlotXAxisHovered())
}

func IsPlotYAxisHovered(yAxis PlotYAxis) bool {

	return bool(C.ImPlot_IsPlotYAxisHovered(C.ImPlotYAxis(yAxis)))
}

func GetPlotMousePos(yAxis PlotYAxis) PlotPoint {

	return C.xGetPlotMousePos(C.ImPlotYAxis(yAxis)).PlotPoint()
}

func GetPlotLimits(yAxis PlotYAxis) PlotLimits {

	return C.xGetPlotLimits(C.ImPlotYAxis(yAxis)).PlotLimits()
}

func IsPlotSelected() bool {

	return bool(C.ImPlot_IsPlotSelected())
}

//CIMGUI_API void ImPlot_GetPlotSelection(ImPlotLimits *pOut,ImPlotYAxis y_axis);
//CIMGUI_API bool ImPlot_IsPlotQueried(void);
//CIMGUI_API void ImPlot_GetPlotQuery(ImPlotLimits *pOut,ImPlotYAxis y_axis);
//CIMGUI_API void ImPlot_SetPlotQuery(const ImPlotLimits query,ImPlotYAxis y_axis);
//CIMGUI_API bool ImPlot_IsSubplotsHovered(void);
//CIMGUI_API bool ImPlot_BeginAlignedPlots(const char* group_id,ImPlotOrientation orientation);
//CIMGUI_API void ImPlot_EndAlignedPlots(void);
//CIMGUI_API void ImPlot_Annotate_Str(double x,double y,const ImVec2 pix_offset,const char* fmt,...);
//CIMGUI_API void ImPlot_Annotate_Vec4(double x,double y,const ImVec2 pix_offset,const ImVec4 color,const char* fmt,...);
//CIMGUI_API void ImPlot_AnnotateV_Str(double x,double y,const ImVec2 pix_offset,const char* fmt,va_list args);
//CIMGUI_API void ImPlot_AnnotateV_Vec4(double x,double y,const ImVec2 pix_offset,const ImVec4 color,const char* fmt,va_list args);
//CIMGUI_API void ImPlot_AnnotateClamped_Str(double x,double y,const ImVec2 pix_offset,const char* fmt,...);
//CIMGUI_API void ImPlot_AnnotateClamped_Vec4(double x,double y,const ImVec2 pix_offset,const ImVec4 color,const char* fmt,...);
//CIMGUI_API void ImPlot_AnnotateClampedV_Str(double x,double y,const ImVec2 pix_offset,const char* fmt,va_list args);
//CIMGUI_API void ImPlot_AnnotateClampedV_Vec4(double x,double y,const ImVec2 pix_offset,const ImVec4 color,const char* fmt,va_list args);
//CIMGUI_API bool ImPlot_DragLineX(const char* id,double* x_value,bool show_label,const ImVec4 col,float thickness);
//CIMGUI_API bool ImPlot_DragLineY(const char* id,double* y_value,bool show_label,const ImVec4 col,float thickness);
//CIMGUI_API bool ImPlot_DragPoint(const char* id,double* x,double* y,bool show_label,const ImVec4 col,float radius);

func SetLegendLocation(location PlotLocation, orientation PlotOrientation, outside bool) {

	C.ImPlot_SetLegendLocation(C.ImPlotLocation(location), C.ImPlotOrientation(orientation), C.bool(outside))
}

func SetMousePosLocation(location PlotLocation) {

	C.ImPlot_SetMousePosLocation(C.ImPlotLocation(location))
}

func IsLegendEntryHovered(labelId string) bool {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	return bool(C.ImPlot_IsLegendEntryHovered(label_id))
}

func BeginLegendPopup(labelId string, mouseButton MouseButton) bool {

	label_id := C.CString(labelId)
	defer C.free(unsafe.Pointer(label_id))
	return bool(C.ImPlot_BeginLegendPopup(label_id, C.ImGuiMouseButton(mouseButton)))
}

func EndLegendPopup() {

	C.ImPlot_EndLegendPopup()
}

func PlotBeginDragDropTarget() bool {

	return bool(C.ImPlot_BeginDragDropTarget())
}

func PlotBeginDragDropTargetX() bool {

	return bool(C.ImPlot_BeginDragDropTargetX())
}

func PlotBeginDragDropTargetY(axis PlotYAxis) bool {

	return bool(C.ImPlot_BeginDragDropTargetY(C.ImPlotYAxis(axis)))
}

func PlotBeginDragDropTargetLegend() bool {

	return bool(C.ImPlot_BeginDragDropTargetLegend())
}

func PlotEndDragDropTarget() {

	C.ImPlot_EndDragDropTarget()
}

//CIMGUI_API bool ImPlot_BeginDragDropSource(ImGuiKeyModFlags key_mods,ImGuiDragDropFlags flags);
//CIMGUI_API bool ImPlot_BeginDragDropSourceX(ImGuiKeyModFlags key_mods,ImGuiDragDropFlags flags);
//CIMGUI_API bool ImPlot_BeginDragDropSourceY(ImPlotYAxis axis,ImGuiKeyModFlags key_mods,ImGuiDragDropFlags flags);
//CIMGUI_API bool ImPlot_BeginDragDropSourceItem(const char* label_id,ImGuiDragDropFlags flags);
//CIMGUI_API void ImPlot_EndDragDropSource(void);
//CIMGUI_API ImPlotStyle* ImPlot_GetStyle(void);
//CIMGUI_API void ImPlot_StyleColorsAuto(ImPlotStyle* dst);
//CIMGUI_API void ImPlot_StyleColorsClassic(ImPlotStyle* dst);
//CIMGUI_API void ImPlot_StyleColorsDark(ImPlotStyle* dst);
//CIMGUI_API void ImPlot_StyleColorsLight(ImPlotStyle* dst);

func PlotPushStyleColorU32(idx PlotCol, col U32) {

	C.ImPlot_PushStyleColor_U32(C.ImPlotCol(idx), C.ImU32(col))
}

func PlotPushStyleColor(idx PlotCol, col Vec4) {

	C.ImPlot_PushStyleColor_Vec4(C.ImPlotCol(idx), col.ImVec4())
}

func PlotPopStyleColor(count int) {

	C.ImPlot_PopStyleColor(C.int(count))
}

func PlotPushStyleVarFloat32(idx PlotStyleVar, val float32) {

	C.ImPlot_PushStyleVar_Float(C.ImPlotStyleVar(idx), C.float(val))
}

func PlotPushStyleVarInt(idx PlotStyleVar, val int) {

	C.ImPlot_PushStyleVar_Int(C.ImPlotStyleVar(idx), C.int(val))
}

func PlotPushStyleVarVec2(idx PlotStyleVar, val Vec2) {

	C.ImPlot_PushStyleVar_Vec2(C.ImPlotStyleVar(idx), val.ImVec2())
}

func PlotPopStyleVar(count int) {

	C.ImPlot_PopStyleVar(C.int(count))
}

func SetNextLineStyle(col Vec4, weight float32) {

	C.ImPlot_SetNextLineStyle(col.ImVec4(), C.float(weight))
}

func SetNextFillStyle(col Vec4, alphaMod float32) {

	C.ImPlot_SetNextFillStyle(col.ImVec4(), C.float(alphaMod))
}

func SetNextMarkerStyle(marker PlotMarker, size float32, fill Vec4, weight float32, outline Vec4) {

	C.ImPlot_SetNextMarkerStyle(C.ImPlotMarker(marker), C.float(size), fill.ImVec4(), C.float(weight), outline.ImVec4())
}

func SetNextErrorBarStyle(col Vec4, size float32, weight float32) {

	C.ImPlot_SetNextErrorBarStyle(col.ImVec4(), C.float(size), C.float(weight))
}

func GetLastItemColor() Vec4 {

	return C.xGetLastItemColor().Vec4()
}

func PlotGetStyleColorName(idx PlotCol) string {

	return C.GoString(C.ImPlot_GetStyleColorName(C.ImPlotCol(idx)))
}

func GetMarkerName(idx PlotMarker) string {

	return C.GoString(C.ImPlot_GetMarkerName(C.ImPlotMarker(idx)))
}

//CIMGUI_API ImPlotColormap ImPlot_AddColormap_Vec4Ptr(const char* name,const ImVec4* cols,int size,bool qual);
//CIMGUI_API ImPlotColormap ImPlot_AddColormap_U32Ptr(const char* name,const ImU32* cols,int size,bool qual);

func GetColormapCount() int {

	return int(C.ImPlot_GetColormapCount())
}

func GetColormapName(cmap PlotColormap) string {

	return C.GoString(C.ImPlot_GetColormapName(C.ImPlotColormap(cmap)))
}

func GetColormapIndex(name string) PlotColormap {

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return PlotColormap(C.ImPlot_GetColormapIndex(cname))
}

func PushColormap(cmap PlotColormap) {

	C.ImPlot_PushColormap_PlotColormap(C.ImPlotColormap(cmap))
}

func PopColormap(count int) {

	C.ImPlot_PopColormap(C.int(count))
}

func PushColormapStr(name string) {

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.ImPlot_PushColormap_Str(cname)
}

func PushColormapStrCS(name CString) {

	C.ImPlot_PushColormap_Str(name.CString())
}

func NextColormapColor() Vec4 {

	return C.xNextColormapColor().Vec4()
}

func GetColormapSize(cmap PlotColormap) int {

	return int(C.ImPlot_GetColormapSize(C.ImPlotColormap(cmap)))
}

func GetColormapColor(idx int, cmap PlotColormap) Vec4 {

	return C.xGetColormapColor(C.int(idx), C.ImPlotColormap(cmap)).Vec4()
}

func SampleColorMap(t float32, cmap PlotColormap) Vec4 {

	return C.xSampleColormap(C.float(t), C.ImPlotColormap(cmap)).Vec4()
}

func ColormapScale(label string, scaleMin, scaleMax float64, size Vec2, cmap PlotColormap, fmt string) {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cfmt := C.CString(fmt)
	defer C.free(unsafe.Pointer(cfmt))
	C.ImPlot_ColormapScale(clabel, C.double(scaleMin), C.double(scaleMax), size.ImVec2(), C.ImPlotColormap(cmap), cfmt)
}

func ColormapScaleCS(label CString, scaleMin, scaleMax float64, size Vec2, cmap PlotColormap, fmt CString) {

	C.ImPlot_ColormapScale(label.CString(), C.double(scaleMin), C.double(scaleMax), size.ImVec2(),
		C.ImPlotColormap(cmap), fmt.CString())
}

//CIMGUI_API bool ImPlot_ColormapSlider(const char* label,float* t,ImVec4* out,const char* format,ImPlotColormap cmap);

func ColormapButton(label string, size Vec2, cmap PlotColormap) bool {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.ImPlot_ColormapButton(clabel, size.ImVec2(), C.ImPlotColormap(cmap)))
}

func ColormapButtonCS(label CString, size Vec2, cmap PlotColormap) bool {

	return bool(C.ImPlot_ColormapButton(label.CString(), size.ImVec2(), C.ImPlotColormap(cmap)))
}

func BustColorCache(plotTitleId string) {

	plot_title_id := C.CString(plotTitleId)
	defer C.free(unsafe.Pointer(plot_title_id))
	C.ImPlot_BustColorCache(plot_title_id)
}

func BustColorCacheCS(plotTitleId CString) {

	C.ImPlot_BustColorCache(plotTitleId.CString())
}

func PlotItemIconVec4(col Vec4) {

	C.ImPlot_ItemIcon_Vec4(col.ImVec4())
}

func PlotItemIconU32(col U32) {

	C.ImPlot_ItemIcon_U32(C.ImU32(col))
}

func ColormapIcon(cmap PlotColormap) {

	C.ImPlot_ColormapIcon(C.ImPlotColormap(cmap))
}

func GetPlotDrawList() *DrawList {

	return (*DrawList)(C.ImPlot_GetPlotDrawList())
}

func PlotPushtPlotClipRect(expand float32) {

	C.ImPlot_PushPlotClipRect(C.float(expand))
}

func PlotPopPlotClipRect() {

	C.ImPlot_PopPlotClipRect()
}

func PlotShowStyleSelector(label string) {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	C.ImPlot_ShowStyleSelector(clabel)
}

func PlotShowColormapSelector(label string) {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	C.ImPlot_ShowColormapSelector(clabel)
}

func PlotShowStyleEditor(ref *PlotStyle) {

	C.ImPlot_ShowStyleEditor((*C.ImPlotStyle)(ref))
}

func PlotShowUserGuide() {

	C.ImPlot_ShowUserGuide()
}

func PlotShowMetricsWindow(pOpen *bool) {

	C.ImPlot_ShowMetricsWindow((*C.bool)(pOpen))
}

func PlotShowDemoWindow(pOpen *bool) {

	C.ImPlot_ShowDemoWindow((*C.bool)(pOpen))
}

func pf32(data []float32) *C.float {

	if len(data) == 0 {
		return nil
	}
	return (*C.float)(&data[0])
}
