#define NOB_IMPLEMENTATION
#include "nob.h"

int main(int argc, char **argv)
{
    NOB_GO_REBUILD_URSELF(argc, argv);

    const char* main_dir = nob_get_current_dir_temp();

    Nob_Cmd cmd = {0};
    const char *c_wasm_path = nob_temp_sprintf("%s/wasm", main_dir);
    nob_set_current_dir(c_wasm_path);

    nob_cmd_append(&cmd, "clang", "-O3", "--target=wasm32", "--no-standard-libraries", "-Wl,--export-all",
                   "-Wl,--no-entry", "-o", "calc.wasm", "calc.c");
    if (!nob_cmd_run_sync(cmd))
        return 1;

    const char* web_path = nob_temp_sprintf("%s/web", main_dir);
    nob_set_current_dir(web_path);

    cmd = (Nob_Cmd){0};
    nob_cmd_append(&cmd, "npm", "install");
    if (!nob_cmd_run_sync(cmd))
        return 1;

    cmd = (Nob_Cmd){0};
    nob_cmd_append(&cmd, "npm", "run", "build");
    if (!nob_cmd_run_sync(cmd))
        return 1;

    nob_set_current_dir(main_dir);
    cmd = (Nob_Cmd){0};
    nob_cmd_append(&cmd, "go", "build", "-o", "bin/server", "cmd/server/main.go");
    if (!nob_cmd_run_sync(cmd))
        return 1;
        
    return 0;
}