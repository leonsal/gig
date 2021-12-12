package main

import (
	"github.com/leonsal/gig"
)

func init() {

	registerDemo(catImPlot, "Stairstep", NewPlotStairsDemo)
}

func NewPlotStairsDemo() gig.IWidget {

	group := gig.NewGroup().AddChildren(
		gig.NewPlot("Plot stairs").AddChildren(
			gig.NewPlotStairs("Data1").SetValuesFloat32(&genSineF32(1, 0, 0.1, 100)[0], 100),
			gig.NewPlotStairs("Data2").SetValuesFloat64(&genSineF64(0.6, 0.1, 0.05, 100)[0], 100),
		),
	)
	return group
}
