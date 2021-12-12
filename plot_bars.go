package gig

import (
	"github.com/leonsal/gig/imgui"
)

type PlotBars struct {
	Widget
	labelId     imgui.CString
	width       float64
	shift       float64
	offset      int
	stride      int
	x           interface{}
	y           interface{}
	count       int
	renderIndex int
}

func NewPlotBars(label string) *PlotBars {

	p := new(PlotBars)
	p.Init(p)
	p.labelId.SetString(label)
	p.width = 0.67
	p.shift = 0
	p.stride = 1
	p.SetRender(func() {
		switch p.renderIndex {
		case indexPlotI32:
			imgui.PlotBarsInt32CS(p.labelId, p.x.(*int32), p.count, p.width, p.shift, p.offset, p.stride)
		case indexPlotF32:
			imgui.PlotBarsFloat32CS(p.labelId, p.x.(*float32), p.count, p.width, p.shift, p.offset, p.stride)
		case indexPlotF64:
			imgui.PlotBarsFloat64CS(p.labelId, p.x.(*float64), p.count, p.width, p.shift, p.offset, p.stride)
		case indexPlot2I32:
			imgui.PlotBarsXYInt32CS(p.labelId, p.x.(*int32), p.y.(*int32), p.count, p.width, p.offset, p.stride)
		case indexPlot2F32:
			imgui.PlotBarsXYFloat32CS(p.labelId, p.x.(*float32), p.y.(*float32), p.count, p.width, p.offset, p.stride)
		case indexPlot2F64:
			imgui.PlotBarsXYFloat64CS(p.labelId, p.x.(*float64), p.y.(*float64), p.count, p.width, p.offset, p.stride)
		}
	})
	return p
}

func (p *PlotBars) SetLabel(label string) *PlotBars {

	p.labelId.SetString(label)
	return p
}

func (p *PlotBars) SetValuesInt32(v *int32, count int) *PlotBars {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotI32
	return p
}

func (p *PlotBars) SetValuesFloat32(v *float32, count int) *PlotBars {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotF32
	return p
}

func (p *PlotBars) SetValuesFloat64(v *float64, count int) *PlotBars {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotF64
	return p
}

func (p *PlotBars) SetValuesXYInt32(x, y *int32, count int) *PlotBars {

	p.x = x
	p.y = y
	p.count = count
	p.renderIndex = indexPlot2I32
	return p
}

func (p *PlotBars) SetValuesXYFloat32(x, y *float32, count int) *PlotBars {

	p.x = x
	p.y = y
	p.count = count
	p.renderIndex = indexPlot2F32
	return p
}

func (p *PlotBars) SetValuesXYFloat64(x, y *float64, count int) *PlotBars {

	p.x = x
	p.y = y
	p.count = count
	p.renderIndex = indexPlot2F64
	return p
}

func (p *PlotBars) SetBarWidth(width float64) *PlotBars {

	p.width = width
	return p
}

func (p *PlotBars) SetShift(shift float64) *PlotBars {

	p.shift = shift
	return p
}

// SetOffset sets the offset from the start of the data in number of elements (default = 0)
func (p *PlotBars) SetOffset(offset int) *PlotBars {

	p.offset = offset
	return p
}

// SetStride sets the stride in number of elements (default = 1)
func (p *PlotBars) SetStride(stride int) *PlotBars {

	p.stride = stride
	return p
}

type PlotBarsH struct {
	Widget
	labelId     imgui.CString
	height      float64
	shift       float64
	offset      int
	stride      int
	x           interface{}
	y           interface{}
	count       int
	renderIndex int
}

func NewPlotBarsH(label string) *PlotBarsH {

	p := new(PlotBarsH)
	p.Init(p)
	p.labelId.SetString(label)
	p.height = 0.67
	p.shift = 0
	p.stride = 1
	p.SetRender(func() {
		switch p.renderIndex {
		case indexPlotI32:
			imgui.PlotBarsHInt32CS(p.labelId, p.x.(*int32), p.count, p.height, p.shift, p.offset, p.stride)
		case indexPlotF32:
			imgui.PlotBarsHFloat32CS(p.labelId, p.x.(*float32), p.count, p.height, p.shift, p.offset, p.stride)
		case indexPlotF64:
			imgui.PlotBarsHFloat64CS(p.labelId, p.x.(*float64), p.count, p.height, p.shift, p.offset, p.stride)
		case indexPlot2I32:
			imgui.PlotBarsHXYInt32CS(p.labelId, p.x.(*int32), p.y.(*int32), p.count, p.height, p.offset, p.stride)
		case indexPlot2F32:
			imgui.PlotBarsHXYFloat32CS(p.labelId, p.x.(*float32), p.y.(*float32), p.count, p.height, p.offset, p.stride)
		case indexPlot2F64:
			imgui.PlotBarsHXYFloat64CS(p.labelId, p.x.(*float64), p.y.(*float64), p.count, p.height, p.offset, p.stride)
		}
	})
	return p
}

func (p *PlotBarsH) SetLabel(label string) *PlotBarsH {

	p.labelId.SetString(label)
	return p
}

func (p *PlotBarsH) SetValuesInt32(v *int32, count int) *PlotBarsH {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotI32
	return p
}

func (p *PlotBarsH) SetValuesFloat32(v *float32, count int) *PlotBarsH {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotF32
	return p
}

func (p *PlotBarsH) SetValuesFloat64(v *float64, count int) *PlotBarsH {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotF64
	return p
}

func (p *PlotBarsH) SetValuesXYInt32(x, y *int32, count int) *PlotBarsH {

	p.x = x
	p.y = y
	p.count = count
	p.renderIndex = indexPlot2I32
	return p
}

func (p *PlotBarsH) SetValuesXYFloat32(x, y *float32, count int) *PlotBarsH {

	p.x = x
	p.y = y
	p.count = count
	p.renderIndex = indexPlot2F32
	return p
}

func (p *PlotBarsH) SetValuesXYFloat64(x, y *float64, count int) *PlotBarsH {

	p.x = x
	p.y = y
	p.count = count
	p.renderIndex = indexPlot2F64
	return p
}

func (p *PlotBarsH) SetBarHeight(height float64) *PlotBarsH {

	p.height = height
	return p
}

func (p *PlotBarsH) SetShift(shift float64) *PlotBarsH {

	p.shift = shift
	return p
}

// SetOffset sets the offset from the start of the data in number of elements (default = 0)
func (p *PlotBarsH) SetOffset(offset int) *PlotBarsH {

	p.offset = offset
	return p
}

// SetStride sets the stride in number of elements (default = 1)
func (p *PlotBarsH) SetStride(stride int) *PlotBarsH {

	p.stride = stride
	return p
}
