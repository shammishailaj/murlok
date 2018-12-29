package murlok

import "net/http"

type app struct {
	server       *http.Server
	pkgconfs     []PackageConfig
	defaultURL   string
	allowedHosts []string
}

func (a *app) WithCustomServer(s *http.Server) Application {
	a.server = s
	return a
}

func (a *app) WithPackageConfig(c PackageConfig) Application {
	a.pkgconfs = append(a.pkgconfs, c)
	return a
}

func (a *app) Run(url string, allowedHosts ...string) error {
	a.defaultURL = url
	a.allowedHosts = allowedHosts
	return a.run()
}
