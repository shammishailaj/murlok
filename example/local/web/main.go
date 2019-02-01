// +build wasm

package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("hello world")

	href := js.Global().
		Get("location").
		Get("href").
		String()
	fmt.Println("href:", href)

	// localURL := js.Global().
	// 	Get("murlok").
	// 	Get("url").
	// 	String()
	// fmt.Println("lcal url:", localURL)

}
