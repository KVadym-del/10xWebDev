MAKE ?= make

.PHONY: all clean build-wasm build-frontend build-server run

NPM ?= npm

all: build-wasm build-frontend build-server

build-wasm:
	$(MAKE) -C wasm

build-frontend:
	cd web 
	$(NPM) install
	$(NPM) run build

build-server:
	go build -o bin/server.exe cmd/server/main.go

clean:
	rm -rf bin/
	rm -rf web/dist/
	$(MAKE) -C wasm clean

run: all
	./bin/server.exe