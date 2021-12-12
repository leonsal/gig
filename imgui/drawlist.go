package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
// static ImVec2 xImDrawList_GetClipRectMin(ImDrawList* self) {
//		ImVec2 out;
//		ImDrawList_GetClipRectMin(&out, self);
//		return out;
// }
// static ImVec2 xImDrawList_GetClipRectMax(ImDrawList* self) {
//		ImVec2 out;
//		ImDrawList_GetClipRectMax(&out, self);
//		return out;
// }
import "C"
import "unsafe"

type DrawList C.ImDrawList

type DrawFlags C.ImDrawFlags

const (
	DrawFlags_None                    DrawFlags = C.ImDrawFlags_None
	DrawFlags_Closed                  DrawFlags = C.ImDrawFlags_Closed
	DrawFlags_RoundCornersTopLeft     DrawFlags = C.ImDrawFlags_RoundCornersTopLeft
	DrawFlags_RoundCornersTopRight    DrawFlags = C.ImDrawFlags_RoundCornersTopRight
	DrawFlags_RoundCornersBottomLeft  DrawFlags = C.ImDrawFlags_RoundCornersBottomLeft
	DrawFlags_RoundCornersBottomRight DrawFlags = C.ImDrawFlags_RoundCornersBottomRight
	DrawFlags_RoundCornersNone        DrawFlags = C.ImDrawFlags_RoundCornersNone
	DrawFlags_RoundCornersTop         DrawFlags = C.ImDrawFlags_RoundCornersTop
	DrawFlags_RoundCornersBottom      DrawFlags = C.ImDrawFlags_RoundCornersBottom
	DrawFlags_RoundCornersLeft        DrawFlags = C.ImDrawFlags_RoundCornersLeft
	DrawFlags_RoundCornersRight       DrawFlags = C.ImDrawFlags_RoundCornersRight
	DrawFlags_RoundCornersAll         DrawFlags = C.ImDrawFlags_RoundCornersAll
	DrawFlags_RoundCornersDefault_    DrawFlags = C.ImDrawFlags_RoundCornersDefault_
	DrawFlags_RoundCornersMask_       DrawFlags = C.ImDrawFlags_RoundCornersMask_
)

func (d *DrawList) PushClipRect(clipRectMin, clipRectMax Vec2, intersectWithCurrentClipRect bool) {

	C.ImDrawList_PushClipRect((*C.ImDrawList)(d), clipRectMin.ImVec2(), clipRectMax.ImVec2(), C.bool(intersectWithCurrentClipRect))
}

func (d *DrawList) PusthClipRectFullScreen() {
	C.ImDrawList_PushClipRectFullScreen((*C.ImDrawList)(d))
}

func (d *DrawList) PopClipRect() {

	C.ImDrawList_PopClipRect((*C.ImDrawList)(d))
}

func (d *DrawList) PushTextureID(textureId TextureID) {

	C.ImDrawList_PushTextureID((*C.ImDrawList)(d), C.ImTextureID(textureId))
}

func (d *DrawList) PopTextureID() {

	C.ImDrawList_PopTextureID((*C.ImDrawList)(d))
}

func (d *DrawList) GetClipRectMin() Vec2 {

	return C.xImDrawList_GetClipRectMin((*C.ImDrawList)(d)).Vec2()
}

func (d *DrawList) GetClipRectMax() Vec2 {

	return C.xImDrawList_GetClipRectMax((*C.ImDrawList)(d)).Vec2()
}

func (d *DrawList) AddLine(p1, p2 Vec2, col U32, thickness float32) {

	C.ImDrawList_AddLine((*C.ImDrawList)(d), p1.ImVec2(), p2.ImVec2(), C.ImU32(col), C.float(thickness))
}

func (d *DrawList) AddRect(pMin, pMax Vec2, col U32, rounding float32, flags DrawFlags, thickness float32) {

	C.ImDrawList_AddRect((*C.ImDrawList)(d), pMin.ImVec2(), pMax.ImVec2(), C.ImU32(col), C.float(rounding),
		C.ImDrawFlags(flags), C.float(thickness))
}

func (d *DrawList) AddRectFilled(pMin, pMax Vec2, col U32, rounding float32, flags DrawFlags) {

	C.ImDrawList_AddRectFilled((*C.ImDrawList)(d), pMin.ImVec2(), pMax.ImVec2(), C.ImU32(col), C.float(rounding), C.ImDrawFlags(flags))
}

