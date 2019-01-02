package murlok

import (
	"os"

	"github.com/maxence-charriere/murlok/internal/mac"
)

func (a *app) run() error {
	var conf PackageConfig

	for _, c := range a.pkgconfs {
		if c.Target() == "macos" {
			conf = c
		}
	}

	if murlokBuild := os.Getenv("MURLOK_BUILD"); len(murlokBuild) != 0 {
		return a.runForBuild(murlokBuild, conf)
	}

	backend := mac.Backend{
		AllowedHosts:      a.allowedHosts,
		BackgroundColor:   a.backgroundColor,
		FrostedBackground: a.frostedBackground,
	}

	return backend.Run()
}
