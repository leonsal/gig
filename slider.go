package gig

import (
	"github.com/leonsal/gig/imgui"
)

type SliderInt32 struct {
	Widget
	label    imgui.CString
	value    []int32
	min      int32
	max      int32
	format   imgui.CString
	flags    imgui.SliderFlags
	onchange func(*SliderInt32)
}

func NewSliderInt32(label string, value int32, min int32, max int32) *SliderInt32 {

	s := new(SliderInt32)
	s.Init(s)
	s.label.SetString(label)
	s.value = []int32{value}
	s.min = min
	s.max = max
	s.SetRender(func() {
		if imgui.SliderInt32NCS(s.label, s.value, &s.min, &s.max, s.format, s.flags) {
			if s.onchange != nil {
				s.onchange(s)
			}
		}
	})
	return s
}

func (s *SliderInt32) SetLabel(label string) *SliderInt32 {

	s.label.SetString(label)
	return s
}

func (s *SliderInt32) Value() int32 {

	return s.value[0]
}

func (s *SliderInt32) SetValue(v int32) *SliderInt32 {

	s.value[0] = v
	return s
}

func (s *SliderInt32) Min() int32 {

	return s.min
}

func (s *SliderInt32) SetMin(min int32) *SliderInt32 {

	s.min = min
	return s
}

func (s *SliderInt32) Max() int32 {

	return s.max
}

func (s *SliderInt32) SetMax(max int32) *SliderInt32 {

	s.max = max
	return s
}

func (s *SliderInt32) SetFormat(format string) *SliderInt32 {

	s.format.SetString(format)
	return s
}

func (s *SliderInt32) Flags() imgui.SliderFlags {

	return s.flags
}

func (s *SliderInt32) SetFlags(flags imgui.SliderFlags) *SliderInt32 {

	s.flags = flags
	return s
}

func (s *SliderInt32) SetOnChange(f func(*SliderInt32)) *SliderInt32 {

	s.onchange = f
	return s
}

type SliderInt32Slice struct {
	Widget
	label    imgui.CString
	values   []int32
	min      int32
	max      int32
	format   imgui.CString
	flags    imgui.SliderFlags
	onchange func(*SliderInt32Slice)
}

func NewSliderInt32Slice(label string, values []int32, min int32, max int32) *SliderInt32Slice {

	s := new(SliderInt32Slice)
	s.Init(s)
	s.label.SetString(label)
	s.values = make([]int32, len(values))
	copy(s.values, values)
	s.min = min
	s.max = max
	s.SetRender(func() {
		if imgui.SliderInt32NCS(s.label, s.values, &s.min, &s.max, s.format, s.flags) {
			if s.onchange != nil {
				s.onchange(s)
			}
		}
	})
	return s
}

func (s *SliderInt32Slice) SetLabel(label string) *SliderInt32Slice {

	s.label.SetString(label)
	return s
}

// Values returns the internal values slice
func (s *SliderInt32Slice) Values() []int32 {

	return s.values
}

// SetValues copy the specified slice to the internal values slice.
func (s *SliderInt32Slice) SetValues(src []int32) *SliderInt32Slice {

	copy(s.values, src)
	return s
}

func (s *SliderInt32Slice) Min() int32 {

	return s.min
}

func (s *SliderInt32Slice) SetMin(min int32) *SliderInt32Slice {

	s.min = min
	return s
}

func (s *SliderInt32Slice) Max() int32 {

	return s.max
}

func (s *SliderInt32Slice) SetMax(max int32) *SliderInt32Slice {

	s.max = max
	return s
}

func (s *SliderInt32Slice) SetFormat(format string) *SliderInt32Slice {

	s.format.SetString(format)
	return s
}

func (s *SliderInt32Slice) Flags() imgui.SliderFlags {

	return s.flags
}

func (s *SliderInt32Slice) SetFlags(flags imgui.SliderFlags) *SliderInt32Slice {

	s.flags = flags
	return s
}

func (s *SliderInt32Slice) SetOnChange(f func(*SliderInt32Slice)) *SliderInt32Slice {

	s.onchange = f
	return s
}

type SliderInt64 struct {
	Widget
	label    imgui.CString
	value    []int64
	min      int64
	max      int64
	format   imgui.CString
	flags    imgui.SliderFlags
	onchange func(*SliderInt64)
}

