package main

import (
	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
)

func init() {

	registerDemo(catImGui, "Button", NewButtonDemo)
}

func NewButtonDemo() gig.IWidget {

	group := gig.NewGroup().AddChildren(

		gig.NewText("Button:"),
		gig.NewButton("Button1").SetOnClick(func(b *gig.Button) {
			log.Info("Button1 clicked")
		}),

		gig.NewSameLine(),
		gig.NewButton("Button2").SetSize(imgui.Vec2{120, 0}).SetOnClick(func(b *gig.Button) {
			log.Info("Button2 clicked")
		}),

		gig.NewSameLine(),
		gig.NewButton("Button3 with repeat").SetSize(imgui.Vec2{-1, 0}).SetRepeat(true).SetOnClick(func(b *gig.Button) {
			log.Info("Button3 clicked")
		}),

		gig.NewButton("Button4").SetSize(imgui.Vec2{0, 40}).SetOnClick(func(b *gig.Button) {
			log.Info("Button4 clicked")
		}),
		gig.NewSameLine(),
		gig.NewButton("Button5").SetSize(imgui.Vec2{120, 40}).SetOnClick(func(b *gig.Button) {
			log.Info("Button5 clicked")
		}),
		gig.NewSameLine(),
		gig.NewButton("Button6 disabled").SetSize(imgui.Vec2{-1, 40}).SetOnClick(func(b *gig.Button) {
			log.Info("Button6 clicked")
		}).
			SetDisabled(true),

		gig.NewSmallButton("Small1").SetOnClick(func(b *gig.SmallButton) {
			log.Info("Small1 clicked")
		}),
		gig.NewSameLine(),
		gig.NewSmallButton("Small2").SetOnClick(func(b *gig.SmallButton) {
			log.Info("Small2 clicked")
		}),

		gig.NewArrowButton(imgui.Dir_Down),
		gig.NewSameLine(),
		gig.NewArrowButton(imgui.Dir_Left),
		gig.NewSameLine(),
		gig.NewArrowButton(imgui.Dir_Right),
		gig.NewSameLine(),
		gig.NewArrowButton(imgui.Dir_Up),

		gig.NewText("Checkbox:"),
		gig.NewCheckbox("Check1").SetOnClick(func(c *gig.Checkbox) {
			log.Info("Check1 clicked: ", c.Value())
		}),
		gig.NewSameLine(),
		gig.NewCheckbox("Check2").SetOnClick(func(c *gig.Checkbox) {
			log.Info("Check2 clicked: ", c.Value())
		}),
		gig.NewSameLine(),
		gig.NewCheckbox("Check3").SetOnClick(func(c *gig.Checkbox) {
			log.Info("Check3 clicked: ", c.Value())
		}),

		gig.NewText("RadioButton:"),
		gig.NewRadioGroup(
			gig.NewRadioButton("Radio1"),
			gig.NewSameLine(),
			gig.NewRadioButton("Radio2"),
			gig.NewSameLine(),
			gig.NewRadioButton("Radio3"),
		).SetOnChange(func(rg *gig.RadioGroup) {
			idx, rb := rg.Active()
			log.Infof("RadioGroup changed: %v/%v", idx, rb.Label())
		}),

		gig.NewText("ProgressBar:"),
		gig.NewProgressBar(0.5),
		gig.NewProgressBar(0.2).SetSize(imgui.Vec2{-1, 40}).SetOverlay(""),
		gig.NewText("Bullet:"),
		gig.NewBullet(),
		gig.NewSameLine(),
		gig.NewBullet(),
		gig.NewSameLine(),
		gig.NewBullet(),
	)

	return group
}
