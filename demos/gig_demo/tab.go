package main

import (
	"fmt"

	"github.com/leonsal/gig"
)

func init() {

	registerDemo(catImGui, "Tab", NewTabDemo)
}

func NewTabDemo() gig.IWidget {

	var tabBar *gig.TabBar
	nextTab := 2
	group := gig.NewGroup().AddChildren(
		gig.NewTabBar().Save(&tabBar).AddChildren(

			gig.NewTabItemButton("Add tab").SetOnClick(func(b *gig.TabItemButton) {
				item := gig.NewTabItem(fmt.Sprintf("Tab%d", nextTab)).
					SetCanClose(true).
					AddChildren(gig.NewText(fmt.Sprintf("Contents of Tab%d", nextTab)))
				tabBar.AddChildren(item)
				nextTab++
			}),

			gig.NewTabItemButton("Show closed").SetOnClick(func(b *gig.TabItemButton) {
				items := tabBar.Children()
				for _, tiw := range items {
					tiw.SetVisible(true)
				}
			}),

			gig.NewTabItem("Tab1").SetCanClose(true).AddChildren(
				gig.NewText("Contents of Tab1"),
			),
		),
	)
	return group
}
