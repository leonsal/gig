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

type DisplayConfig struct {
	Title      string
	Width      int
	Height     int
	MSAA       int
	Fullscreen bool
	Opengl     OpenglConfig
}

type OpenglConfig struct {
	ES bool
}

type Display struct {
	di C.displayImpl
}

// DisplayInit creates and returns a native window display
func DisplayInit(cfg *DisplayConfig) (Display, error) {

	var cc C.display_config_t
	if cfg == nil {
		cc.title = C.CString("")
		cc.width = 800
		cc.height = 600
		cc.msaa = 0
		cc.fullscreen = false
		cc.opengl.es = false
	} else {
		cc.title = C.CString(cfg.Title)
		cc.width = C.int(cfg.Width)
		cc.height = C.int(cfg.Height)
		cc.msaa = C.int(cfg.MSAA)
		cc.fullscreen = C.bool(cfg.Fullscreen)
		cc.opengl.es = C.bool(cfg.Opengl.ES)
	}
	defer C.free(unsafe.Pointer(cc.title))
	var err C.int

	d := C.display_init(&cc, &err)
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

func (d Display) SetShouldClose(close bool) {

	C.display_set_should_close(d.di, C.bool(close))
}

func (d Display) CreateTexture() TextureID {

	return TextureID(C.display_create_texture(d.di))
}

func (d Display) DeleteTexture(tid TextureID) {

	C.display_delete_texture(d.di, C.ImTextureID(tid))
}

func (d Display) TransferTexture(tid TextureID, width, height int, pix []color.RGBA) {

	C.display_transfer_texture(d.di, C.ImTextureID(tid), C.int(width), C.int(height), unsafe.Pointer(&pix[0].R))
}
