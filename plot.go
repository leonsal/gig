package gig

import (
	"github.com/leonsal/gig/imgui"
)

type Plot struct {
	plotBase
	title   imgui.CString
	xlabel  imgui.CString
	ylabel  imgui.CString
	y2label imgui.CString
	y3label imgui.CString
	size    imgui.Vec2
	flags   imgui.PlotFlags
	xflags  imgui.PlotAxisFlags
	yflags  imgui.PlotAxisFlags
	y2flags imgui.PlotAxisFlags
	y3flags imgui.PlotAxisFlags
}

func NewPlot(title string) *Plot {

	p := new(Plot)
	p.Init(title)
	return p
}

func (p *Plot) Init(title string) {

	p.Widget.Init(p)
	p.title.SetString(title)
	p.size = imgui.Vec2{-1, 0}
	p.y2flags = imgui.PlotAxisFlags_NoGridLines
	p.y3flags = imgui.PlotAxisFlags_NoGridLines

	p.SetRender(func() {
		if imgui.BeginPlotCS(p.title, p.xlabel, p.ylabel, p.size, p.flags, p.xflags, p.yflags, p.y2flags, p.y3flags, p.y2label, p.y3label) {
			p.RenderChildren()
			imgui.EndPlot()
		}
	})
}

func (p *Plot) SetTitle(title string) *Plot {

	p.title.SetString(title)
	return p
}

func (p *Plot) Title() string {

	return p.title.String()
}

func (p *Plot) SetXLabel(label string) *Plot {

	p.xlabel.SetString(label)
	return p
}

func (p *Plot) SetYLabel(label string) *Plot {

	p.ylabel.SetString(label)
	return p
}

func (p *Plot) SetSize(size imgui.Vec2) *Plot {

	p.size = size
	return p
}

func (p *Plot) SetFlags(flags imgui.PlotFlags) *Plot {

	p.flags = flags
	return p
}

func (p *Plot) Flags() imgui.PlotFlags {

	return p.flags
}

func (p *Plot) SetXAxisFlags(flags imgui.PlotAxisFlags) *Plot {

	p.xflags = flags
	return p
}

func (p *Plot) SetYAxisFlags(flags imgui.PlotAxisFlags) *Plot {

	p.yflags = flags
	return p
}

func (p *Plot) SetTicksX(values []float64, labels []string) *Plot {

	p.ClearTicksXRange()
	tx := p.TicksX()
	if tx != nil {
		tx.Set(values, labels)
		return p
	}
	p.preList = append(p.preList, NewSetNextPlotTicksX(values, labels).GetWidget())
	return p
}

func (p *Plot) TicksX() *SetNextPlotTicksX {

	for _, it := range p.preList {
		tx, ok := it.super.(*SetNextPlotTicksX)
		if ok {
			return tx
		}
	}
	return nil
}

func (p *Plot) ClearTicksX() bool {

	for idx, it := range p.preList {
		_, ok := it.super.(*SetNextPlotTicksX)
		if ok {
			p.delPreListItem(idx)
			return true
		}
	}
	return false
}

func (p *Plot) TicksXRange() *SetNextPlotTicksXRange {

	for _, it := range p.preList {
		tx, ok := it.super.(*SetNextPlotTicksXRange)
		if ok {
			return tx
		}
	}
	return nil
}

func (p *Plot) SetTicksXRange(min, max float64, nticks int, labels []string) *Plot {

	p.ClearTicksX()
	txr := p.TicksXRange()
	if txr != nil {
		txr.Set(min, max, nticks, labels)
		return p
	}
	p.appendPreList(NewSetNextPlotTicksXRange(min, max, nticks, labels))
	return p
}

func (p *Plot) ClearTicksXRange() bool {

	for idx, it := range p.preList {
		_, ok := it.super.(*SetNextPlotTicksXRange)
		if ok {
			p.delPreListItem(idx)
			return true
		}
	}
	return false
}

func (p *Plot) SetTicksY(values []float64, labels []string) *Plot {

	ty := p.TicksY()
	if ty != nil {
		ty.Set(values, labels)
		return p
	}
	p.appendPreList(NewSetNextPlotTicksY(values, labels))
	return p
}

