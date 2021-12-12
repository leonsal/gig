package gig

import (
	"github.com/leonsal/gig/imgui"
)

type PlotErrorBars struct {
	plotBase
	labelId     imgui.CString
	offset      int
	stride      int
	x           interface{}
	y           interface{}
	e1          interface{}
	e2          interface{}
	count       int
	renderIndex int
}

func NewPlotErrorBars(labelId string) *PlotErrorBars {

	p := new(PlotErrorBars)
	p.Init(p)
	p.labelId.SetString(labelId)
	p.offset = 0
	p.stride = 1
	p.SetRender(func() {
		switch p.renderIndex {
		case indexPlotI32:
			imgui.PlotErrorBarsInt32CS(p.labelId, p.x.(*int32), p.y.(*int32), p.e1.(*int32), p.count, p.offset, p.stride)
		case indexPlotF32:
			imgui.PlotErrorBarsFloat32CS(p.labelId, p.x.(*float32), p.y.(*float32), p.e1.(*float32), p.count, p.offset, p.stride)
		case indexPlotF64:
			imgui.PlotErrorBarsFloat64CS(p.labelId, p.x.(*float64), p.y.(*float64), p.e2.(*float64), p.count, p.offset, p.stride)
		case indexPlot2I32:
			imgui.PlotErrorBars2Int32CS(p.labelId, p.x.(*int32), p.y.(*int32), p.e1.(*int32), p.e2.(*int32), p.count, p.offset, p.stride)
		case indexPlot2F32:
			imgui.PlotErrorBars2Float32CS(p.labelId, p.x.(*float32), p.y.(*float32), p.e1.(*float32), p.e2.(*float32), p.count, p.offset, p.stride)
		case indexPlot2F64:
			imgui.PlotErrorBars2Float64CS(p.labelId, p.x.(*float64), p.y.(*float64), p.e1.(*float64), p.e2.(*float64), p.count, p.offset, p.stride)
		}
	})
	return p
}

func (p *PlotErrorBars) SetLabel(label string) *PlotErrorBars {

	p.labelId.SetString(label)
	return p
}

func (p *PlotErrorBars) SetValuesInt32(x, y, er *int32, count int) *PlotErrorBars {

	p.x = x
	p.y = y
	p.e1 = er
	p.count = count
	p.renderIndex = indexPlotI32
	return p
}

func (p *PlotErrorBars) SetValuesFloat32(x, y, er *float32, count int) *PlotErrorBars {

	p.x = x
	p.y = y
	p.e1 = er
	p.count = count
	p.renderIndex = indexPlotF32
	return p
}

func (p *PlotErrorBars) SetValuesFloat64(x, y, er *float64, count int) *PlotErrorBars {

	p.x = x
	p.y = y
	p.e1 = er
	p.count = count
	p.renderIndex = indexPlotF64
	return p
}

func (p *PlotErrorBars) SetValues2Int32(x, y, neg, pos *int32, count int) *PlotErrorBars {

	p.x = x
	p.y = y
	p.e1 = neg
	p.e2 = pos
	p.count = count
	p.renderIndex = indexPlot2I32
	return p
}

func (p *PlotErrorBars) SetValues2Float32(x, y, neg, pos *float32, count int) *PlotErrorBars {

	p.x = x
	p.y = y
	p.e1 = neg
	p.e2 = pos
	p.count = count
	p.renderIndex = indexPlot2F32
	return p
}

func (p *PlotErrorBars) SetValues2Float64(x, y, neg, pos *float64, count int) *PlotErrorBars {

	p.x = x
	p.y = y
	p.e1 = neg
	p.e2 = pos
	p.count = count
	p.renderIndex = indexPlot2F32
	return p
}

// SetOffset sets the offset from the start of the data in number of elements (default = 0)
func (p *PlotErrorBars) SetOffset(offset int) *PlotErrorBars {

	p.offset = offset
	return p
}

// SetStride sets the stride in number of elements (default = 1)
func (p *PlotErrorBars) SetStride(stride int) *PlotErrorBars {

	p.stride = stride
	return p
}