func (d *DrawList) AddRectFilledMultiColor(pMin, pMax Vec2, colUprLeft, colUprRight, colBotRight, colBotLeft U32) {

	C.ImDrawList_AddRectFilledMultiColor((*C.ImDrawList)(d), pMin.ImVec2(), pMax.ImVec2(), C.ImU32(colUprLeft), C.ImU32(colUprRight),
		C.ImU32(colBotRight), C.ImU32(colBotLeft))
}

func (d *DrawList) AddQuad(p1, p2, p3, p4 Vec2, col U32, thickness float32) {

	C.ImDrawList_AddQuad((*C.ImDrawList)(d), p1.ImVec2(), p2.ImVec2(), p3.ImVec2(), p4.ImVec2(), C.ImU32(col), C.float(thickness))
}

func (d *DrawList) AddQuadFilled(p1, p2, p3, p4 Vec2, col U32) {

	C.ImDrawList_AddQuadFilled((*C.ImDrawList)(d), p1.ImVec2(), p2.ImVec2(), p3.ImVec2(), p4.ImVec2(), C.ImU32(col))
}

func (d *DrawList) AddTriangle(p1, p2, p3 Vec2, col U32, thickness float32) {

	C.ImDrawList_AddTriangle((*C.ImDrawList)(d), p1.ImVec2(), p2.ImVec2(), p3.ImVec2(), C.ImU32(col), C.float(thickness))
}

func (d *DrawList) AddTriangleFilled(p1, p2, p3 Vec2, col U32) {

	C.ImDrawList_AddTriangleFilled((*C.ImDrawList)(d), p1.ImVec2(), p2.ImVec2(), p3.ImVec2(), C.ImU32(col))
}

func (d *DrawList) AddCircle(center Vec2, radius float32, col U32, numSegments int, thickness float32) {

	C.ImDrawList_AddCircle((*C.ImDrawList)(d), center.ImVec2(), C.float(radius), C.ImU32(col), C.int(numSegments), C.float(thickness))
}

func (d *DrawList) AddCircleFilled(center Vec2, radius float32, col U32, numSegments int) {

	C.ImDrawList_AddCircleFilled((*C.ImDrawList)(d), center.ImVec2(), C.float(radius), C.ImU32(col), C.int(numSegments))
}

func (d *DrawList) AddNgon(center Vec2, radius float32, col U32, numSegments int, thickness float32) {

	C.ImDrawList_AddNgon((*C.ImDrawList)(d), center.ImVec2(), C.float(radius), C.ImU32(col), C.int(numSegments), C.float(thickness))
}

func (d *DrawList) AddNgonFilled(center Vec2, radius float32, col U32, numSegments int) {

	C.ImDrawList_AddNgonFilled((*C.ImDrawList)(d), center.ImVec2(), C.float(radius), C.ImU32(col), C.int(numSegments))
}

func (d *DrawList) AddText(pos Vec2, col U32, text string) {

	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.ImDrawList_AddText_Vec2((*C.ImDrawList)(d), pos.ImVec2(), C.ImU32(col), ctext, nil)
}

func (d *DrawList) AddTextFont(font *Font, fontSize float32, pos Vec2, col U32, text string, wrapWidth float32, cpuFineClipRect Vec4) {

	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	var cpu_fine_clip_rect C.ImVec4
	C.ImDrawList_AddText_FontPtr((*C.ImDrawList)(d), (*C.ImFont)(font), C.float(fontSize), pos.ImVec2(), C.ImU32(col),
		ctext, nil, C.float(wrapWidth), &cpu_fine_clip_rect)
}

func (d *DrawList) AddPolyline(points []Vec2, col U32, flags DrawFlags, thickness float32) {

	C.ImDrawList_AddPolyline((*C.ImDrawList)(d), (*C.ImVec2)(unsafe.Pointer(&points[0])), C.int(len(points)), C.ImU32(col),
		C.ImDrawFlags(flags), C.float(thickness))
}

func (d *DrawList) AddConvexPolyline(points []Vec2, col U32) {

	C.ImDrawList_AddConvexPolyFilled((*C.ImDrawList)(d), (*C.ImVec2)(unsafe.Pointer(&points[0])), C.int(len(points)), C.ImU32(col))
}

