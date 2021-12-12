package main

import (
	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
)

func init() {

	registerDemo(catImGui, "Custom rendering", NewCustomRenderingDemo)
}

func NewCustomRenderingDemo() gig.IWidget {

	text := gig.NewText("Gradients")
	ivb1 := gig.NewInvisibleButton(imgui.Vec2{0, 0})
	ivb2 := gig.NewInvisibleButton(imgui.Vec2{0, 0})
	group := gig.NewGroup().SetPreRender(func(w gig.IWidget) {

		text.Render()

		dl := imgui.GetWindowDrawList()
		gradientSize := imgui.Vec2{imgui.CalcItemWidth(), imgui.GetFrameHeight()}

		p0 := imgui.GetCursorScreenPos()
		p1 := imgui.Vec2{p0.X + gradientSize.X, p0.Y + gradientSize.Y}
		col_a := imgui.GetColorU32Vec4(imgui.Vec4{0, 0, 0, 255})
		col_b := imgui.GetColorU32Vec4(imgui.Vec4{255, 255, 255, 255})
		dl.AddRectFilledMultiColor(p0, p1, col_a, col_b, col_b, col_a)
		ivb1.SetSize(gradientSize).Render()

		p0 = imgui.GetCursorScreenPos()
		p1 = imgui.Vec2{p0.X + gradientSize.X, p0.Y + gradientSize.Y}
		col_a = imgui.GetColorU32Vec4(imgui.Vec4{0, 255, 0, 255})
		col_b = imgui.GetColorU32Vec4(imgui.Vec4{255, 0, 0, 255})
		dl.AddRectFilledMultiColor(p0, p1, col_a, col_b, col_b, col_a)
		ivb2.SetSize(gradientSize).Render()
	})

	return group
}
