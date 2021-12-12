package main

import (
	"github.com/leonsal/gig"
)

func init() {

	registerDemo(catImPlot, "Line", NewPlotLineDemo)
}

func NewPlotLineDemo() gig.IWidget {

	group := gig.NewGroup().AddChildren(

		gig.NewPlot("Plot lines").AddChildren(
			gig.NewPlotLine("Sine1").SetValuesFloat32(&genSineF32(1, 0, 0.05, 400)[0], 400),
			gig.NewPlotLine("Sine2").SetValuesFloat32(&genSineF32(0.5, 0.5, 0.02, 400)[0], 400),
			gig.NewPlotLine("Sine3").SetValuesFloat64(&genSineF64(0.3, -0.5, 0.01, 400)[0], 400),
			gig.NewPlotLine("XY").SetValuesXYFloat32(
				&[]float32{0, 100, 150, 310, 400}[0], &[]float32{1, 0.5, -1, 0, 0.3}[0], 5),
		),

		gig.NewPlot("Plot filled").AddChildren(
			gig.NewPlotShaded("Sine1").SetValuesFloat32(&genSineF32(1, 0, 0.05, 400)[0], 400),
			gig.NewPlotShaded("Sine2").SetValuesFloat32(&genSineF32(0.2, 0, 0.03, 400)[0], 400),
		),

		gig.NewPlot("Plot shaded").AddChildren(
			gig.NewPlotShaded("Sine1").SetValuesXYYFloat32(
				&genSeqF32(0.03, 400)[0],
				&genSineF32(0.2, 0, 0.03, 400)[0],
				&genSineF32(0.2, 0.3, 0.03, 400)[0],
				400,
			),
		),
	)
	return group
}
