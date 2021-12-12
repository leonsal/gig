package gig

import (
	"github.com/leonsal/gig/imgui"
)

type DragInt32 struct {
	Widget
	label    imgui.CString
	value    int32
	speed    float32
	min      int32
	max      int32
	format   imgui.CString
	flags    imgui.SliderFlags
	onchange func(*DragInt32)
}

func NewDragInt32(label string, value int32) *DragInt32 {

	d := new(DragInt32)
	d.Init(d)
	d.label.SetString(label)
	d.value = value
	d.speed = 1.0
	d.SetRender(func() {
		if imgui.DragIntCS(d.label, &d.value, d.speed, d.min, d.max, d.format, d.flags) {
			if d.onchange != nil {
				d.onchange(d)
			}
		}
	})
	return d
}

func (d *DragInt32) SetLabel(label string) *DragInt32 {

	d.label.SetString(label)
	return d
}

func (d *DragInt32) Value() int32 {

	return d.value
}

func (d *DragInt32) SetValue(v int32) *DragInt32 {

	d.value = v
	return d
}

func (d *DragInt32) Speed() float32 {

	return d.speed
}

func (d *DragInt32) SetSpeed(speed float32) *DragInt32 {

	d.speed = speed
	return d
}

func (d *DragInt32) Min() int32 {

	return d.min
}

func (d *DragInt32) SetMin(min int32) *DragInt32 {

	d.min = min
	return d
}

func (d *DragInt32) Max() int32 {

	return d.max
}

func (d *DragInt32) SetMax(max int32) *DragInt32 {

	d.max = max
	return d
}

func (d *DragInt32) SetFormat(format string) *DragInt32 {

	d.format.SetString(format)
	return d
}

func (d *DragInt32) Flags() imgui.SliderFlags {

	return d.flags
}

func (d *DragInt32) SetFlags(flags imgui.SliderFlags) *DragInt32 {

	d.flags = flags
	return d
}

func (d *DragInt32) SetOnChange(f func(*DragInt32)) *DragInt32 {

	d.onchange = f
	return d
}

type DragInt32Slice struct {
	Widget
	label    imgui.CString
	values   []int32
	speed    float32
	min      *int32
	max      *int32
	format   imgui.CString
	flags    imgui.SliderFlags
	onchange func(*DragInt32Slice)
}

func NewDragInt32Slice(label string, values []int32) *DragInt32Slice {

	d := new(DragInt32Slice)
	d.Init(d)
	d.label.SetString(label)
	d.values = values
	d.speed = 1.0
	d.SetRender(func() {
		if imgui.DragInt32N(d.label, &d.values[0], len(d.values), d.speed, d.min, d.max, d.format, d.flags) {
			if d.onchange != nil {
				d.onchange(d)
			}
		}
	})
	return d
}

func (d *DragInt32Slice) SetLabel(label string) *DragInt32Slice {

	d.label.SetString(label)
	return d
}

func (d *DragInt32Slice) Values() []int32 {

	return d.values
}

func (d *DragInt32Slice) SetValues(v []int32) *DragInt32Slice {

	d.values = v
	return d
}

func (d *DragInt32Slice) Speed() float32 {

	return d.speed
}

func (d *DragInt32Slice) SetSpeed(speed float32) *DragInt32Slice {

	d.speed = speed
	return d
}

func (d *DragInt32Slice) SetMin(min int32) *DragInt32Slice {

	if d.min == nil {
		d.min = new(int32)
	}
	*d.min = min
	return d
}

func (d *DragInt32Slice) SetMax(max int32) *DragInt32Slice {

	if d.max == nil {
		d.max = new(int32)
	}
	*d.max = max
	return d
}

func (d *DragInt32Slice) SetFormat(format string) *DragInt32Slice {

	d.format.SetString(format)
	return d
}

func (d *DragInt32Slice) Flags() imgui.SliderFlags {

	return d.flags
}

func (d *DragInt32Slice) SetFlags(flags imgui.SliderFlags) *DragInt32Slice {

	d.flags = flags
	return d
}

func (d *DragInt32Slice) SetOnChange(f func(*DragInt32Slice)) *DragInt32Slice {

	d.onchange = f
	return d
}

type DragInt32Range2 struct {
	Widget
	label      imgui.CString
	currentMin int32
	currentMax int32
	speed      float32
	min        int32
	max        int32
	format     imgui.CString
	formatMax  imgui.CString
	flags      imgui.SliderFlags
	onchange   func(*DragInt32Range2)
}

