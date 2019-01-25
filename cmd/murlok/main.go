package main

//go:generate go run templates/main.go

import (
	"context"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/segmentio/conf"
)

var (
	verbose bool
)

func main() {
	ld := conf.Loader{
		Name: "goapp",
		Args: os.Args[1:],
		Commands: []conf.Command{
			{Name: "build", Help: "Build and package a program."},
			{Name: "run", Help: "Build and package and run program."},
			{Name: "clean", Help: "Delete packaged program."},
			{Name: "update", Help: "Update repository to the latest version."},
			{Name: "help", Help: "Show the help."},
		},
	}

	ctx, cancel := ctxWithSignals(context.Background(), os.Interrupt)
	defer cancel()

	switch cmd, args := conf.LoadWith(nil, ld); cmd {
	case "build":
		build(ctx, args)

	case "run":
		run(ctx, args)

	case "clean":
		clean(ctx, args)

	case "help":
		ld.PrintHelp(nil)

	default:
		panic("unreachable")
	}
}

func ctxWithSignals(parent context.Context, s ...os.Signal) (ctx context.Context, cancel func()) {
	ctx, cancel = context.WithCancel(parent)
	sigc := make(chan os.Signal)
	signal.Notify(sigc, s...)

	go func() {
		defer close(sigc)
		<-sigc
		cancel()
	}()

	return ctx, cancel
}

func stringWithDefault(value, defaultValue string) string {
	if len(value) == 0 {
		return defaultValue
	}

	return value
}

func intWithDefault(value, defaultValue int) int {
	if value == 0 {
		return defaultValue
	}

	return value
}

func trimExt(path string) string {
	return strings.TrimSuffix(path, filepath.Ext(path))
}

func generateTemplatedFile(path, tmpl string, data interface{}) error {
	t, err := template.New("").Parse(strings.TrimSpace(tmpl))
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return t.Execute(f, data)
}
