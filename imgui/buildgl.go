//go:build !vulkan

package imgui

// #cgo CFLAGS:   -I./libs -I./libs/cimgui -I./libs/cimgui/imgui -I./libs/cimplot -I./libs/display
// #cgo CXXFLAGS: -I./libs -I./libs/cimgui -I./libs/cimgui/imgui -I./libs/cimgui/imgui/backends -I./libs/cimplot -I./libs/cimplot/implot -I./libs/display
// #cgo linux   LDFLAGS: -L${SRCDIR}/libs -lcimguigl -lGL -lglfw3 -lm -ldl -lX11 -lstdc++
// #cgo windows LDFLAGS: -L./libs -lglfw -lstdc++ -lopengl32 -lgdi32 -limm32
import "C"
