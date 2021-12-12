package main

import (
	"github.com/leonsal/gig"
)

func init() {

	registerDemo(catImPlot, "Histogram", NewPlotHistogramDemo)
}

func NewPlotHistogramDemo() gig.IWidget {

	group := gig.NewGroup().AddChildren(
		gig.NewRadioGroup(
			gig.NewRadioButton("Sqrt"), gig.NewSameLine(),
			gig.NewRadioButton("Sturges"), gig.NewSameLine(),
			gig.NewRadioButton("Rice"), gig.NewSameLine(),
			gig.NewRadioButton("Scot"), gig.NewSameLine(),
			gig.NewRadioButton("N bins"),
			gig.NewGroup().AddChildren(
				gig.NewSameLine(),
				gig.NewSetNextItemWidth(200),
				gig.NewSliderInt32("##Bins", 50, 1, 100),
			).SetVisible(false).SaveRef("binSlider"),

			gig.NewCheckbox("Density").SetOnClick(func(c *gig.Checkbox) {

			}),
			gig.NewSameLine(),
			gig.NewCheckbox("Cumulative").SetOnClick(func(c *gig.Checkbox) {

			}),
			gig.NewSameLine(),
			gig.NewCheckbox("Range").SetOnClick(func(c *gig.Checkbox) {

			}),
		).SetOnChange(func(rb *gig.RadioGroup) {
			idx, _ := rb.Active()
			binSlider := gig.Ref("binSlider").(*gig.Group)
			binSlider.SetVisible(false)
			switch idx {
			case 4:
				binSlider.SetVisible(true)

			}

		}),
	)
	return group
}