func NewDragIntRange2(label string, cmin, cmax int32) *DragInt32Range2 {

	d := new(DragInt32Range2)
	d.Init(d)
	d.label.SetString(label)
	d.currentMin = cmin
	d.currentMax = cmax
	d.speed = 1.0
	d.SetRender(func() {
		if imgui.DragIntRange2CS(d.label, &d.currentMin, &d.currentMax, d.speed, d.min, d.max, d.format, d.formatMax, d.flags) {
			if d.onchange != nil {
				d.onchange(d)
			}
		}
	})
	return d
}

func (d *DragInt32Range2) SetLabel(label string) *DragInt32Range2 {

	d.label.SetString(label)
	return d
}

func (d *DragInt32Range2) CurrentMin() int32 {

	return d.currentMin
}

func (d *DragInt32Range2) SetCurrentMin(cmin int32) *DragInt32Range2 {

	d.currentMin = cmin
	return d
}

func (d *DragInt32Range2) CurrentMax() int32 {

	return d.currentMax
}

func (d *DragInt32Range2) SetCurrentMax(cmax int32) *DragInt32Range2 {

	d.currentMax = cmax
	return d
}

func (d *DragInt32Range2) SetSpeed(speed float32) *DragInt32Range2 {

	d.speed = speed
	return d
}

func (d *DragInt32Range2) Min() int32 {

	return d.min
}

func (d *DragInt32Range2) SetMin(min int32) *DragInt32Range2 {

	d.min = min
	return d
}

func (d *DragInt32Range2) Max() int32 {

	return d.max
}

func (d *DragInt32Range2) SetMax(max int32) *DragInt32Range2 {

	d.max = max
	return d
}

func (d *DragInt32Range2) SetFormat(format string) *DragInt32Range2 {

	d.format.SetString(format)
	return d
}

func (d *DragInt32Range2) SetFormatMax(format string) *DragInt32Range2 {

	d.formatMax.SetString(format)
	return d
}

func (d *DragInt32Range2) SetFlags(flags imgui.SliderFlags) *DragInt32Range2 {

	d.flags = flags
	return d
}

func (d *DragInt32Range2) SetOnChange(f func(*DragInt32Range2)) *DragInt32Range2 {

	d.onchange = f
	return d
}

type DragInt64 struct {
	Widget
	label    imgui.CString
	value    int64
	speed    float32
	min      *int64
	max      *int64
	format   imgui.CString
	flags    imgui.SliderFlags
	onchange func(*DragInt64)
}

func NewDragInt64(label string, value int64) *DragInt64 {

	d := new(DragInt64)
	d.Init(d)
	d.label.SetString(label)
	d.value = value
	d.speed = 1.0
	d.SetRender(func() {
		if imgui.DragInt64N(d.label, &d.value, 1, d.speed, d.min, d.max, d.format, d.flags) {
			if d.onchange != nil {
				d.onchange(d)
			}
		}
	})
	return d
}

func (d *DragInt64) SetLabel(label string) *DragInt64 {

	d.label.SetString(label)
	return d
}

func (d *DragInt64) Value() int64 {

	return d.value
}

func (d *DragInt64) SetValue(v int64) *DragInt64 {

	d.value = v
	return d
}

func (d *DragInt64) Speed() float32 {

	return d.speed
}

func (d *DragInt64) SetSpeed(speed float32) *DragInt64 {

	d.speed = speed
	return d
}

func (d *DragInt64) SetMin(min int64) *DragInt64 {

	if d.min == nil {
		d.min = new(int64)
	}
	*d.min = min
	return d
}

func (d *DragInt64) SetMax(max int64) *DragInt64 {

	if d.max == nil {
		d.max = new(int64)
	}
	*d.max = max
	return d
}

func (d *DragInt64) SetFormat(format string) *DragInt64 {

	d.format.SetString(format)
	return d
}

func (d *DragInt64) Flags() imgui.SliderFlags {

	return d.flags
}

func (d *DragInt64) SetFlags(flags imgui.SliderFlags) *DragInt64 {

	d.flags = flags
	return d
}

func (d *DragInt64) SetOnChange(f func(*DragInt64)) *DragInt64 {

	d.onchange = f
	return d
}