func (p *Plot) TicksY() *SetNextPlotTicksY {

	for _, it := range p.preList {
		ty, ok := it.super.(*SetNextPlotTicksY)
		if ok {
			return ty
		}
	}
	return nil
}

func (p *Plot) ClearTicksY() bool {

	for idx, it := range p.preList {
		_, ok := it.super.(*SetNextPlotTicksY)
		if ok {
			p.delPreListItem(idx)
			return true
		}
	}
	return false
}

func (p *Plot) SetTicksYRange(min, max float64, nticks int, labels []string) *Plot {

	tyr := p.TicksYRange()
	if tyr != nil {
		tyr.Set(min, max, nticks, labels)
		return p
	}
	p.appendPreList(NewSetNextPlotTicksYRange(min, max, nticks, labels))
	return p
}

func (p *Plot) TicksYRange() *SetNextPlotTicksYRange {

	for _, it := range p.preList {
		tyr, ok := it.super.(*SetNextPlotTicksYRange)
		if ok {
			return tyr
		}
	}
	return nil
}
func (p *Plot) LimitsX() *SetNextPlotLimitsX {

	for _, it := range p.preList {
		lx, ok := it.super.(*SetNextPlotLimitsX)
		if ok {
			return lx
		}
	}
	return nil
}

func (p *Plot) SetLimitsX(xmin, xmax float64, cond imgui.Cond) *Plot {

	lx := p.LimitsX()
	if lx != nil {
		lx.Set(xmin, xmax, cond)
		return p
	}
	p.appendPreList(NewSetNextPlotLimitsX(xmin, xmax, cond))
	return p
}

func (p *Plot) ClearLimitsX() bool {

	for idx, it := range p.preList {
		_, ok := it.super.(*SetNextPlotLimitsX)
		if ok {
			p.delPreListItem(idx)
			return true
		}
	}
	return false
}

func (p *Plot) LimitsY() *SetNextPlotLimitsY {

	for _, it := range p.preList {
		ly, ok := it.super.(*SetNextPlotLimitsY)
		if ok {
			return ly
		}
	}
	return nil
}

func (p *Plot) SetLimitsY(ymin, ymax float64, cond imgui.Cond, yaxis imgui.PlotYAxis) *Plot {

	ly := p.LimitsY()
	if ly != nil {
		ly.Set(ymin, ymax, cond)
		return p
	}
	p.appendPreList(NewSetNextPlotLimitsY(ymin, ymax, cond))
	return p
}

func (p *Plot) ClearLimitsY() bool {

	for idx, it := range p.preList {
		_, ok := it.super.(*SetNextPlotLimitsY)
		if ok {
			p.delPreListItem(idx)
			return true
		}
	}
	return false
}

func (p *Plot) Limits() *SetNextPlotLimits {

	for _, it := range p.preList {
		l, ok := it.super.(*SetNextPlotLimits)
		if ok {
			return l
		}
	}
	return nil
}

func (p *Plot) SetLimits(xmin, xmax, ymin, ymax float64, cond imgui.Cond) *Plot {

	p.ClearLimitsX()
	p.ClearLimitsY()
	l := p.Limits()
	if l != nil {
		l.Set(xmin, xmax, ymin, ymax, cond)
		return p
	}
	p.appendPreList(NewSetNextPlotLimits(xmin, xmax, ymin, ymax, cond))
	return p
}

func (p *Plot) LegendLocation() *SetLegendLocation {

	for _, c := range p.Children() {
		lc, ok := c.super.(*SetLegendLocation)
		if ok {
			return lc
		}
	}
	return nil
}

func (p *Plot) SetLegendLocation(location imgui.PlotLocation, orientation imgui.PlotOrientation, outside bool) *Plot {

	lc := p.LegendLocation()
	if lc != nil {
		lc.Set(location, orientation, outside)
		return p
	}
	p.AddChildren(NewSetLegendLocation(location, orientation, outside))
	return p
}

// plotBase is the base Widget for all plot objects
type plotBase struct {
	Widget
}

func (pb *plotBase) SetStyleColor(idx imgui.PlotCol, color imgui.Vec4) *plotBase {

	for _, it := range pb.preList {
		sc, ok := it.super.(*PlotPushStyleColor)
		if ok && sc.idx == idx {
			sc.col = color
			return pb
		}
	}
	pb.preList = append(pb.preList, NewPlotPushStyleColor(idx, color).GetWidget())
	pb.postList = append(pb.postList, NewPlotPopStyleColor(1).GetWidget())
	return pb
}

