package imgui

// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"

type Style C.ImGuiStyle

func GetStyle() *Style {

	return (*Style)(C.igGetStyle())
}

func (s *Style) GetAlpha() float32 {

	return float32(s.Alpha)
}

func (s *Style) SetAlpha(alpha float32) {

	s.Alpha = C.float(alpha)
}

func (s *Style) GetWindowPadding() Vec2 {

	return s.WindowPadding.Vec2()
}

func (s *Style) SetWindowPadding(padding Vec2) {

	s.WindowPadding = padding.ImVec2()
}

func (s *Style) GetWindowRounding() float32 {

	return float32(s.WindowRounding)
}

func (s *Style) SetWindowRounding(rounding float32) {

	s.WindowRounding = C.float(rounding)
}

func (s *Style) GetWindowBorderSize() float32 {

	return float32(s.WindowBorderSize)
}

func (s *Style) SetWindowBorderSize(size float32) {

	s.WindowBorderSize = C.float(size)
}

func (s *Style) GetWindowMinSize() Vec2 {

	return s.WindowMinSize.Vec2()
}

func (s *Style) SetWindowMinSize(size Vec2) {

	s.WindowMinSize = size.ImVec2()
}

func (s *Style) GetWindowTitleAlign() Vec2 {

	return s.WindowTitleAlign.Vec2()
}

func (s *Style) SetWindowTitleAlign(align Vec2) {

	s.WindowTitleAlign = align.ImVec2()
}

func (s *Style) GetWindowMenuButtonPosition() Dir {

	return Dir(s.WindowMenuButtonPosition)
}

func (s *Style) SetWindowMenuButtonPosition(pos Dir) {

	s.WindowMenuButtonPosition = C.ImGuiDir(pos)
}

func (s *Style) GetChildRounding() float32 {

	return float32(s.ChildRounding)
}

func (s *Style) SetChildRounding(rounding float32) {

	s.ChildRounding = C.float(rounding)
}

func (s *Style) GetChildBorderSize() float32 {

	return float32(s.ChildBorderSize)
}

func (s *Style) SetChildBorderSize(size float32) {

	s.ChildBorderSize = C.float(size)
}

func (s *Style) GetPopupRounding() float32 {

	return float32(s.PopupRounding)
}

func (s *Style) SetPopupRounding(rounding float32) {

	s.PopupRounding = C.float(rounding)
}

func (s *Style) GetPopupBorderSize() float32 {

	return float32(s.PopupBorderSize)
}

func (s *Style) SetPopupBorderSize(size float32) {

	s.PopupBorderSize = C.float(size)
}

func (s *Style) GetFramePadding() Vec2 {

	return s.FramePadding.Vec2()
}

func (s *Style) SetFramePadding(padding Vec2) {

	s.FramePadding = padding.ImVec2()
}

func (s *Style) GetFrameRounding() float32 {

	return float32(s.FrameRounding)
}

func (s *Style) SetFrameRounding(rounding float32) {

	s.FrameRounding = C.float(rounding)
}

func (s *Style) GetFrameBorderSize() float32 {

	return float32(s.FrameBorderSize)
}

func (s *Style) SetFrameBorderSize(size float32) {

	s.FrameBorderSize = C.float(size)
}

func (s *Style) GetItemSpacing() Vec2 {

	return s.ItemSpacing.Vec2()
}

func (s *Style) SetItemSpacing(spacing Vec2) {

	s.ItemSpacing = spacing.ImVec2()
}

func (s *Style) GetItemInnerSpacing() Vec2 {

	return s.ItemInnerSpacing.Vec2()
}

func (s *Style) SetItemInnerSpacing(spacing Vec2) {

	s.ItemInnerSpacing = spacing.ImVec2()
}

func (s *Style) GetCellPadding() Vec2 {

	return s.CellPadding.Vec2()
}

func (s *Style) SetCellPadding(padding Vec2) {

	s.CellPadding = padding.ImVec2()
}

func (s *Style) GetTouchExtraPadding() Vec2 {

	return s.TouchExtraPadding.Vec2()
}

func (s *Style) SetTouchExtraPadding(padding Vec2) {

	s.TouchExtraPadding = padding.ImVec2()
}

func (s *Style) GetIndentSpacing() float32 {

	return float32(s.IndentSpacing)
}

func (s *Style) SetIndentSpacing(spacing float32) {

	s.IndentSpacing = C.float(spacing)
}

func (s *Style) GetColumnsMinSpacing() float32 {

	return float32(s.ColumnsMinSpacing)
}

func (s *Style) SetColumnsMinSpacing(spacing float32) {

	s.ColumnsMinSpacing = C.float(spacing)
}

func (s *Style) GetScrollbarSize() float32 {

	return float32(s.ScrollbarSize)
}

func (s *Style) SetScrollbarSize(size float32) {

	s.ScrollbarSize = C.float(size)
}

func (s *Style) GetScrollbarRounding() float32 {

	return float32(s.ScrollbarRounding)
}

func (s *Style) SetScrollbarRounding(rounding float32) {

	s.ScrollbarRounding = C.float(rounding)
}

func (s *Style) GetGrabMinSize() float32 {

	return float32(s.GrabMinSize)
}

