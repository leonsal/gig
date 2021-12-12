package gig

import (
	"github.com/leonsal/gig/imgui"
)

type PlotHeatmap struct {
	Widget
	labelId     imgui.CString
	v           interface{}
	rows        int
	cols        int
	scaleMin    float64
	scaleMax    float64
	labelFmt    imgui.CString
	boundsMin   imgui.PlotPoint
	boundsMax   imgui.PlotPoint
	renderIndex int
}

func NewPlotHeatmap(label string) *PlotHeatmap {

	p := new(PlotHeatmap)
	p.Init(p)
	p.labelId.SetString(label)
	p.scaleMin = 0
	p.scaleMax = 0
	p.labelFmt.SetString("%.1f")
	p.boundsMin = imgui.PlotPoint{0, 0}
	p.boundsMax = imgui.PlotPoint{1, 1}
	p.SetRender(func() {
		switch p.renderIndex {
		case indexPlotI32:
			imgui.PlotHeatmapInt32CS(p.labelId, p.v.(*int32), p.rows, p.cols, p.scaleMin, p.scaleMax, p.labelFmt, p.boundsMin, p.boundsMax)
		case indexPlotF32:
			imgui.PlotHeatmapFloat32CS(p.labelId, p.v.(*float32), p.rows, p.cols, p.scaleMin, p.scaleMax, p.labelFmt, p.boundsMin, p.boundsMax)
		case indexPlotF64:
			imgui.PlotHeatmapFloat64CS(p.labelId, p.v.(*float64), p.rows, p.cols, p.scaleMin, p.scaleMax, p.labelFmt, p.boundsMin, p.boundsMax)
		}
	})
	return p
}

func (p *PlotHeatmap) Save(pp **PlotHeatmap) *PlotHeatmap {

	*pp = p
	return p
}

func (p *PlotHeatmap) SetLabel(label string) *PlotHeatmap {

	p.labelId.SetString(label)
	return p
}

func (p *PlotHeatmap) SetValuesInt32(v *int32, rows, cols int) *PlotHeatmap {

	p.v = v
	p.rows = rows
	p.cols = cols
	p.renderIndex = indexPlotI32
	return p
}

func (p *PlotHeatmap) SetValuesFloat32(v *float32, rows, cols int) *PlotHeatmap {

	p.v = v
	p.rows = rows
	p.cols = cols
	p.renderIndex = indexPlotF32
	return p
}

func (p *PlotHeatmap) SetValuesFloat64(v *float64, rows, cols int) *PlotHeatmap {

	p.v = v
	p.rows = rows
	p.cols = cols
	p.renderIndex = indexPlotF64
	return p
}

func (p *PlotHeatmap) SetScale(min, max float64) *PlotHeatmap {

	p.scaleMin = min
	p.scaleMax = max
	return p
}

func (p *PlotHeatmap) SetLabelFmt(fmt string) *PlotHeatmap {

	p.labelFmt.SetString(fmt)
	return p
}

func (p *PlotHeatmap) SetBounds(min, max imgui.PlotPoint) *PlotHeatmap {

	p.boundsMin = min
	p.boundsMax = max
	return p
}
