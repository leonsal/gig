package imgui

// #include <time.h>
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
// #include "cimplot.h"
import "C"

const PlotAuto = -1

var PlotAutoCol = Vec4{0, 0, 0, -1}

type PlotFlags C.ImPlotFlags

const (
	PlotFlags_None        PlotFlags = C.ImPlotFlags_None
	PlotFlags_NoTitle     PlotFlags = C.ImPlotFlags_NoTitle
	PlotFlags_NoLegend    PlotFlags = C.ImPlotFlags_NoLegend
	PlotFlags_NoMenus     PlotFlags = C.ImPlotFlags_NoMenus
	PlotFlags_NoBoxSelect PlotFlags = C.ImPlotFlags_NoBoxSelect
	PlotFlags_NoMousePos  PlotFlags = C.ImPlotFlags_NoMousePos
	PlotFlags_NoHighlight PlotFlags = C.ImPlotFlags_NoHighlight
	PlotFlags_NoChild     PlotFlags = C.ImPlotFlags_NoChild
	PlotFlags_Equal       PlotFlags = C.ImPlotFlags_Equal
	PlotFlags_YAxis2      PlotFlags = C.ImPlotFlags_YAxis2
	PlotFlags_YAxis3      PlotFlags = C.ImPlotFlags_YAxis3
	PlotFlags_Query       PlotFlags = C.ImPlotFlags_Query
	PlotFlags_Crosshairs  PlotFlags = C.ImPlotFlags_Crosshairs
	PlotFlags_AntiAliased PlotFlags = C.ImPlotFlags_AntiAliased
	PlotFlags_CanvasOnly  PlotFlags = C.ImPlotFlags_CanvasOnly
)

type PlotAxisFlags C.ImPlotAxisFlags

const (
	PlotAxisFlags_None          PlotAxisFlags = C.ImPlotAxisFlags_None
	PlotAxisFlags_NoLabel       PlotAxisFlags = C.ImPlotAxisFlags_NoLabel
	PlotAxisFlags_NoGridLines   PlotAxisFlags = C.ImPlotAxisFlags_NoGridLines
	PlotAxisFlags_NoTickMarks   PlotAxisFlags = C.ImPlotAxisFlags_NoTickMarks
	PlotAxisFlags_NoTickLabels  PlotAxisFlags = C.ImPlotAxisFlags_NoTickLabels
	PlotAxisFlags_Foreground    PlotAxisFlags = C.ImPlotAxisFlags_Foreground
	PlotAxisFlags_LogScale      PlotAxisFlags = C.ImPlotAxisFlags_LogScale
	PlotAxisFlags_Time          PlotAxisFlags = C.ImPlotAxisFlags_Time
	PlotAxisFlags_Invert        PlotAxisFlags = C.ImPlotAxisFlags_Invert
	PlotAxisFlags_NoInitialFit  PlotAxisFlags = C.ImPlotAxisFlags_NoInitialFit
	PlotAxisFlags_AutoFit       PlotAxisFlags = C.ImPlotAxisFlags_AutoFit
	PlotAxisFlags_RangeFit      PlotAxisFlags = C.ImPlotAxisFlags_RangeFit
	PlotAxisFlags_LockMin       PlotAxisFlags = C.ImPlotAxisFlags_LockMin
	PlotAxisFlags_LockMax       PlotAxisFlags = C.ImPlotAxisFlags_LockMax
	PlotAxisFlags_Lock          PlotAxisFlags = C.ImPlotAxisFlags_Lock
	PlotAxisFlags_NoDecorations PlotAxisFlags = C.ImPlotAxisFlags_NoDecorations
)

type PlotSubplotFlags C.ImPlotSubplotFlags

const (
	PlotSubplotFlags_None       PlotSubplotFlags = C.ImPlotSubplotFlags_None
	PlotSubplotFlags_NoTitle    PlotSubplotFlags = C.ImPlotSubplotFlags_NoTitle
	PlotSubplotFlags_NoLegend   PlotSubplotFlags = C.ImPlotSubplotFlags_NoLegend
	PlotSubplotFlags_NoMenus    PlotSubplotFlags = C.ImPlotSubplotFlags_NoMenus
	PlotSubplotFlags_NoResize   PlotSubplotFlags = C.ImPlotSubplotFlags_NoResize
	PlotSubplotFlags_NoAlign    PlotSubplotFlags = C.ImPlotSubplotFlags_NoAlign
	PlotSubplotFlags_ShareItems PlotSubplotFlags = C.ImPlotSubplotFlags_ShareItems
	PlotSubplotFlags_LinkRows   PlotSubplotFlags = C.ImPlotSubplotFlags_LinkRows
	PlotSubplotFlags_LinkCols   PlotSubplotFlags = C.ImPlotSubplotFlags_LinkCols
	PlotSubplotFlags_LinkAllX   PlotSubplotFlags = C.ImPlotSubplotFlags_LinkAllX
	PlotSubplotFlags_LinkAllY   PlotSubplotFlags = C.ImPlotSubplotFlags_LinkAllY
	PlotSubplotFlags_ColMajor   PlotSubplotFlags = C.ImPlotSubplotFlags_ColMajor
)

