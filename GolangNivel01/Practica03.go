package main

import (
	"fmt"
)

var n int = 42
var m string = "Juanito"
var p bool = true

func main() {

	s := fmt.Sprintf("%v\t%v\t%v", n, m, p)

	fmt.Println(s)

}
