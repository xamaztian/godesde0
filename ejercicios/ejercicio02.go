package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GeneraTabla() {
	var err error
	var numero3 int
	var bl bool
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
		fmt.Printf("%d * %d = %d \n", numero3, i, numero3*i)
	}

}
