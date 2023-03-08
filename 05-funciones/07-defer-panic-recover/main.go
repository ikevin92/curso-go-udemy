package main

import (
	"fmt"
	"os"
)

func main() {

	// Defer Function
	defer func() {
		// la funcion recover() permite recuperar el error que se produjo
		if error := recover(); error != nil {
			fmt.Println("Ups al parecer el programa no finalizo de forma correcta: ", error)
		}
	}()

	if file, error := os.Open("hoa.txt"); error != nil {
		panic("No fue posible abrir el archivo")
	} else {
		defer func() {
			fmt.Println("Cerrando el archivo")
			file.Close()
		}()
		contenido := make([]byte, 254)
		long, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[:long])
		fmt.Println(contenidoArchivo)
	}

}
