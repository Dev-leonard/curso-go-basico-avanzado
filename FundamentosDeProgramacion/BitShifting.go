package main

import (
	"fmt"
)

const (
	_ = iota   // _ este signo elimina el cero
	kb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
	tb = 1 << (iota * 10)

)
func main ()  {
j := 4
fmt.Printf("%d\t\t%b\n", j, j)
s := j << 1
fmt.Printf("%d\t\t%b\n", s, s)


/*kb := 1024
gb := kb * 1024
tb := gb * 1024*/
// estos datos pueden ser interpretados como aparecen en el const

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", gb, gb)
	fmt.Printf("%d\t\t%b\n", tb, tb)
}