type PlotCol C.ImPlotCol

const (
	PlotCol_Line          = C.ImPlotCol_Line
	PlotCol_Fill          = C.ImPlotCol_Fill
	PlotCol_MarkerOutline = C.ImPlotCol_MarkerOutline
	PlotCol_MarkerFill    = C.ImPlotCol_MarkerFill
	PlotCol_ErrorBar      = C.ImPlotCol_ErrorBar
	PlotCol_FrameBg       = C.ImPlotCol_FrameBg
	PlotCol_PlotBg        = C.ImPlotCol_PlotBg
	PlotCol_PlotBorder    = C.ImPlotCol_PlotBorder
	PlotCol_LegendBg      = C.ImPlotCol_LegendBg
	PlotCol_LegendBorder  = C.ImPlotCol_LegendBorder
	PlotCol_LegendText    = C.ImPlotCol_LegendText
	PlotCol_TitleText     = C.ImPlotCol_TitleText
	PlotCol_InlayText     = C.ImPlotCol_InlayText
	PlotCol_XAxis         = C.ImPlotCol_XAxis
	PlotCol_XAxisGrid     = C.ImPlotCol_XAxisGrid
	PlotCol_YAxis         = C.ImPlotCol_YAxis
	PlotCol_YAxisGrid     = C.ImPlotCol_YAxisGrid
	PlotCol_YAxis2        = C.ImPlotCol_YAxis2
	PlotCol_YAxisGrid2    = C.ImPlotCol_YAxisGrid2
	PlotCol_YAxis3        = C.ImPlotCol_YAxis3
	PlotCol_YAxisGrid3    = C.ImPlotCol_YAxisGrid3
	PlotCol_Selection     = C.ImPlotCol_Selection
	PlotCol_Query         = C.ImPlotCol_Query
	PlotCol_Crosshairs    = C.ImPlotCol_Crosshairs
	PlotCol_COUNT         = C.ImPlotCol_COUNT
)

type PlotStyleVar C.ImPlotStyleVar

const (
	PlotStyleVar_LineWeight         = C.ImPlotStyleVar_LineWeight
	PlotStyleVar_Marker             = C.ImPlotStyleVar_Marker
	PlotStyleVar_MarkerSize         = C.ImPlotStyleVar_MarkerSize
	PlotStyleVar_MarkerWeight       = C.ImPlotStyleVar_MarkerWeight
	PlotStyleVar_FillAlpha          = C.ImPlotStyleVar_FillAlpha
	PlotStyleVar_ErrorBarSize       = C.ImPlotStyleVar_ErrorBarSize
	PlotStyleVar_ErrorBarWeight     = C.ImPlotStyleVar_ErrorBarWeight
	PlotStyleVar_DigitalBitHeight   = C.ImPlotStyleVar_DigitalBitHeight
	PlotStyleVar_DigitalBitGap      = C.ImPlotStyleVar_DigitalBitGap
	PlotStyleVar_PlotBorderSize     = C.ImPlotStyleVar_PlotBorderSize
	PlotStyleVar_MinorAlpha         = C.ImPlotStyleVar_MinorAlpha
	PlotStyleVar_MajorTickLen       = C.ImPlotStyleVar_MajorTickLen
	PlotStyleVar_MinorTickLen       = C.ImPlotStyleVar_MinorTickLen
	PlotStyleVar_MajorTickSize      = C.ImPlotStyleVar_MajorTickSize
	PlotStyleVar_MinorTickSize      = C.ImPlotStyleVar_MinorTickSize
	PlotStyleVar_MajorGridSize      = C.ImPlotStyleVar_MajorGridSize
	PlotStyleVar_MinorGridSize      = C.ImPlotStyleVar_MinorGridSize
	PlotStyleVar_PlotPadding        = C.ImPlotStyleVar_PlotPadding
	PlotStyleVar_LabelPadding       = C.ImPlotStyleVar_LabelPadding
	PlotStyleVar_LegendPadding      = C.ImPlotStyleVar_LegendPadding
	PlotStyleVar_LegendInnerPadding = C.ImPlotStyleVar_LegendInnerPadding
	PlotStyleVar_LegendSpacing      = C.ImPlotStyleVar_LegendSpacing
	PlotStyleVar_MousePosPadding    = C.ImPlotStyleVar_MousePosPadding
	PlotStyleVar_AnnotationPadding  = C.ImPlotStyleVar_AnnotationPadding
	PlotStyleVar_FitPadding         = C.ImPlotStyleVar_FitPadding
	PlotStyleVar_PlotDefaultSize    = C.ImPlotStyleVar_PlotDefaultSize
	PlotStyleVar_PlotMinSize        = C.ImPlotStyleVar_PlotMinSize
	PlotStyleVar_COUNT              = C.ImPlotStyleVar_COUNT
)

type PlotMarker C.ImPlotMarker

