// +build !js

package ui

import (
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

// Handler is a http handler that serves UI components created with this
// package.
type Handler struct {
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

	handler := http.FileServer(http.Dir(h.webDir))
	handler = newGzipHandler(handler)
	h.fileHandler = handler
}
