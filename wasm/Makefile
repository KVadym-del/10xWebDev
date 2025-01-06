.PHONY: all clean

C ?= clang

GENERIC_WASM_FLAGS ?= -O3 --target=wasm32 --no-standard-libraries -Wl,--export-all -Wl,--no-entry

all: calc.wasm

calc.wasm: calc.c calc.h
	$(C) $(GENERIC_WASM_FLAGS) -o calc.wasm calc.c 

clean:
	rm -f calc.wasm