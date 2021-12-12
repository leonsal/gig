package gig

import (
	"github.com/leonsal/gig/imgui"
)

type PlotScatter struct {
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

func NewPlotScatter(label string) *PlotScatter {

	p := new(PlotScatter)
	p.Init(p)
	p.labelId.SetString(label)
	p.xscale = 1
	p.stride = 1
	p.SetRender(func() {
		switch p.renderIndex {
		case indexPlotI32:
			imgui.PlotScatterInt32CS(p.labelId, p.x.(*int32), p.count, p.xscale, p.x0, p.offset, p.stride)
		case indexPlotF32:
			imgui.PlotScatterFloat32CS(p.labelId, p.x.(*float32), p.count, p.xscale, p.x0, p.offset, p.stride)
		case indexPlotF64:
			imgui.PlotScatterFloat64CS(p.labelId, p.x.(*float64), p.count, p.xscale, p.x0, p.offset, p.stride)
		case indexPlot2I32:
			imgui.PlotScatterXYInt32CS(p.labelId, p.x.(*int32), p.y.(*int32), p.count, p.offset, p.stride)
		case indexPlot2F32:
			imgui.PlotScatterXYFloat32CS(p.labelId, p.x.(*float32), p.y.(*float32), p.count, p.offset, p.stride)
		case indexPlot2F64:
			imgui.PlotScatterXYFloat64CS(p.labelId, p.x.(*float64), p.y.(*float64), p.count, p.offset, p.stride)
		}
	})
	return p
}

func (p *PlotScatter) SetLabel(label string) *PlotScatter {

	p.labelId.SetString(label)
	return p
}

func (p *PlotScatter) SetValuesInt32(v *int32, count int) *PlotScatter {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotI32
	return p
}

func (p *PlotScatter) SetValuesFloat32(v *float32, count int) *PlotScatter {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotF32
	return p
}

func (p *PlotScatter) SetValuesFloat64(v *float64, count int) *PlotScatter {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotF64
	return p
}

func (p *PlotScatter) SetValuesXYInt32(x, y *int32, count int) *PlotScatter {

	p.x = x
	p.y = y
	p.count = count
	p.renderIndex = indexPlot2I32
	return p
}

func (p *PlotScatter) SetValuesXYFloat32(x, y *float32, count int) *PlotScatter {

	p.x = x
	p.y = y
	p.count = count
	p.renderIndex = indexPlot2F32
	return p
}

func (p *PlotScatter) SetValuesXYFloat64(x, y *float64, count int) *PlotScatter {

	p.x = x
	p.y = y
	p.count = count
	p.renderIndex = indexPlot2F64
	return p
}

// SetXscale sets the scale of the X axis (default = 1)
func (p *PlotScatter) SetXscale(xscale float64) *PlotScatter {

	p.xscale = xscale
	return p
}

// SetX0 sets the X coordinate of the start of the plot (default = 0)
func (p *PlotScatter) SetX0(x0 float64) *PlotScatter {

	p.x0 = x0
	return p
}

// SetOffset sets the offset from the start of the data in number of elements (default = 0)
func (p *PlotScatter) SetOffset(offset int) *PlotScatter {

	p.offset = offset
	return p
}

// SetStride sets the stride in number of elements (default = 1)
func (p *PlotScatter) SetStride(stride int) *PlotScatter {

	p.stride = stride
	return p
}
