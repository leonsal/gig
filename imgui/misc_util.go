package imgui

// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"

func IsRectVisible(size Vec2) bool {

	return bool(C.igIsRectVisible_Nil(size.ImVec2()))
}

func IsRectVisibleMinMax(rectMin, rectMax Vec2) bool {

	return bool(C.igIsRectVisible_Vec2(rectMin.ImVec2(), rectMax.ImVec2()))
}

func GetTime() float64 {

	return float64(C.igGetTime())
}

func GetFrameCount() int {

	return int(C.igGetFrameCount())
}

func GetBackgroundDrawList() *DrawList {

	return (*DrawList)(C.igGetBackgroundDrawList_Nil())
}

func GetForegroundDrawList() *DrawList {

	return (*DrawList)(C.igGetForegroundDrawList_Nil())
}

func GetDrawListSharedData() *DrawListSharedData {

	return (*DrawListSharedData)(C.igGetDrawListSharedData())
}

func GetStyleColorName(idx Col) string {

	return C.GoString(C.igGetStyleColorName(C.ImGuiCol(idx)))
}

func SetStateStorage(storage *Storage) {

	C.igSetStateStorage((*C.ImGuiStorage)(storage))
}

func GetStateStorage() *Storage {

	return (*Storage)(C.igGetStateStorage())
}

// TODO?
//CIMGUI_API void igCalcListClipping(int items_count,float items_height,int* out_items_display_start,int* out_items_display_end);

func BeginChildFrame(id ID, size Vec2, flags WindowFlags) bool {

	return bool(C.igBeginChildFrame(C.ImGuiID(id), size.ImVec2(), C.ImGuiWindowFlags(flags)))
}

func EndChildFrame() {

	C.igEndChildFrame()
}
