package main

import "github.com/maxence-charriere/murlok"

func main() {
	murlok.AllowHosts(
		"www.theverge.com",
		"store.theverge.com",
		"https://the-verge.myshopify.com/",
	)

	murlok.Run("https://www.theverge.com")
}
