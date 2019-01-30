package main

import (
	"net/http"

	"github.com/maxence-charriere/murlok"
	"github.com/maxence-charriere/murlok/ui"
)

func main() {
	http.Handle("/", &ui.Handler{
		WebDir: murlok.WebDir,
	})

	murlok.Run("/logo.png")
}
