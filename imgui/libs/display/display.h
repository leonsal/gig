#pragma once

#ifdef __cplusplus
extern "C" {
#endif

// Pointer to Display implementation
typedef void* displayImpl;

// Common Display configuration
typedef struct display_config {
    const char* title;
    int         width;
    int         height;
    int         msaa;
    double      ev_timeout;
    bool        fullscreen;
    struct {
        bool es;
    } opengl;
} display_config_t;

// Common Display methods
displayImpl display_init(display_config_t* cfg, int* error);
void display_destroy(displayImpl di);
void display_upload_fonts(displayImpl di);
void display_size(displayImpl di, int* width, int* height);
void display_set_title(displayImpl di, const char* title);
bool display_start_frame(displayImpl di);
void display_end_frame(displayImpl di);
void display_update(displayImpl di);
void display_set_should_close(displayImpl di, bool close);
ImTextureID display_create_texture(displayImpl di);
void display_delete_texture(displayImpl di, ImTextureID tid);
void display_transfer_texture(displayImpl di, ImTextureID tid, int width, int height, const void* data);

#ifdef __cplusplus
}
#endif