func (d *DrawList) AddBezierCubic(p1, p2, p3, p4 Vec2, col U32, thickness float32, numSegments int) {

	C.ImDrawList_AddBezierCubic((*C.ImDrawList)(d), p1.ImVec2(), p2.ImVec2(), p3.ImVec2(), p4.ImVec2(), C.ImU32(col),
		C.float(thickness), C.int(numSegments))
}

func (d *DrawList) AddBezierQuadratic(p1, p2, p3 Vec2, col U32, thickness float32, numSegments int) {

	C.ImDrawList_AddBezierQuadratic((*C.ImDrawList)(d), p1.ImVec2(), p2.ImVec2(), p3.ImVec2(), C.ImU32(col),
		C.float(thickness), C.int(numSegments))
}

func (d *DrawList) AddImage(userTextureID TextureID, pMin, pMax Vec2, uvMin, uvMax Vec2, col U32) {

	C.ImDrawList_AddImage((*C.ImDrawList)(d), C.ImTextureID(userTextureID), pMin.ImVec2(), pMax.ImVec2(),
		uvMin.ImVec2(), uvMax.ImVec2(), C.ImU32(col))
}

func (d *DrawList) AddImageQuad(userTextureID TextureID, p1, p2, p3, p4, uv1, uv2, uv3, uv4 Vec2, col U32) {

	C.ImDrawList_AddImageQuad((*C.ImDrawList)(d), C.ImTextureID(userTextureID), p1.ImVec2(), p2.ImVec2(), p3.ImVec2(), p4.ImVec2(),
		uv1.ImVec2(), uv2.ImVec2(), uv3.ImVec2(), uv4.ImVec2(), C.ImU32(col))
}

func (d *DrawList) AddImageRounded(userTextureID TextureID, pMin, pMax Vec2, uvMin, uvMax Vec2, col U32, rounding float32, flags DrawFlags) {

	C.ImDrawList_AddImageRounded((*C.ImDrawList)(d), C.ImTextureID(userTextureID), pMin.ImVec2(), pMax.ImVec2(),
		uvMin.ImVec2(), uvMax.ImVec2(), C.ImU32(col), C.float(rounding), C.ImDrawFlags(flags))
}

func (d *DrawList) PathClear() {

	C.ImDrawList_PathClear((*C.ImDrawList)(d))
}

func (d *DrawList) PathLineTo(pos Vec2) {

	C.ImDrawList_PathLineTo((*C.ImDrawList)(d), pos.ImVec2())
}

func (d *DrawList) PathLineToMergeDuplicate(pos Vec2) {

	C.ImDrawList_PathLineToMergeDuplicate((*C.ImDrawList)(d), pos.ImVec2())
}

func (d *DrawList) PathFillConvex(col U32) {

	C.ImDrawList_PathFillConvex((*C.ImDrawList)(d), C.ImU32(col))
}

func (d *DrawList) PathStroke(col U32, flags DrawFlags, thickness float32) {

	C.ImDrawList_PathStroke((*C.ImDrawList)(d), C.ImU32(col), C.ImDrawFlags(flags), C.float(thickness))
}

func (d *DrawList) PathArcTo(center Vec2, radius, aMin, aMax float32, numSegments int) {

	C.ImDrawList_PathArcTo((*C.ImDrawList)(d), center.ImVec2(), C.float(radius), C.float(aMin), C.float(aMax),
		C.int(numSegments))
}

func (d *DrawList) PathArcToFast(center Vec2, radius, aMin12, aMax12 int) {

	C.ImDrawList_PathArcToFast((*C.ImDrawList)(d), center.ImVec2(), C.float(radius), C.int(aMin12), C.int(aMin12))
}

func (d *DrawList) PathBezierCubicCurveTo(p2, p3, p4 Vec2, numSegments int) {

	C.ImDrawList_PathBezierCubicCurveTo((*C.ImDrawList)(d), p2.ImVec2(), p3.ImVec2(), p4.ImVec2(), C.int(numSegments))
}

func (d *DrawList) PathBezierQuadraticCurveTo(p2, p3 Vec2, numSegments int) {

	C.ImDrawList_PathBezierQuadraticCurveTo((*C.ImDrawList)(d), p2.ImVec2(), p3.ImVec2(), C.int(numSegments))
}
