package core

import (
	"bytes"
	"html/template"
)

// BridgeJS generates the javascript string to evaluate in order to setup the
// murlok object in a web view.
func BridgeJS(localServerURL string) string {
	var b bytes.Buffer

	tmpl := template.Must(template.New("bridge.js").Parse(bridgeJS))
	tmpl.Execute(&b, struct {
		LocalServerURL string
	}{
		LocalServerURL: localServerURL,
	})

	return b.String()
}
