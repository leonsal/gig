package main

import (
	"github.com/leonsal/gig"
)

func main() {

	// Creates view
	view, err := gig.ViewInit("Gig Demo", 800, 600)
	if err != nil {
		panic(err)
	}

	mainMenu := gig.NewMainMenuBar().AddChildren(
		gig.NewMenu("Menu1").AddChildren(
			gig.NewMenuItem("Option1").SetEnabled(false),
			gig.NewMenuItem("Option2"),
			gig.NewSeparator(),
		).SaveRef("menu1"),

		gig.NewMenu("Demos").AddChildren(

			gig.NewMenuItem("Text").SetOnClick(func(mi *gig.MenuItem) {
				win := gig.Ref("textDemo").(*gig.Window)
				vis := !win.Visible()
				win.SetVisible(vis)
				mi.SetSelected(vis)
			}),

			gig.NewMenuItem("Button").SetOnClick(func(mi *gig.MenuItem) {
				win := gig.Ref("buttonDemo").(*gig.Window)
				vis := !win.Visible()
				win.SetVisible(vis)
				mi.SetSelected(vis)
			}),

			gig.NewMenuItem("Input").SetOnClick(func(mi *gig.MenuItem) {
				win := gig.Ref("inputDemo").(*gig.Window)
				vis := !win.Visible()
				win.SetVisible(vis)
				mi.SetSelected(vis)
			}),

			gig.NewMenuItem("Drag").SetOnClick(func(mi *gig.MenuItem) {
				win := gig.Ref("dragDemo").(*gig.Window)
				vis := !win.Visible()
				win.SetVisible(vis)
				mi.SetSelected(vis)
			}),

			gig.NewMenuItem("Slider").SetOnClick(func(mi *gig.MenuItem) {
				win := gig.Ref("sliderDemo").(*gig.Window)
				vis := !win.Visible()
				win.SetVisible(vis)
				mi.SetSelected(vis)
			}),

			gig.NewMenuItem("List").SetOnClick(func(mi *gig.MenuItem) {
				win := gig.Ref("listDemo").(*gig.Window)
				vis := !win.Visible()
				win.SetVisible(vis)
				mi.SetSelected(vis)
			}),

			gig.NewMenuItem("Tab").SetOnClick(func(mi *gig.MenuItem) {
				win := gig.Ref("tabDemo").(*gig.Window)
				vis := !win.Visible()
				win.SetVisible(vis)
				mi.SetSelected(vis)
			}),

			gig.NewMenuItem("Tree").SetOnClick(func(mi *gig.MenuItem) {
				win := gig.Ref("treeDemo").(*gig.Window)
				vis := !win.Visible()
				win.SetVisible(vis)
				mi.SetSelected(vis)
			}),

			gig.NewMenuItem("Popup").SetOnClick(func(mi *gig.MenuItem) {
				win := gig.Ref("popupDemo").(*gig.Window)
				vis := !win.Visible()
				win.SetVisible(vis)
				mi.SetSelected(vis)
			}),

			gig.NewMenuItem("Table").SetOnClick(func(mi *gig.MenuItem) {
				win := gig.Ref("tableDemo").(*gig.Window)
				vis := !win.Visible()
				win.SetVisible(vis)
				mi.SetSelected(vis)
			}),

			gig.NewMenuItem("Plot").SetOnClick(func(mi *gig.MenuItem) {
				win := gig.Ref("plotDemo").(*gig.Window)
				vis := !win.Visible()
				win.SetVisible(vis)
				mi.SetSelected(vis)
			}),
		),

		gig.NewMenu("C++ demos").AddChildren(
			gig.NewMenuItem("ImGui Demo").SetOnClick(func(mi *gig.MenuItem) {
				win := gig.Ref("imguiDemoWindow").(*gig.ImGuiDemoWindow)
				vis := !win.Visible()
				win.SetVisible(vis)
				mi.SetSelected(vis)
			}),

			gig.NewMenuItem("ImGui Metrics").SetOnClick(func(mi *gig.MenuItem) {
				win := gig.Ref("imguiMetricsWindow").(*gig.ImGuiMetricsWindow)
				vis := !win.Visible()
				win.SetVisible(vis)
				mi.SetSelected(vis)
			}),

			gig.NewMenuItem("ImPlot Demo").SetOnClick(func(mi *gig.MenuItem) {
				win := gig.Ref("implotDemoWindow").(*gig.ImPlotDemoWindow)
				vis := !win.Visible()
				win.SetVisible(vis)
				mi.SetSelected(vis)
			}),
		),
	)

	// Gig demos
	view.AddChildren(mainMenu)
	view.AddChildren(NewTextDemo())
	view.AddChildren(NewButtonDemo())
	view.AddChildren(NewInputDemo())
	view.AddChildren(NewDragDemo())
	view.AddChildren(NewSliderDemo())
	view.AddChildren(NewListDemo())
	view.AddChildren(NewTabDemo())
	view.AddChildren(NewTreeDemo())
	view.AddChildren(NewPopupDemo())
	view.AddChildren(NewTableDemo())
	view.AddChildren(NewPlotDemo())

	// ImGui C++ demo windows
	view.AddChildren(gig.NewImGuiDemoWindow().
		SetVisible(false).
		SaveRef("imguiDemoWindow"),
	)
	view.AddChildren(gig.NewImGuiMetricsWindow().
		SetVisible(false).
		SaveRef("imguiMetricsWindow"),
	)
	view.AddChildren(gig.NewImPlotDemoWindow().
		SetVisible(false).
		SaveRef("implotDemoWindow"),
	)

	view.Run()
}
