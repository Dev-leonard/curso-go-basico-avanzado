package main

import (
	"fmt"
)

type numero int
var r numero

func main () {
	fmt.Println(r)
	fmt.Printf("El tipo de r es: %T\n", r)
	r = 42
	fmt.Println(r)
}