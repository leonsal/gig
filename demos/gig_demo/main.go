package main

import (
	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
	"github.com/leonsal/gig/logtable"
)

// demoInfo describes a demo
type demoInfo struct {
	category catType
	name     string
	fn       func() gig.IWidget
}

// Maps demo category to slice of demoInfo's of that category
var mapDemos = map[catType][]demoInfo{}

// Package global logger
var log = logtable.Get()

type catType string

const (
	catImGui   catType = "ImGui"
	catImPlot  catType = "ImPlot"
	catWidgets catType = "Widgets"
)

func main() {

	// Creates App
	view, err := gig.ViewInit("Gig Demo", 800, 600)
	if err != nil {
		panic(err)
	}

	// Builds UI and run
	buildUI(view)
	view.Run()
}

func buildUI(a *gig.View) {

	mainMenu := gig.NewMainMenuBar().AddChildren(
		gig.NewMenu("Menu").AddChildren(
			gig.NewMenuItem("Show Gig Demos").SetOnClick(func(mi *gig.MenuItem) {
				gig.Ref("splitter").(*gig.WindowVSplitter).SetVisible(true)
				gig.Ref("imguiDemo").(*gig.ImGuiDemoWindow).SetVisible(false)
				gig.Ref("implotDemo").(*gig.ImPlotDemoWindow).SetVisible(false)
			}),
			gig.NewSeparator(),
			gig.NewMenuItem("Show C++ ImGui Demo Window").SetOnClick(func(mi *gig.MenuItem) {
				gig.Ref("splitter").(*gig.WindowVSplitter).SetVisible(false)
				gig.Ref("imguiDemo").(*gig.ImGuiDemoWindow).SetVisible(true)
			}),
			gig.NewMenuItem("Show C++ ImPlot Demo Window").SetOnClick(func(mi *gig.MenuItem) {
				gig.Ref("splitter").(*gig.WindowVSplitter).SetVisible(false)
				gig.Ref("implotDemo").(*gig.ImPlotDemoWindow).SetVisible(true)
			}),
			gig.NewSeparator(),
			gig.NewMenuItem("Clear log").SetOnClick(func(mi *gig.MenuItem) {
				logtable.Get().Clear()
			}),
			gig.NewMenuItem("Exit").SetOnClick(func(mi *gig.MenuItem) {
				a.Close()
			}),
		),
	)
	a.AddChildren(mainMenu)

	// Empty right window which will be populated with selected demo
	rightWin := gig.NewWindow("##right")

	// Builds left window with demos categories and names
	leftWin := gig.NewWindow("left").
		SetFlags(imgui.WindowFlags_NoTitleBar).
		AddChildren(
			newCategory(catImGui, rightWin),
			newCategory(catImPlot, rightWin),
			newCategory(catWidgets, rightWin),
		)

	// Creates bottom window for log
	logWin := gig.NewWindow("Log").
		SetFlags(imgui.WindowFlags_NoTitleBar).
		AddChildren(logtable.Get())

	// Creates window splitter for the left and right windows
	hsplitter := gig.NewWindowHSplitter(leftWin, rightWin).
		SetRatio(0.2)

	// Creates window splitter for the top hsplitter and the bottom log window
	vsplitter := gig.NewWindowVSplitter(hsplitter, logWin).
		SetMaximize(true).
		SetRatio(0.8).
		SaveRef("splitter")
	a.AddChildren(vsplitter)

	// Add C++ demo windows (invisible)
	a.AddChildren(
		gig.NewImGuiDemoWindow().SaveRef("imguiDemo").SetVisible(false),
		gig.NewImPlotDemoWindow().SaveRef("implotDemo").SetVisible(false),
	)
}

// newCategory returns a collapsing header with all demos of the specified category
func newCategory(cat catType, win *gig.Window) *gig.CollapsingHeader {

	tree := gig.NewCollapsingHeader(string(cat))
	tree.SetFlags(imgui.TreeNodeFlags_DefaultOpen)
	demos := mapDemos[cat]
	for _, d := range demos {
		demo := d
		sel := gig.NewSelectable(demo.name)
		sel.SetOnClick(func(s *gig.Selectable) {
			win.ClearChildren()
			// TODO currently it is recreating the widgets. Cache the create widget ?
			win.AddChildren(demo.fn())
			win.SetTitle(demo.name)
		})
		tree.AddChildren(sel)
	}
	return tree
}

// registerDemo is used by demos to register themselves
func registerDemo(category catType, name string, f func() gig.IWidget) {

	mapDemos[category] = append(mapDemos[category], demoInfo{category: category, name: name, fn: f})
}
