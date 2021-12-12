cmake_minimum_required(VERSION 3.0 FATAL_ERROR)

project(libcimgui)

set(CMAKE_CXX_STANDARD 11)
set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${CMAKE_SOURCE_DIR})
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

include_directories(./)
include_directories(cimgui)
include_directories(cimgui/imgui)
include_directories(cimgui/imgui/backends)
include_directories(cimplot/implot)
include_directories(glfw-3.3.5/include)

set(CIMGUI_SOURCES
    cimgui/cimgui.cpp
    cimgui/imgui/imgui.cpp 
    cimgui/imgui/imgui_demo.cpp
    cimgui/imgui/imgui_draw.cpp
    cimgui/imgui/imgui_tables.cpp
    cimgui/imgui/imgui_widgets.cpp
)
set(CIMPLOT_SOURCES
    cimplot/cimplot.cpp
    cimplot/implot/implot.cpp 
    cimplot/implot/implot_demo.cpp
    cimplot/implot/implot_items.cpp
)

# GLFW dependency subproject
set(GLFW_BUILD_DOCS OFF CACHE BOOL "" FORCE)
set(GLFW_BUILD_TESTS OFF CACHE BOOL "" FORCE)
set(GLFW_BUILD_EXAMPLES OFF CACHE BOOL "" FORCE)
add_subdirectory(glfw-3.3.5 glfw)

set(DISPLAYGL_SOURCES
    display/displaygl.cpp
    cimgui/imgui/backends/imgui_impl_opengl3.cpp
    cimgui/imgui/backends/imgui_impl_glfw.cpp
)

set(DISPLAYVK_SOURCES
    display/displayvk.cpp
    cimgui/imgui/backends/imgui_impl_vulkan.cpp
    cimgui/imgui/backends/imgui_impl_glfw.cpp
)

# For windows use local copy of vulkan headers from SDK
if (WIN32) 
    include_directories(vulkan)
endif()

add_library(cimguigl STATIC ${CIMGUI_SOURCES} ${CIMPLOT_SOURCES} ${DISPLAYGL_SOURCES})
add_library(cimguivk STATIC ${CIMGUI_SOURCES} ${CIMPLOT_SOURCES} ${DISPLAYVK_SOURCES})

install(
    TARGETS cimguigl cimguivk
)
