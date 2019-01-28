// +build darwin

package mac

import "github.com/pkg/errors"

var backend *Backend

// Backend represents a backend that performs MacOS operations. It implements
// the murlok.Backend interface.
type Backend struct {
	// The allowed hosts.
	AllowedHosts map[string]struct{}

	// The javascript to evaluate in order setup the murlok object in a web
	// view.
	BridgeJS string

	// The default URL to be loaded when a window is created.
	DefaultURL string

	// The function called before the app is closed.
	Finalize func()

	// The function to write logs.
	Logf func(string, ...interface{})

	// The function used to create a default window.
	NewDefaultWindow func(string)

	// The url loaded when the menu bar "Preferences" button is clicked.
	SettingsURL string

	// The function to execute debug scoped instructions.
	WhenDebug func(func())
}

// Run launches NSApplication. It satisfies the murlok.Backend interface.
func (b *Backend) Run() error {
	backend = b

	golang.Handle("app.Running", onRun)
	golang.Handle("app.Reopened", onReopen)
	golang.Handle("app.Debug", onDebug)
	golang.Handle("app.Error", onError)
	golang.Handle("app.Windows.NewDefault", onNewDefaultWindow)

	return platform.Call("app.Run", nil, struct {
		AllowedHosts map[string]struct{}
		BridgeJS     string
		DefaultURL   string
		SettingsURL  string `json:",omitempty"`
	}{
		AllowedHosts: b.AllowedHosts,
		BridgeJS:     b.BridgeJS,
		DefaultURL:   b.DefaultURL,
		SettingsURL:  b.SettingsURL,
	})
}

// Call satisfies the murlok.Backend interface.
func (b *Backend) Call(method string, out, in interface{}) error {
	return platform.Call(method, out, in)
}

func onRun(in map[string]interface{}) {
	backend.NewDefaultWindow("")
}

func onReopen(in map[string]interface{}) {
	if hasVisibleWindows := in["HasVisibleWindows"].(bool); !hasVisibleWindows {
		backend.NewDefaultWindow("")
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

func onNewDefaultWindow(in map[string]interface{}) {
	backend.NewDefaultWindow(in["URL"].(string))
}
