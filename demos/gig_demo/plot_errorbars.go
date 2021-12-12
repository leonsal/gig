package main

import (
	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
)

func init() {

	registerDemo(catImPlot, "ErrorBars", NewPlotErrorBarsDemo)
}

func NewPlotErrorBarsDemo() gig.IWidget {

	xs := []float32{1, 2, 3, 4, 5}
	bar := []float32{1, 2, 5, 3, 4}
	lin1 := []float32{8, 8, 9, 7, 8}
	lin2 := []float32{6, 7, 6, 9, 6}
	err1 := []float32{0.2, 0.4, 0.2, 0.6, 0.4}
	err2 := []float32{0.4, 0.2, 0.4, 0.8, 0.6}
	err3 := []float32{0.09, 0.14, 0.09, 0.12, 0.16}
	err4 := []float32{0.02, 0.08, 0.15, 0.05, 0.2}

	group := gig.NewGroup().AddChildren(
		gig.NewPlot("Vertical").SetLimits(0, 6, 0, 10, imgui.Cond_Always).
			AddChildren(
				gig.NewPlotBars("Bar").SetValuesXYFloat32(&xs[0], &bar[0], len(xs)).SetItemWidth(0.5),
				gig.NewPlotErrorBars("Bar").SetValuesFloat32(&xs[0], &bar[0], &err1[0], len(xs)),

				gig.NewSetNextErrorBarStyle().SetColor(imgui.GetColormapColor(1, imgui.PlotAuto)).SetSize(0),
				gig.NewPlotErrorBars("Line").SetValues2Float32(&xs[0], &lin1[0], &err1[0], &err2[0], len(xs)),
				gig.NewSetNextMarkerStyle(imgui.PlotMarker_Circle),
				gig.NewPlotLine("Line").SetValuesXYFloat32(&xs[0], &lin1[0], len(xs)),

				gig.NewPlotPushStyleColor(imgui.PlotCol_ErrorBar, imgui.GetColormapColor(2, imgui.PlotAuto)),
				gig.NewPlotErrorBars("Scatter").SetValuesFloat32(&xs[0], &lin2[0], &err2[0], len(xs)),
				gig.NewPlotErrorBarsH("Scatter").SetValues2Float32(&xs[0], &lin2[0], &err3[0], &err4[0], len(xs)),
				gig.NewPlotPopStyleColor(1),
				gig.NewPlotScatter("Scatter").SetValuesXYFloat32(&xs[0], &lin2[0], len(xs)),
			),
	)
	return group
}
