// +build !js

// Package murlok is a Go library to build gui apps.
// /!\ To be continued.
package murlok

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"time"
)

var (
	// DefaultLogger is the logger used to write logs.
	DefaultLogger Logger = func(f string, a ...interface{}) {
		fmt.Print("‣ ")
		fmt.Printf(f, a...)
		fmt.Println()
	}

	// DefaultWindow describes the appearance of the default window.
	DefaultWindow Window

	// Finalize is the function called before the app is closed. It can be set
	// to perform any final cleanup.
	Finalize func()

	// MacOS is the configuration used by the murlok command line tool to build
	// a MacOS package.
	MacOS MacOSConfig

	// Server is the server used to serve local requests.
	Server = &http.Server{Addr: ":0"}

	allowedHosts = make(map[string]struct{})
	backend      Backend
	target       string
	verbose      string
	whenDebug    = func(func()) {}
)

// Backend is the interface that describes a backend that handles the platform
// specific operations.
type Backend interface {
	// Runs the backend.
	Run() error

	// Calls the named method with the given input.
	Call(method string, out, in interface{}) error
}

// Logger describes a function that writes logs.
type Logger func(string, ...interface{})

func init() {
	runtime.LockOSThread()
}

// Log logs the given arguments separated by a space.
func Log(args ...interface{}) {
	format := ""

	for range args {
		format += "%v "
	}

	format = strings.TrimSpace(format)
	DefaultLogger(format, args...)
}

// Logf logs the given arguments according to the specified format.
func Logf(format string, args ...interface{}) {
	DefaultLogger(format, args...)
}

// EnableDebug enables or disable debug mode.
func EnableDebug(v bool) {
	if v {
		whenDebug = func(f func()) { f() }
	} else {
		whenDebug = func(func()) {}
	}
}

// WhenDebug calls the given function when debug mode is enabled.
func WhenDebug(f func()) {
	whenDebug(f)
}

// AllowHosts authorized url with the given hosts to be loaded in web views.
// Unauthorized url are loaded in the operating system default browser.
func AllowHosts(hosts ...string) {
	for _, host := range hosts {
		allowedHosts[host] = struct{}{}
	}
}

// Run runs the application and shows a web view that loads the given url
func Run(rawurl string) {
	EnableDebug(verbose == "true")

	defaultWindowURL, err := url.Parse(rawurl)
	if err != nil {
		Log("parsing url failed:", err)
		os.Exit(1)
	}

	DefaultWindow.URL = defaultWindowURL.String()
	AllowHosts(defaultWindowURL.Host)

	if murlokBuild := os.Getenv("MURLOK_BUILD"); len(murlokBuild) != 0 {
		build(murlokBuild)
		return
	}

	http.HandleFunc("/murlok", rpc)

	listener, err := net.Listen("tcp", Server.Addr)
	if err != nil {
		Logf("listening on %s failed: %s", Server.Addr, err)
		os.Exit(1)
	}

	go func() {
		Server.Serve(listener)
	}()

	port := listener.Addr().(*net.TCPAddr).Port
	localServerURL := fmt.Sprintf("http://localhost:%v", port)
	AllowHosts(localServerURL)

	backend = newBackend(localServerURL)
	if err = backend.Run(); err != nil {
		Logf("running %T failed: %s", backend, err)
		os.Exit(1)
	}
}

func build(path string) {
	var packageConfig interface{}

	switch target {
	case "macos":
		packageConfig = MacOS

	default:
		Log("no operating system targeted")
		return
	}

	f, err := os.Create(path)
	if err != nil {
		Log(err)
		return
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	if err := enc.Encode(packageConfig); err != nil {
		Log(err)
	}
}

func newDefaultWindow(url string) {
	WhenDebug(func() {
		b, _ := json.MarshalIndent(DefaultWindow, "", "    ")
		Log("creating window", string(b))
	})

	c := DefaultWindow

	if url != "" {
		c.URL = url
	}

	if err := backend.Call("windows.New", nil, c); err != nil {
		Log("creating view failed:", err)
	}
}

func finalize() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := Server.Shutdown(ctx); err != nil {
		Log("shutting down server failed:", err)
	}

	if Finalize != nil {
		finalize()
	}
}
