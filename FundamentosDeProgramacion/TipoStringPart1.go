package main

import (
	"fmt"
	"runtime"
)

var f int
var g uint8

func main ()  {

	f := 42
	g := uint8(f)

	fmt.Println(f)
	fmt.Printf("%T\n", g)

	fmt.Println(runtime.GOOS)  // muestra el sistema operativo
	fmt.Println(runtime.GOARCH) // muestra los bits


}