package gig

import (
	"bytes"
	"runtime"
	"runtime/cgo"

	"github.com/leonsal/gig/imgui"
)

type InputText struct {
	Widget
	label    imgui.CString
	hint     imgui.CString
	flags    imgui.InputTextFlags
	onchange func(*InputText)
	handle   cgo.Handle
	buf      []byte
}

func NewInputText(label string) *InputText {

	it := new(InputText)
	it.Init(it)
	it.label.SetString(label)
	it.SetMaxLen(80)
	runtime.SetFinalizer(it, func(id *InputText) {
		it.handle.Delete()
	})
	it.SetRender(func() {
		if it.hint.Len() == 0 {
			if imgui.InputTextCS(it.label, it.buf, it.flags, it.handle) {
				if it.onchange != nil {
					it.onchange(it)
				}
			}
		} else {
			if imgui.InputTextWithHintCS(it.label, it.hint, it.buf, it.flags, it.handle) {
				if it.onchange != nil {
					it.onchange(it)
				}
			}
		}
	})
	return it
}

func (it *InputText) Save(dst **InputText) *InputText {

	*dst = it
	return it
}

func (it *InputText) SetLabel(label string) *InputText {

	it.label.SetString(label)
	return it
}

func (it *InputText) SetHint(hint string) *InputText {

	it.hint.SetString(hint)
	return it
}

func (it *InputText) Flags() imgui.InputTextFlags {

	return it.flags
}

func (it *InputText) SetFlags(flags imgui.InputTextFlags) *InputText {

	it.flags = flags
	return it
}

func (it *InputText) Text() string {

	idx := bytes.IndexByte(it.buf, 0)
	if idx < 0 {
		return ""
	}
	return string(it.buf[:idx])
}

func (it *InputText) SetText(text string) *InputText {

	copy(it.buf[0:len(it.buf)-1], text)
	pos := len(text)
	if pos >= len(it.buf) {
		pos = len(it.buf) - 1
	}
	it.buf[pos] = 0
	return it
}

func (it *InputText) MaxLen() int {

	return len(it.buf)
}

func (it *InputText) SetMaxLen(size int) *InputText {

	if size < 0 {
		panic("InputText.SetBufSize(): invalid buf size")
	}
	it.buf = make([]byte, size+1)
	it.buf[0] = 0
	return it
}

func (it *InputText) SetOnChange(f func(*InputText)) *InputText {

	it.onchange = f
	return it
}

func (it *InputText) SetCallback(cb imgui.InputTextCallback) *InputText {

	if it.handle != 0 {
		it.handle.Delete()
	}
	it.handle = cgo.NewHandle(cb)
	return it
}

type InputTextMultiline struct {
	Widget
	label    imgui.CString
	flags    imgui.InputTextFlags
	size     imgui.Vec2
	onchange func(*InputTextMultiline)
	handle   cgo.Handle
	buf      []byte
}

func NewInputTextMultiline(label string) *InputTextMultiline {

	it := new(InputTextMultiline)
	it.Init(it)
	it.label.SetString(label)
	it.SetMaxLen(200)
	runtime.SetFinalizer(it, func(id *InputTextMultiline) {
		it.handle.Delete()
	})
	it.SetRender(func() {
		if imgui.InputTextMultilineCS(it.label, it.buf, it.size, it.flags, it.handle) {
			if it.onchange != nil {
				it.onchange(it)
			}
		}
	})
	return it
}

func (it *InputTextMultiline) SetLabel(label string) *InputTextMultiline {

	it.label.SetString(label)
	return it
}

func (it *InputTextMultiline) Size() imgui.Vec2 {

	return it.size
}

func (it *InputTextMultiline) SetSize(size imgui.Vec2) *InputTextMultiline {

	it.size = size
	return it
}

func (it *InputTextMultiline) Flags() imgui.InputTextFlags {

	return it.flags
}

func (it *InputTextMultiline) SetFlags(flags imgui.InputTextFlags) *InputTextMultiline {

	it.flags = flags
	return it
}

func (it *InputTextMultiline) Text() string {

	idx := bytes.IndexByte(it.buf, 0)
	if idx < 0 {
		return ""
	}
	return string(it.buf[:idx])
}

