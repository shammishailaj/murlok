// +build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Create("template.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "// Code generated by go generate; DO NOT EDIT.")
	fmt.Fprintln(f, "package core")

	gen := []struct {
		Var      string
		Filename string
	}{
		{Var: "bridgeJS", Filename: "bridge.js"},
	}

	for _, g := range gen {
		b, err := ioutil.ReadFile(g.Filename)
		if err != nil {
			panic(err)
		}

		fmt.Fprintln(f)
		fmt.Fprintf(f, "const %s = `%s`", g.Var, b)
	}
}