func (pb *plotBase) SetStyleVar(idx imgui.PlotStyleVar, value interface{}) *plotBase {

	for _, it := range pb.preList {
		sc, ok := it.super.(*PlotPushStyleVar)
		if ok && sc.idx == idx {
			sc.value = value
			return pb
		}
	}
	pb.preList = append(pb.preList, NewPlotPushStyleVar(idx, value).GetWidget())
	pb.postList = append(pb.postList, NewPlotPopStyleVar(1).GetWidget())
	return pb
}

type SetNextErrorBarStyle struct {
	Widget
	color  imgui.Vec4
	size   float32
	weight float32
}

func NewSetNextErrorBarStyle() *SetNextErrorBarStyle {

	s := new(SetNextErrorBarStyle)
	s.Init(s)
	s.color = imgui.PlotAutoCol
	s.size = imgui.PlotAuto
	s.weight = imgui.PlotAuto
	s.SetRender(func() {
		imgui.SetNextErrorBarStyle(s.color, s.size, s.weight)
	})
	return s
}

func (s *SetNextErrorBarStyle) SetColor(color imgui.Vec4) *SetNextErrorBarStyle {

	s.color = color
	return s
}

func (s *SetNextErrorBarStyle) SetSize(size float32) *SetNextErrorBarStyle {

	s.size = size
	return s
}

func (s *SetNextErrorBarStyle) SetWeight(weight float32) *SetNextErrorBarStyle {

	s.weight = weight
	return s
}

type SetNextLineStyle struct {
	Widget
	col    imgui.Vec4
	weight float32
}

func NewSetNextLineStyle(col imgui.Vec4, weight float32) *SetNextLineStyle {

	s := new(SetNextLineStyle)
	s.Init(s)
	s.col = col
	s.weight = weight
	s.SetRender(func() {
		imgui.SetNextLineStyle(s.col, s.weight)
	})
	return s
}

func (s *SetNextLineStyle) SetColor(col imgui.Vec4) *SetNextLineStyle {

	s.col = col
	return s
}

func (s *SetNextLineStyle) SetWeight(weight float32) *SetNextLineStyle {

	s.weight = weight
	return s
}

type SetNextMarkerStyle struct {
	Widget
	marker  imgui.PlotMarker
	size    float32
	fill    imgui.Vec4
	weight  float32
	outline imgui.Vec4
}

func NewSetNextMarkerStyle(marker imgui.PlotMarker) *SetNextMarkerStyle {

	s := new(SetNextMarkerStyle)
	s.Init(s)
	s.marker = marker
	s.size = imgui.PlotAuto
	s.fill = imgui.PlotAutoCol
	s.weight = imgui.PlotAuto
	s.outline = imgui.PlotAutoCol
	s.SetRender(func() {
		imgui.SetNextMarkerStyle(s.marker, s.size, s.fill, s.weight, s.outline)
	})
	return s
}

func (s *SetNextMarkerStyle) SetMarker(marker imgui.PlotMarker) *SetNextMarkerStyle {

	s.marker = marker
	return s
}

func (s *SetNextMarkerStyle) SetSize(size float32) *SetNextMarkerStyle {

	s.size = size
	return s
}

func (s *SetNextMarkerStyle) SetFill(fill imgui.Vec4) *SetNextMarkerStyle {

	s.fill = fill
	return s
}

func (s *SetNextMarkerStyle) SetWeight(weight float32) *SetNextMarkerStyle {

	s.weight = weight
	return s
}

func (s *SetNextMarkerStyle) SetOutline(outline imgui.Vec4) *SetNextMarkerStyle {

	s.outline = outline
	return s
}

type PushColormap struct {
	Widget
	cmap imgui.PlotColormap
}

func NewPushColormap(cmap imgui.PlotColormap) *PushColormap {

	c := new(PushColormap)
	c.Init(c)
	c.cmap = cmap
	c.SetRender(func() {
		imgui.PushColormap(c.cmap)
	})
	return c
}

func (c *PushColormap) Set(cmap imgui.PlotColormap) *PushColormap {

	c.cmap = cmap
	return c
}