type DragInt64Slice struct {
	Widget
	label    imgui.CString
	values   []int64
	speed    float32
	min      *int64
	max      *int64
	format   imgui.CString
	flags    imgui.SliderFlags
	onchange func(*DragInt64Slice)
}

func NewDragInt64Slice(label string, values []int64) *DragInt64Slice {

	d := new(DragInt64Slice)
	d.Init(d)
	d.label.SetString(label)
	d.values = make([]int64, len(values))
	d.values = values
	d.speed = 1.0
	d.SetRender(func() {
		if imgui.DragInt64N(d.label, &d.values[0], len(d.values), d.speed, d.min, d.max, d.format, d.flags) {
			if d.onchange != nil {
				d.onchange(d)
			}
		}
	})
	return d
}

func (d *DragInt64Slice) SetLabel(label string) *DragInt64Slice {

	d.label.SetString(label)
	return d
}

func (d *DragInt64Slice) Values() []int64 {

	return d.values
}

func (d *DragInt64Slice) SetValues(v []int64) *DragInt64Slice {

	d.values = v
	return d
}

func (d *DragInt64Slice) Speed() float32 {

	return d.speed
}

func (d *DragInt64Slice) SetSpeed(speed float32) *DragInt64Slice {

	d.speed = speed
	return d
}

func (d *DragInt64Slice) SetMin(min int64) *DragInt64Slice {

	if d.min == nil {
		d.min = new(int64)
	}
	*d.min = min
	return d
}

func (d *DragInt64Slice) SetMax(max int64) *DragInt64Slice {

	if d.max == nil {
		d.max = new(int64)
	}
	*d.max = max
	return d
}

func (d *DragInt64Slice) SetFormat(format string) *DragInt64Slice {

	d.format.SetString(format)
	return d
}

func (d *DragInt64Slice) Flags() imgui.SliderFlags {

	return d.flags
}

func (d *DragInt64Slice) SetFlags(flags imgui.SliderFlags) *DragInt64Slice {

	d.flags = flags
	return d
}

func (d *DragInt64Slice) SetOnChange(f func(*DragInt64Slice)) *DragInt64Slice {

	d.onchange = f
	return d
}

type DragFloat32 struct {
	Widget
	label    imgui.CString
	value    float32
	speed    float32
	min      float32
	max      float32
	format   imgui.CString
	flags    imgui.SliderFlags
	onchange func(*DragFloat32)
}

func NewDragFloat32(label string, value float32) *DragFloat32 {

	d := new(DragFloat32)
	d.Init(d)
	d.label.SetString(label)
	d.value = value
	d.speed = 1.0
	d.format.SetString("%.3f")
	d.SetRender(func() {
		if imgui.DragFloatCS(d.label, &d.value, d.speed, d.min, d.max, d.format, d.flags) {
			if d.onchange != nil {
				d.onchange(d)
			}
		}
	})
	return d
}

func (d *DragFloat32) SetLabel(label string) *DragFloat32 {

	d.label.SetString(label)
	return d
}

func (d *DragFloat32) Value() float32 {

	return float32(d.value)
}

func (d *DragFloat32) SetValue(v float32) *DragFloat32 {

	d.value = v
	return d
}

func (d *DragFloat32) Speed() float32 {

	return d.speed
}

func (d *DragFloat32) SetSpeed(speed float32) *DragFloat32 {

	d.speed = speed
	return d
}

func (d *DragFloat32) Min() float32 {

	return d.min
}

func (d *DragFloat32) SetMin(min float32) *DragFloat32 {

	d.min = min
	return d
}

func (d *DragFloat32) Max() float32 {

	return d.min
}

func (d *DragFloat32) SetMax(max float32) *DragFloat32 {

	d.max = max
	return d
}

func (d *DragFloat32) SetFormat(format string) *DragFloat32 {

	d.format.SetString(format)
	return d
}

func (d *DragFloat32) Flags() imgui.SliderFlags {

	return d.flags
}

func (d *DragFloat32) SetFlags(flags imgui.SliderFlags) *DragFloat32 {

	d.flags = flags
	return d
}

func (d *DragFloat32) SetOnChange(f func(*DragFloat32)) *DragFloat32 {

	d.onchange = f
	return d
}

type DragFloat32Slice struct {
	Widget
	label    imgui.CString
	values   []float32
	speed    float32
	min      *float32
	max      *float32
	format   imgui.CString
	flags    imgui.SliderFlags
	onchange func(*DragFloat32Slice)
}

