package gig

import (
	"github.com/leonsal/gig/imgui"
)

type PlotLine struct {
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

const (
	indexPlotI32  = 0
	indexPlotF32  = 1
	indexPlotF64  = 2
	indexPlot2I32 = 3
	indexPlot2F32 = 4
	indexPlot2F64 = 5
	indexPlot3I32 = 6
	indexPlot3F32 = 7
	indexPlot3F64 = 8
)

func NewPlotLine(label string) *PlotLine {

	p := new(PlotLine)
	p.Init(p)
	p.labelId.SetString(label)
	p.xscale = 1
	p.stride = 1
	p.SetRender(func() {
		switch p.renderIndex {
		case indexPlotI32:
			imgui.PlotLineInt32CS(p.labelId, p.x.(*int32), p.count, p.xscale, p.x0, p.offset, p.stride)
		case indexPlotF32:
			imgui.PlotLineFloat32CS(p.labelId, p.x.(*float32), p.count, p.xscale, p.x0, p.offset, p.stride)
		case indexPlotF64:
			imgui.PlotLineFloat64CS(p.labelId, p.x.(*float64), p.count, p.xscale, p.x0, p.offset, p.stride)
		case indexPlot2I32:
			imgui.PlotLineXYInt32CS(p.labelId, p.x.(*int32), p.y.(*int32), p.count, p.offset, p.stride)
		case indexPlot2F32:
			imgui.PlotLineXYFloat32CS(p.labelId, p.x.(*float32), p.y.(*float32), p.count, p.offset, p.stride)
		case indexPlot2F64:
			imgui.PlotLineXYFloat64CS(p.labelId, p.x.(*float64), p.y.(*float64), p.count, p.offset, p.stride)
		}
	})
	return p
}

func (p *PlotLine) SetLabel(label string) *PlotLine {

	p.labelId.SetString(label)
	return p
}

func (p *PlotLine) SetValuesInt32(v *int32, count int) *PlotLine {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotI32
	return p
}

func (p *PlotLine) SetValuesFloat32(v *float32, count int) *PlotLine {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotF32
	return p
}

func (p *PlotLine) SetValuesFloat64(v *float64, count int) *PlotLine {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotF64
	return p
}

func (p *PlotLine) SetValuesXYInt32(x, y *int32, count int) *PlotLine {

	p.x = x
	p.y = y
	p.count = count
	p.renderIndex = indexPlot2I32
	return p
}

func (p *PlotLine) SetValuesXYFloat32(x, y *float32, count int) *PlotLine {

	p.x = x
	p.y = y
	p.count = count
	p.renderIndex = indexPlot2F32
	return p
}

func (p *PlotLine) SetValuesXYFloat64(x, y *float64, count int) *PlotLine {

	p.x = x
	p.y = y
	p.count = count
	p.renderIndex = indexPlot2F64
	return p
}

// SetXscale sets the scale of the X axis (default = 1)
func (p *PlotLine) SetXscale(xscale float64) *PlotLine {

	p.xscale = xscale
	return p
}

// SetX0 sets the X coordinate of the start of the plot (default = 0)
func (p *PlotLine) SetX0(x0 float64) *PlotLine {

	p.x0 = x0
	return p
}

// SetOffset sets the offset from the start of the data in number of elements (default = 0)
func (p *PlotLine) SetOffset(offset int) *PlotLine {

	p.offset = offset
	return p
}

// SetStride sets the stride in number of elements (default = 1)
func (p *PlotLine) SetStride(stride int) *PlotLine {

	p.stride = stride
	return p
}

type PlotShaded struct {
	Widget
	labelId     imgui.CString
	yref        float64
	x0          float64
	xscale      float64
	offset      int
	stride      int
	x           interface{}
	y1          interface{}
	y2          interface{}
	count       int
	renderIndex int
}

func NewPlotShaded(label string) *PlotShaded {

	p := new(PlotShaded)
	p.Init(p)
	p.labelId.SetString(label)
	p.xscale = 1
	p.stride = 1
	p.SetRender(func() {
		switch p.renderIndex {
		case indexPlotI32:
			imgui.PlotShadedInt32CS(p.labelId, p.x.(*int32), p.count, p.yref, p.xscale, p.x0, p.offset, p.stride)
		case indexPlotF32:
			imgui.PlotShadedFloat32CS(p.labelId, p.x.(*float32), p.count, p.yref, p.xscale, p.x0, p.offset, p.stride)
		case indexPlotF64:
			imgui.PlotShadedFloat64CS(p.labelId, p.x.(*float64), p.count, p.yref, p.xscale, p.x0, p.offset, p.stride)
		case indexPlot2I32:
			imgui.PlotShadedXYInt32CS(p.labelId, p.x.(*int32), p.y1.(*int32), p.count, p.yref, p.offset, p.stride)
		case indexPlot2F32:
			imgui.PlotShadedXYFloat32CS(p.labelId, p.x.(*float32), p.y1.(*float32), p.count, p.yref, p.offset, p.stride)
		case indexPlot2F64:
			imgui.PlotShadedXYFloat64CS(p.labelId, p.x.(*float64), p.y1.(*float64), p.count, p.yref, p.offset, p.stride)
		case indexPlot3I32:
			imgui.PlotShadedXYYInt32CS(p.labelId, p.x.(*int32), p.y1.(*int32), p.y2.(*int32), p.count, p.offset, p.stride)
		case indexPlot3F32:
			imgui.PlotShadedXYYFloat32CS(p.labelId, p.x.(*float32), p.y1.(*float32), p.y2.(*float32), p.count, p.offset, p.stride)
		case indexPlot3F64:
			imgui.PlotShadedXYYFloat64CS(p.labelId, p.x.(*float64), p.y1.(*float64), p.y2.(*float64), p.count, p.offset, p.stride)
		}
	})
	return p
}

func (p *PlotShaded) SetLabel(label string) *PlotShaded {

	p.labelId.SetString(label)
	return p
}

func (p *PlotShaded) SetValuesInt32(v *int32, count int) *PlotShaded {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotI32
	return p
}

func (p *PlotShaded) SetValuesFloat32(v *float32, count int) *PlotShaded {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotF32
	return p
}

func (p *PlotShaded) SetValuesFloat64(v *float64, count int) *PlotShaded {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotF64
	return p
}

func (p *PlotShaded) SetValuesXYInt32(x, y *int32, count int) *PlotShaded {

	p.x = x
	p.y1 = y
	p.count = count
	p.renderIndex = indexPlot2I32
	return p
}

func (p *PlotShaded) SetValuesXYFloat32(x, y *float32, count int) *PlotShaded {

	p.x = x
	p.y1 = y
	p.count = count
	p.renderIndex = indexPlot2F32
	return p
}

func (p *PlotShaded) SetValuesXYFloat64(x, y *float64, count int) *PlotShaded {

	p.x = x
	p.y1 = y
	p.count = count
	p.renderIndex = indexPlot2F64
	return p
}

func (p *PlotShaded) SetValuesXYYInt32(x, y1, y2 *int32, count int) *PlotShaded {

	p.x = x
	p.y1 = y1
	p.y2 = y2
	p.count = count
	p.renderIndex = indexPlot3I32
	return p
}

func (p *PlotShaded) SetValuesXYYFloat32(x, y1, y2 *float32, count int) *PlotShaded {

	p.x = x
	p.y1 = y1
	p.y2 = y2
	p.count = count
	p.renderIndex = indexPlot3F32
	return p
}

func (p *PlotShaded) SetValuesXYYFloat64(x, y1, y2 *float64, count int) *PlotShaded {

	p.x = x
	p.y1 = y1
	p.y2 = y2
	p.count = count
	p.renderIndex = indexPlot3F64
	return p
}

func (p *PlotShaded) SetYref(yref float64) *PlotShaded {

	p.yref = yref
	return p
}

// SetXscale sets the scale of the X axis (default = 1)
func (p *PlotShaded) SetXscale(xscale float64) *PlotShaded {

	p.xscale = xscale
	return p
}

// SetX0 sets the X coordinate of the start of the plot (default = 0)
func (p *PlotShaded) SetX0(x0 float64) *PlotShaded {

	p.x0 = x0
	return p
}

// SetOffset sets the offset from the start of the data in number of elements (default = 0)
func (p *PlotShaded) SetOffset(offset int) *PlotShaded {

	p.offset = offset
	return p
}

// SetStride sets the stride in number of elements (default = 1)
func (p *PlotShaded) SetStride(stride int) *PlotShaded {

	p.stride = stride
	return p
}
