package main

import (
	"github.com/leonsal/gig"
)

func main() {

	// Creates view
	view, err := gig.ViewInit("Gig Hello", 800, 600)
	if err != nil {
		panic(err)
	}

	view.AddChildren(
		gig.NewWindow("Hello").AddChildren(
			gig.NewText("Hello"),
		),
	)

	view.Run()
}