func (it *InputTextMultiline) SetText(text string) *InputTextMultiline {

	copy(it.buf[0:len(it.buf)-1], text)
	return it
}

func (it *InputTextMultiline) MaxLen() int {

	return len(it.buf)
}

func (it *InputTextMultiline) SetMaxLen(size int) *InputTextMultiline {

	if size < 0 {
		size = 0
	}
	it.buf = make([]byte, size+1)
	it.buf[0] = 0
	return it
}

func (it *InputTextMultiline) SetOnChange(f func(*InputTextMultiline)) *InputTextMultiline {

	it.onchange = f
	return it
}

func (it *InputTextMultiline) SetCallback(cb imgui.InputTextCallback) *InputTextMultiline {

	if it.handle != 0 {
		it.handle.Delete()
	}
	it.handle = cgo.NewHandle(cb)
	return it
}

type InputInt32 struct {
	Widget
	label     imgui.CString
	value     int32
	step      int
	step_fast int
	flags     imgui.InputTextFlags
	onchange  func(*InputInt32)
}

func NewInputInt32(label string, v int32) *InputInt32 {

	in := new(InputInt32)
	in.Init(in)
	in.label.SetString(label)
	in.value = v
	in.step = 0
	in.step_fast = 0
	in.SetRender(func() {
		if imgui.InputIntCS(in.label, &in.value, in.step, in.step_fast, in.flags) {
			if in.onchange != nil {
				in.onchange(in)
			}
		}
	})
	return in
}

func (in *InputInt32) Value() int32 {

	return in.value
}

func (in *InputInt32) SetValue(v int32) *InputInt32 {

	in.value = v
	return in
}

func (in *InputInt32) Step() int {

	return in.step
}

func (in *InputInt32) SetStep(s int) *InputInt32 {

	in.step = s
	return in
}

func (in *InputInt32) StepFast() int {

	return in.step_fast
}

func (in *InputInt32) SetStepFast(s int) *InputInt32 {

	in.step_fast = s
	return in
}

func (in *InputInt32) Flags() imgui.InputTextFlags {

	return in.flags
}

func (in *InputInt32) SetFlags(flags imgui.InputTextFlags) *InputInt32 {

	in.flags = flags
	return in
}

func (in *InputInt32) SetOnChange(f func(*InputInt32)) *InputInt32 {

	in.onchange = f
	return in
}

type InputInt32Slice struct {
	Widget
	label    imgui.CString
	values   []int32
	step     int32
	stepFast int32
	format   imgui.CString
	flags    imgui.InputTextFlags
	onchange func(*InputInt32Slice)
}

func NewInputInt32Slice(label string, v []int32) *InputInt32Slice {

	if len(v) < 1 {
		panic("NewInputInt32Slice: length of slice must be >= 1")
	}
	in := new(InputInt32Slice)
	in.Init(in)
	in.values = make([]int32, len(v))
	copy(in.values, v)
	in.label.SetString(label)
	ptr := func(v *int32) *int32 {
		if *v == 0 {
			return nil
		}
		return v
	}
	in.SetRender(func() {
		if imgui.InputInt32N(in.label, &in.values[0], len(in.values), ptr(&in.step), ptr(&in.stepFast), in.format, in.flags) {
			if in.onchange != nil {
				in.onchange(in)
			}
		}
	})
	return in
}

func (in *InputInt32Slice) Values() []int32 {

	return in.values
}

func (in *InputInt32Slice) SetValues(values []int32) *InputInt32Slice {

	copy(in.values, values)
	return in
}

func (in *InputInt32Slice) Step(step int32) int32 {

	return in.step
}

func (in *InputInt32Slice) SetStep(step int32) *InputInt32Slice {

	in.step = step
	return in
}

func (in *InputInt32Slice) StepFast(stepFast int32) int32 {

	return in.stepFast
}

func (in *InputInt32Slice) SetStepFast(stepFast int32) *InputInt32Slice {

	in.stepFast = stepFast
	return in
}

func (in *InputInt32Slice) Format() string {

	return in.format.String()
}

func (in *InputInt32Slice) SetFormat(f string) *InputInt32Slice {

	in.format.SetString(f)
	return in
}

