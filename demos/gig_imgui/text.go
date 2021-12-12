package main

import (
	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
)

func NewTextDemo() gig.IWidget {

	win := gig.NewWindow("Text").AddChildren(
		gig.NewText("Simple text"),
		gig.NewTextColored(imgui.ColorName("Red"), "Colored text"),
		gig.NewTextDisabled("Disabled text"),
		gig.NewTextWrapped("Wrapped text: Long text line which will be wrapped at the window borders"),
		gig.NewLabelText("Label text 1", "text 1"),
		gig.NewLabelText("Label text 2", "text 2"),
		gig.NewBulletText("Bullet text 1"),
		gig.NewBulletText("Bullet text 2"),
	).SetVisible(false).SaveRef("textDemo")

	return win
}
