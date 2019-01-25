// +build darwin

package mac

import "github.com/pkg/errors"

var backend *Backend

// Backend represents a backend that performs MacOS operations. It implements
// the murlok.Backend interface.
type Backend struct {
	// The allowed hosts.
	AllowedHosts map[string]struct{}

	// The function called before the app is closed.
	Finalize func()

	// The local server endpoint.
	LocalServerEndpoint string

	// The function to write logs.
	Logf func(string, ...interface{})

	// The function used to create a default window.
	NewDefaultWindow func()

	// The function to execute debug scoped instructions.
	WhenDebug func(func())
}

// Run launches NSApplication. It satisfies the murlok.Backend interface.
func (b *Backend) Run() error {
	backend = b

	golang.Handle("app.OnRun", onRun)
	golang.Handle("app.OnReopen", onReopen)
	golang.Handle("app.Debug", onDebug)
	golang.Handle("app.Error", onError)

	return platform.Call("app.Run", nil, struct {
		LocalServerEndpoint string
		AllowedHosts        map[string]struct{}
	}{
		LocalServerEndpoint: b.LocalServerEndpoint,
		AllowedHosts:        b.AllowedHosts,
	})
}

// Call satisfies the murlok.Backend interface.
func (b *Backend) Call(method string, out, in interface{}) error {
	return platform.Call(method, out, in)
}

func onRun(in map[string]interface{}) {
	backend.NewDefaultWindow()
}

func onReopen(in map[string]interface{}) {
	if hasVisibleWindows := in["HasVisibleWindows"].(bool); !hasVisibleWindows {
		backend.NewDefaultWindow()
	}
}

func onDebug(in map[string]interface{}) {
	backend.WhenDebug(func() {
		backend.Logf("%s", in["Msg"])
	})
}

func onError(in map[string]interface{}) {
	backend.Logf("%s", errors.Errorf("%s", in["Msg"]))
}