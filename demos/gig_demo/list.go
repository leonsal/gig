package main

import (
	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
)

func init() {

	registerDemo(catImGui, "List", NewListDemo)
}

func NewListDemo() gig.IWidget {

	var list1 *gig.ListBox
	var list2 *gig.ListBox
	group := gig.NewGroup().AddChildren(

		gig.NewText("ListBox"), gig.NewSameLine(), gig.NewSpacing(), gig.NewSameLine(),
		gig.NewButton("Unselect").SetOnClick(func(b *gig.Button) {
			list1.UnselectAll()
		}),

		gig.NewListBox("List1").AddStrings(
			"Item1", "Item2", "Item3", "Item4", "Item5", "Item6",
		).SetOnChange(func(l *gig.ListBox) {
			log.Info("ListBox1 changed:", l.SelectedIndex())
		}).Save(&list1),

		gig.NewText("ListBox multi"), gig.NewSameLine(), gig.NewSpacing(), gig.NewSameLine(),
		gig.NewButton("Unselect").SetOnClick(func(b *gig.Button) {
			list2.UnselectAll()
		}),
		gig.NewListBox("List2").AddItems(
			gig.NewSelectable("Item1"),
			gig.NewSelectable("Item2"),
			gig.NewSelectable("Item3"),
			gig.NewSelectable("Item4"),
			gig.NewSelectable("Item5"),
			gig.NewSelectable("Item6"),
		).
			Save(&list2).
			SetMulti(true).
			SetOnChange(func(l *gig.ListBox) {
				log.Info("ListBox2 changed:", l.SelectedIndex())
			}).
			SetPreRender(func(w gig.IWidget) {
				w.(*gig.ListBox).SetSize(imgui.Vec2{-1, 5 * imgui.GetTextLineHeightWithSpacing()})
			}),

		gig.NewText("ComboBox"),
		gig.NewCombo("Combo1").AddChildren(
			gig.NewSelectable("Item1"),
			gig.NewSelectable("Item2"),
			gig.NewSelectable("Item3"),
			gig.NewSelectable("Item4"),
			gig.NewSelectable("Item5"),
			gig.NewSelectable("Item6"),
		).SetSelectedIndex(2).SetOnChange(func(c *gig.Combo) {
			log.Info("Combo1 changed:", c.SelectedIndex())
		}),

		gig.NewCombo("").AddStrings(
			"Item1", "Item2", "Item3", "Item4", "Item5", "Item6",
		).
			SetSelectedIndex(0).
			SetOnChange(func(c *gig.Combo) {
				log.Info("Combo2 changed:", c.SelectedIndex())
			}).
			SetPreRender(func(w gig.IWidget) {
				width := imgui.GetWindowContentRegionMax().X - imgui.GetWindowContentRegionMin().X
				w.(*gig.Combo).SetItemWidth(width)
			}),
	)
	return group
}
