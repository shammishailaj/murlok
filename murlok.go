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

	// Gives a background color to views.
	WithBackgroundColor(color string) Application

	// Gives a frosted background to views. It overrides background color
	// setting when set.
	WithFrostedBackground() Application

	// Run starts the app.
	Run(url string, allowedHosts ...string) error
}

// PackageConfig is the interface that descibes a configuration to build a
// package.
type PackageConfig interface {
	// The targetted operating system.
	Target() string

	// Dumps the configuration in the given writer.
	Dump(io.Writer)
}

// Backend is the interface that describes a backend that handles the platform
// specific operations.
type Backend interface {
	// Runs the backend.
	Run() error

	// Calls the named method with the given input.
	Call(method string, out, in interface{}) error
}

// App returns the application.
func App() Application {
	return &defaultApp
}
