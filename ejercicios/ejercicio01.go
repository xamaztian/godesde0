package ejercicios

import (
	"strconv"
)

func CalculoEjercicioUno(valor string) (int, string) {
	var num int
	num, err := strconv.Atoi(valor)

	if err != nil {
		return 0, err.Error()
	}

	if num >= 100 {
		return num, "Es mayor o igual a 100"
	} else {
		return num, "Es menor a 100"
	}
}
