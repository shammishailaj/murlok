package murlok

import "github.com/maxence-charriere/murlok/internal/mac"

func newBackend(localServerURL string) Backend {
	switch target {
	case "macos":
		return &mac.Backend{
			AllowedHosts:     allowedHosts,
			Finalize:         finalize,
			LocalServerURL:   localServerURL,
			Logf:             Logf,
			NewDefaultWindow: newDefaultWindow,
			WhenDebug:        WhenDebug,
		}

	default:
		return nil
	}
}
