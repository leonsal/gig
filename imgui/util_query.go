package imgui

// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"

func IsItemHovered(flags HoveredFlags) bool {

	return bool(C.igIsItemHovered(C.ImGuiHoveredFlags(flags)))
}

func IsItemActive() bool {

	return bool(C.igIsItemActive())
}

func IsItemFocused() bool {

	return bool(C.igIsItemFocused())
}

func IsItemClicked(mouseButton MouseButton) bool {

	return bool(C.igIsItemClicked(C.ImGuiMouseButton(mouseButton)))
}

func IsItemVisible() bool {

	return bool(C.igIsItemVisible())
}

func IsItemEdited() bool {

	return bool(C.igIsItemEdited())
}

func IsItemActivated() bool {

	return bool(C.igIsItemActivated())
}

func IsItemDeactivated() bool {

	return bool(C.igIsItemDeactivated())
}

func IsItemDeactivatedAfterEdit() bool {

	return bool(C.igIsItemDeactivatedAfterEdit())
}

func IsItemToggledOpen() bool {

	return bool(C.igIsItemToggledOpen())
}

func IsAnyItemHovered() bool {

	return bool(C.igIsAnyItemHovered())
}

func IsAnyItemActive() bool {

	return bool(C.igIsAnyItemActive())
}

func IsAnyItemFocused() bool {

	return bool(C.igIsAnyItemFocused())
}

func GetItemRectMin() Vec2 {

	var out C.ImVec2
	C.igGetItemRectMin(&out)
	return out.Vec2()
}

func GetItemRectMax() Vec2 {

	var out C.ImVec2
	C.igGetItemRectMax(&out)
	return out.Vec2()
}

func GetItemRectSize() Vec2 {

	var out C.ImVec2
	C.igGetItemRectSize(&out)
	return out.Vec2()
}

func SetItemAllowOverlap() {

	C.igSetItemAllowOverlap()
}
