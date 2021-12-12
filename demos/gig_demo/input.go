package main

import (
	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
)

func init() {

	registerDemo(catImGui, "Input", NewInputDemo)
}

func NewInputDemo() gig.IWidget {

	group := gig.NewGroup().AddChildren(

		gig.NewInputText("InputText").
			SetMaxLen(40).
			SetOnChange(func(it *gig.InputText) {
				log.Info("InputText:", it.Text())
			}).
			SetFlags(imgui.InputTextFlags_CallbackCompletion),
		//SetCallback(func(cd *imgui.InputTextCallbackData) int {
		//	log.Printf("callback:%+v\n", cd)
		//	return 0
		//}),

		gig.NewInputText("InputTextWithHint").
			SetHint("Hint").
			SetOnChange(func(it *gig.InputText) {
				log.Info("InputText with hint:", it.Text())
			}),

		gig.NewInputTextMultiline("InputTextMultiline").
			SetOnChange(func(it *gig.InputTextMultiline) {
				log.Info("InputTextMultiline:", it.Text())
			}),

		gig.NewSeparator(),
		gig.NewText("InputInt32:"),
		gig.NewInputInt32("InputInt32", 1).
			SetStep(0).
			SetOnChange(func(it *gig.InputInt32) {
				log.Info("InputInt:", it.Value())
			}),

		gig.NewInputInt32Slice("InputInt32Slice (1)", []int32{1}).
			SetOnChange(func(it *gig.InputInt32Slice) {
				log.Info("InputInt32Slice (1):", it.Values())
			}),

		gig.NewInputInt32Slice("InputInt32Slice (2)", []int32{1, 2}).
			SetStep(0).
			SetOnChange(func(it *gig.InputInt32Slice) {
				log.Info("InputInt32Slice (2):", it.Values())
			}),

		gig.NewInputInt32Slice("InputInt32Slice (3)", []int32{1, 2, 3}).
			SetStep(0).
			SetOnChange(func(it *gig.InputInt32Slice) {
				log.Info("InputInt32Slice (3):", it.Values())
			}),

		gig.NewInputInt32Slice("InputInt32Slice (4)", []int32{1, 2, 3, 4}).
			SetStep(0).
			SetOnChange(func(it *gig.InputInt32Slice) {
				log.Info("InputInt32Slice (4):", it.Values())
			}),

		gig.NewSeparator(),
		gig.NewInputInt64("InputInt64", 10_000_000_000).
			SetStep(10.0).
			SetOnChange(func(it *gig.InputInt64) {
				log.Info("InputInt64:", it.Value())
			}),

		gig.NewInputInt64Slice("InputInt64Slice (1)", []int64{1}).
			SetStep(0).
			SetOnChange(func(it *gig.InputInt64Slice) {
				log.Info("InputInt64Slice (1):", it.Values())
			}),

		gig.NewInputInt64Slice("InputInt64Slice (2)", []int64{1, 2}).
			SetStep(0).
			SetOnChange(func(it *gig.InputInt64Slice) {
				log.Info("InputInt64Slice (2):", it.Values())
			}),

		gig.NewInputInt64Slice("InputInt64Slice (3)", []int64{1, 2, 3}).
			SetStep(0).
			SetOnChange(func(it *gig.InputInt64Slice) {
				log.Info("InputInt64Slice (3):", it.Values())
			}),

		gig.NewInputInt64Slice("InputInt64Slice (4)", []int64{1, 2, 3, 4}).
			SetStep(0).
			SetOnChange(func(it *gig.InputInt64Slice) {
				log.Info("InputInt64Slice (4):", it.Values())
			}),

		gig.NewSeparator(),
		gig.NewInputFloat32("InputFloat32", 1.0).
			SetOnChange(func(it *gig.InputFloat32) {
				log.Info("InputFloat32:", it.Value())
			}),

		gig.NewInputFloat32("InputFloat32 (step/format)", 2.0).
			SetStep(10.0).
			SetFormat("US$%6.2f").
			SetOnChange(func(it *gig.InputFloat32) {
				log.Info("InputFloat32 (step/format):", it.Value())
			}),

		gig.NewInputFloat32Slice("InputFloat32Slice (2)", []float32{1.1, 2.2}).
			SetOnChange(func(it *gig.InputFloat32Slice) {
				log.Info("InputFloat32Slice (2):", it.Values())
			}),

		gig.NewInputFloat32Slice("InputFloat32Slice (3)", []float32{1.1, 2.2, 3.3}).
			SetOnChange(func(it *gig.InputFloat32Slice) {
				log.Info("InputFloat32Slice (3):", it.Values())
			}),

		gig.NewInputFloat32Slice("InputFloat32Slice (4)", []float32{1.1, 2.2, 3.3, 4.4}).
			SetOnChange(func(it *gig.InputFloat32Slice) {
				log.Info("InputFloat32Slice (4):", it.Values())
			}),

		gig.NewSeparator(),
		gig.NewInputFloat64("InputFloat64", 12_345_678_900).
			SetOnChange(func(it *gig.InputFloat64) {
				log.Info("InputFloat64", it.Value())
			}),

		gig.NewInputFloat64Slice("InputFloat64Slice (2)", []float64{1.1, 2.2}).
			SetOnChange(func(it *gig.InputFloat64Slice) {
				log.Info("InputFloat64Slice (2):", it.Values())
			}),

		gig.NewInputFloat64Slice("InputFloat64Slice (3)", []float64{1.1, 2.2, 3.3}).
			SetOnChange(func(it *gig.InputFloat64Slice) {
				log.Info("InputFloat64Slice (3):", it.Values())
			}),

		gig.NewInputFloat64Slice("InputFloat64Slice (4)", []float64{1.1, 2.2, 3.3, 4.4}).
			SetOnChange(func(it *gig.InputFloat64Slice) {
				log.Info("InputFloat64Slice (4):", it.Values())
			}),
	)
	return group
}
