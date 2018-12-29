// Package murlok is a Go library to build gui apps.
// /!\ To be continued.
package murlok

import (
	"io"
	"net/http"
)

var (
	target     string
	defaultApp app
)

// Application is the interface that describes an app.
type Application interface {
	// Sets a custom servers.
	WithCustomServer(*http.Server) Application

	// Sets a package configuration.
	WithPackageConfig(PackageConfig) Application

	// Run starts the app.
	Run(url string, allowedHosts ...string) error
}

// PackageConfig is the interface that descibes a configuration to build a
// package.
type PackageConfig interface {
	// Dumps the configuration in the given writer.
	Dump(io.Writer)
}

// App returns the application.
func App() Application {
	return &defaultApp
}
