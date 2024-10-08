//go:build js && wasm

package main

import (
	"fmt"
	wasm "nknovh-wasm"
)

func main() {
	fmt.Println("Go Wasm loaded")
	c := new(wasm.CLIENT)
	c.RegisterJSFuncs()
	c.Init()
	c.Run()
	<-make(chan bool)
}
