package main

import (
	"fmt"

	"github.com/xamaztian/godesde0/ejercicios"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(321458)
	fmt.Println(estado)
	fmt.Println(texto)*/

	numero, texto := ejercicios.CalculoEjercicioUno("a98")

	fmt.Println("El número es ", numero)
	fmt.Println(texto)
}
