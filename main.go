package main

import (
	"fmt"

	"github.com/xamaztian/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(321458)
	fmt.Println(estado)
	fmt.Println(texto)
}
