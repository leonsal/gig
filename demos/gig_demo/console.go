package main

import (
	"strings"

	"github.com/leonsal/gig"
	"github.com/leonsal/gig/console"
	"github.com/leonsal/gig/imgui"
)

func init() {

	registerDemo(catWidgets, "Console", NewConsoleDemo)
}

func NewConsoleDemo() gig.IWidget {

	help := []string{
		"'help' shows this help",
		"'clear' clears the console",
		"Up Arrow loads the previous history line",
		"Down Arrow loads the next history line",
	}
	color := imgui.ColorName("Gray")
	group := gig.NewGroup().AddChildren(
		console.New().AddLines(help, color).SetOnEnter(func(c *console.Console) {
			line := c.Input().Text()
			c.AddLine(line, color)
			c.AddHistory(line)
			c.Input().SetText("")
			cmd := strings.Trim(line, "")
			if cmd == "clear" {
				c.Clear()
			} else if cmd == "help" {
				c.AddLines(help, color)
			}
		}),
	)
	return group
}