func NewSliderInt64(label string, value int64, min int64, max int64) *SliderInt64 {

	s := new(SliderInt64)
	s.Init(s)
	s.label.SetString(label)
	s.value = []int64{value}
	s.min = min
	s.max = max
	s.SetRender(func() {
		if imgui.SliderInt64NCS(s.label, s.value, &s.min, &s.max, s.format, s.flags) {
			if s.onchange != nil {
				s.onchange(s)
			}
		}
	})
	return s
}

func (s *SliderInt64) SetLabel(label string) *SliderInt64 {

	s.label.SetString(label)
	return s
}

func (s *SliderInt64) Value() int64 {

	return s.value[0]
}

func (s *SliderInt64) SetValue(v int64) *SliderInt64 {

	s.value[0] = v
	return s
}

func (s *SliderInt64) Min() int64 {

	return s.min
}

func (s *SliderInt64) SetMin(min int64) *SliderInt64 {

	s.min = min
	return s
}

func (s *SliderInt64) Max() int64 {

	return s.max
}

func (s *SliderInt64) SetMax(max int64) *SliderInt64 {

	s.max = max
	return s
}

func (s *SliderInt64) SetFormat(format string) *SliderInt64 {

	s.format.SetString(format)
	return s
}

func (s *SliderInt64) Flags() imgui.SliderFlags {

	return s.flags
}

func (s *SliderInt64) SetFlags(flags imgui.SliderFlags) *SliderInt64 {

	s.flags = flags
	return s
}

func (s *SliderInt64) SetOnChange(f func(*SliderInt64)) *SliderInt64 {

	s.onchange = f
	return s
}

type SliderInt64Slice struct {
	Widget
	label    imgui.CString
	values   []int64
	min      int64
	max      int64
	format   imgui.CString
	flags    imgui.SliderFlags
	onchange func(*SliderInt64Slice)
}

func NewSliderInt64Slice(label string, values []int64, min int64, max int64) *SliderInt64Slice {

	s := new(SliderInt64Slice)
	s.Init(s)
	s.label.SetString(label)
	s.values = make([]int64, len(values))
	copy(s.values, values)
	s.min = min
	s.max = max
	s.SetRender(func() {
		if imgui.SliderInt64NCS(s.label, s.values, &s.min, &s.max, s.format, s.flags) {
			if s.onchange != nil {
				s.onchange(s)
			}
		}
	})
	return s
}

func (s *SliderInt64Slice) SetLabel(label string) *SliderInt64Slice {

	s.label.SetString(label)
	return s
}

// Values returns the internal values slice
func (s *SliderInt64Slice) Values() []int64 {

	return s.values
}

// SetValues copy the specified slice to the internal values slice.
func (s *SliderInt64Slice) SetValues(src []int64) *SliderInt64Slice {

	copy(s.values, src)
	return s
}

func (s *SliderInt64Slice) Min() int64 {

	return s.min
}

func (s *SliderInt64Slice) SetMin(min int64) *SliderInt64Slice {

	s.min = min
	return s
}

func (s *SliderInt64Slice) Max() int64 {

	return s.max
}

func (s *SliderInt64Slice) SetMax(max int64) *SliderInt64Slice {

	s.max = max
	return s
}

func (s *SliderInt64Slice) SetFormat(format string) *SliderInt64Slice {

	s.format.SetString(format)
	return s
}

func (s *SliderInt64Slice) Flags() imgui.SliderFlags {

	return s.flags
}

func (s *SliderInt64Slice) SetFlags(flags imgui.SliderFlags) *SliderInt64Slice {

	s.flags = flags
	return s
}

func (s *SliderInt64Slice) SetOnChange(f func(*SliderInt64Slice)) *SliderInt64Slice {

	s.onchange = f
	return s
}

type SliderFloat32 struct {
	Widget
	label    imgui.CString
	value    []float32
	min      float32
	max      float32
	format   imgui.CString
	flags    imgui.SliderFlags
	onchange func(*SliderFloat32)
}

func NewSliderFloat32(label string, value float32, min float32, max float32) *SliderFloat32 {

	s := new(SliderFloat32)
	s.Init(s)
	s.label.SetString(label)
	s.value = []float32{value}
	s.min = min
	s.max = max
	s.SetRender(func() {
		if imgui.SliderFloat32NCS(s.label, s.value, &s.min, &s.max, s.format, s.flags) {
			if s.onchange != nil {
				s.onchange(s)
			}
		}
	})
	return s
}

func (s *SliderFloat32) SetLabel(label string) *SliderFloat32 {

	s.label.SetString(label)
	return s
}

