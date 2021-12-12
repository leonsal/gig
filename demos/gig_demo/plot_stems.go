package main

import (
	"math"

	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
)

func init() {

	registerDemo(catImPlot, "Stems", NewPlotStemsDemo)
}

func NewPlotStemsDemo() gig.IWidget {

	const count = 51
	xs := make([]float64, count)
	ys1 := make([]float64, count)
	ys2 := make([]float64, count)
	for i := 0; i < count; i++ {
		xs[i] = float64(i) * 0.02
		ys1[i] = 1 + 0.5*math.Sin(25*xs[i])*math.Cos(2*xs[i])
		ys2[i] = 0.5 + 0.25*math.Sin(10*xs[i])*math.Sin(xs[i])
	}

	group := gig.NewGroup().AddChildren(
		gig.NewPlot("Plot stems").AddChildren(
			gig.NewPlotStems("Stems 1").SetValuesXYFloat64(&xs[0], &ys1[0], count),
			gig.NewSetNextLineStyle(imgui.Vec4{1, 0.5, 0, 0.75}, imgui.PlotAuto),
			gig.NewSetNextMarkerStyle(imgui.PlotMarker_Square).SetSize(5).SetFill(imgui.Vec4{1, 0.5, 0, 0.25}),
			gig.NewPlotStems("Stems 2").SetValuesXYFloat64(&xs[0], &ys2[0], count),
		),
	)
	return group
}