type PopColormap struct {
	Widget
	count int
}

func NewPopColormap(count int) *PopColormap {

	c := new(PopColormap)
	c.Init(c)
	c.count = count
	c.SetRender(func() {
		imgui.PopColormap(c.count)
	})
	return c
}

type PlotPushStyleColor struct {
	Widget
	idx imgui.PlotCol
	col imgui.Vec4
}

func NewPlotPushStyleColor(idx imgui.PlotCol, col imgui.Vec4) *PlotPushStyleColor {

	s := new(PlotPushStyleColor)
	s.Init(s)
	s.idx = idx
	s.col = col
	s.SetRender(func() {
		imgui.PlotPushStyleColor(s.idx, s.col)
	})
	return s
}

func (s *PlotPushStyleColor) Set(idx imgui.PlotCol, col imgui.Vec4) *PlotPushStyleColor {

	s.idx = idx
	s.col = col
	return s
}

type PlotPopStyleColor struct {
	Widget
	count int
}

func NewPlotPopStyleColor(count int) *PlotPopStyleColor {

	s := new(PlotPopStyleColor)
	s.Init(s)
	s.count = count
	s.SetRender(func() {
		imgui.PlotPopStyleColor(s.count)
	})
	return s
}

type PlotPushStyleVar struct {
	Widget
	idx   imgui.PlotStyleVar
	value interface{}
}

func NewPlotPushStyleVar(idx imgui.PlotStyleVar, value interface{}) *PlotPushStyleVar {

	w := new(PlotPushStyleVar)
	w.Init(w)
	w.idx = idx
	w.value = value
	w.SetRender(func() {
		switch v := w.value.(type) {
		case int:
			imgui.PlotPushStyleVarInt(idx, v)
		case float32:
			imgui.PlotPushStyleVarFloat32(idx, v)
		case imgui.Vec2:
			imgui.PlotPushStyleVarVec2(idx, v)
		default:
			panic("Invalid value type for PlotStyleVar")
		}
	})
	return w
}

func (w *PlotPushStyleVar) Set(idx imgui.PlotStyleVar, value interface{}) *PlotPushStyleVar {

	w.idx = idx
	w.value = value
	return w
}

type PlotPopStyleVar struct {
	Widget
	count int
}

func NewPlotPopStyleVar(count int) *PlotPopStyleVar {

	w := new(PlotPopStyleVar)
	w.Init(w)
	w.count = count
	w.SetRender(func() {
		imgui.PlotPopStyleVar(w.count)
	})
	return w
}

type SetNextPlotTicksX struct {
	Widget
	values      []float64
	labels      imgui.CStrArr
	keepDefault bool
}

func NewSetNextPlotTicksX(values []float64, labels []string) *SetNextPlotTicksX {

	w := new(SetNextPlotTicksX)
	w.Init(w)
	w.Set(values, labels)
	w.SetRender(func() {
		imgui.SetNextPlotTicksXCS(w.values, w.labels, w.keepDefault)
	})
	return w
}

func (w *SetNextPlotTicksX) Set(values []float64, labels []string) *SetNextPlotTicksX {

	w.values = values
	w.labels.Clear()
	w.labels.Append(labels...)
	return w
}

func (w *SetNextPlotTicksX) SetKeepDefault(keepDefault bool) *SetNextPlotTicksX {

	w.keepDefault = keepDefault
	return w
}

type SetNextPlotTicksXRange struct {
	Widget
	min         float64
	max         float64
	nticks      int
	labels      imgui.CStrArr
	keepDefault bool
}

func NewSetNextPlotTicksXRange(min, max float64, nticks int, labels []string) *SetNextPlotTicksXRange {

	w := new(SetNextPlotTicksXRange)
	w.Init(w)
	w.Set(min, max, nticks, labels)
	w.SetRender(func() {
		imgui.SetNextPlotTicksXRangeCS(w.min, w.max, w.nticks, w.labels, w.keepDefault)
	})
	return w
}

func (w *SetNextPlotTicksXRange) Set(min, max float64, nticks int, labels []string) *SetNextPlotTicksXRange {

	w.min = min
	w.max = max
	w.nticks = nticks
	w.labels.Append(labels...)
	return w
}

