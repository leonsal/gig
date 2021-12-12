package gig

import (
	"github.com/leonsal/gig/imgui"
)

type PlotPieChart struct {
	Widget
	labelIds    imgui.CStrArr
	labelFmt    imgui.CString
	v           interface{}
	count       int
	x           float64
	y           float64
	radius      float64
	angle0      float64
	normalize   bool
	renderIndex int
}

func NewPlotPieChart(labels []string) *PlotPieChart {

	p := new(PlotPieChart)
	p.Init(p)
	p.labelIds.Append(labels...)
	p.labelFmt.SetString("%.1f")
	p.angle0 = 90
	p.normalize = false
	p.SetRender(func() {
		switch p.renderIndex {
		case indexPlotI32:
			imgui.PlotPieChartInt32CS(p.labelIds, p.v.(*int32), p.count, p.x, p.y, p.radius, p.normalize, p.labelFmt, p.angle0)
		case indexPlotF32:
			imgui.PlotPieChartFloat32CS(p.labelIds, p.v.(*float32), p.count, p.x, p.y, p.radius, p.normalize, p.labelFmt, p.angle0)
		case indexPlotF64:
			imgui.PlotPieChartFloat64CS(p.labelIds, p.v.(*float64), p.count, p.x, p.y, p.radius, p.normalize, p.labelFmt, p.angle0)
		}
	})
	return p
}

func (p *PlotPieChart) Save(pp **PlotPieChart) *PlotPieChart {

	*pp = p
	return p
}

func (p *PlotPieChart) SetLabels(labels []string) *PlotPieChart {

	p.labelIds.Clear()
	p.labelIds.Append(labels...)
	return p
}

func (p *PlotPieChart) SetPos(x, y, radius float64) *PlotPieChart {

	p.x = x
	p.y = y
	p.radius = radius
	return p
}

func (p *PlotPieChart) SetNormalize(normalize bool) *PlotPieChart {

	p.normalize = normalize
	return p
}

func (p *PlotPieChart) SetAngle0(angle0 float64) *PlotPieChart {

	p.angle0 = angle0
	return p
}

func (p *PlotPieChart) SetLabelFmt(fmt string) *PlotPieChart {

	p.labelFmt.SetString(fmt)
	return p
}

func (p *PlotPieChart) SetValuesInt32(v *int32, count int) *PlotPieChart {

	p.v = v
	p.count = count
	p.renderIndex = indexPlotI32
	return p
}

func (p *PlotPieChart) SetValuesFloat32(v *float32, count int) *PlotPieChart {

	p.v = v
	p.count = count
	p.renderIndex = indexPlotF32
	return p
}

func (p *PlotPieChart) SetValuesFloat64(v *float64, count int) *PlotPieChart {

	p.v = v
	p.count = count
	p.renderIndex = indexPlotF64
	return p
}
