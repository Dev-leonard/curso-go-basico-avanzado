package main

import (
	"fmt"
)

const (
	h = iota  // genera numeros constantes en orden consecutivo
	l
	k
	m = iota // no importa si hay otra descripcion en la ruta
	n
	p
)

func main ()  {

	fmt.Println(h)
	fmt.Println(l)
	fmt.Println(k)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(p)
}