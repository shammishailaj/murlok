package main

import (
	"fmt"
	"os"
)

var (
	greenColor   = "\033[92m"
	redColor     = "\033[91m"
	orangeColor  = "\033[93m"
	defaultColor = "\033[00m"
)

func printVerbose(format string, v ...interface{}) {
	if verbose {
		format = "‣ " + format
		fmt.Printf(format, v...)
		fmt.Println()
	}
}

func printSuccess(format string, v ...interface{}) {
	fmt.Print(greenColor)
	format = "✔ " + format
	fmt.Printf(format, v...)
	fmt.Println(defaultColor)
}

func printErr(format string, v ...interface{}) {
	fmt.Print(redColor)
	format = "x " + format
	fmt.Printf(format, v...)
	fmt.Println(defaultColor)
}

func fail(format string, v ...interface{}) {
	printErr(format, v...)
	os.Exit(-1)
}
