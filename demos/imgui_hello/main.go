package main

import (
	"github.com/leonsal/gig/imgui"
)

func main() {

	// Creates the default display window
	cfg := imgui.DisplayConfig{
		Title:      "ImGui Hello",
		Width:      600,
		Height:     400,
		MSAA:       0,
		Fullscreen: false,
	}
	disp, err := imgui.DisplayInit(&cfg)
	if err != nil {
		panic(err)
	}

	// Render loop
	open := true
	for {
		close := disp.StartFrame()
		if close {
			break
		}

		if open {
			if imgui.Begin("Window", &open, imgui.WindowFlags_None) {
				imgui.Text("Hello, ImGui")
			}
			imgui.End()
		}

		disp.EndFrame()
	}
	disp.Destroy()
}
