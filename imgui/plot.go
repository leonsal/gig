package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"
import "unsafe"

func PlotLines(label string, values []float32, valuesOffset int, overlayText string, scaleMin, scaleMax float32, graphSize Vec2, stride int) {

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	overlay_text := C.CString(overlayText)
	defer C.free(unsafe.Pointer(overlay_text))
	C.igPlotLines_FloatPtr(clabel, (*C.float)(&values[0]), C.int(len(values)), C.int(valuesOffset),
		overlay_text, C.float(scaleMin), C.float(scaleMax), graphSize.ImVec2(), C.int(stride))
}

// TODO?
//CIMGUI_API void igPlotLines_FnFloatPtr(const char* label,float(*values_getter)(void* data,int idx),void* data,int values_count,int values_offset,const char* overlay_text,float scale_min,float scale_max,ImVec2 graph_size);

func PlotHistogram(label string, values []float32, valuesOffset int, overlayText string, scaleMin, scaleMax float32, graphSize Vec2, stride int) {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	overlay_text := C.CString(overlayText)
	defer C.free(unsafe.Pointer(overlay_text))
	C.igPlotHistogram_FloatPtr(clabel, (*C.float)(&values[0]), C.int(len(values)), C.int(valuesOffset),
		overlay_text, C.float(scaleMin), C.float(scaleMax), graphSize.ImVec2(), C.int(stride))
}

// TODO?
//CIMGUI_API void igPlotHistogram_FnFloatPtr(const char* label,float(*values_getter)(void* data,int idx),void* data,int values_count,int values_offset,const char* overlay_text,float scale_min,float scale_max,ImVec2 graph_size);
