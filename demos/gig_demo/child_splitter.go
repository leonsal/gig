package main

import (
	"github.com/leonsal/gig"
)

func init() {

	registerDemo(catWidgets, "ChildSplitter", NewChildSplitterDemo)
}

func NewChildSplitterDemo() gig.IWidget {

	group := gig.NewGroup().AddChildren(

		gig.NewChildHSplitter(
			gig.NewChildWindow().AddChildren(gig.NewText("Child Window 1")),
			gig.NewChildVSplitter(
				gig.NewChildWindow().AddChildren(gig.NewText("Child Window 2")),
				gig.NewChildHSplitter(
					gig.NewChildWindow().AddChildren(gig.NewText("Child Window 2.1")),
					gig.NewChildWindow().AddChildren(gig.NewText("Child Window 2.2")),
				),
			),
		),
	)
	return group
}
