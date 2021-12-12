package imgui

// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
import "C"

func GetKeyIndex(imguiKey Key) int {

	return int(C.igGetKeyIndex(C.ImGuiKey(imguiKey)))
}

func IsKeyDown(userKeyIndex int) bool {

	return bool(C.igIsKeyDown(C.int(userKeyIndex)))
}

func IsKeyPressed(userKeyIndex int, repeat bool) bool {

	return bool(C.igIsKeyPressed(C.int(userKeyIndex), C.bool(repeat)))
}

func IsKeyReleased(userKeyIndex int) bool {

	return bool(C.igIsKeyReleased(C.int(userKeyIndex)))
}

func IsKeyPressedAmount(keyIndex int, repeatDelay, rate float32) int {

	return int(C.igGetKeyPressedAmount(C.int(keyIndex), C.float(repeatDelay), C.float(rate)))
}

func CaptureKeyboardFromApp(wantCaptureKeyboardValue bool) {

	C.igCaptureKeyboardFromApp(C.bool(wantCaptureKeyboardValue))
}
