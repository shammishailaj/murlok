// +build darwin

package mac

// Backend represents a backend that performs MacOS operations.
type Backend struct {
	// The allowed hosts.
	AllowedHosts []string

	// The function called when the app terminates.
	Close func() error
}

// Run launches NSApplication.
func (b *Backend) Run() error {
	return platform.Call("app.Run", nil, struct{}{})
}