func (in *InputInt32Slice) Flags() imgui.InputTextFlags {

	return in.flags
}

func (in *InputInt32Slice) SetFlags(flags imgui.InputTextFlags) *InputInt32Slice {

	in.flags = flags
	return in
}

func (in *InputInt32Slice) SetOnChange(f func(*InputInt32Slice)) *InputInt32Slice {

	in.onchange = f
	return in
}

// InputInt64 is a widget for input of an int64 value
type InputInt64 struct {
	Widget
	label    imgui.CString
	value    int64
	step     int64
	stepFast int64
	format   imgui.CString
	flags    imgui.InputTextFlags
	onchange func(*InputInt64)
}

func NewInputInt64(label string, v int64) *InputInt64 {

	in := new(InputInt64)
	in.Init(in)
	in.label.SetString(label)
	in.value = v
	ptr := func(v *int64) *int64 {
		if *v == 0 {
			return nil
		}
		return v
	}
	in.SetRender(func() {
		if imgui.InputInt64N(in.label, &in.value, 1, ptr(&in.step), ptr(&in.stepFast), in.format, in.flags) {
			if in.onchange != nil {
				in.onchange(in)
			}
		}
	})
	return in
}

func (in *InputInt64) Value() int64 {

	return in.value
}

func (in *InputInt64) SetValue(v int64) *InputInt64 {

	in.value = v
	return in
}

func (in *InputInt64) Step() int64 {

	return in.step
}

func (in *InputInt64) SetStep(s int64) *InputInt64 {

	in.step = s
	return in
}

func (in *InputInt64) StepFast() int64 {

	return in.stepFast
}

func (in *InputInt64) SetStepFast(s int64) *InputInt64 {

	in.stepFast = s
	return in
}

func (in *InputInt64Slice) Format() string {

	return in.format.String()
}

func (in *InputInt64Slice) SetFormat(f string) *InputInt64Slice {

	in.format.SetString(f)
	return in
}

func (in *InputInt64) Flags() imgui.InputTextFlags {

	return in.flags
}

func (in *InputInt64) SetFlags(flags imgui.InputTextFlags) *InputInt64 {

	in.flags = flags
	return in
}

func (in *InputInt64) SetOnChange(f func(*InputInt64)) *InputInt64 {

	in.onchange = f
	return in
}

type InputInt64Slice struct {
	Widget
	label    imgui.CString
	values   []int64
	step     int64
	stepFast int64
	format   imgui.CString
	flags    imgui.InputTextFlags
	onchange func(*InputInt64Slice)
}

func NewInputInt64Slice(label string, v []int64) *InputInt64Slice {

	if len(v) < 1 {
		panic("NewInputInt64Slice: length of slice must be >= 1")
	}
	in := new(InputInt64Slice)
	in.Init(in)
	in.values = make([]int64, len(v))
	copy(in.values, v)
	in.label.SetString(label)
	ptr := func(v *int64) *int64 {
		if *v == 0 {
			return nil
		}
		return v
	}
	in.SetRender(func() {
		if imgui.InputInt64N(in.label, &in.values[0], len(in.values), ptr(&in.step), ptr(&in.stepFast), in.format, in.flags) {
			if in.onchange != nil {
				in.onchange(in)
			}
		}
	})
	return in
}

func (in *InputInt64Slice) Values() []int64 {

	return in.values
}

func (in *InputInt64Slice) SetValues(values []int64) *InputInt64Slice {

	copy(in.values, values)
	return in
}

func (in *InputInt64Slice) Step(step int64) int64 {

	return in.step
}

func (in *InputInt64Slice) SetStep(step int64) *InputInt64Slice {

	in.step = step
	return in
}

func (in *InputInt64Slice) StepFast(stepFast int64) int64 {

	return in.stepFast
}

func (in *InputInt64Slice) SetStepFast(stepFast int64) *InputInt64Slice {

	in.stepFast = stepFast
	return in
}

func (in *InputInt64Slice) Flags() imgui.InputTextFlags {

	return in.flags
}

func (in *InputInt64Slice) SetFlags(flags imgui.InputTextFlags) *InputInt64Slice {

	in.flags = flags
	return in
}

