// +build darwin

package mac

import (
	"fmt"
)

// Backend represents a backend that performs MacOS operations. It implements
// the murlok.Backend interface.
type Backend struct {
	// The allowed hosts.
	AllowedHosts []string

	// The views background color.
	BackgroundColor string

	// Enables views frosted background.
	FrostedBackground bool

	// The function called when the app terminates.
	Close func() error
}

// Run launches NSApplication. It satisfies the murlok.Backend interface.
func (b *Backend) Run() error {
	golang.Handle("app.OnRun", onRun)
	golang.Handle("app.OnReopen", onReopen)

	return platform.Call("app.Run", nil, struct{}{})
}

// Call satisfies the murlok.Backend interface.
func (b *Backend) Call(method string, out, in interface{}) error {
	return platform.Call(method, out, in)
}

func onRun(in map[string]interface{}) {
	fmt.Println("running")
}

func onReopen(in map[string]interface{}) {
	hasVisibleWindows := in["HasVisibleWindows"].(bool)
	fmt.Println("reopened - has windows:", hasVisibleWindows)
}
