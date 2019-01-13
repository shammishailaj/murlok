package murlok

import "github.com/maxence-charriere/murlok/internal/mac"

func newBackend(localServerEndpoint string) Backend {
	switch target {
	case "macos":
		return &mac.Backend{
			AllowedHosts:        allowedHosts,
			Finalize:            finalize,
			LocalServerEndpoint: localServerEndpoint,
			Logf:                Logf,
			NewDefaultWindow:    newDefaultWindow,
			WhenDebug:           WhenDebug,
		}

	default:
		return nil
	}
}
