package murlok

import "github.com/maxence-charriere/murlok/internal/mac"

func newBackend(host string) Backend {
	switch target {
	case "macos":
		return &mac.Backend{
			Host:           host,
			AllowedHosts:   allowedHosts,
			NewDefaultView: newDefaultView,
			Finalize:       finalize,
		}

	default:
		return nil
	}
}
