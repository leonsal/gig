## Gig

This Go module implements a binding for [ImGui](https://github.com/ocornut/imgui) and
[ImPlot](https://github.com/epezent/implot) C++ libraries, using [GLFW](https://www.glfw.org) 
for creating native platform windows and supporting OpenGL3 and Vulkan backends.
It currently needs Go 1.17+ and builds on Linux and Windows 10.

The module is composed of Go packages and C/C++ libraries.
To build it is necessary to have `gcc`, `g++`, `make`, `CMake`
and the following required tools and library dependencies installed:

## Debian/Ubuntu dependencies

```
$ sudo apt-get install xorg-dev libgl1-mesa-dev libvulkan-dev
```

## Windows 10 dependencies

We tested the Windows 10 build using a recent version of the [TDM-GCC](https://jmeubank.github.io/tdm-gcc/download/) toolchain.
After installing the toolchain, create a copy of `mingw32-make.exe` and rename it to `make.exe` in the same `\bin` folder.
Also you need to install [git](https://git-scm.com/downloads) and [cmake](https://cmake.org/download/). 

## Building

The following commands download the repository, builds the C/C++ libraries
and install them. In Linux the libraries are installed in `/usr/local/lib`
and in Window in `C:\giglibs`.

```
>git clone https://github.com/leonsal/gig
>cd gig
>make   # precede with 'sudo' in Linux as the libraries are installed in /usr/local/lib

```

## Usage

The following application is a simple example of how to use GIG.
It first initializes a `Display` (native window) and then executes the render loop.
Each iteration of this loop should begin with `StartFrame()` and terminate with `EndFrame()`.
Between these two calls you can draw your user interface using ImGui and ImPlot functions.

```go
package main

import (
	"github.com/leonsal/gig/imgui"
)

func main() {

	// Creates the default display window
	cfg := imgui.DisplayConfig{
		Title:      "ImGui Hello",
		Width:      600,
		Height:     400,
		MSAA:       0,
		Fullscreen: false,
	}
	disp, err := imgui.DisplayInit(&cfg)
	if err != nil {
		panic(err)
	}

	// Render loop
	open := true
	for {
		close := disp.StartFrame()
		if close {
			break
		}

		if open {
			if imgui.Begin("Window", &open, imgui.WindowFlags_None) {
				imgui.Text("Hello, ImGui")
			}
			imgui.End()
		}

		disp.EndFrame()
	}
	disp.Destroy()
}

```

To build this example:

```
>cd gig/demos/imgui_hello
>go install 
```

The default build will use the `OpenGL3` backend. To use `Vulkan`:

```
>go install -tags vulkan
```