func (s *SliderFloat32) Value() float32 {

	return s.value[0]
}

func (s *SliderFloat32) SetValue(v float32) *SliderFloat32 {

	s.value[0] = v
	return s
}

func (s *SliderFloat32) Min() float32 {

	return s.min
}

func (s *SliderFloat32) SetMin(min float32) *SliderFloat32 {

	s.min = min
	return s
}

func (s *SliderFloat32) Max() float32 {

	return s.max
}

func (s *SliderFloat32) SetMax(max float32) *SliderFloat32 {

	s.max = max
	return s
}

func (s *SliderFloat32) SetFormat(format string) *SliderFloat32 {

	s.format.SetString(format)
	return s
}

func (s *SliderFloat32) Flags() imgui.SliderFlags {

	return s.flags
}

func (s *SliderFloat32) SetFlags(flags imgui.SliderFlags) *SliderFloat32 {

	s.flags = flags
	return s
}

func (s *SliderFloat32) SetOnChange(f func(*SliderFloat32)) *SliderFloat32 {

	s.onchange = f
	return s
}

type SliderFloat32Slice struct {
	Widget
	label    imgui.CString
	values   []float32
	min      float32
	max      float32
	format   imgui.CString
	flags    imgui.SliderFlags
	onchange func(*SliderFloat32Slice)
}

func NewSliderFloat32Slice(label string, values []float32, min float32, max float32) *SliderFloat32Slice {

	s := new(SliderFloat32Slice)
	s.Init(s)
	s.label.SetString(label)
	s.values = make([]float32, len(values))
	copy(s.values, values)
	s.min = min
	s.max = max
	s.SetRender(func() {
		if imgui.SliderFloat32NCS(s.label, s.values, &s.min, &s.max, s.format, s.flags) {
			if s.onchange != nil {
				s.onchange(s)
			}
		}
	})
	return s
}

func (s *SliderFloat32Slice) SetLabel(label string) *SliderFloat32Slice {

	s.label.SetString(label)
	return s
}

// Values returns the internal values slice
func (s *SliderFloat32Slice) Values() []float32 {

	return s.values
}

// SetValues copy the specified slice to the internal values slice.
func (s *SliderFloat32Slice) SetValues(src []float32) *SliderFloat32Slice {

	copy(s.values, src)
	return s
}

func (s *SliderFloat32Slice) Min() float32 {

	return s.min
}

func (s *SliderFloat32Slice) SetMin(min float32) *SliderFloat32Slice {

	s.min = min
	return s
}

func (s *SliderFloat32Slice) Max() float32 {

	return s.max
}

func (s *SliderFloat32Slice) SetMax(max float32) *SliderFloat32Slice {

	s.max = max
	return s
}

func (s *SliderFloat32Slice) SetFormat(format string) *SliderFloat32Slice {

	s.format.SetString(format)
	return s
}

func (s *SliderFloat32Slice) Flags() imgui.SliderFlags {

	return s.flags
}

func (s *SliderFloat32Slice) SetFlags(flags imgui.SliderFlags) *SliderFloat32Slice {

	s.flags = flags
	return s
}

func (s *SliderFloat32Slice) SetOnChange(f func(*SliderFloat32Slice)) *SliderFloat32Slice {

	s.onchange = f
	return s
}

type SliderFloat64 struct {
	Widget
	label    imgui.CString
	value    []float64
	min      float64
	max      float64
	format   imgui.CString
	flags    imgui.SliderFlags
	onchange func(*SliderFloat64)
}

func NewSliderFloat64(label string, value float64, min float64, max float64) *SliderFloat64 {

	s := new(SliderFloat64)
	s.Init(s)
	s.label.SetString(label)
	s.value = []float64{value}
	s.min = min
	s.max = max
	s.SetRender(func() {
		if imgui.SliderFloat64NCS(s.label, s.value, &s.min, &s.max, s.format, s.flags) {
			if s.onchange != nil {
				s.onchange(s)
			}
		}
	})
	return s
}

func (s *SliderFloat64) SetLabel(label string) *SliderFloat64 {

	s.label.SetString(label)
	return s
}

func (s *SliderFloat64) Value() float64 {

	return s.value[0]
}

func (s *SliderFloat64) SetValue(v float64) *SliderFloat64 {

	s.value[0] = v
	return s
}

func (s *SliderFloat64) Min() float64 {

	return s.min
}

