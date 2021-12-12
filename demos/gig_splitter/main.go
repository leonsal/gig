package main

import (
	"github.com/leonsal/gig"
)

func main() {

	// Creates App
	view, err := gig.ViewInit("Gig Splitter", 800, 600)
	if err != nil {
		panic(err)
	}

	hs := gig.NewWindowHSplitter(
		gig.NewWindowVSplitter(
			gig.NewWindow("Window 1").AddChildren(
				gig.NewChildHSplitter(
					gig.NewChildWindow().AddChildren(gig.NewText("Child 1.1")),
					gig.NewChildVSplitter(
						gig.NewChildWindow().AddChildren(gig.NewText("Child 1.2")),
						gig.NewChildWindow().AddChildren(gig.NewText("Child 1.3")),
					),
				),
			),
			gig.NewWindowHSplitter(
				gig.NewWindow("Window 1.1").AddChildren(
					gig.NewChildHSplitter(
						gig.NewChildWindow().AddChildren(gig.NewText("Child 1.1.1")),
						gig.NewChildWindow().AddChildren(gig.NewText("Child 1.1.2")),
					),
				),
				gig.NewWindow("Window 1.2").AddChildren(
					gig.NewChildVSplitter(
						gig.NewChildWindow().AddChildren(gig.NewText("Child 1.2.1")),
						gig.NewChildWindow().AddChildren(gig.NewText("Child 1.2.2")),
					),
				),
			),
		),
		gig.NewWindowHSplitter(
			gig.NewWindow("Window 2").AddChildren(
				gig.NewChildVSplitter(
					gig.NewChildWindow().AddChildren(gig.NewText("Child 2.1")),
					gig.NewChildWindow().AddChildren(gig.NewText("Child 2.2")),
				),
			),
			gig.NewWindowVSplitter(
				gig.NewWindow("Window 2.1").AddChildren(
					gig.NewChildHSplitter(
						gig.NewChildWindow().AddChildren(gig.NewText("Child 2.1.1")),
						gig.NewChildWindow().AddChildren(gig.NewText("Child 2.1.2")),
					),
				),
				gig.NewWindow("Window 2.2").AddChildren(
					gig.NewChildVSplitter(
						gig.NewChildWindow().AddChildren(gig.NewText("Child 2.2.1")),
						gig.NewChildWindow().AddChildren(gig.NewText("Child 2.2.2")),
					),
				),
			),
		),
	).SetMaximize(true)

	view.AddChildren(hs)
	view.Run()
}
