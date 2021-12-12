package main

import (
	"fmt"

	"github.com/leonsal/gig"
)

func NewTreeDemo() gig.IWidget {

	var header *gig.CollapsingHeader
	win := gig.NewWindow("Tree").AddChildren(

		gig.NewTreeNode("TreeNode1").
			AddChildren(
				gig.NewText("item1.1"),
				gig.NewTreeNode("TreeNode1.1").
					AddChildren(
						gig.NewTreeNode("TreeNode1.1.1").
							AddChildren(
								gig.NewText("item1.1.1"),
							),
						gig.NewTreeNode("TreeNode1.1.2").
							AddChildren(
								gig.NewText("item1.1.2"),
							),
					),
				gig.NewText("item1.2"),
			),

		gig.NewTreeNode("TreeNode2").
			AddChildren(
				gig.NewTreeNode("TreeNode2.1").
					AddChildren(
						gig.NewTreeNode("TreeNode2.1.1").
							AddChildren(
								gig.NewText("item2.2.1"),
							),
						gig.NewTreeNode("TreeNode2.1.2").
							AddChildren(
								gig.NewText("item2.1.2"),
							),
						gig.NewCheckbox("check2"),
					),
				gig.NewTreeNode("TreeNode2.2").
					AddChildren(
						gig.NewTreeNode("TreeNode2.2.1").
							AddChildren(
								gig.NewText("item2.1.1"),
							),
						gig.NewTreeNode("TreeNode2.2.2").
							AddChildren(
								gig.NewText("item2.2.2"),
							),
					),
			),

		gig.NewCollapsingHeader("").
			SetLabel("CollapsingHeader1").
			Call(func(hi gig.IWidget) {
				h := hi.(*gig.CollapsingHeader)
				for i := 0; i < 10; i++ {
					h.AddChildren(gig.NewText(fmt.Sprintf("item1.%d\n", i)))
				}
			}),
		gig.NewButton("Show closed header").
			SetOnClick(func(b *gig.Button) {
				header.SetVisible(true)
			}),
		gig.NewCollapsingHeader("").
			SetLabel("CollapsingHeader2 (can close)").
			SetCanClose(true).
			Save(&header).
			Call(func(hi gig.IWidget) {
				h := hi.(*gig.CollapsingHeader)
				for i := 0; i < 5; i++ {
					h.AddChildren(gig.NewText(fmt.Sprintf("item2.%d\n", i)))
				}
			}),
	).SetVisible(false).SaveRef("treeDemo")
	return win
}
