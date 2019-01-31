// +build !wasm

package app

import (
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

// Handler is a http handler that serves UI components created with this
// package.
type Handler struct {
	// The app author.
	Author string

	// The app description.
	Description string

	// The app keywords.
	Keywords []string

	// The app name.
	Name string

	// The path of the go web assembly file to serve. Default is app.wasm.
	Wasm string

	// The function that returns the path of the web directory.
	WebDir func() string

	once        sync.Once
	webDir      string
	fileHandler http.Handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.once.Do(h.initFileHandler)

	path := filepath.Join(h.webDir, r.URL.Path)
	if _, err := os.Stat(path); err == nil {
		h.fileHandler.ServeHTTP(w, r)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func (h *Handler) initFileHandler() {
	h.webDir = h.WebDir()

	if h.Wasm == "" {
		h.Wasm = "app.wasm"
	}

	handler := http.FileServer(http.Dir(h.webDir))
	handler = newGzipHandler(handler)
	h.fileHandler = handler
}