func NewDragFloat32Slice(label string, values []float32) *DragFloat32Slice {

	d := new(DragFloat32Slice)
	d.Init(d)
	d.label.SetString(label)
	d.values = values
	d.speed = 1.0
	d.SetRender(func() {
		if imgui.DragFloat32N(d.label, &d.values[0], len(d.values), d.speed, d.min, d.max, d.format, d.flags) {
			if d.onchange != nil {
				d.onchange(d)
			}
		}
	})
	return d
}

func (d *DragFloat32Slice) SetLabel(label string) *DragFloat32Slice {

	d.label.SetString(label)
	return d
}

func (d *DragFloat32Slice) Values() []float32 {

	return d.values
}

func (d *DragFloat32Slice) SetValues(values []float32) *DragFloat32Slice {

	d.values = values
	return d
}

func (d *DragFloat32Slice) Speed() float32 {

	return d.speed
}

func (d *DragFloat32Slice) SetSpeed(speed float32) *DragFloat32Slice {

	d.speed = speed
	return d
}

func (d *DragFloat32Slice) SetMin(min float32) *DragFloat32Slice {

	if d.min == nil {
		d.min = new(float32)
	}
	*d.min = min
	return d
}

func (d *DragFloat32Slice) SetMax(max float32) *DragFloat32Slice {

	if d.max == nil {
		d.max = new(float32)
	}
	*d.max = max
	return d
}

func (d *DragFloat32Slice) SetFormat(format string) *DragFloat32Slice {

	d.format.SetString(format)
	return d
}

func (d *DragFloat32Slice) Flags() imgui.SliderFlags {

	return d.flags
}

func (d *DragFloat32Slice) SetFlags(flags imgui.SliderFlags) *DragFloat32Slice {

	d.flags = flags
	return d
}

func (d *DragFloat32Slice) SetOnChange(f func(*DragFloat32Slice)) *DragFloat32Slice {

	d.onchange = f
	return d
}

type DragFloat32Range2 struct {
	Widget
	label      imgui.CString
	currentMin float32
	currentMax float32
	speed      float32
	min        float32
	max        float32
	format     imgui.CString
	formatMax  imgui.CString
	flags      imgui.SliderFlags
	onchange   func(*DragFloat32Range2)
}

func NewDragFloat32Range2(label string, cmin, cmax float32) *DragFloat32Range2 {

	d := new(DragFloat32Range2)
	d.Init(d)
	d.label.SetString(label)
	d.currentMin = cmin
	d.currentMax = cmax
	d.speed = 1.0
	d.SetRender(func() {
		if imgui.DragFloatRange2CS(d.label, &d.currentMin, &d.currentMax, d.speed,
			d.min, d.max, d.format, d.formatMax, d.flags) {
			if d.onchange != nil {
				d.onchange(d)
			}
		}
	})
	return d
}

func (d *DragFloat32Range2) SetLabel(label string) *DragFloat32Range2 {

	d.label.SetString(label)
	return d
}

func (d *DragFloat32Range2) CurrentMin() float32 {

	return d.currentMin
}

func (d *DragFloat32Range2) SetCurrentMin(cmin float32) *DragFloat32Range2 {

	d.currentMin = cmin
	return d
}

func (d *DragFloat32Range2) CurrentMax() float32 {

	return d.currentMax
}

func (d *DragFloat32Range2) SetCurrentMax(cmax float32) *DragFloat32Range2 {

	d.currentMax = cmax
	return d
}

func (d *DragFloat32Range2) Speed() float32 {

	return d.speed
}

func (d *DragFloat32Range2) SetSpeed(speed float32) *DragFloat32Range2 {

	d.speed = speed
	return d
}

func (d *DragFloat32Range2) SetMin(min float32) *DragFloat32Range2 {

	d.min = min
	return d
}

func (d *DragFloat32Range2) SetMax(max float32) *DragFloat32Range2 {

	d.max = max
	return d
}

func (d *DragFloat32Range2) SetFormat(format string) *DragFloat32Range2 {

	d.format.SetString(format)
	return d
}

func (d *DragFloat32Range2) SetFormatMax(format string) *DragFloat32Range2 {

	d.formatMax.SetString(format)
	return d
}

func (d *DragFloat32Range2) Flags() imgui.SliderFlags {

	return d.flags
}

func (d *DragFloat32Range2) SetFlags(flags imgui.SliderFlags) *DragFloat32Range2 {

	d.flags = flags
	return d
}

