package handler

import (
	"net/http"
	"path/filepath"
)

type Handler struct {
	*http.ServeMux
}

func New() *Handler {
	mux := http.NewServeMux()
	h := &Handler{mux}

	fs := http.FileServer(http.Dir("web/public"))
	mux.Handle("/", fs)

	mux.HandleFunc("/wasm/calc.wasm", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/wasm")
		http.ServeFile(w, r, filepath.Join("wasm", "calc.wasm"))
	})

	mux.HandleFunc("/web/dist/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("web", "dist", "bundle.js"))
	})

	return h
}