func (s *Style) SetGrabMinSize(size float32) {

	s.GrabMinSize = C.float(size)
}

func (s *Style) GetGrabRounding() float32 {

	return float32(s.GrabRounding)
}

func (s *Style) SetGrabRounding(rounding float32) {

	s.GrabRounding = C.float(rounding)
}

func (s *Style) GetLogSliderDeadzone() float32 {

	return float32(s.LogSliderDeadzone)
}

func (s *Style) SetLogSliderDeadzone(zone float32) {

	s.LogSliderDeadzone = C.float(zone)
}

func (s *Style) GetTabRounding() float32 {

	return float32(s.TabRounding)
}

func (s *Style) SetTabRounding(rounding float32) {

	s.TabRounding = C.float(rounding)
}

func (s *Style) GetTabBorderSize() float32 {

	return float32(s.TabBorderSize)
}

func (s *Style) SetTabBorderSize(size float32) {

	s.TabBorderSize = C.float(size)
}

func (s *Style) GetTabMinWidthForCloseButton() float32 {

	return float32(s.TabMinWidthForCloseButton)
}

func (s *Style) SetTabMinWidthForCloseButton(width float32) {

	s.TabMinWidthForCloseButton = C.float(width)
}

func (s *Style) GetColorButtonPosition() Dir {

	return Dir(s.ColorButtonPosition)
}

func (s *Style) SetColorButtonPosition(pos Dir) {

	s.ColorButtonPosition = C.ImGuiDir(pos)
}

func (s *Style) GetButtonTextAlign() Vec2 {

	return s.ButtonTextAlign.Vec2()
}

func (s *Style) SetButtonTextAlign(align Vec2) {

	s.ButtonTextAlign = align.ImVec2()
}

func (s *Style) GetSelectableTextAlign() Vec2 {

	return s.SelectableTextAlign.Vec2()
}

func (s *Style) SetSelectableTextAlign(align Vec2) {

	s.SelectableTextAlign = align.ImVec2()
}

func (s *Style) GetDisplayWindowPadding() Vec2 {

	return s.DisplayWindowPadding.Vec2()
}

func (s *Style) SetDisplayWindowPadding(padding Vec2) {

	s.DisplayWindowPadding = padding.ImVec2()
}

func (s *Style) GetDisplaySafeAreaPadding() Vec2 {

	return s.DisplaySafeAreaPadding.Vec2()
}

func (s *Style) SetDisplaySafeAreaPadding(padding Vec2) {

	s.DisplaySafeAreaPadding = padding.ImVec2()
}

func (s *Style) GetMouseCursorScale() float32 {

	return float32(s.MouseCursorScale)
}

func (s *Style) SetMouseCursorScale(scale float32) {

	s.MouseCursorScale = C.float(scale)
}

func (s *Style) GetAntiAliasedLines() bool {

	return bool(s.AntiAliasedLines)
}

func (s *Style) SetAntiAliasedLines(state bool) {

	s.AntiAliasedLines = C.bool(state)
}

func (s *Style) GetAntiAliasedLinesUseTex() bool {

	return bool(s.AntiAliasedLinesUseTex)
}

func (s *Style) SetAntiAliasedLinesUseTex(state bool) {

	s.AntiAliasedLinesUseTex = C.bool(state)
}

func (s *Style) GetAntiAliasedFill() bool {

	return bool(s.AntiAliasedFill)
}

func (s *Style) SetAntiAliasedFill(state bool) {

	s.AntiAliasedFill = C.bool(state)
}

func (s *Style) GetCurveTessellationTol() float32 {

	return float32(s.CurveTessellationTol)
}

func (s *Style) SetCurveTessellationTol(tol float32) {

	s.CurveTessellationTol = C.float(tol)
}

func (s *Style) GetCircleTessellationMaxError() float32 {

	return float32(s.CircleTessellationMaxError)
}

func (s *Style) SetCircleTessellationMaxError(maxerror float32) {

	s.CircleTessellationMaxError = C.float(maxerror)
}

func GetFont() *Font {

	return (*Font)(C.igGetFont())
}

func GetFontSize() float32 {

	return float32(C.igGetFontSize())
}

func GetFontTexUvWhitePixel() Vec2 {

	var out C.ImVec2
	C.igGetFontTexUvWhitePixel(&out)
	return out.Vec2()
}

func GetColorU32Col(idx Col, alphaMul float32) U32 {

	return U32(C.igGetColorU32_Col(C.ImGuiCol(idx), C.float(alphaMul)))
}

func GetColorU32Vec4(col Vec4) U32 {

	return U32(C.igGetColorU32_Vec4(col.ImVec4()))
}

func GetColorU32(col U32) U32 {

	return U32(C.igGetColorU32_U32(C.ImU32(col)))
}

func GetStyleColorVec4(idx Col) Vec4 {

	return C.igGetStyleColorVec4(C.ImGuiCol(idx)).Vec4()
}

func StyleColorsDark(dst *Style) {

	C.igStyleColorsDark((*C.ImGuiStyle)(dst))
}

func StyleColorsLight(dst *Style) {

	C.igStyleColorsLight((*C.ImGuiStyle)(dst))
}

func StyleColorsClassic(dst *Style) {

	C.igStyleColorsClassic((*C.ImGuiStyle)(dst))
}