func (in *InputInt64Slice) SetOnChange(f func(*InputInt64Slice)) *InputInt64Slice {

	in.onchange = f
	return in
}

type InputFloat32 struct {
	Widget
	label     imgui.CString
	value     float32
	step      float32
	step_fast float32
	format    imgui.CString
	flags     imgui.InputTextFlags
	onchange  func(*InputFloat32)
}

func NewInputFloat32(label string, v float32) *InputFloat32 {

	in := new(InputFloat32)
	in.Init(in)
	in.label.SetString(label)
	in.value = v
	in.format.SetString("%.3f")
	in.SetRender(func() {
		if imgui.InputFloatCS(in.label, &in.value, in.step, in.step_fast, in.format, in.flags) {
			if in.onchange != nil {
				in.onchange(in)
			}
		}
	})
	return in
}

func (in *InputFloat32) Value() float32 {

	return float32(in.value)
}

func (in *InputFloat32) SetValue(v float32) *InputFloat32 {

	in.value = v
	return in
}

func (in *InputFloat32) Step() float32 {

	return in.step
}

func (in *InputFloat32) SetStep(s float32) *InputFloat32 {

	in.step = s
	return in
}

func (in *InputFloat32) StepFast() float32 {

	return in.step_fast
}

func (in *InputFloat32) SetStepFast(s float32) *InputFloat32 {

	in.step_fast = s
	return in
}

func (in *InputFloat32) Format() string {

	return in.format.String()
}

func (in *InputFloat32) SetFormat(f string) *InputFloat32 {

	in.format.SetString(f)
	return in
}

func (in *InputFloat32) Flags() imgui.InputTextFlags {

	return in.flags
}

func (in *InputFloat32) SetFlags(flags imgui.InputTextFlags) *InputFloat32 {

	in.flags = flags
	return in
}

func (in *InputFloat32) SetOnChange(f func(*InputFloat32)) *InputFloat32 {

	in.onchange = f
	return in
}

type InputFloat32Slice struct {
	Widget
	label    imgui.CString
	values   []float32
	step     float32
	stepFast float32
	format   imgui.CString
	flags    imgui.InputTextFlags
	onchange func(*InputFloat32Slice)
}

func NewInputFloat32Slice(label string, v []float32) *InputFloat32Slice {

	if len(v) < 1 {
		panic("NewInputFloat32Slice: length of slice must be >= 1")
	}
	in := new(InputFloat32Slice)
	in.Init(in)
	in.values = make([]float32, len(v))
	copy(in.values, v)
	in.label.SetString(label)
	in.format.SetString("%.3f")
	ptr := func(v *float32) *float32 {
		if *v == 0 {
			return nil
		}
		return v
	}
	in.SetRender(func() {
		if imgui.InputFloat32N(in.label, &in.values[0], len(in.values), ptr(&in.step), ptr(&in.stepFast), in.format, in.flags) {
			if in.onchange != nil {
				in.onchange(in)
			}
		}
	})
	return in
}

func (in *InputFloat32Slice) Values() []float32 {

	return in.values
}

func (in *InputFloat32Slice) SetValues(values []float32) *InputFloat32Slice {

	copy(in.values, values)
	return in
}

func (in *InputFloat32Slice) Step() float32 {

	return in.step
}

func (in *InputFloat32Slice) SetStep(s float32) *InputFloat32Slice {

	in.step = s
	return in
}

func (in *InputFloat32Slice) StepFast() float32 {

	return in.stepFast
}

func (in *InputFloat32Slice) SetStepFast(s float32) *InputFloat32Slice {

	in.stepFast = s
	return in
}

func (in *InputFloat32Slice) Format() string {

	return in.format.String()
}

func (in *InputFloat32Slice) SetFormat(f string) *InputFloat32Slice {

	in.format.SetString(f)
	return in
}

func (in *InputFloat32Slice) Flags() imgui.InputTextFlags {

	return in.flags
}

func (in *InputFloat32Slice) SetFlags(flags imgui.InputTextFlags) *InputFloat32Slice {

	in.flags = flags
	return in
}

