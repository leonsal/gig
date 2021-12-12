package main

import (
	"github.com/leonsal/gig"
)

func init() {

	registerDemo(catImGui, "Drag", NewDragDemo)
}

func NewDragDemo() gig.IWidget {

	group := gig.NewGroup().AddChildren(

		gig.NewText("DragInt32:"),
		gig.NewDragInt32("DragInt32", 1).
			SetOnChange(func(d *gig.DragInt32) {
				log.Info("DragInt32:", d.Value())
			}),

		gig.NewDragInt32Slice("DragInt32Slice (2) (-100/100)", []int32{1, 2}).
			SetOnChange(func(d *gig.DragInt32Slice) {
				log.Info("DragInt32Slice (2):", d.Values())
			}),

		gig.NewDragInt32Slice("DragInt32Slice (3) (-200/300)", []int32{1, 2, 3}).
			SetMin(-200).SetMax(300).
			SetOnChange(func(d *gig.DragInt32Slice) {
				log.Info("DragInt32Slice (3):", d.Values())
			}),

		gig.NewDragInt32Slice("DragInt32Slice (4) (-300/400)", []int32{1, 2, 3, 4}).
			SetMin(-300).SetMax(400).
			SetOnChange(func(d *gig.DragInt32Slice) {
				log.Info("DragInt32Slice (4):", d.Values())
			}),

		gig.NewSeparator(),
		gig.NewText("DragInt64:"),
		gig.NewDragInt64("DragInt64", 10_123_456_789).
			SetOnChange(func(d *gig.DragInt64) {
				log.Info("DragInt64:", d.Value())
			}),

		gig.NewDragInt64Slice("DragInt64Slice (2)", []int64{1, 2}).
			SetMin(-10_000_000_000).SetMax(10_000_000_000).
			SetOnChange(func(d *gig.DragInt64Slice) {
				log.Info("DragInt64Slice (2):", d.Values())
			}),

		gig.NewDragInt64Slice("DragInt64Slice (3)", []int64{1, 2, 3}).
			SetMin(-20_000_000_000).SetMax(20_000_000_000).
			SetOnChange(func(d *gig.DragInt64Slice) {
				log.Info("DragInt64Slice (3):", d.Values())
			}),

		gig.NewDragInt64Slice("DragInt64Slice (4)", []int64{1, 2, 3, 4}).
			SetMin(-30_000_000_000).SetMax(30_000_000_000).
			SetOnChange(func(d *gig.DragInt64Slice) {
				log.Info("DragInt64Slice (4):", d.Values())
			}),

		gig.NewDragIntRange2("DragIntRange2", 0, 100).
			SetOnChange(func(d *gig.DragInt32Range2) {
				log.Info("DragInt32Range2:", d.CurrentMin(), d.CurrentMax())
			}),

		gig.NewSeparator(),
		gig.NewText("DragFloat32:"),
		gig.NewDragFloat32("DragFloat32", 1).
			SetOnChange(func(d *gig.DragFloat32) {
				log.Info("DragFloat32:", d.Value())
			}),

		gig.NewDragFloat32Slice("DragFloat32Slice (2)", []float32{1.1, 2.2}).
			SetOnChange(func(d *gig.DragFloat32Slice) {
				log.Info("DragFloat32Slice (2):", d.Values())
			}),

		gig.NewDragFloat32Slice("DragFloat32Slice(3)", []float32{1.1, 2.2, 3.3}).
			SetOnChange(func(d *gig.DragFloat32Slice) {
				log.Info("DragFloat32Slice (3):", d.Values())
			}),

		gig.NewDragFloat32Slice("DragFloat32Slice (3)", []float32{1.1, 2.2, 3.3, 4.4}).
			SetOnChange(func(d *gig.DragFloat32Slice) {
				log.Info("DragFloat32Slice (4):", d.Values())
			}),

		gig.NewDragFloat32Range2("DragFloa32Range2 (0-200)", 0, 100).
			SetMin(0).SetMax(200).
			SetOnChange(func(d *gig.DragFloat32Range2) {
				log.Info("DragFloat32Range2:", d.CurrentMin(), d.CurrentMax())
			}),

		gig.NewSeparator(),
		gig.NewText("DragFloat64:"),
		gig.NewDragFloat64("DragFloat64", 1).
			SetOnChange(func(d *gig.DragFloat64) {
				log.Info("DragFloat64:", d.Value())
			}),

		gig.NewDragFloat64Slice("DragFloat64Slice (2)", []float64{1.1, 2.2}).
			SetOnChange(func(d *gig.DragFloat64Slice) {
				log.Info("DragFloat64Slice (2):", d.Values())
			}),

		gig.NewDragFloat64Slice("DragFloa64Slice (3)", []float64{1.1, 2.2, 3.3}).
			SetOnChange(func(d *gig.DragFloat64Slice) {
				log.Info("DragFloat64Slice (3):", d.Values())
			}),

		gig.NewDragFloat64Slice("DragFloat64Slice (4)", []float64{1.1, 2.2, 3.3, 4.4}).
			SetOnChange(func(d *gig.DragFloat64Slice) {
				log.Info("DragFloat64Slice (4):", d.Values())
			}),
	)
	return group
}
