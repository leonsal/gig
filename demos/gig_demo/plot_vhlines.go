package main

import (
	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
)

func init() {

	registerDemo(catImPlot, "Infinite Lines", NewPlotVHLinesDemo)
}

func NewPlotVHLinesDemo() gig.IWidget {

	vals := []float64{0.25, 0.5, 0.75}

	group := gig.NewGroup().AddChildren(
		gig.NewPlot("##Infinite").
			SetSize(imgui.Vec2{-1, 0}).
			SetXAxisFlags(imgui.PlotAxisFlags_NoInitialFit).
			SetYAxisFlags(imgui.PlotAxisFlags_NoInitialFit).
			AddChildren(
				gig.NewPlotVLines("VLines").SetValuesFloat64(&vals[0], len(vals)),
				gig.NewPlotHLines("HLines").SetValuesFloat64(&vals[0], len(vals)),
			),
	)
	return group
}