func (d *DragFloat32Range2) SetOnChange(f func(*DragFloat32Range2)) *DragFloat32Range2 {

	d.onchange = f
	return d
}

type DragFloat64 struct {
	Widget
	label    imgui.CString
	value    float64
	speed    float32
	min      *float64
	max      *float64
	format   imgui.CString
	flags    imgui.SliderFlags
	onchange func(*DragFloat64)
}

func NewDragFloat64(label string, value float64) *DragFloat64 {

	d := new(DragFloat64)
	d.Init(d)
	d.label.SetString(label)
	d.value = value
	d.speed = 1.0
	d.SetRender(func() {
		if imgui.DragFloat64N(d.label, &d.value, 1, d.speed, d.min, d.max, d.format, d.flags) {
			if d.onchange != nil {
				d.onchange(d)
			}
		}
	})
	return d
}

func (d *DragFloat64) SetLabel(label string) *DragFloat64 {

	d.label.SetString(label)
	return d
}

func (d *DragFloat64) Value() float64 {

	return d.value
}

func (d *DragFloat64) SetValue(v float64) *DragFloat64 {

	d.value = v
	return d
}

func (d *DragFloat64) Speed() float32 {

	return d.speed
}

func (d *DragFloat64) SetSpeed(speed float32) *DragFloat64 {

	d.speed = speed
	return d
}

func (d *DragFloat64) SetMin(min float64) *DragFloat64 {

	if d.min == nil {
		d.min = new(float64)
	}
	*d.min = min
	return d
}

func (d *DragFloat64) SetMax(max float64) *DragFloat64 {

	if d.max == nil {
		d.max = new(float64)
	}
	*d.max = max
	return d
}

func (d *DragFloat64) SetFormat(format string) *DragFloat64 {

	d.format.SetString(format)
	return d
}

func (d *DragFloat64) Flags() imgui.SliderFlags {

	return d.flags
}

func (d *DragFloat64) SetFlags(flags imgui.SliderFlags) *DragFloat64 {

	d.flags = flags
	return d
}

func (d *DragFloat64) SetOnChange(f func(*DragFloat64)) *DragFloat64 {

	d.onchange = f
	return d
}

type DragFloat64Slice struct {
	Widget
	label    imgui.CString
	values   []float64
	speed    float32
	min      *float64
	max      *float64
	format   imgui.CString
	flags    imgui.SliderFlags
	onchange func(*DragFloat64Slice)
}

func NewDragFloat64Slice(label string, values []float64) *DragFloat64Slice {

	d := new(DragFloat64Slice)
	d.Init(d)
	d.label.SetString(label)
	d.values = values
	d.speed = 1.0
	d.SetRender(func() {
		if imgui.DragFloat64N(d.label, &d.values[0], len(d.values), d.speed, d.min, d.max, d.format, d.flags) {
			if d.onchange != nil {
				d.onchange(d)
			}
		}
	})
	return d
}

func (d *DragFloat64Slice) SetLabel(label string) *DragFloat64Slice {

	d.label.SetString(label)
	return d
}

func (d *DragFloat64Slice) Values() []float64 {

	return d.values
}

func (d *DragFloat64Slice) SetValues(values []float64) *DragFloat64Slice {

	d.values = values
	return d
}

func (d *DragFloat64Slice) Speed() float32 {

	return d.speed
}

func (d *DragFloat64Slice) SetSpeed(speed float32) *DragFloat64Slice {

	d.speed = speed
	return d
}

func (d *DragFloat64Slice) SetMin(min float64) *DragFloat64Slice {

	if d.min == nil {
		d.min = new(float64)
	}
	*d.min = min
	return d
}

func (d *DragFloat64Slice) SetMax(max float64) *DragFloat64Slice {

	if d.max == nil {
		d.max = new(float64)
	}
	*d.max = max
	return d
}

func (d *DragFloat64Slice) SetFormat(format string) *DragFloat64Slice {

	d.format.SetString(format)
	return d
}

func (d *DragFloat64Slice) Flags() imgui.SliderFlags {

	return d.flags
}

func (d *DragFloat64Slice) SetFlags(flags imgui.SliderFlags) *DragFloat64Slice {

	d.flags = flags
	return d
}

func (d *DragFloat64Slice) SetOnChange(f func(*DragFloat64Slice)) *DragFloat64Slice {

	d.onchange = f
	return d
}
