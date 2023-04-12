package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GeneraTabla() string {
	var err error
	var numero3 int
	var bl bool
	var texto string

	bl = true
	scanner := bufio.NewScanner(os.Stdin)
	for bl {
		fmt.Println("Ingrese un número")

		if scanner.Scan() {
			numero3, err = strconv.Atoi(scanner.Text())

			if err != nil {
				fmt.Println("Formato no reconocido, error :" + err.Error())
			} else {
				bl = false
			}
		}
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d * %d = %d \n", numero3, i, numero3*i)
	}

	return texto
}
