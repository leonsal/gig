package main

import (
	"math"
	"math/rand"

	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
)

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

func NewPlotDemo() gig.IWidget {

	genSineF32 := func(amplitude, offset, nf float64, count int) []float32 {

		const period = 2 * math.Pi
		phase := 0.0
		delta := period * nf
		out := make([]float32, count)
		for i := 0; i < count; i++ {
			out[i] = float32(amplitude*math.Sin(phase) + offset)
			phase += delta
			for phase > period {
				phase -= period
			}
		}
		return out
	}

	genSineF64 := func(amplitude, offset, nf float64, count int) []float64 {

		const period = 2 * math.Pi
		phase := 0.0
		delta := period * nf
		out := make([]float64, count)
		for i := 0; i < count; i++ {
			out[i] = amplitude*math.Sin(phase) + offset
			phase += delta
			for phase > period {
				phase -= period
			}
		}
		return out
	}

	genSeqF32 := func(nf float64, count int) []float32 {

		delta := 2 * math.Pi * nf
		out := make([]float32, count)
		for i := 0; i < count; i++ {
			out[i] = float32(delta * float64(i))
		}
		return out
	}

	win := gig.NewWindow("Plot").AddChildren(

		gig.NewPlot("Plot lines").AddChildren(
			gig.NewPlotLine("Sine1").SetValuesFloat32(&genSineF32(1, 0, 0.05, 400)[0], 400),
			gig.NewPlotLine("Sine2").SetValuesFloat32(&genSineF32(0.5, 0.5, 0.02, 400)[0], 400),
			gig.NewPlotLine("Sine3").SetValuesFloat64(&genSineF64(0.3, -0.5, 0.01, 400)[0], 400),
			gig.NewPlotLine("XY").SetValuesXYFloat32(
				&[]float32{0, 100, 150, 310, 400}[0], &[]float32{1, 0.5, -1, 0, 0.3}[0], 5),
		),

		gig.NewPlot("Plot filled").AddChildren(
			gig.NewPlotShaded("Sine1").SetValuesFloat32(&genSineF32(1, 0, 0.05, 400)[0], 400),
			gig.NewPlotShaded("Sine2").SetValuesFloat32(&genSineF32(0.2, 0, 0.03, 400)[0], 400),
		),

		gig.NewPlot("Plot shaded").AddChildren(
			gig.NewPlotShaded("Sine1").SetValuesXYYFloat32(
				&genSeqF32(0.03, 400)[0],
				&genSineF32(0.2, 0, 0.03, 400)[0],
				&genSineF32(0.2, 0.3, 0.03, 400)[0],
				400,
			),
		),

		gig.NewPlot("Plot scatter").AddChildren(
			gig.NewPlotScatter("Data1").Call(func(w gig.IWidget) {
				ps := w.(*gig.PlotScatter)
				count := 100
				xs := make([]float32, count)
				ys := make([]float32, count)
				for i := 0; i < count; i++ {
					xs[i] = float32(i) * 0.01
					ys[i] = xs[i] + 0.1*rand.Float32()
				}
				ps.SetValuesXYFloat32(&xs[0], &ys[0], len(xs))
			}),
			gig.NewPlotScatter("Data2").Call(func(w gig.IWidget) {
				ps := w.(*gig.PlotScatter)
				count := 50
				xs := make([]float32, count)
				ys := make([]float32, count)
				for i := 0; i < count; i++ {
					xs[i] = 0.25 + 0.2*rand.Float32()
					ys[i] = 0.75 + 0.2*rand.Float32()
				}
				ps.SetValuesXYFloat32(&xs[0], &ys[0], len(xs))
			}),
		),

		gig.NewPlot("Realtime scrolling buffer").
			SetFlags(imgui.PlotAxisFlags_NoTickLabels).
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
			SetFlags(imgui.PlotAxisFlags_NoTickLabels).
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
	).SetVisible(false).SaveRef("plotDemo")

	return win
}
