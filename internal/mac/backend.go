// +build darwin

package mac

var backend *Backend

// Backend represents a backend that performs MacOS operations. It implements
// the murlok.Backend interface.
type Backend struct {
	// The local server host.
	Host string

	// The allowed hosts.
	AllowedHosts map[string]struct{}

	// The function used to create a default window.
	NewDefaultWindow func()

	// The function called before the app is closed.
	Finalize func()
}

// Run launches NSApplication. It satisfies the murlok.Backend interface.
func (b *Backend) Run() error {
	backend = b

	golang.Handle("app.OnRun", onRun)
	golang.Handle("app.OnReopen", onReopen)

	return platform.Call("app.Run", nil, struct{}{})
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
