package main

import "github.com/maxence-charriere/murlok"

func main() {
	// Allows the addresses with hosts listed below to be loaded into the
	// webview.
	murlok.AllowHosts(
		"app.segment.com",
		"segment.com",
	)

	// Launches the webview and loads the given remote url.
	murlok.Run("https://app.segment.com")
}
