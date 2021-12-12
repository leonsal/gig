package main

import (
	"log"

	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
)

func NewInputDemo() gig.IWidget {

	win := gig.NewWindow("Input").AddChildren(

		gig.NewText("InputText:"),
		gig.NewInputText("Label").
			SetMaxLen(40).
			SetOnChange(func(it *gig.InputText) {
				log.Println("InputText:", it.Text())
			}).
			SetFlags(imgui.InputTextFlags_CallbackCompletion),
		//SetCallback(func(cd *imgui.InputTextCallbackData) int {
		//	log.Printf("callback:%+v\n", cd)
		//	return 0
		//}),

		gig.NewText("InputText with hint:"),
		gig.NewInputText("Label").
			SetHint("Hint").
			SetOnChange(func(it *gig.InputText) {
				log.Println("InputText with hint:", it.Text())
			}),

		gig.NewText("InputTextMultiline:"),
		gig.NewInputTextMultiline("Label").
			SetOnChange(func(it *gig.InputTextMultiline) {
				log.Println("InputTextMultiline:", it.Text())
			}),

		gig.NewSeparator(),
		gig.NewText("InputInt32:"),
		gig.NewInputInt32("Label", 1).
			SetStep(0).
			SetOnChange(func(it *gig.InputInt32) {
				log.Println("InputInt:", it.Value())
			}),

		gig.NewText("InputInt32Slice (1):"),
		gig.NewInputInt32Slice("Label", []int32{1}).
			SetOnChange(func(it *gig.InputInt32Slice) {
				log.Println("InputInt32Slice (1):", it.Values())
			}),

		gig.NewText("InputInt32Slice (2):"),
		gig.NewInputInt32Slice("Label", []int32{1, 2}).
			SetStep(0).
			SetOnChange(func(it *gig.InputInt32Slice) {
				log.Println("InputInt32Slice (2):", it.Values())
			}),

		gig.NewText("InputInt32Slice (3):"),
		gig.NewInputInt32Slice("Label", []int32{1, 2, 3}).
			SetStep(0).
			SetOnChange(func(it *gig.InputInt32Slice) {
				log.Println("InputInt32Slice (3):", it.Values())
			}),

		gig.NewText("InputInt32Slice (4):"),
		gig.NewInputInt32Slice("Label", []int32{1, 2, 3, 4}).
			SetStep(0).
			SetOnChange(func(it *gig.InputInt32Slice) {
				log.Println("InputInt32Slice (4):", it.Values())
			}),

		gig.NewSeparator(),
		gig.NewText("InputInt64:"),
		gig.NewInputInt64("Label", 10_000_000_000).
			SetStep(10.0).
			SetOnChange(func(it *gig.InputInt64) {
				log.Println("InputInt64:", it.Value())
			}),

		gig.NewText("InputInt64Slice (1):"),
		gig.NewInputInt64Slice("Label", []int64{1}).
			SetStep(0).
			SetOnChange(func(it *gig.InputInt64Slice) {
				log.Println("InputInt64Slice (1):", it.Values())
			}),

		gig.NewText("InputInt64Slice (2):"),
		gig.NewInputInt64Slice("Label", []int64{1, 2}).
			SetStep(0).
			SetOnChange(func(it *gig.InputInt64Slice) {
				log.Println("InputInt64Slice (2):", it.Values())
			}),

		gig.NewText("InputInt64Slice (3):"),
		gig.NewInputInt64Slice("Label", []int64{1, 2, 3}).
			SetStep(0).
			SetOnChange(func(it *gig.InputInt64Slice) {
				log.Println("InputInt64Slice (3):", it.Values())
			}),

		gig.NewText("InputInt64Slice (4):"),
		gig.NewInputInt64Slice("Label", []int64{1, 2, 3, 4}).
			SetStep(0).
			SetOnChange(func(it *gig.InputInt64Slice) {
				log.Println("InputInt64Slice (4):", it.Values())
			}),

		gig.NewSeparator(),
		gig.NewText("InputFloat32:"),
		gig.NewInputFloat32("Label", 1.0).
			SetOnChange(func(it *gig.InputFloat32) {
				log.Println("InputFloat32:", it.Value())
			}),

		gig.NewText("InputFloat32 step/format:"),
		gig.NewInputFloat32("Label", 2.0).
			SetStep(10.0).
			SetFormat("US$%6.2f").
			SetOnChange(func(it *gig.InputFloat32) {
				log.Println("InputFloat32 step/format:", it.Value())
			}),

		gig.NewText("InputFloat32Slice (2):"),
		gig.NewInputFloat32Slice("Label", []float32{1.1, 2.2}).
			SetOnChange(func(it *gig.InputFloat32Slice) {
				log.Println("InputFloat32Slice (2):", it.Values())
			}),

		gig.NewText("InputFloat32Slice (3):"),
		gig.NewInputFloat32Slice("Label", []float32{1.1, 2.2, 3.3}).
			SetOnChange(func(it *gig.InputFloat32Slice) {
				log.Println("InputFloat32Slice (3):", it.Values())
			}),

		gig.NewText("InputFloat32Slice (4):"),
		gig.NewInputFloat32Slice("Label", []float32{1.1, 2.2, 3.3, 4.4}).
			SetOnChange(func(it *gig.InputFloat32Slice) {
				log.Println("InputFloat32Slice (4):", it.Values())
			}),

		gig.NewSeparator(),
		gig.NewText("InputFloat64:"),
		gig.NewInputFloat64("Label", 12_345_678_900).
			SetOnChange(func(it *gig.InputFloat64) {
				log.Println("InputFloat64", it.Value())
			}),

		gig.NewText("InputFloat64Slice (2):"),
		gig.NewInputFloat64Slice("Label", []float64{1.1, 2.2}).
			SetOnChange(func(it *gig.InputFloat64Slice) {
				log.Println("InputFloat64Slice (2):", it.Values())
			}),

		gig.NewText("InputFloat64Slice (3):"),
		gig.NewInputFloat64Slice("Label", []float64{1.1, 2.2, 3.3}).
			SetOnChange(func(it *gig.InputFloat64Slice) {
				log.Println("InputFloat64Slice (3):", it.Values())
			}),

		gig.NewText("InputFloat64Slice (4):"),
		gig.NewInputFloat64Slice("Label", []float64{1.1, 2.2, 3.3, 4.4}).
			SetOnChange(func(it *gig.InputFloat64Slice) {
				log.Println("InputFloat64Slice (4):", it.Values())
			}),
	).SetVisible(false).SaveRef("inputDemo")

	return win
}
