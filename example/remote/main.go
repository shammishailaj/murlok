package main

import "github.com/maxence-charriere/murlok"

func main() {
	murlok.AllowHosts(
		"dropbox.com",
		"www.dropbox.com",
		"api-3bdc2f77.duosecurity.com",
	)
	murlok.Run("https://paper.dropbox.com")
}
