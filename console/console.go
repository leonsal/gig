package console

import (
	"sync"

	"github.com/leonsal/gig"
	"github.com/leonsal/gig/imgui"
)

type consoleLines struct {
	gig.ChildWindow
	scrollBottom bool
}

func newConsoleLines() *consoleLines {

	l := new(consoleLines)
	l.Init(l)
	l.SetRender(func() {
		if imgui.BeginChildID(imgui.ID(l.ID()), l.Size(), l.Border(), l.Flags()) {
			l.RenderChildren()
			if l.scrollBottom {
				imgui.SetScrollHereY(1.0)
				l.scrollBottom = false
			}
			imgui.EndChild()
		}
	})
	return l
}

type Console struct {
	gig.Widget
	lines      *consoleLines
	input      *gig.InputText
	changed    bool
	onenter    func(*Console)
	history    []string
	histPos    int
	maxHistory int
	mut        sync.Mutex
}

func New() *Console {

	c := new(Console)
	c.Init(c)

	// Creates child window which contains the console lines
	c.lines = newConsoleLines()
	c.lines.SetFlags(imgui.WindowFlags_HorizontalScrollbar)

	// Creates input text for entering line
	sep := gig.NewSeparator()
	c.input = gig.NewInputText("").SetOnChange(func(ti *gig.InputText) {
		if c.onenter != nil {
			c.changed = true
			c.onenter(c)
		}
	}).SetMaxLen(256)

	// Sets input text callback for processing history
	c.input.SetFlags(imgui.InputTextFlags_EnterReturnsTrue | imgui.InputTextFlags_CallbackHistory)
	c.input.SetCallback(func(cd imgui.InputTextCallbackData) int {
		return c.cbHistory(cd)
	})
	c.SetMaxHistory(50)
	c.histPos = -1

	c.SetRender(func() {

		// Render the child window with console lines
		// Reserve enough left-over height for 1 separator + 1 input text
		reserveHeight := imgui.GetStyle().GetItemSpacing().Y + imgui.GetFrameHeightWithSpacing()
		c.lines.SetSize(imgui.Vec2{0, -reserveHeight})
		c.lines.Render()

		// Render the separator and input widget
		sep.Render()
		width := imgui.GetWindowContentRegionMax().X - imgui.GetWindowContentRegionMin().X
		c.input.SetItemWidth(width)
		c.input.Render()

		// Set keyboard focus if needed
		imgui.SetItemDefaultFocus()
		if c.changed {
			imgui.SetKeyboardFocusHere(-1)
			c.changed = false
		}
	})
	return c
}

func (c *Console) Save(pc **Console) *Console {

	*pc = c
	return c
}

func (c *Console) Clear() {

	c.mut.Lock()
	defer c.mut.Unlock()
	c.lines.ClearChildren()
}

func (c *Console) AddLine(line string, color imgui.Vec4) *Console {

	c.mut.Lock()
	defer c.mut.Unlock()
	c.lines.AddChildren(gig.NewTextColored(color, line))
	c.lines.scrollBottom = true
	return c
}

func (c *Console) AddLines(lines []string, color imgui.Vec4) *Console {

	c.mut.Lock()
	defer c.mut.Unlock()
	for _, line := range lines {
		c.lines.AddChildren(gig.NewTextColored(color, line))
	}
	return c
}

// AddHistory adds the specified line to the lines history
func (c *Console) AddHistory(line string) *Console {

	if len(c.history) >= c.maxHistory {
		copy(c.history, c.history[1:])
		c.history = c.history[:len(c.history)-1]
	}
	c.history = append(c.history, line)
	c.histPos = len(c.history) - 1
	return c
}

// Clear history clears the lines history
func (c *Console) ClearHistory() *Console {

	c.history = c.history[:0]
	c.histPos = -1
	return c
}

// MaxHistory returns the maximum number of history lines
func (c *Console) MaxHistory() int {

	return c.maxHistory
}

// SetMaxHistory sets the maximum number of history lines
func (c *Console) SetMaxHistory(lines int) *Console {

	c.maxHistory = lines
	return c
}

// Input returns the InputText widget of the Console
func (c *Console) Input() *gig.InputText {

	return c.input
}

func (c *Console) SetOnEnter(f func(c *Console)) *Console {

	c.onenter = f
	return c
}

func (c *Console) cbHistory(cd imgui.InputTextCallbackData) int {

	buflen := cd.BufTextLen()
	if cd.EventKey() == imgui.Key_UpArrow {
		if c.histPos >= 0 {
			cd.DeleteChars(0, buflen)
			cd.InsertChars(0, c.history[c.histPos])
			if c.histPos > 0 {
				c.histPos--
			}
		}
		return 0
	}
	if cd.EventKey() == imgui.Key_DownArrow {
		if c.histPos+1 < len(c.history) {
			c.histPos++
			cd.DeleteChars(0, buflen)
			cd.InsertChars(0, c.history[c.histPos])
		}
	}
	return 0
}
