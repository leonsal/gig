//go:build vulkan

package imgui

// #cgo CFLAGS:   -I./libs -I./libs/cimgui -I./libs/cimgui/imgui -I./libs/cimplot -I./libs/display
// #cgo CXXFLAGS: -I./libs -I./libs/cimgui -I./libs/cimgui/imgui -I./libs/cimgui/imgui/backends -I./libs/cimplot -I./libs/cimplot/implot -I./libs/display
// #cgo linux   LDFLAGS: -L/usr/local/lib -lgigvk -lglfw3 -lvulkan -lm -ldl -lX11 -lstdc++
// #cgo windows LDFLAGS: -LC:/giglibs -lgigvk -lglfw3 -lvulkan-1 -lstdc++ -lgdi32 -limm32
import "C"
