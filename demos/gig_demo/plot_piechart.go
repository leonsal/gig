package main

import (
	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
)

func init() {

	registerDemo(catImPlot, "Pie Chart", NewPlotPieChartDemo)
}

func NewPlotPieChartDemo() gig.IWidget {

	labels1 := []string{"Frogs", "Hogs", "Dogs", "Logs"}
	data1 := []float32{0.15, 0.30, 0.2, 0.05}
	labels2 := []string{"A", "B", "C", "D", "E"}
	data2 := []int32{1, 1, 2, 3, 5}

	var check *gig.Group
	var pie1 *gig.PlotPieChart
	group := gig.NewGroup().AddChildren(
		gig.NewDragFloat32Slice("Values", data1).
			SetSpeed(0.01).SetMin(0).SetMax(1).SetOnChange(func(d *gig.DragFloat32Slice) {
			total := float32(0)
			for _, d := range d.Values() {
				total += d
			}
			if total < 1 {
				check.SetVisible(true)
			} else {
				check.SetVisible(false)
			}
		}).SetItemWidth(250),
		gig.NewGroup().AddChildren(
			gig.NewSameLine(),
			gig.NewCheckbox("Normalize").SetOnClick(func(c *gig.Checkbox) {
				pie1.SetNormalize(c.Value())
			}),
		).Save(&check),
		gig.NewPlot("##Pie1").
			SetLimits(0, 1, 0, 1, imgui.Cond_Always).
			SetSize(imgui.Vec2{250, 250}).
			SetFlags(imgui.PlotFlags_Equal|imgui.PlotFlags_NoMousePos).
			SetXAxisFlags(imgui.PlotAxisFlags_NoDecorations).
			SetYAxisFlags(imgui.PlotAxisFlags_NoDecorations).
			AddChildren(
				gig.NewPlotPieChart(labels1).
					Save(&pie1).
					SetValuesFloat32(&data1[0], len(data1)).
					SetPos(0.5, 0.5, 0.4).
					SetLabelFmt("%.2f"),
			),

		gig.NewSameLine(),

		gig.NewPushColormap(imgui.PlotColormap_Pastel),
		gig.NewPlot("##Pie2").
			SetLimits(0, 1, 0, 1, imgui.Cond_Always).
			SetSize(imgui.Vec2{250, 250}).
			SetFlags(imgui.PlotFlags_Equal|imgui.PlotFlags_NoMousePos).
			SetXAxisFlags(imgui.PlotAxisFlags_NoDecorations).
			SetYAxisFlags(imgui.PlotAxisFlags_NoDecorations).
			AddChildren(
				gig.NewPlotPieChart(labels2).
					SetValuesInt32(&data2[0], len(data2)).
					SetPos(0.5, 0.5, 0.4).
					SetNormalize(true).
					SetLabelFmt("%.0f").
					SetAngle0(180),
			),
		gig.NewPopColormap(1),
	)
	return group
}
