package imgui

// #include <stdlib.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimgui.h"
// #include "display.h"
import "C"
import (
	"fmt"
	"image/color"
	"unsafe"
)

// DisplayConfig describes the display (native window) configuration
type DisplayConfig struct {
	Evtimeout  float64
	Fullscreen bool
	Opengl     OpenglConfig
}

type OpenglConfig struct {
	ES   bool
	MSAA int
}

type Display struct {
	di C.displayImpl
}

// DisplayInit creates and returns a native window display
func DisplayInit(title string, width, height int, cfg *DisplayConfig) (Display, error) {

	var cc C.display_config_t
	if cfg == nil {
		cc.ev_timeout = 0
		cc.fullscreen = false
		cc.opengl.es = false
		cc.opengl.msaa = 0
	} else {
		cc.ev_timeout = C.double(cfg.Evtimeout)
		cc.fullscreen = C.bool(cfg.Fullscreen)
		cc.opengl.es = C.bool(cfg.Opengl.ES)
		cc.opengl.msaa = C.int(0)
	}

	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	var err C.int
	d := C.display_init(ctitle, C.int(width), C.int(height), &cc, &err)
	if err != 0 {
		return Display{}, fmt.Errorf("Error:%d returned by backend", err)
	}
	return Display{d}, nil
}

// Destroy destroys the display
func (d Display) Destroy() {

	C.display_destroy(d.di)
	d.di = nil
}

func (d Display) UploadFonts() {

	C.display_upload_fonts(d.di)
}

// Size returns the current size of the display window
func (d Display) Size() (width int, height int) {

	var cwidth C.int
	var cheight C.int
	C.display_size(d.di, &cwidth, &cheight)
	return int(cwidth), int(cheight)
}

// SetTitle sets the title of the display window
func (d Display) SetTitle(title string) {

	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	C.display_set_title(d.di, ctitle)
}

// StartFrame starts a new rendering frame
func (d Display) StartFrame() bool {

	return bool(C.display_start_frame(d.di))
}

// EndFrame ends a previously started rendering frame
func (d Display) EndFrame() {

	C.display_end_frame(d.di)
}

// Returns if the native window should be closed
func (d Display) SetShouldClose(close bool) {

	C.display_set_should_close(d.di, C.bool(close))
}

// Creates a texture for the current backend and return its ImGui ID
func (d Display) CreateTexture() TextureID {

	return TextureID(C.display_create_texture(d.di))
}

// Deletes a previously created texture by its ID
func (d Display) DeleteTexture(tid TextureID) {

	C.display_delete_texture(d.di, C.ImTextureID(tid))
}

// Transfer texture with the specified id, width, height and pixels to the backend.
func (d Display) TransferTexture(tid TextureID, width, height int, pix []color.RGBA) {

	C.display_transfer_texture(d.di, C.ImTextureID(tid), C.int(width), C.int(height), unsafe.Pointer(&pix[0].R))
}
