package murlok

import (
	"github.com/maxence-charriere/murlok/internal/core"
	"github.com/maxence-charriere/murlok/internal/mac"
)

func newBackend(localServerURL string) Backend {
	switch target {
	case "macos":
		return &mac.Backend{
			AllowedHosts:     allowedHosts,
			BridgeJS:         core.BridgeJS(localServerURL),
			Finalize:         finalize,
			Logf:             Logf,
			NewDefaultWindow: newDefaultWindow,
			SettingsURL:      SettingsURL,
			WhenDebug:        WhenDebug,
		}

	default:
		return nil
	}
}