func (s *SliderFloat64) SetMin(min float64) *SliderFloat64 {

	s.min = min
	return s
}

func (s *SliderFloat64) Max() float64 {

	return s.max
}

func (s *SliderFloat64) SetMax(max float64) *SliderFloat64 {

	s.max = max
	return s
}

func (s *SliderFloat64) SetFormat(format string) *SliderFloat64 {

	s.format.SetString(format)
	return s
}

func (s *SliderFloat64) Flags() imgui.SliderFlags {

	return s.flags
}

func (s *SliderFloat64) SetFlags(flags imgui.SliderFlags) *SliderFloat64 {

	s.flags = flags
	return s
}

func (s *SliderFloat64) SetOnChange(f func(*SliderFloat64)) *SliderFloat64 {

	s.onchange = f
	return s
}

type SliderFloat64Slice struct {
	Widget
	label    imgui.CString
	values   []float64
	min      float64
	max      float64
	format   imgui.CString
	flags    imgui.SliderFlags
	onchange func(*SliderFloat64Slice)
}

func NewSliderFloat64Slice(label string, values []float64, min float64, max float64) *SliderFloat64Slice {

	s := new(SliderFloat64Slice)
	s.Init(s)
	s.label.SetString(label)
	s.values = make([]float64, len(values))
	copy(s.values, values)
	s.min = min
	s.max = max
	s.SetRender(func() {
		if imgui.SliderFloat64NCS(s.label, s.values, &s.min, &s.max, s.format, s.flags) {
			if s.onchange != nil {
				s.onchange(s)
			}
		}
	})
	return s
}

func (s *SliderFloat64Slice) SetLabel(label string) *SliderFloat64Slice {

	s.label.SetString(label)
	return s
}

// Values returns the internal values slice
func (s *SliderFloat64Slice) Values() []float64 {

	return s.values
}

// SetValues copy the specified slice to the internal values slice.
func (s *SliderFloat64Slice) SetValues(src []float64) *SliderFloat64Slice {

	copy(s.values, src)
	return s
}

func (s *SliderFloat64Slice) Min() float64 {

	return s.min
}

func (s *SliderFloat64Slice) SetMin(min float64) *SliderFloat64Slice {

	s.min = min
	return s
}

func (s *SliderFloat64Slice) Max() float64 {

	return s.max
}

func (s *SliderFloat64Slice) SetMax(max float64) *SliderFloat64Slice {

	s.max = max
	return s
}

func (s *SliderFloat64Slice) SetFormat(format string) *SliderFloat64Slice {

	s.format.SetString(format)
	return s
}

func (s *SliderFloat64Slice) Flags() imgui.SliderFlags {

	return s.flags
}

func (s *SliderFloat64Slice) SetFlags(flags imgui.SliderFlags) *SliderFloat64Slice {

	s.flags = flags
	return s
}

func (s *SliderFloat64Slice) SetOnChange(f func(*SliderFloat64Slice)) *SliderFloat64Slice {

	s.onchange = f
	return s
}

//func SliderAngleCS(label CString, vRad *float32, vDegreesMin, vDegreesMax float32, format CString, flags SliderFlags) bool {

type SliderAngle struct {
	Widget
	label    imgui.CString
	vrad     float32
	min      float32
	max      float32
	format   imgui.CString
	flags    imgui.SliderFlags
	onchange func(*SliderAngle)
}

func NewSliderAngle(label string, vrad float32) *SliderAngle {

	s := new(SliderAngle)
	s.Init(s)
	s.label.SetString(label)
	s.vrad = vrad
	s.min = -360.0
	s.max = 360.0
	s.format.SetString("%.0f deg")
	s.SetRender(func() {
		if imgui.SliderAngleCS(s.label, &vrad, s.min, s.max, s.format, s.flags) {
			if s.onchange != nil {
				s.onchange(s)
			}
		}
	})
	return s
}

func (s *SliderAngle) SetLabel(label string) *SliderAngle {

	s.label.SetString(label)
	return s
}

func (s *SliderAngle) Value() float32 {

	return s.vrad
}

func (s *SliderAngle) SetValue(vrad float32) *SliderAngle {

	s.vrad = vrad
	return s
}

func (s *SliderAngle) SetMin(min float32) *SliderAngle {

	s.min = min
	return s
}

func (s *SliderAngle) SetMax(max float32) *SliderAngle {

	s.max = max
	return s
}

func (s *SliderAngle) SetFormat(format string) *SliderAngle {

	s.format.SetString(format)
	return s
}

