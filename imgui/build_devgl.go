//go:build devgl

package imgui

// #cgo CFLAGS:   -I./libs -I./libs/cimgui -I./libs/cimgui/imgui -I./libs/cimplot -I./libs/display
// #cgo CXXFLAGS: -I./libs -I./libs/cimgui -I./libs/cimgui/imgui -I./libs/cimgui/imgui/backends -I./libs/cimplot -I./libs/cimplot/implot -I./libs/display
// #cgo linux   LDFLAGS: -L/usr/lib/local -L/usr/lib -L./libs -lcimguigl -lglfw -lGL -lm -ldl -lX11 -lstdc++
// #cgo windows LDFLAGS: -L./libs -lcimguigl -lglfw -lstdc++ -lopengl32 -lgdi32 -limm32
import "C"