const (
	PlotMarker_None     = C.ImPlotMarker_None
	PlotMarker_Circle   = C.ImPlotMarker_Circle
	PlotMarker_Square   = C.ImPlotMarker_Square
	PlotMarker_Diamond  = C.ImPlotMarker_Diamond
	PlotMarker_Up       = C.ImPlotMarker_Up
	PlotMarker_Down     = C.ImPlotMarker_Down
	PlotMarker_Left     = C.ImPlotMarker_Left
	PlotMarker_Right    = C.ImPlotMarker_Right
	PlotMarker_Cross    = C.ImPlotMarker_Cross
	PlotMarker_Plus     = C.ImPlotMarker_Plus
	PlotMarker_Asterisk = C.ImPlotMarker_Asterisk
	PlorMarker_COUNT    = C.ImPlotMarker_COUNT
)

type PlotColormap C.ImPlotColormap

const (
	PlotColormap_Deep     PlotColormap = C.ImPlotColormap_Deep
	PlotColormap_Dark     PlotColormap = C.ImPlotColormap_Dark
	PlotColormap_Pastel   PlotColormap = C.ImPlotColormap_Pastel
	PlotColormap_Paired   PlotColormap = C.ImPlotColormap_Paired
	PlotColormap_Viridis  PlotColormap = C.ImPlotColormap_Viridis
	PlotColormap_Plasma   PlotColormap = C.ImPlotColormap_Plasma
	PlotColormap_Hot      PlotColormap = C.ImPlotColormap_Hot
	PlotColormap_Cool     PlotColormap = C.ImPlotColormap_Cool
	PlotColormap_Pink     PlotColormap = C.ImPlotColormap_Pink
	PlotColormap_Jet      PlotColormap = C.ImPlotColormap_Jet
	PlotColormap_Twilight PlotColormap = C.ImPlotColormap_Twilight
	PlotColormap_RdBu     PlotColormap = C.ImPlotColormap_RdBu
	PlotColormap_BrBG     PlotColormap = C.ImPlotColormap_BrBG
	PlotColormap_PiYG     PlotColormap = C.ImPlotColormap_PiYG
	PlotColormap_Spectral PlotColormap = C.ImPlotColormap_Spectral
	PlotColormap_Greys    PlotColormap = C.ImPlotColormap_Greys
)

type PlotLocation C.ImPlotLocation

const (
	PlotLocation_Center    = C.ImPlotLocation_Center
	PlotLocation_North     = C.ImPlotLocation_North
	PlotLocation_South     = C.ImPlotLocation_South
	PlotLocation_West      = C.ImPlotLocation_West
	PlotLocation_East      = C.ImPlotLocation_East
	PlotLocation_NorthWest = C.ImPlotLocation_NorthWest
	PlotLocation_NorthEast = C.ImPlotLocation_NorthEast
	PlotLocation_SouthWest = C.ImPlotLocation_SouthWest
	PlotLocation_SouthEast = C.ImPlotLocation_SouthEast
)

type PlotOrientation C.ImPlotOrientation

const (
	PlotOrientation_Horizontal = C.ImPlotOrientation_Horizontal
	PlotOrientation_Vertical   = C.ImPlotOrientation_Vertical
)

type PlotYAxis C.ImPlotYAxis

const (
	PlotYAxis_1 = C.ImPlotYAxis_1
	PlotYAxis_2 = C.ImPlotYAxis_2
	PlotYAxis_3 = C.ImPlotYAxis_3
)

type PlotBin C.ImPlotBin

const (
	PlotBin_Sqrt    = C.ImPlotBin_Sqrt
	PlotBin_Sturges = C.ImPlotBin_Sturges
	PlotBin_Rice    = C.ImPlotBin_Rice
	PlotBin_Scott   = C.ImPlotBin_Scott
)

type PlotPoint struct {
	X float64
	Y float64
}

func (p PlotPoint) ImPlotPoint() C.ImPlotPoint {

	return C.ImPlotPoint{C.double(p.X), C.double(p.Y)}
}

func (p C.ImPlotPoint) PlotPoint() PlotPoint {

	return PlotPoint{float64(p.x), float64(p.y)}
}

type PlotRange struct {
	Min float64
	Max float64
}

func (r PlotRange) ImPlotRange() C.ImPlotRange {

	return C.ImPlotRange{C.double(r.Min), C.double(r.Max)}
}

type PlotStyle C.ImPlotStyle

type PlotLimits struct {
	X PlotRange
	Y PlotRange
}

func (pl PlotLimits) ImPlotLimits() C.ImPlotLimits {

	return C.ImPlotLimits{
		C.ImPlotRange{C.double(pl.X.Min), C.double(pl.X.Max)},
		C.ImPlotRange{C.double(pl.Y.Min), C.double(pl.Y.Max)},
	}
}

func (pl C.ImPlotLimits) PlotLimits() PlotLimits {

	return PlotLimits{
		PlotRange{float64(pl.X.Min), float64(pl.X.Max)},
		PlotRange{float64(pl.Y.Min), float64(pl.Y.Max)},
	}
}

type PlotAxis C.ImPlotAxis
