package main

import (
	"net/http"

	"github.com/maxence-charriere/murlok"
	"github.com/maxence-charriere/murlok/app"
)

func main() {
	http.Handle("/", &app.Handler{
		WebDir: murlok.WebDir,
	})

	murlok.Run("/")
}
