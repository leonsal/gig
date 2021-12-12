package main

import (
	"math/rand"

	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
)

func init() {

	registerDemo(catImPlot, "Heatmap", NewPlotHeatmapDemo)
}

func NewPlotHeatmapDemo() gig.IWidget {

	values1 := []float32{
		0.8, 2.4, 2.5, 3.9, 0.0, 4.0, 0.0,
		2.4, 0.0, 4.0, 1.0, 2.7, 0.0, 0.0,
		1.1, 2.4, 0.8, 4.3, 1.9, 4.4, 0.0,
		0.6, 0.0, 0.3, 0.0, 3.1, 0.0, 0.0,
		0.7, 1.7, 0.6, 2.6, 2.2, 6.2, 0.0,
		1.3, 1.2, 0.0, 0.0, 0.0, 3.2, 5.1,
		0.1, 2.0, 0.0, 1.4, 0.0, 1.9, 6.3,
	}
	xlabels := []string{"C1", "C2", "C3", "C4", "C5", "C6", "C7"}
	ylabels := []string{"R1", "R2", "R3", "R4", "R5", "R6", "R7"}
	scaleMin := float32(0.0)
	scaleMax := float32(6.3)

	const size = 200
	values2 := make([]float64, size*size)

	cmap := imgui.PlotColormap_Viridis
	axesFlags := imgui.PlotAxisFlags_Lock | imgui.PlotAxisFlags_NoGridLines | imgui.PlotAxisFlags_NoTickMarks
	group := gig.NewGroup().AddChildren(
		gig.NewColormapButton(imgui.GetColormapName(cmap)).
			SetSize(imgui.Vec2{225, 0}).
			SetColormap(cmap).
			SetOnClick(func(b *gig.ColormapButton) {
				cmap = imgui.PlotColormap((int(cmap) + 1) % imgui.GetColormapCount())
				b.SetLabel(imgui.GetColormapName(cmap))
				b.SetColormap(cmap)
				gig.Ref("pushColormap").(*gig.PushColormap).Set(cmap)
				imgui.BustColorCache("##Heatmap1")
				imgui.BustColorCache("##Heatmap2")
			}),
		gig.NewSameLine(),
		gig.NewLabelText("##Colormap Index", "%s", "Change Colormap"),
		gig.NewSetNextItemWidth(225),
		gig.NewDragFloat32Range2("Min / Max", scaleMin, scaleMax).SetSpeed(0.01).SetMin(-20).SetMax(20).
			SetOnChange(func(d *gig.DragFloat32Range2) {
				min := d.CurrentMin()
				max := d.CurrentMax()
				gig.Ref("heat").(*gig.PlotHeatmap).SetScale(float64(min), float64(max))
				gig.Ref("colormapScale").(*gig.ColormapScale).SetScale(float64(min), float64(max))
			}),
		gig.NewPushColormap(cmap).SaveRef("pushColormap"),
		gig.NewPlot("##Heatmap1").
			SetTicksXRange(0+1.0/14.0, 1-1.0/14.0, 7, xlabels).
			SetTicksYRange(1-1.0/14.0, 0+1.0/14.0, 7, ylabels).
			SetSize(imgui.Vec2{225, 225}).
			SetFlags(imgui.PlotFlags_NoLegend|imgui.PlotFlags_NoMousePos).
			SetXAxisFlags(axesFlags).
			SetYAxisFlags(axesFlags).
			AddChildren(
				gig.NewPlotHeatmap("heat").
					SetValuesFloat32(&values1[0], 7, 7).
					SetScale(float64(scaleMin), float64(scaleMax)).
					SaveRef("heat"),
			),
		gig.NewSameLine(),
		gig.NewColormapScale("##Heatscale", float64(scaleMin), float64(scaleMax)).SetSize(imgui.Vec2{60, 225}).
			SaveRef("colormapScale"),
		gig.NewSameLine(),
		gig.NewFunc(func() {
			rand.Seed(int64(imgui.GetTime() * 1_000_000))
			for i := 0; i < size*size; i++ {
				values2[i] = rand.Float64()
			}
		}),
		gig.NewPlot("##Heatmap2").
			SetLimits(-1, 1, -1, 1, imgui.Cond_Once).
			SetSize(imgui.Vec2{225, 225}).
			SetXAxisFlags(imgui.PlotAxisFlags_NoDecorations).
			SetYAxisFlags(imgui.PlotAxisFlags_NoDecorations).
			AddChildren(
				gig.NewPlotHeatmap("heat1").
					SetValuesFloat64(&values2[0], size, size).
					SetScale(0, 1).
					SetLabelFmt(""),
				gig.NewPlotHeatmap("heat2").
					SetValuesFloat64(&values2[0], size, size).
					SetScale(0, 1).
					SetBounds(imgui.PlotPoint{-1, -1}, imgui.PlotPoint{0, 0}).
					SetLabelFmt(""),
			),
		gig.NewPopColormap(1),
	)
	return group
}
