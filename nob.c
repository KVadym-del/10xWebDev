#include <Windows.h>

#define NOB_IMPLEMENTATION
#include "nob.h"

#include "string.h"

int main(int argc, char **argv)
{
    NOB_GO_REBUILD_URSELF(argc, argv);
    Nob_Cmd cmd = {0};

    const char *c_wasm_path = nob_temp_sprintf("%s\\wasm", nob_get_current_dir_temp());
    nob_set_current_dir(c_wasm_path);
    printf("Current path: %s\n", nob_get_current_dir_temp());

    nob_cmd_append(&cmd, "clang", "-O3", "--target=wasm32", "--no-standard-libraries", "-Wl,--export-all",
                   "-Wl,--no-entry", "-o", "calc.wasm", "calc.c");
    if (!nob_cmd_run_sync(cmd))
        return 1;
    return 0;
}