func (in *InputFloat32Slice) SetOnChange(f func(*InputFloat32Slice)) *InputFloat32Slice {

	in.onchange = f
	return in
}

type InputFloat64 struct {
	Widget
	label    imgui.CString
	value    float64
	step     float64
	stepFast float64
	format   imgui.CString
	flags    imgui.InputTextFlags
	onchange func(*InputFloat64)
}

func NewInputFloat64(label string, v float64) *InputFloat64 {

	in := new(InputFloat64)
	in.Init(in)
	in.label.SetString(label)
	in.value = v
	in.format.SetString("%.6f")
	in.SetRender(func() {
		if imgui.InputDoubleCS(in.label, &in.value, in.step, in.stepFast, in.format, in.flags) {
			if in.onchange != nil {
				in.onchange(in)
			}
		}
	})
	return in
}

func (in *InputFloat64) Value() float64 {

	return in.value
}

func (in *InputFloat64) SetValue(v float64) *InputFloat64 {

	in.value = v
	return in
}

func (in *InputFloat64) Step() float64 {

	return in.step
}

func (in *InputFloat64) SetStep(s float64) *InputFloat64 {

	in.step = s
	return in
}

func (in *InputFloat64) StepFast() float64 {

	return in.stepFast
}

func (in *InputFloat64) SetStepFast(s float64) *InputFloat64 {

	in.stepFast = s
	return in
}

func (in *InputFloat64) Format() string {

	return in.format.String()
}

func (in *InputFloat64) SetFormat(f string) *InputFloat64 {

	in.format.SetString(f)
	return in
}

func (in *InputFloat64) Flags() imgui.InputTextFlags {

	return in.flags
}

func (in *InputFloat64) SetFlags(flags imgui.InputTextFlags) *InputFloat64 {

	in.flags = flags
	return in
}

func (in *InputFloat64) SetOnChange(f func(*InputFloat64)) *InputFloat64 {

	in.onchange = f
	return in
}

type InputFloat64Slice struct {
	Widget
	label    imgui.CString
	values   []float64
	step     float64
	stepFast float64
	format   imgui.CString
	flags    imgui.InputTextFlags
	onchange func(*InputFloat64Slice)
}

func NewInputFloat64Slice(label string, v []float64) *InputFloat64Slice {

	if len(v) < 1 {
		panic("NewInputFloat64Slice: length of slice must be >= 1")
	}
	in := new(InputFloat64Slice)
	in.Init(in)
	in.values = make([]float64, len(v))
	copy(in.values, v)
	in.label.SetString(label)
	in.format.SetString("%.3f")
	ptr := func(v *float64) *float64 {
		if *v == 0 {
			return nil
		}
		return v
	}
	in.SetRender(func() {
		if imgui.InputFloat64N(in.label, &in.values[0], len(in.values), ptr(&in.step), ptr(&in.stepFast), in.format, in.flags) {
			if in.onchange != nil {
				in.onchange(in)
			}
		}
	})
	return in
}

func (in *InputFloat64Slice) Values() []float64 {

	return in.values
}

func (in *InputFloat64Slice) SetValues(values []float64) *InputFloat64Slice {

	copy(in.values, values)
	return in
}

func (in *InputFloat64Slice) Step() float64 {

	return in.step
}

func (in *InputFloat64Slice) SetStep(s float64) *InputFloat64Slice {

	in.step = s
	return in
}

func (in *InputFloat64Slice) StepFast() float64 {

	return in.stepFast
}

func (in *InputFloat64Slice) SetStepFast(s float64) *InputFloat64Slice {

	in.stepFast = s
	return in
}

func (in *InputFloat64Slice) Format() string {

	return in.format.String()
}

func (in *InputFloat64Slice) SetFormat(f string) *InputFloat64Slice {

	in.format.SetString(f)
	return in
}

func (in *InputFloat64Slice) Flags() imgui.InputTextFlags {

	return in.flags
}

func (in *InputFloat64Slice) SetFlags(flags imgui.InputTextFlags) *InputFloat64Slice {

	in.flags = flags
	return in
}

func (in *InputFloat64Slice) SetOnChange(f func(*InputFloat64Slice)) *InputFloat64Slice {

	in.onchange = f
	return in
}