type SetNextPlotTicksY struct {
	Widget
	values      []float64
	labels      imgui.CStrArr
	keepDefault bool
	yaxis       imgui.PlotYAxis
}

func NewSetNextPlotTicksY(values []float64, labels []string) *SetNextPlotTicksY {

	w := new(SetNextPlotTicksY)
	w.Init(w)
	w.values = values
	w.labels.Append(labels...)
	w.yaxis = imgui.PlotYAxis_1
	w.SetRender(func() {
		imgui.SetNextPlotTicksYCS(w.values, w.labels, w.keepDefault, w.yaxis)
	})
	return w
}

func (w *SetNextPlotTicksY) Set(values []float64, labels []string) *SetNextPlotTicksY {

	w.values = values
	w.labels.Clear()
	w.labels.Append(labels...)
	return w
}

func (w *SetNextPlotTicksY) SetKeepDefault(keepDefault bool) *SetNextPlotTicksY {

	w.keepDefault = keepDefault
	return w
}

func (w *SetNextPlotTicksY) SetYAxis(yaxis imgui.PlotYAxis) *SetNextPlotTicksY {

	w.yaxis = yaxis
	return w
}

type SetNextPlotTicksYRange struct {
	Widget
	min         float64
	max         float64
	nticks      int
	labels      imgui.CStrArr
	keepDefault bool
	yaxis       imgui.PlotYAxis
}

func NewSetNextPlotTicksYRange(min, max float64, nticks int, labels []string) *SetNextPlotTicksYRange {

	w := new(SetNextPlotTicksYRange)
	w.Init(w)
	w.min = min
	w.max = max
	w.nticks = nticks
	w.labels.Append(labels...)
	w.yaxis = imgui.PlotYAxis_1
	w.SetRender(func() {
		imgui.SetNextPlotTicksYRangeCS(w.min, w.max, w.nticks, w.labels, w.keepDefault, w.yaxis)
	})
	return w
}

func (w *SetNextPlotTicksYRange) Set(min, max float64, nticks int, labels []string) *SetNextPlotTicksYRange {

	w.min = min
	w.max = max
	w.nticks = nticks
	w.labels.Clear()
	w.labels.Append(labels...)
	return w
}

func (w *SetNextPlotTicksYRange) SetKeepDefault(keepDefault bool) *SetNextPlotTicksYRange {

	w.keepDefault = keepDefault
	return w
}

func (w *SetNextPlotTicksYRange) SetYAxis(yaxis imgui.PlotYAxis) *SetNextPlotTicksYRange {

	w.yaxis = yaxis
	return w
}

type SetNextPlotLimitsX struct {
	Widget
	min  float64
	max  float64
	cond imgui.Cond
}

func NewSetNextPlotLimitsX(min, max float64, cond imgui.Cond) *SetNextPlotLimitsX {

	w := new(SetNextPlotLimitsX)
	w.Init(w)
	w.Set(min, max, cond)
	w.SetRender(func() {
		imgui.SetNextPlotLimitsX(w.min, w.max, w.cond)
	})
	return w
}

func (w *SetNextPlotLimitsX) Set(min, max float64, cond imgui.Cond) *SetNextPlotLimitsX {

	w.min = min
	w.max = max
	w.cond = cond
	return w
}

type SetNextPlotLimitsY struct {
	Widget
	min   float64
	max   float64
	cond  imgui.Cond
	yaxis imgui.PlotYAxis
}

func NewSetNextPlotLimitsY(min, max float64, cond imgui.Cond) *SetNextPlotLimitsY {

	w := new(SetNextPlotLimitsY)
	w.Init(w)
	w.min = min
	w.max = max
	w.cond = cond
	w.yaxis = imgui.PlotYAxis_1
	w.SetRender(func() {
		imgui.SetNextPlotLimitsY(w.min, w.max, w.cond, w.yaxis)
	})
	return w
}

func (w *SetNextPlotLimitsY) Set(min, max float64, cond imgui.Cond) *SetNextPlotLimitsY {

	w.min = min
	w.max = max
	w.cond = cond
	return w
}

func (w *SetNextPlotLimitsY) SetYAxis(yaxis imgui.PlotYAxis) *SetNextPlotLimitsY {

	w.yaxis = yaxis
	return w
}

