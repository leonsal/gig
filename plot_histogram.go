package gig

import (
	"github.com/leonsal/gig/imgui"
)

type PlotHistogram struct {
	Widget
	labelId     imgui.CString
	values      interface{}
	count       int
	bins        int
	cumulative  bool
	density     bool
	prange      imgui.PlotRange
	outliers    bool
	barScale    float64
	renderIndex int
}

func NewPlotHistogram(label string) *PlotHistogram {

	p := new(PlotHistogram)
	p.Init(p)
	p.labelId.SetString(label)
	p.bins = imgui.PlotBin_Sturges
	p.cumulative = false
	p.density = false
	p.outliers = true
	p.barScale = 1.0
	p.SetRender(func() {
		switch p.renderIndex {
		case indexPlotI32:
			imgui.PlotHistogramInt32CS(p.labelId, p.values.(*int32), p.count, p.bins, p.cumulative, p.density, p.prange, p.outliers, p.barScale)
		case indexPlotF32:
			imgui.PlotHistogramFloat32CS(p.labelId, p.values.(*float32), p.count, p.bins, p.cumulative, p.density, p.prange, p.outliers, p.barScale)
		case indexPlotF64:
			imgui.PlotHistogramFloat64CS(p.labelId, p.values.(*float64), p.count, p.bins, p.cumulative, p.density, p.prange, p.outliers, p.barScale)
		}
	})
	return p
}

func (p *PlotHistogram) Save(pp **PlotHistogram) *PlotHistogram {

	*pp = p
	return p
}

func (p *PlotHistogram) SetLabel(label string) *PlotHistogram {

	p.labelId.SetString(label)
	return p
}

func (p *PlotHistogram) SetValuesInt32(v *int32, count, bins int) *PlotHistogram {

	p.values = v
	p.count = count
	p.bins = bins
	p.renderIndex = indexPlotI32
	return p
}

func (p *PlotHistogram) SetValuesFloat32(v *float32, count, bins int) *PlotHistogram {

	p.values = v
	p.count = count
	p.bins = bins
	p.renderIndex = indexPlotF32
	return p
}

func (p *PlotHistogram) SetValuesFloat64(v *float32, count, bins int) *PlotHistogram {

	p.values = v
	p.count = count
	p.bins = bins
	p.renderIndex = indexPlotF64
	return p
}

func (p *PlotHistogram) SetCumulative(c bool) *PlotHistogram {

	p.cumulative = c
	return p
}

func (p *PlotHistogram) SetDensity(density bool) *PlotHistogram {

	p.density = density
	return p
}

func (p *PlotHistogram) SetOutliers(outliers bool) *PlotHistogram {

	p.outliers = outliers
	return p
}

func (p *PlotHistogram) SetRange(prange imgui.PlotRange) *PlotHistogram {

	p.prange = prange
	return p
}

func (p *PlotHistogram) SetBarScale(barScale float64) *PlotHistogram {

	p.barScale = barScale
	return p
}
