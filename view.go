package gig

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/leonsal/gig/imgui"
)

type View struct {
	disp     imgui.Display
	group    *Group
	run      bool
	memStats runtime.MemStats
}

// Command line options
var (
	opLogMemallocs = flag.Bool("log_memallocs", false, "Logs memory allocations to the console")
	opOpengles     = flag.Bool("opengles", false, "Use OpenGL ES instead of OpenGL")
)

// Singleton application
var defaultView *View

func ViewInit(title string, width, height int) (*View, error) {

	if defaultView != nil {
		return nil, fmt.Errorf("View already initialized")
	}

	// All GUI functions must run in the same thread
	runtime.LockOSThread()

	flag.Parse()

	// Creates the default display window
	cfg := imgui.DisplayConfig{
		Evtimeout:  0,
		Fullscreen: false,
		Opengl:     imgui.OpenglConfig{ES: *opOpengles},
	}
	disp, err := imgui.DisplayInit(title, width, height, &cfg)
	if err != nil {
		return nil, err
	}

	v := new(View)
	v.disp = disp
	v.group = NewGroup()
	v.run = true
	defaultView = v
	return defaultView, nil
}

func GetView() *View {

	return defaultView
}

func (v *View) AddChildren(c ...IWidget) {

	v.group.AddChildren(c...)
}

func (v *View) Display() imgui.Display {

	return v.disp
}

func (v *View) Run() {

	if !v.run {
		return
	}
	if *opLogMemallocs {
		runtime.ReadMemStats(&v.memStats)
	}

	for v.run {
		close := v.disp.StartFrame()
		if close {
			break
		}
		v.group.Render()
		v.disp.EndFrame()

		if *opLogMemallocs {
			var ms runtime.MemStats
			runtime.ReadMemStats(&ms)
			//log.Printf("Mallocs:%v Frees:%v Live:%v TotalAlloc: %v\n", ms.Mallocs, ms.Frees, ms.Mallocs-ms.Frees, ms.TotalAlloc)
			if ms.Mallocs != v.memStats.Mallocs {
				print(".")
				v.memStats = ms
			}
		}
	}
	v.disp.Destroy()
	runtime.UnlockOSThread()
}

func (v *View) Close() {

	v.run = false
}
