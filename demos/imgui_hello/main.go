package main

import (
	"github.com/leonsal/gig/imgui"
)

func main() {

	disp, err := imgui.DisplayInit("ImGui Hello", 600, 400, nil)
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
