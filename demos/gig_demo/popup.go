package main

import (
	"fmt"

	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
)

func init() {

	registerDemo(catImGui, "Popup", NewPopupDemo)
}

func NewPopupDemo() gig.IWidget {

	var popup *gig.Popup
	var selText *gig.Text
	var popupModal *gig.PopupModal
	group := gig.NewGroup().AddChildren(

		gig.NewButton("Selected:").
			SetOnClick(func(b *gig.Button) {
				popup.Open()
			}).
			SetTooltip("Click on the button to show popup"),
		gig.NewSameLine(),
		gig.NewText("?").Save(&selText),
		gig.NewPopup().
			Save(&popup).
			AddChildren(
				gig.NewSelectable("item1").
					SetOnClick(func(s *gig.Selectable) {
						selText.SetText("item1")
					}),
				gig.NewSelectable("item2").
					SetOnClick(func(s *gig.Selectable) {
						selText.SetText("item2")
					}),
				gig.NewSelectable("item3").
					SetOnClick(func(s *gig.Selectable) {
						selText.SetText("item3")
					}),
			),

		gig.NewSeparator(),
		gig.NewButton("Open Popup modal").
			SetOnClick(func(b *gig.Button) {
				popupModal.Open()
			}),
		gig.NewPopupModal("Popup modal").
			Save(&popupModal).
			AddChildren(
				gig.NewText("popup modal contents"),
				gig.NewButton("OK").SetOnClick(func(b *gig.Button) {
					fmt.Println("popup modal OK")
				}),
				gig.NewSameLine(),
				gig.NewButton("Cancel").SetOnClick(func(b *gig.Button) {
					fmt.Println("popup modal Cancel")
				}),
				gig.NewSameLine(),
				gig.NewButton("Close").SetOnClick(func(b *gig.Button) {
					fmt.Println("popup modal Close")
					popupModal.Close()
				}),
			),

		gig.NewSeparator(),
		gig.NewButton("Button with context").
			SetOnContext(imgui.HoveredFlags_None, func(w *gig.Widget) { popup.Open() }).
			SetTooltip("Right click to show context popup"),
	)
	return group
}
