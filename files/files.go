package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/xamaztian/godesde0/ejercicios"
)

var fileName string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.GeneraTabla()

	archivo, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error al crear el archivo : " + err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()

}

func SumaTabla() {
	var texto string = ejercicios.GeneraTabla()

	if !Append(fileName, texto) {
		fmt.Println("Error al concatenar contenido")
	}

}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error durante el Append" + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)

	if err != nil {
		fmt.Println("Error durante el Append/Append" + err.Error())
		return false
	}

	arch.Close()

	return true
}

func LeoArchivo() {
	arch, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo: " + err.Error())
		return
	}

	scanner := bufio.NewScanner(arch)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println(">  " + registro)
	}

	arch.Close()
}
