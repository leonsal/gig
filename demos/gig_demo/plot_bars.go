package main

import (
	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
)

func init() {

	registerDemo(catImPlot, "Bars", NewPlotBarsDemo)
}

func NewPlotBarsDemo() gig.IWidget {

	midtm := [10]int32{83, 67, 23, 89, 83, 78, 91, 82, 85, 90}
	final := [10]float32{80, 62, 56, 99, 55, 78, 88, 78, 90, 100}
	grade := [10]float64{80, 69, 52, 92, 72, 78, 75, 76, 89, 95}
	labels := []string{"S1", "S2", "S3", "S4", "S5", "S6", "S7", "S8", "S9", "S10"}
	positions := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	group := gig.NewGroup().AddChildren(
		gig.NewPlot("Vertical").
			SetXLabel("Student").SetYLabel("Score").
			SetLimits(-0.5, 9.5, 0, 110, imgui.Cond_Always).
			SetTicksX(positions, labels).
			SetLegendLocation(imgui.PlotLocation_South, imgui.PlotOrientation_Horizontal, false).
			AddChildren(
				gig.NewPlotBars("Midterm Exam").SetValuesInt32(&midtm[0], len(midtm)).SetBarWidth(0.2).SetShift(-0.2),
				gig.NewPlotBars("Final Exam").SetValuesFloat32(&final[0], len(final)).SetBarWidth(0.2),
				gig.NewPlotBars("Course Grade").SetValuesFloat64(&grade[0], len(grade)).SetBarWidth(0.2).SetShift(0.2),
			),
		gig.NewPlot("Horizontal").
			SetXLabel("Score").SetYLabel("Student").
			SetLimits(0, 110, -0.5, 9.5, imgui.Cond_Always).
			SetTicksY(positions, labels).
			SetLegendLocation(imgui.PlotLocation_West, imgui.PlotOrientation_Vertical, false).
			AddChildren(
				gig.NewPlotBarsH("Midterm Exam").SetValuesInt32(&midtm[0], len(midtm)).SetBarHeight(0.2).SetShift(-0.2),
				gig.NewPlotBarsH("Final Exam").SetValuesFloat32(&final[0], len(final)).SetBarHeight(0.2),
				gig.NewPlotBarsH("Course Grade").SetValuesFloat64(&grade[0], len(grade)).SetBarHeight(0.2).SetShift(0.2),
			),
	)
	return group
}
