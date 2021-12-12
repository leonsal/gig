package main

import (
	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
)

func init() {

	registerDemo(catImGui, "Color", NewColorDemo)
}

func NewColorDemo() gig.IWidget {

	group := gig.NewGroup().AddChildren(

		gig.NewText("ColorEdit:"),
		gig.NewColorEdit3("ColorEdit", imgui.Vec4{1, 0, 0, 1}).SetOnChange(func(c *gig.ColorEdit3) {
			log.Info("ColorEdit3: ", c.Color())
		}),
		gig.NewColorEdit4("ColorEdit4", imgui.Vec4{0, 1, 1, 0.5}).SetOnChange(func(c *gig.ColorEdit4) {
			log.Info("ColorEdit4: ", c.Color())
		}),

		gig.NewText("ColorPicker:"),
		gig.NewColorPicker3("ColorPicker3", imgui.Vec4{0, 0, 1, 1}).SetOnChange(func(c *gig.ColorPicker3) {
			log.Info("ColorPicker3: ", c.Color())
		}),
		gig.NewColorPicker4("ColorPicker4", imgui.Vec4{0, 1, 0, 1}).SetOnChange(func(c *gig.ColorPicker4) {
			log.Info("ColorPicker4: ", c.Color())
		}),

		gig.NewText("ColorButton:"),
		gig.NewColorButton("ColorButton", imgui.Vec4{1, 1, 0, 1}).SetOnClick(func(b *gig.ColorButton) {
			log.Info("ColorButton clicked", b.Color())
		}),
	)
	return group
}
