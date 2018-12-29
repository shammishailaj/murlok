package murlok

import "github.com/maxence-charriere/murlok/internal/mac"

func (a *app) run() error {
	backend := mac.Backend{}
	return backend.Run()
}
