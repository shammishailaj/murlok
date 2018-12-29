// +build darwin

package mac

// Backend represents a backend that performs MacOS operations.
type Backend struct {
	AllowedHosts []string
}

// Run launches NSApplication.
func (b *Backend) Run() error {
	panic("not implemented")
}
