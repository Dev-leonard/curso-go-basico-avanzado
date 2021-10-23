package main

import (
	"CursoGo/Calculadorabasica/Operadores"
	"bufio"
	"fmt"
	"os"
	"strings"
)
 func main() {
	 scanner := bufio.NewScanner(os.Stdin)
	 fmt.Println("LA CALCULADORA BASICA")
	 scanner.Scan()
	 operacion := scanner.Text()

	 resultado := 0
	 if strings.Contains(operacion, "+") {
		 resultado = Operadores.Sumar(operacion)
	 } else if strings.Contains(operacion, "-"){
		 resultado = Operadores.Restar(operacion)
	 } else if strings.Contains(operacion, "*"){
		 resultado = Operadores.Multiplicar(operacion)
	 } else if strings.Contains(operacion, "/"){
		 resultado = Operadores.Dividir(operacion)
	 } else {
		 fmt.Println("Error: El Operdaor está mal Ingresado,")
		 fmt.Println("o ¡¡Sólo debes realiar con un Operador!!")
		 fmt.Println("o ¡¡ Este Operador no esta implementado!!")
	 }

	 fmt.Println(resultado)
 }