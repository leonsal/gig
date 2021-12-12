package imgui

// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
// static ImVec2 xGetCursorPos(void) {
// 		ImVec2 pos;
//		igGetCursorPos(&pos);
//		return pos;
// }
// static ImVec2 xGetCursorStartPos(void) {
// 		ImVec2 pos;
//		igGetCursorStartPos(&pos);
//		return pos;
// }
// static ImVec2 xGetCursorScreenPos(void) {
// 		ImVec2 pos;
//		igGetCursorScreenPos(&pos);
//		return pos;
// }
import "C"

func Separator() {

	C.igSeparator()
}

func SameLine(offsetFromStartX, spacing float32) {

	C.igSameLine(C.float(offsetFromStartX), C.float(spacing))
}

func NewLine() {

	C.igNewLine()
}

func Spacing() {

	C.igSpacing()
}

func Dummy(size Vec2) {

	C.igDummy(size.ImVec2())
}

func Indent(indentW float32) {

	C.igIndent(C.float(indentW))
}

func Unindent(indentW float32) {

	C.igUnindent(C.float(indentW))
}

func BeginGroup() {

	C.igBeginGroup()
}

func EndGroup() {

	C.igEndGroup()
}

func GetCursorPos() Vec2 {

	return C.xGetCursorPos().Vec2()
}

func GetCursorPosX() float32 {

	return float32(C.igGetCursorPosX())
}

func GetCursorPosY() float32 {

	return float32(C.igGetCursorPosY())
}

func SetCursorPos(localPos Vec2) {

	C.igSetCursorPos(localPos.ImVec2())
}

func SetCursorPosX(localX float32) {

	C.igSetCursorPosX(C.float(localX))
}

func SetCursorPosY(localY float32) {

	C.igSetCursorPosY(C.float(localY))
}

func GetCursorStartPos() Vec2 {

	return C.xGetCursorStartPos().Vec2()
}

func GetCursorScreenPos() Vec2 {

	return C.xGetCursorScreenPos().Vec2()
}

func SetCursorScreenPos(pos Vec2) {

	C.igSetCursorScreenPos(pos.ImVec2())
}

func AlignTextToFramePadding() {

	C.igAlignTextToFramePadding()
}

func GetTextLineHeight() float32 {

	return float32(C.igGetTextLineHeight())
}

func GetTextLineHeightWithSpacing() float32 {

	return float32(C.igGetTextLineHeightWithSpacing())
}

func GetFrameHeight() float32 {

	return float32(C.igGetFrameHeight())
}

func GetFrameHeightWithSpacing() float32 {

	return float32(C.igGetFrameHeightWithSpacing())
}
