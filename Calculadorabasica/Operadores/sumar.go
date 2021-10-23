package Operadores

import (
	"fmt"
	"strconv"
	"strings"
)
//Sumar realiza la suma
func Sumar(operacion string) int {
	valores := strings.Split(operacion, "+")
	resultado := 0

	for i := 0; i < len(valores); i++ {
		num, error := strconv.Atoi(valores[i])

		if error != nil {
			fmt.Println(error)
			fmt.Println("Error: tiene que ingresar un numero entero")
			fmt.Println("o ¡¡Sólo debes realizar con un operador!!")
		} else {
			resultado += num
		}
	}
	return resultado
}