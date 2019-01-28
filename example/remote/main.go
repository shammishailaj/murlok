package main

import "github.com/maxence-charriere/murlok"

func main() {
	murlok.DefaultWindow = murlok.Window{
		BackgroundColor: "#ffffff",
		TextColor:       "#000000",
		TitleBarColor:   "#ffffff",
	}

	murlok.AllowHosts(
		// "www.theverge.com",
		// "store.theverge.com",
		// "https://the-verge.myshopify.com/",
		"fiddle.jshell.net",
	)

	// murlok.Run("https://www.theverge.com")
	murlok.Run("https://paper.dropbox.com")
	// murlok.Run("https://www.w3schools.com/jsref/tryit.asp?filename=tryjsref_alert")
	// murlok.Run("https://www.w3schools.com/jsref/tryit.asp?filename=tryjsref_confirm")
	// murlok.Run("https://www.w3schools.com/jsref/tryit.asp?filename=tryjsref_prompt")
	// murlok.Run("https://jsfiddle.net/Luvo4axh/5/")
}
