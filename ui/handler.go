// +build !js

package ui

import (
	"fmt"
	"net/http"
)

// Handler is a http handler that serves UI components created with this
// package.
type Handler struct {
	// The name of the component to load when no path is specified.
	DefaultCompo string

	// The path of the web directory.
	WebDir string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("url:", r.URL)
	fmt.Println("path:", r.URL.Path)

	w.WriteHeader(http.StatusNotImplemented)
}
