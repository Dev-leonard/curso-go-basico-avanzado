package main

import (
	"fmt"
	"runtime"
)

var a int
var b uint8

func main ()  {

	a := 42
	b := uint8(a)

	fmt.Println(a)
	fmt.Printf("%T\n", b)

	fmt.Println(runtime.GOOS)  // muestra el sistema operativo
	fmt.Println(runtime.GOARCH) // muestra los bits


}