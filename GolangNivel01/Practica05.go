package main

import (
	"fmt"
)

type numeroR int
var R numeroR
var L int

func main () {
	fmt.Println(R)
	fmt.Printf("El tipo de R es: %T\n", R)
	R = 42
	fmt.Println(R)

	L = int (R)
	fmt.Println(L)
	fmt.Printf("%T", L)
}