type PlotErrorBarsH struct {
	plotBase
	labelId     imgui.CString
	offset      int
	stride      int
	x           interface{}
	y           interface{}
	e1          interface{}
	e2          interface{}
	count       int
	renderIndex int
}

func NewPlotErrorBarsH(labelId string) *PlotErrorBarsH {

	p := new(PlotErrorBarsH)
	p.Init(p)
	p.labelId.SetString(labelId)
	p.offset = 0
	p.stride = 1
	p.SetRender(func() {
		switch p.renderIndex {
		case indexPlotI32:
			imgui.PlotErrorBarsHInt32CS(p.labelId, p.x.(*int32), p.y.(*int32), p.e1.(*int32), p.count, p.offset, p.stride)
		case indexPlotF32:
			imgui.PlotErrorBarsHFloat32CS(p.labelId, p.x.(*float32), p.y.(*float32), p.e1.(*float32), p.count, p.offset, p.stride)
		case indexPlotF64:
			imgui.PlotErrorBarsHFloat64CS(p.labelId, p.x.(*float64), p.y.(*float64), p.e2.(*float64), p.count, p.offset, p.stride)
		case indexPlot2I32:
			imgui.PlotErrorBarsH2Int32CS(p.labelId, p.x.(*int32), p.y.(*int32), p.e1.(*int32), p.e2.(*int32), p.count, p.offset, p.stride)
		case indexPlot2F32:
			imgui.PlotErrorBarsH2Float32CS(p.labelId, p.x.(*float32), p.y.(*float32), p.e1.(*float32), p.e2.(*float32), p.count, p.offset, p.stride)
		case indexPlot2F64:
			imgui.PlotErrorBarsH2Float64CS(p.labelId, p.x.(*float64), p.y.(*float64), p.e1.(*float64), p.e2.(*float64), p.count, p.offset, p.stride)
		}
	})
	return p
}

func (p *PlotErrorBarsH) SetLabel(label string) *PlotErrorBarsH {

	p.labelId.SetString(label)
	return p
}

func (p *PlotErrorBarsH) SetValuesInt32(x, y, er *int32, count int) *PlotErrorBarsH {

	p.x = x
	p.y = y
	p.e1 = er
	p.count = count
	p.renderIndex = indexPlotI32
	return p
}

func (p *PlotErrorBarsH) SetValuesFloat32(x, y, er *float32, count int) *PlotErrorBarsH {

	p.x = x
	p.y = y
	p.e1 = er
	p.count = count
	p.renderIndex = indexPlotF32
	return p
}

func (p *PlotErrorBarsH) SetValuesFloat64(x, y, er *float64, count int) *PlotErrorBarsH {

	p.x = x
	p.y = y
	p.e1 = er
	p.count = count
	p.renderIndex = indexPlotF64
	return p
}

func (p *PlotErrorBarsH) SetValues2Int32(x, y, neg, pos *int32, count int) *PlotErrorBarsH {

	p.x = x
	p.y = y
	p.e1 = neg
	p.e2 = pos
	p.count = count
	p.renderIndex = indexPlot2I32
	return p
}

func (p *PlotErrorBarsH) SetValues2Float32(x, y, neg, pos *float32, count int) *PlotErrorBarsH {

	p.x = x
	p.y = y
	p.e1 = neg
	p.e2 = pos
	p.count = count
	p.renderIndex = indexPlot2F32
	return p
}

func (p *PlotErrorBarsH) SetValues2Float64(x, y, neg, pos *float64, count int) *PlotErrorBarsH {

	p.x = x
	p.y = y
	p.e1 = neg
	p.e2 = pos
	p.count = count
	p.renderIndex = indexPlot2F64
	return p
}

// SetOffset sets the offset from the start of the data in number of elements (default = 0)
func (p *PlotErrorBarsH) SetOffset(offset int) *PlotErrorBarsH {

	p.offset = offset
	return p
}

// SetStride sets the stride in number of elements (default = 1)
func (p *PlotErrorBarsH) SetStride(stride int) *PlotErrorBarsH {

	p.stride = stride
	return p
}
