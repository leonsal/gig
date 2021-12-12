package gig

import (
	"github.com/leonsal/gig/imgui"
)

type PlotVLines struct {
	Widget
	labelId     imgui.CString
	offset      int
	stride      int
	x           interface{}
	count       int
	renderIndex int
}

func NewPlotVLines(label string) *PlotVLines {

	p := new(PlotVLines)
	p.Init(p)
	p.labelId.SetString(label)
	p.stride = 1
	p.SetRender(func() {
		switch p.renderIndex {
		case indexPlotI32:
			imgui.PlotVLinesInt32CS(p.labelId, p.x.(*int32), p.count, p.offset, p.stride)
		case indexPlotF32:
			imgui.PlotVLinesFloat32CS(p.labelId, p.x.(*float32), p.count, p.offset, p.stride)
		case indexPlotF64:
			imgui.PlotVLinesFloat64CS(p.labelId, p.x.(*float64), p.count, p.offset, p.stride)
		}
	})
	return p
}

func (p *PlotVLines) SetLabel(label string) *PlotVLines {

	p.labelId.SetString(label)
	return p
}

func (p *PlotVLines) SetValuesInt32(v *int32, count int) *PlotVLines {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotI32
	return p
}

func (p *PlotVLines) SetValuesFloat32(v *float32, count int) *PlotVLines {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotF32
	return p
}

func (p *PlotVLines) SetValuesFloat64(v *float64, count int) *PlotVLines {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotF64
	return p
}

// SetOffset sets the offset from the start of the data in number of elements (default = 0)
func (p *PlotVLines) SetOffset(offset int) *PlotVLines {

	p.offset = offset
	return p
}

// SetStride sets the stride in number of elements (default = 1)
func (p *PlotVLines) SetStride(stride int) *PlotVLines {

	p.stride = stride
	return p
}

type PlotHLines struct {
	Widget
	labelId     imgui.CString
	offset      int
	stride      int
	x           interface{}
	count       int
	renderIndex int
}

func NewPlotHLines(label string) *PlotHLines {

	p := new(PlotHLines)
	p.Init(p)
	p.labelId.SetString(label)
	p.stride = 1
	p.SetRender(func() {
		switch p.renderIndex {
		case indexPlotI32:
			imgui.PlotHLinesInt32CS(p.labelId, p.x.(*int32), p.count, p.offset, p.stride)
		case indexPlotF32:
			imgui.PlotHLinesFloat32CS(p.labelId, p.x.(*float32), p.count, p.offset, p.stride)
		case indexPlotF64:
			imgui.PlotHLinesFloat64CS(p.labelId, p.x.(*float64), p.count, p.offset, p.stride)
		}
	})
	return p
}

func (p *PlotHLines) SetLabel(label string) *PlotHLines {

	p.labelId.SetString(label)
	return p
}

func (p *PlotHLines) SetValuesInt32(v *int32, count int) *PlotHLines {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotI32
	return p
}

func (p *PlotHLines) SetValuesFloat32(v *float32, count int) *PlotHLines {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotF32
	return p
}

func (p *PlotHLines) SetValuesFloat64(v *float64, count int) *PlotHLines {

	p.x = v
	p.count = count
	p.renderIndex = indexPlotF64
	return p
}

// SetOffset sets the offset from the start of the data in number of elements (default = 0)
func (p *PlotHLines) SetOffset(offset int) *PlotHLines {

	p.offset = offset
	return p
}

// SetStride sets the stride in number of elements (default = 1)
func (p *PlotHLines) SetStride(stride int) *PlotHLines {

	p.stride = stride
	return p
}