type SetNextPlotLimits struct {
	Widget
	xmin float64
	xmax float64
	ymin float64
	ymax float64
	cond imgui.Cond
}

func NewSetNextPlotLimits(xmin, xmax, ymin, ymax float64, cond imgui.Cond) *SetNextPlotLimits {

	w := new(SetNextPlotLimits)
	w.Init(w)
	w.Set(xmin, xmax, ymin, ymax, cond)
	w.SetRender(func() {
		imgui.SetNextPlotLimits(w.xmin, w.xmax, w.ymin, w.ymax, w.cond)
	})
	return w
}

func (w *SetNextPlotLimits) Set(xmin, xmax, ymin, ymax float64, cond imgui.Cond) *SetNextPlotLimits {

	w.xmin = xmin
	w.xmax = xmax
	w.ymin = ymin
	w.ymax = ymax
	w.cond = cond
	return w
}

type SetLegendLocation struct {
	Widget
	location    imgui.PlotLocation
	orientation imgui.PlotOrientation
	outside     bool
}

func NewSetLegendLocation(location imgui.PlotLocation, orientation imgui.PlotOrientation, outside bool) *SetLegendLocation {

	w := new(SetLegendLocation)
	w.Init(w)
	w.Set(location, orientation, outside)
	w.SetRender(func() {
		imgui.SetLegendLocation(w.location, w.orientation, w.outside)
	})
	return w
}

func (w *SetLegendLocation) Set(location imgui.PlotLocation, orientation imgui.PlotOrientation, outside bool) *SetLegendLocation {

	w.location = location
	w.orientation = orientation
	w.outside = outside
	return w
}

type ColormapButton struct {
	Widget
	label   imgui.CString
	size    imgui.Vec2
	cmap    imgui.PlotColormap
	onclick func(*ColormapButton)
}

func NewColormapButton(label string) *ColormapButton {

	b := new(ColormapButton)
	b.Init(b)
	b.label.SetString(label)
	b.cmap = imgui.PlotAuto
	b.SetRender(func() {
		if imgui.ColormapButtonCS(b.label, b.size, b.cmap) {
			if b.onclick != nil {
				b.onclick(b)
			}
		}
	})
	return b
}

func (b *ColormapButton) SetLabel(label string) *ColormapButton {

	b.label.SetString(label)
	return b
}

func (b *ColormapButton) SetSize(size imgui.Vec2) *ColormapButton {

	b.size = size
	return b
}

func (b *ColormapButton) SetColormap(cmap imgui.PlotColormap) *ColormapButton {

	b.cmap = cmap
	return b
}

func (b *ColormapButton) SetOnClick(f func(b *ColormapButton)) *ColormapButton {

	b.onclick = f
	return b
}

type ColormapScale struct {
	Widget
	label    imgui.CString
	scaleMin float64
	scaleMax float64
	size     imgui.Vec2
	cmap     imgui.PlotColormap
	fmt      imgui.CString
}

//IMPLOT_API void ColormapScale(const char* label, double scale_min, double scale_max, const ImVec2& size = ImVec2(0,0),
//ImPlotColormap cmap = IMPLOT_AUTO, const char* fmt = "%g");

func NewColormapScale(label string, scaleMin, scaleMax float64) *ColormapScale {

	c := new(ColormapScale)
	c.Init(c)
	c.label.SetString(label)
	c.scaleMin = scaleMin
	c.scaleMax = scaleMax
	c.cmap = imgui.PlotAuto
	c.fmt.SetString("%g")
	c.SetRender(func() {
		imgui.ColormapScaleCS(c.label, c.scaleMin, c.scaleMax, c.size, c.cmap, c.fmt)
	})
	return c
}

func (c *ColormapScale) SetScale(min, max float64) *ColormapScale {

	c.scaleMin = min
	c.scaleMax = max
	return c
}

func (c *ColormapScale) SetSize(size imgui.Vec2) *ColormapScale {

	c.size = size
	return c
}

func (c *ColormapScale) SetColormap(cmap imgui.PlotColormap) *ColormapScale {

	c.cmap = cmap
	return c
}

func (c *ColormapScale) SetFormat(fmt string) *ColormapScale {

	c.fmt.SetString(fmt)
	return c
}