func (s *SliderAngle) Flags() imgui.SliderFlags {

	return s.flags
}

func (s *SliderAngle) SetFlags(flags imgui.SliderFlags) *SliderAngle {

	s.flags = flags
	return s
}

func (s *SliderAngle) SetOnChange(f func(*SliderAngle)) *SliderAngle {

	s.onchange = f
	return s
}

type VSliderInt32 struct {
	Widget
	label    imgui.CString
	size     imgui.Vec2
	value    int32
	min      int32
	max      int32
	format   imgui.CString
	flags    imgui.SliderFlags
	onchange func(*VSliderInt32)
}

func NewVSliderInt32(label string, size imgui.Vec2, value, min, max int32) *VSliderInt32 {

	s := new(VSliderInt32)
	s.Init(s)
	s.label.SetString(label)
	s.size = size
	s.value = value
	s.min = min
	s.max = max
	s.SetRender(func() {
		if imgui.VSliderInt32CS(s.label, s.size, &s.value, &s.min, &s.max, s.format, s.flags) {
			if s.onchange != nil {
				s.onchange(s)
			}
		}
	})
	return s
}

func (d *VSliderInt32) SetLabel(label string) *VSliderInt32 {

	d.label.SetString(label)
	return d
}

func (d *VSliderInt32) SetSize(size imgui.Vec2) *VSliderInt32 {

	d.size = size
	return d
}

func (d *VSliderInt32) Value() int32 {

	return d.value
}

func (d *VSliderInt32) SetValue(v int32) *VSliderInt32 {

	d.value = v
	return d
}

func (d *VSliderInt32) Min() int32 {

	return d.min
}

func (d *VSliderInt32) SetMin(min int32) *VSliderInt32 {

	d.min = min
	return d
}

func (d *VSliderInt32) Max() int32 {

	return d.max
}

func (d *VSliderInt32) SetMax(max int32) *VSliderInt32 {

	d.max = max
	return d
}

func (d *VSliderInt32) SetFormat(format string) *VSliderInt32 {

	d.format.SetString(format)
	return d
}

func (d *VSliderInt32) Flags() imgui.SliderFlags {

	return d.flags
}

func (d *VSliderInt32) SetFlags(flags imgui.SliderFlags) *VSliderInt32 {

	d.flags = flags
	return d
}

func (d *VSliderInt32) SetOnChange(f func(*VSliderInt32)) *VSliderInt32 {

	d.onchange = f
	return d
}

type VSliderInt64 struct {
	Widget
	label    imgui.CString
	size     imgui.Vec2
	value    int64
	min      int64
	max      int64
	format   imgui.CString
	flags    imgui.SliderFlags
	onchange func(*VSliderInt64)
}

func NewVSliderInt64(label string, size imgui.Vec2, value, min, max int64) *VSliderInt64 {

	s := new(VSliderInt64)
	s.Init(s)
	s.label.SetString(label)
	s.size = size
	s.value = value
	s.min = min
	s.max = max
	s.SetRender(func() {
		if imgui.VSliderInt64CS(s.label, s.size, &s.value, &s.min, &s.max, s.format, s.flags) {
			if s.onchange != nil {
				s.onchange(s)
			}
		}
	})
	return s
}

func (d *VSliderInt64) SetLabel(label string) *VSliderInt64 {

	d.label.SetString(label)
	return d
}

func (d *VSliderInt64) SetSize(size imgui.Vec2) *VSliderInt64 {

	d.size = size
	return d
}

func (d *VSliderInt64) Value() int64 {

	return d.value
}

func (d *VSliderInt64) SetValue(v int64) *VSliderInt64 {

	d.value = v
	return d
}

func (d *VSliderInt64) Min() int64 {

	return d.min
}

func (d *VSliderInt64) SetMin(min int64) *VSliderInt64 {

	d.min = min
	return d
}

func (d *VSliderInt64) Max() int64 {

	return d.max
}

func (d *VSliderInt64) SetMax(max int64) *VSliderInt64 {

	d.max = max
	return d
}

func (d *VSliderInt64) SetFormat(format string) *VSliderInt64 {

	d.format.SetString(format)
	return d
}

func (d *VSliderInt64) Flags() imgui.SliderFlags {

	return d.flags
}

func (d *VSliderInt64) SetFlags(flags imgui.SliderFlags) *VSliderInt64 {

	d.flags = flags
	return d
}

