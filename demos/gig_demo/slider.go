package main

import (
	"math"

	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
)

func init() {

	registerDemo(catImGui, "Slider", NewSliderDemo)
}

func NewSliderDemo() gig.IWidget {

	group := gig.NewGroup().AddChildren(

		gig.NewText("SliderInt32:"),
		gig.NewSliderInt32("SliderInt32", 1, 0, 100).SetOnChange(func(s *gig.SliderInt32) {
			log.Info("SliderInt32:", s.Value())
		}),

		gig.NewSliderInt32Slice("SliderInt32Slice (1)", []int32{1}, 0, 10).SetOnChange(func(s *gig.SliderInt32Slice) {
			log.Info("SliderInt32Slice (1):", s.Values())
		}),

		gig.NewSliderInt32Slice("SliderInt32Slice (2)", []int32{1, 2}, 0, 10).SetOnChange(func(s *gig.SliderInt32Slice) {
			log.Info("SliderInt32Slice (2):", s.Values())
		}),

		gig.NewSliderInt32Slice("SliderInt32Slice (3)", []int32{1, 2, 3}, 200, 300).SetOnChange(func(s *gig.SliderInt32Slice) {
			log.Info("SliderInt32Slice (3):", s.Values())
		}),

		gig.NewSliderInt32Slice("SliderInt32Slice (4)", []int32{1, 2, 3, 4}, -300, 400).SetOnChange(func(s *gig.SliderInt32Slice) {
			log.Info("SliderInt32Slice (4):", s.Values())
		}),

		gig.NewSeparator(),
		gig.NewText("SliderInt64:"),
		gig.NewSliderInt64("SliderInt64", 1, -10_000_000_000, 10_000_000_000).SetOnChange(func(s *gig.SliderInt64) {
			log.Info("SliderInt64:", s.Value())
		}),
		gig.NewSliderInt64Slice("SliderInt64Slice (1)", []int64{1}, 0, 10).SetOnChange(func(s *gig.SliderInt64Slice) {
			log.Info("SliderInt64Slice (1):", s.Values())
		}),

		gig.NewSliderInt64Slice("SliderInt64Slice (2)", []int64{1, 2}, 0, 10).SetOnChange(func(s *gig.SliderInt64Slice) {
			log.Info("SliderInt64Slice (2):", s.Values())
		}),

		gig.NewSliderInt64Slice("SliderInt64Slice (3)", []int64{1, 2, 3}, 200, 300).SetOnChange(func(s *gig.SliderInt64Slice) {
			log.Info("SliderInt64Slice (3):", s.Values())
		}),

		gig.NewSliderInt64Slice("SliderInt64Slice (4)", []int64{1, 2, 3, 4}, -300, 400).SetOnChange(func(s *gig.SliderInt64Slice) {
			log.Info("SliderInt64Slice (4):", s.Values())
		}),

		gig.NewSeparator(),
		gig.NewText("SliderFloat32:"),
		gig.NewSliderFloat32("SliderFloat32", 1, -100.0, 200.0).SetOnChange(func(s *gig.SliderFloat32) {
			log.Info("SliderFloat32:", s.Value())
		}),
		gig.NewSliderFloat32("SliderFloat32", 1, -1_000_000_000, 1_000_000_000).SetOnChange(func(s *gig.SliderFloat32) {
			log.Info("SliderFloat32:", s.Value())
		}),
		gig.NewSliderFloat32Slice("SliderFloat32Slice (1)", []float32{1}, 0, 10).SetOnChange(func(s *gig.SliderFloat32Slice) {
			log.Info("SliderFloat32Slice (1):", s.Values())
		}),

		gig.NewSliderFloat32Slice("SliderFloat32Slice (2)", []float32{1, 2}, 0, 10).SetOnChange(func(s *gig.SliderFloat32Slice) {
			log.Info("SliderFloat32Slice (2):", s.Values())
		}),

		gig.NewSliderFloat32Slice("SliderFloat32Slice (3)", []float32{1, 2, 3}, 200, 300).SetOnChange(func(s *gig.SliderFloat32Slice) {
			log.Info("SliderFloat32Slice (3):", s.Values())
		}),

		gig.NewSliderFloat32Slice("SliderFloat32Slice (4)", []float32{1, 2, 3, 4}, -300, 400).SetOnChange(func(s *gig.SliderFloat32Slice) {
			log.Info("SliderFloat32Slice (4):", s.Values())
		}),

		gig.NewSeparator(),
		gig.NewText("SliderFloat64:"),
		gig.NewSliderFloat64("SliderFloat64", 1, -100.0, 200.0).SetOnChange(func(s *gig.SliderFloat64) {
			log.Info("SliderFloat64:", s.Value())
		}),
		gig.NewSliderFloat64("SliderFloat64", 1, -1_000_000_000, 1_000_000_000).SetOnChange(func(s *gig.SliderFloat64) {
			log.Info("SliderFloat64:", s.Value())
		}),
		gig.NewSliderFloat64Slice("SliderFloat64Slice (1)", []float64{1}, 0, 10).SetOnChange(func(s *gig.SliderFloat64Slice) {
			log.Info("SliderFloat64Slice (1):", s.Values())
		}),

		gig.NewSliderFloat64Slice("SliderFloat64Slice (2)", []float64{1, 2}, 0, 10).SetOnChange(func(s *gig.SliderFloat64Slice) {
			log.Info("SliderFloat64Slice (2):", s.Values())
		}),

		gig.NewSliderFloat64Slice("SliderFloat64Slice (3)", []float64{1, 2, 3}, 200, 300).SetOnChange(func(s *gig.SliderFloat64Slice) {
			log.Info("SliderFloat64Slice (3):", s.Values())
		}),

		gig.NewSliderFloat64Slice("SliderFloat64Slice (4)", []float64{1, 2, 3, 4}, -300, 400).SetOnChange(func(s *gig.SliderFloat64Slice) {
			log.Info("SliderFloat64Slice (4):", s.Values())
		}),

		gig.NewSeparator(),
		gig.NewText("SliderAngle:"),
		gig.NewSliderAngle("SliderAngle", math.Pi/2).SetOnChange(func(s *gig.SliderAngle) {
			log.Info("SliderAngle():", s.Value())
		}),

		gig.NewSeparator(),
		gig.NewText("VSlider:"),
		gig.NewVSliderInt32("Int32", imgui.Vec2{100, 200}, 0, -1_000_000_000, 1_000_000_000).SetOnChange(func(s *gig.VSliderInt32) {
			log.Info("VSliderInt32():", s.Value())
		}),
		gig.NewSameLine(),
		gig.NewVSliderInt64("Int64", imgui.Vec2{100, 200}, 0, -10_123_456_789, 10_123_456_789).SetOnChange(func(s *gig.VSliderInt64) {
			log.Info("VSliderInt64():", s.Value())
		}),
		gig.NewSameLine(),
		gig.NewVSliderFloat32("Float32", imgui.Vec2{140, 200}, 0, -10_123_456_789, 10_123_456_789).SetOnChange(func(s *gig.VSliderFloat32) {
			log.Info("VSliderFloat32():", s.Value())
		}),
		gig.NewSameLine(),
		gig.NewVSliderFloat64("Float64", imgui.Vec2{160, 200}, 0, -100_123_456_789, 100_123_456_789).SetOnChange(func(s *gig.VSliderFloat64) {
			log.Info("VSliderFloat64():", s.Value())
		}),
	)
	return group
}
