#include <stdlib.h>
#include <stdio.h>
#include "imgui_impl_glfw.h"
#include "imgui_impl_opengl3.h"
#include <GLFW/glfw3.h>
#include "imgui.h"
#include "implot.h"
#include "display.h"

typedef struct displayGL {
    GLFWwindow* win;
    double      ev_timeout;
    ImVec4      clear_color;
} displayGL;

// Default Display
static displayGL* displayDef;

static void glfw_error_callback(int error, const char* description) {
    fprintf(stderr, "Glfw Error %d: %s\n", error, description);
}

// Initializes the default display and return pointer to its state
displayImpl display_init(display_config_t* cfg, int* error) {

    if (displayDef != nullptr) {
        *error = 1;
        return nullptr;
    }

    // Initialize GLFW
    glfwSetErrorCallback(glfw_error_callback);
    if (!glfwInit()) {
        fprintf(stderr, "Error initializing GLFW");
        *error = 2;
        return nullptr;
    }


    // Determines the api version from configuration
    int clientApi;
    const char* gGlVersion;
    int vmajor;
    int vminor;
    if (cfg->opengl.es) {
        clientApi = GLFW_OPENGL_ES_API;
        gGlVersion ="#version 300 es";
        vmajor = 3;
        vminor = 1;
    } else {
        clientApi = GLFW_OPENGL_API;
        gGlVersion ="#version 330";
        vmajor = 3;
        vminor = 3;
    }

    printf("api:%d v:%s major:%d, minor:%d\n", clientApi, gGlVersion, vmajor, vminor);
    // Set GLFW hints
    glfwWindowHint(GLFW_CLIENT_API, clientApi);
    glfwWindowHint(GLFW_CONTEXT_VERSION_MAJOR, vmajor);
    glfwWindowHint(GLFW_CONTEXT_VERSION_MINOR, vminor);
    if (!cfg->opengl.es) {
        glfwWindowHint(GLFW_OPENGL_PROFILE, GLFW_OPENGL_CORE_PROFILE);  // 3.2+ only
    }

    if (cfg->msaa > 0 && cfg->msaa <= 16) {
        glfwWindowHint(GLFW_SAMPLES, cfg->msaa);
    }

    // Create window
    GLFWmonitor* monitor = nullptr;
    if (cfg->fullscreen) {
        monitor = glfwGetPrimaryMonitor();
    }
    auto win = glfwCreateWindow(cfg->width, cfg->height, cfg->title, monitor, nullptr);
    if (win == nullptr) {
        fprintf(stderr, "Error creating GLFW Window");
        *error = 3;
        return nullptr;
    }

    glfwMakeContextCurrent(win);
    glfwSwapInterval(1); // Enable vsync

    // Setup Dear ImGui context
    IMGUI_CHECKVERSION();
    ImGui::CreateContext();
    ImGuiIO& io = ImGui::GetIO();
    io.ConfigFlags |= ImGuiConfigFlags_NavEnableKeyboard;  // Enable Keyboard Controls

    // Setup ImPlot context
    ImPlot::CreateContext();

    // Setup Dear ImGui style
    ImGui::StyleColorsDark();

    // Setup Platform/Renderer bindings
    ImGui_ImplGlfw_InitForOpenGL(win, true);
    ImGui_ImplOpenGL3_Init(gGlVersion);

    // Creates display object
    auto disp = new(displayGL);
    disp->win = win;
    disp->clear_color = ImVec4{0.45f, 0.55f, 0.60f, 1.00f};

    if (cfg->ev_timeout >= 0 && cfg->ev_timeout <= 1.0) {
        disp->ev_timeout = cfg->ev_timeout;
    }
    displayDef = disp;
    return disp;
}

void display_destroy(displayImpl di) {

    auto d = reinterpret_cast<displayGL*>(di);
    ImGui_ImplOpenGL3_Shutdown();
    ImGui_ImplGlfw_Shutdown();
    ImGui::DestroyContext();
    ImPlot::DestroyContext();
    glfwDestroyWindow(d->win);

    // TODO: only for the last display
    glfwTerminate();
    delete d;
}

void display_upload_fonts(displayImpl di) {

}

// Returns the current size of the display window
void display_size(displayImpl di, int* width, int* height) {

    auto d = reinterpret_cast<displayGL*>(di);
    glfwGetWindowSize(d->win, width, height);
}

// Sets the display window title
void display_set_title(displayImpl di, const char* title) {

    auto d = reinterpret_cast<displayGL*>(di);
    glfwSetWindowTitle(d->win, title);
}

// Starts the frame or returns true if the window should be closed
bool display_start_frame(displayImpl di) {

    // Checks if user requested window close
    auto d = reinterpret_cast<displayGL*>(di);
    if (glfwWindowShouldClose(d->win)) {
        return true;
    }

    // Poll and handle events (inputs, window resize, etc.)
    // Blocks if no events for the specified timeout
    glfwWaitEventsTimeout(d->ev_timeout);

    // Start the Dear ImGui frame
    ImGui_ImplOpenGL3_NewFrame();
    ImGui_ImplGlfw_NewFrame();
    ImGui::NewFrame();
    return false;
}

// Ends rendering the frame
void display_end_frame(displayImpl di) {

    auto d = reinterpret_cast<displayGL*>(di);
    ImGui::Render();
    int display_w, display_h;
    glfwGetFramebufferSize(d->win, &display_w, &display_h);
    glViewport(0, 0, display_w, display_h);
    glClearColor(d->clear_color.x, d->clear_color.y, d->clear_color.z, d->clear_color.w);
    glClear(GL_COLOR_BUFFER_BIT);
    ImGui_ImplOpenGL3_RenderDrawData(ImGui::GetDrawData());
    glfwSwapBuffers(d->win);
}

void display_update(displayImpl di) {

    glfwPostEmptyEvent();
}

void display_set_should_close(displayImpl di, bool close) {

    auto d = reinterpret_cast<displayGL*>(di);
    glfwSetWindowShouldClose(d->win, close);
}

// Creates and returns an OpenGL texture idenfifier
ImTextureID display_create_texture(displayImpl di) {

    // Create a OpenGL texture identifier
    GLuint image_texture;
    glGenTextures(1, &image_texture);
    glBindTexture(GL_TEXTURE_2D, image_texture);

    // Setup filtering parameters for display
    glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR);
    glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_LINEAR);

    return (void*)(intptr_t)image_texture;
}

// Deletes previously created texture
void display_delete_texture(displayImpl di, ImTextureID tid) {

    GLuint tex = reinterpret_cast<intptr_t>(tid);
    glDeleteTextures(1, &tex); 
}

// Transfer data for the specified texture
void display_transfer_texture(displayImpl di, ImTextureID tid, int width, int height, const void* data) {

    GLuint tex = reinterpret_cast<intptr_t>(tid);
    glBindTexture(GL_TEXTURE_2D, tex);
    glPixelStorei(GL_UNPACK_ROW_LENGTH, 0);
    glTexImage2D(GL_TEXTURE_2D, 0, GL_RGBA, width, height, 0, GL_RGBA, GL_UNSIGNED_BYTE, data);
}


