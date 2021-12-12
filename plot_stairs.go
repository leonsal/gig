package gig

import (
	"github.com/leonsal/gig/imgui"
)

type PlotStairs struct {
	Widget
	labelId     imgui.CString
	x0          float64
	xscale      float64
	offset      int
	stride      int
	x           interface{}
	y           interface{}
	count       int
	renderIndex int
}

func NewPlotStairs(label string) *PlotStairs {

	p := new(PlotStairs)
	p.Init(p)
	p.labelId.SetString(label)
	p.xscale = 1
	p.stride = 1
	p.SetRender(func() {
		switch p.renderIndex {
		case indexPlotI32:
			imgui.PlotStairsInt32CS(p.labelId, p.x.(*int32), p.count, p.xscale, p.x0, p.offset, p.stride)
		case indexPlotF32:
			imgui.PlotStairsFloat32CS(p.labelId, p.x.(*float32), p.count, p.xscale, p.x0, p.offset, p.stride)
		case indexPlotF64:
			imgui.PlotStairsFloat64CS(p.labelId, p.x.(*float64), p.count, p.xscale, p.x0, p.offset, p.stride)
		case indexPlot2I32:
			imgui.PlotStairsXYInt32CS(p.labelId, p.x.(*int32), p.y.(*int32), p.count, p.offset, p.stride)
		case indexPlot2F32:
			imgui.PlotStairsXYFloat32CS(p.labelId, p.x.(*float32), p.y.(*float32), p.count, p.offset, p.stride)
		case indexPlot2F64:
			imgui.PlotStairsXYFloat64CS(p.labelId, p.x.(*float64), p.y.(*float64), p.count, p.offset, p.stride)
		}
	})
	return p
}

func (p *PlotStairs) SetLabel(label string) *PlotStairs {

	p.labelId.SetString(label)
	return p
}

func (p *PlotStairs) SetValuesInt32(v *int32, count int) *PlotStairs {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotI32
	return p
}

func (p *PlotStairs) SetValuesFloat32(v *float32, count int) *PlotStairs {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotF32
	return p
}

func (p *PlotStairs) SetValuesFloat64(v *float64, count int) *PlotStairs {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotF64
	return p
}

func (p *PlotStairs) SetValuesXYInt32(x, y *int32, count int) *PlotStairs {

	p.x = x
	p.y = y
	p.count = count
	p.renderIndex = indexPlot2I32
	return p
}

func (p *PlotStairs) SetValuesXYFloat32(x, y *float32, count int) *PlotStairs {

	p.x = x
	p.y = y
	p.count = count
	p.renderIndex = indexPlot2F32
	return p
}

func (p *PlotStairs) SetValuesXYFloat64(x, y *float64, count int) *PlotStairs {

	p.x = x
	p.y = y
	p.count = count
	p.renderIndex = indexPlot2F64
	return p
}

// SetXscale sets the scale of the X axis (default = 1)
func (p *PlotStairs) SetXscale(xscale float64) *PlotStairs {

	p.xscale = xscale
	return p
}

// SetX0 sets the X coordinate of the start of the plot (default = 0)
func (p *PlotStairs) SetX0(x0 float64) *PlotStairs {

	p.x0 = x0
	return p
}

// SetOffset sets the offset from the start of the data in number of elements (default = 0)
func (p *PlotStairs) SetOffset(offset int) *PlotStairs {

	p.offset = offset
	return p
}

// SetStride sets the stride in number of elements (default = 1)
func (p *PlotStairs) SetStride(stride int) *PlotStairs {

	p.stride = stride
	return p
}