func (d *VSliderInt64) SetOnChange(f func(*VSliderInt64)) *VSliderInt64 {

	d.onchange = f
	return d
}

type VSliderFloat32 struct {
	Widget
	label    imgui.CString
	size     imgui.Vec2
	value    float32
	min      float32
	max      float32
	format   imgui.CString
	flags    imgui.SliderFlags
	onchange func(*VSliderFloat32)
}

func NewVSliderFloat32(label string, size imgui.Vec2, value, min, max float32) *VSliderFloat32 {

	s := new(VSliderFloat32)
	s.Init(s)
	s.label.SetString(label)
	s.size = size
	s.value = value
	s.min = min
	s.max = max
	s.SetRender(func() {
		if imgui.VSliderFloat32CS(s.label, s.size, &s.value, &s.min, &s.max, s.format, s.flags) {
			if s.onchange != nil {
				s.onchange(s)
			}
		}
	})
	return s
}

func (d *VSliderFloat32) SetLabel(label string) *VSliderFloat32 {

	d.label.SetString(label)
	return d
}

func (d *VSliderFloat32) SetSize(size imgui.Vec2) *VSliderFloat32 {

	d.size = size
	return d
}

func (d *VSliderFloat32) Value() float32 {

	return d.value
}

func (d *VSliderFloat32) SetValue(v float32) *VSliderFloat32 {

	d.value = v
	return d
}

func (d *VSliderFloat32) Min() float32 {

	return d.min
}

func (d *VSliderFloat32) SetMin(min float32) *VSliderFloat32 {

	d.min = min
	return d
}

func (d *VSliderFloat32) Max() float32 {

	return d.max
}

func (d *VSliderFloat32) SetMax(max float32) *VSliderFloat32 {

	d.max = max
	return d
}

func (d *VSliderFloat32) SetFormat(format string) *VSliderFloat32 {

	d.format.SetString(format)
	return d
}

func (d *VSliderFloat32) Flags() imgui.SliderFlags {

	return d.flags
}

func (d *VSliderFloat32) SetFlags(flags imgui.SliderFlags) *VSliderFloat32 {

	d.flags = flags
	return d
}

func (d *VSliderFloat32) SetOnChange(f func(*VSliderFloat32)) *VSliderFloat32 {

	d.onchange = f
	return d
}

type VSliderFloat64 struct {
	Widget
	label    imgui.CString
	size     imgui.Vec2
	value    float64
	min      float64
	max      float64
	format   imgui.CString
	flags    imgui.SliderFlags
	onchange func(*VSliderFloat64)
}

func NewVSliderFloat64(label string, size imgui.Vec2, value, min, max float64) *VSliderFloat64 {

	s := new(VSliderFloat64)
	s.Init(s)
	s.label.SetString(label)
	s.size = size
	s.value = value
	s.min = min
	s.max = max
	s.SetRender(func() {
		if imgui.VSliderFloat64CS(s.label, s.size, &s.value, &s.min, &s.max, s.format, s.flags) {
			if s.onchange != nil {
				s.onchange(s)
			}
		}
	})
	return s
}

func (d *VSliderFloat64) SetLabel(label string) *VSliderFloat64 {

	d.label.SetString(label)
	return d
}

func (d *VSliderFloat64) SetSize(size imgui.Vec2) *VSliderFloat64 {

	d.size = size
	return d
}

func (d *VSliderFloat64) Value() float64 {

	return d.value
}

func (d *VSliderFloat64) SetValue(v float64) *VSliderFloat64 {

	d.value = v
	return d
}

func (d *VSliderFloat64) Min() float64 {

	return d.min
}

func (d *VSliderFloat64) SetMin(min float64) *VSliderFloat64 {

	d.min = min
	return d
}

func (d *VSliderFloat64) Max() float64 {

	return d.max
}

func (d *VSliderFloat64) SetMax(max float64) *VSliderFloat64 {

	d.max = max
	return d
}

func (d *VSliderFloat64) SetFormat(format string) *VSliderFloat64 {

	d.format.SetString(format)
	return d
}

func (d *VSliderFloat64) Flags() imgui.SliderFlags {

	return d.flags
}

func (d *VSliderFloat64) SetFlags(flags imgui.SliderFlags) *VSliderFloat64 {

	d.flags = flags
	return d
}

func (d *VSliderFloat64) SetOnChange(f func(*VSliderFloat64)) *VSliderFloat64 {

	d.onchange = f
	return d
}
