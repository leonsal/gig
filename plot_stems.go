package gig

import (
	"github.com/leonsal/gig/imgui"
)

type PlotStems struct {
	Widget
	labelId     imgui.CString
	yref        float64
	x0          float64
	xscale      float64
	offset      int
	stride      int
	x           interface{}
	y           interface{}
	count       int
	renderIndex int
}

func NewPlotStems(label string) *PlotStems {

	p := new(PlotStems)
	p.Init(p)
	p.labelId.SetString(label)
	p.xscale = 1
	p.stride = 1
	p.SetRender(func() {
		switch p.renderIndex {
		case indexPlotI32:
			imgui.PlotStemsInt32CS(p.labelId, p.x.(*int32), p.count, p.yref, p.xscale, p.x0, p.offset, p.stride)
		case indexPlotF32:
			imgui.PlotStemsFloat32CS(p.labelId, p.x.(*float32), p.count, p.yref, p.xscale, p.x0, p.offset, p.stride)
		case indexPlotF64:
			imgui.PlotStemsFloat64CS(p.labelId, p.x.(*float64), p.count, p.yref, p.xscale, p.x0, p.offset, p.stride)
		case indexPlot2I32:
			imgui.PlotStemsXYInt32CS(p.labelId, p.x.(*int32), p.y.(*int32), p.count, p.yref, p.offset, p.stride)
		case indexPlot2F32:
			imgui.PlotStemsXYFloat32CS(p.labelId, p.x.(*float32), p.y.(*float32), p.count, p.yref, p.offset, p.stride)
		case indexPlot2F64:
			imgui.PlotStemsXYFloat64CS(p.labelId, p.x.(*float64), p.y.(*float64), p.count, p.yref, p.offset, p.stride)
		}
	})
	return p
}

func (p *PlotStems) SetLabel(label string) *PlotStems {

	p.labelId.SetString(label)
	return p
}

func (p *PlotStems) SetValuesInt32(v *int32, count int) *PlotStems {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotI32
	return p
}

func (p *PlotStems) SetValuesFloat32(v *float32, count int) *PlotStems {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotF32
	return p
}

func (p *PlotStems) SetValuesFloat64(v *float64, count int) *PlotStems {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotF64
	return p
}

func (p *PlotStems) SetValuesXYInt32(x, y *int32, count int) *PlotStems {

	p.x = x
	p.y = y
	p.count = count
	p.renderIndex = indexPlot2I32
	return p
}

func (p *PlotStems) SetValuesXYFloat32(x, y *float32, count int) *PlotStems {

	p.x = x
	p.y = y
	p.count = count
	p.renderIndex = indexPlot2F32
	return p
}

func (p *PlotStems) SetValuesXYFloat64(x, y *float64, count int) *PlotStems {

	p.x = x
	p.y = y
	p.count = count
	p.renderIndex = indexPlot2F64
	return p
}

func (p *PlotStems) SetYref(yref float64) *PlotStems {

	p.yref = yref
	return p
}

// SetXscale sets the scale of the X axis (default = 1)
func (p *PlotStems) SetXscale(xscale float64) *PlotStems {

	p.xscale = xscale
	return p
}

// SetX0 sets the X coordinate of the start of the plot (default = 0)
func (p *PlotStems) SetX0(x0 float64) *PlotStems {

	p.x0 = x0
	return p
}

// SetOffset sets the offset from the start of the data in number of elements (default = 0)
func (p *PlotStems) SetOffset(offset int) *PlotStems {

	p.offset = offset
	return p
}

// SetStride sets the stride in number of elements (default = 1)
func (p *PlotStems) SetStride(stride int) *PlotStems {

	p.stride = stride
	return p
}
