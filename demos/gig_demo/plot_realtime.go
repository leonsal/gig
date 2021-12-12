package main

import (
	"math"

	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
)

func init() {

	registerDemo(catImPlot, "Realtime", NewPlotRealtimeDemo)
}

type ScrollingBuffer struct {
	maxSize int
	offset  int
	data    []imgui.Vec2
}

func NewScrollingBuffer(size int) *ScrollingBuffer {

	return &ScrollingBuffer{
		maxSize: size,
		offset:  0,
		data:    make([]imgui.Vec2, 0, size),
	}
}

func (sb *ScrollingBuffer) AddPoint(x, y float32) {

	if len(sb.data) < sb.maxSize {
		sb.data = append(sb.data, imgui.Vec2{x, y})
	} else {
		sb.data[sb.offset] = imgui.Vec2{x, y}
		sb.offset = (sb.offset + 1) % sb.maxSize
	}
}

type RollingBuffer struct {
	span float32
	data []imgui.Vec2
}

func NewRollingBuffer(size int) *RollingBuffer {

	return &RollingBuffer{
		span: 10,
		data: make([]imgui.Vec2, 0, size),
	}
}

func (rb *RollingBuffer) AddPoint(x, y float32) {

	xmod := float32(math.Mod(float64(x), float64(rb.span)))
	if len(rb.data) != 0 && xmod < rb.data[len(rb.data)-1].X {
		rb.data = rb.data[:0]
	}
	rb.data = append(rb.data, imgui.Vec2{xmod, y})
}

func (rb *RollingBuffer) Px() *float32 {

	if len(rb.data) == 0 {
		return nil
	}
	return &rb.data[0].X
}

func (rb *RollingBuffer) Py() *float32 {

	if len(rb.data) == 0 {
		return nil
	}
	return &rb.data[0].Y
}

type realtimeState struct {
	rdata1  *RollingBuffer
	rdata2  *RollingBuffer
	sdata1  *ScrollingBuffer
	sdata2  *ScrollingBuffer
	t       float32
	history float32
}

var rtData = realtimeState{
	rdata1:  NewRollingBuffer(2000),
	rdata2:  NewRollingBuffer(2000),
	sdata1:  NewScrollingBuffer(2000),
	sdata2:  NewScrollingBuffer(2000),
	t:       0,
	history: 10,
}

func realtimeStep() {

	mouse := imgui.GetMousePos()
	rtData.t += imgui.GetIO().DeltaTime()
	rtData.sdata1.AddPoint(rtData.t, mouse.X*0.0005)
	rtData.rdata1.AddPoint(rtData.t, mouse.X*0.0005)
	rtData.sdata2.AddPoint(rtData.t, mouse.Y*0.0005)
	rtData.rdata2.AddPoint(rtData.t, mouse.Y*0.0005)
}

func NewPlotRealtimeDemo() gig.IWidget {

	group := gig.NewGroup().AddChildren(

		gig.NewPlot("Realtime scrolling buffer").
			SetXAxisFlags(imgui.PlotAxisFlags_NoTickLabels).
			SetYAxisFlags(imgui.PlotAxisFlags_NoTickLabels).
			SetLimitsY(0, 1, imgui.Cond_Once, imgui.PlotYAxis_1).
			SetPreRender(func(w gig.IWidget) {
				p := w.(*gig.Plot)
				realtimeStep()
				p.SetLimitsX(float64(rtData.t-10), float64(rtData.t), imgui.Cond_Always)
			}).
			AddChildren(

				gig.NewPlotShaded("Mouse X").SetStride(2).
					SetPreRender(func(w gig.IWidget) {
						pl := w.(*gig.PlotShaded)
						pl.SetValuesXYFloat32(&rtData.sdata1.data[0].X, &rtData.sdata1.data[0].Y, len(rtData.sdata1.data))
						pl.SetOffset(rtData.sdata1.offset)
					}),

				gig.NewPlotLine("Mouse Y").SetStride(2).
					SetPreRender(func(w gig.IWidget) {
						pl := w.(*gig.PlotLine)
						pl.SetValuesXYFloat32(&rtData.sdata2.data[0].X, &rtData.sdata2.data[0].Y, len(rtData.sdata2.data))
						pl.SetOffset(rtData.sdata2.offset)
					}),
			),
		gig.NewPlot("Realtime rolling buffer").
			SetXAxisFlags(imgui.PlotAxisFlags_NoTickLabels).
			SetYAxisFlags(imgui.PlotAxisFlags_NoTickLabels).
			SetLimitsY(0, 1, imgui.Cond_Once, imgui.PlotYAxis_1).
			SetPreRender(func(w gig.IWidget) {
				w.(*gig.Plot).SetLimitsX(0, float64(rtData.history), imgui.Cond_Always)
			}).
			AddChildren(
				gig.NewPlotLine("Mouse X").SetStride(2).
					SetPreRender(func(w gig.IWidget) {
						w.(*gig.PlotLine).SetValuesXYFloat32(rtData.rdata1.Px(), rtData.rdata1.Py(), len(rtData.rdata1.data))
					}),
				gig.NewPlotLine("Mouse Y").SetStride(2).
					SetPreRender(func(w gig.IWidget) {
						w.(*gig.PlotLine).SetValuesXYFloat32(rtData.rdata2.Px(), rtData.rdata2.Py(), len(rtData.rdata2.data))
					}),
			),
	)
	return group
}
