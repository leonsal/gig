package main

import (
	"math/rand"

	"github.com/leonsal/gig"
)

func init() {

	registerDemo(catImPlot, "Scatter", NewPlotScatterDemo)
}

func NewPlotScatterDemo() gig.IWidget {

	group := gig.NewGroup().AddChildren(

		gig.NewPlot("Plot scatter").AddChildren(
			gig.NewPlotScatter("Data1").Call(func(w gig.IWidget) {
				ps := w.(*gig.PlotScatter)
				count := 100
				xs := make([]float32, count)
				ys := make([]float32, count)
				for i := 0; i < count; i++ {
					xs[i] = float32(i) * 0.01
					ys[i] = xs[i] + 0.1*rand.Float32()
				}
				ps.SetValuesXYFloat32(&xs[0], &ys[0], len(xs))
			}),
			gig.NewPlotScatter("Data2").Call(func(w gig.IWidget) {
				ps := w.(*gig.PlotScatter)
				count := 50
				xs := make([]float32, count)
				ys := make([]float32, count)
				for i := 0; i < count; i++ {
					xs[i] = 0.25 + 0.2*rand.Float32()
					ys[i] = 0.75 + 0.2*rand.Float32()
				}
				ps.SetValuesXYFloat32(&xs[0], &ys[0], len(xs))
			}),
		),
	)
	return group
}
