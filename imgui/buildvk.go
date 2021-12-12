//go:build vulkan

package imgui

// #cgo CFLAGS:   -I./libs -I./libs/cimgui -I./libs/cimgui/imgui -I./libs/cimplot -I./libs/display
// #cgo CXXFLAGS: -I./libs -I./libs/cimgui -I./libs/cimgui/imgui -I./libs/cimgui/imgui/backends -I./libs/cimplot -I./libs/cimplot/implot -I./libs/display
// #cgo linux   LDFLAGS: -L./libs -lvulkan -lglfw3 -lm -ldl -lX11 -lstdc++
// #cgo windows LDFLAGS: -L./libs -L${SRCDIR}/libs/vulkan/lib -lglfw3 -lvulkan-1 -lstdc++ -lgdi32 -limm32
import "C"
