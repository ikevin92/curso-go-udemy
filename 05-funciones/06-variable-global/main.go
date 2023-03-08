package main

import "fmt"

// Variables globales
var mensaje string

func funcion1() {
	mensaje = "Hola desde la funcion 1"
	fmt.Println(mensaje)
}

func funcion2() {
	mensaje = "Hola desde la funcion 2"
	fmt.Println(mensaje)
}

func main() {
	mensaje = "Hola desde el main"
	fmt.Println(mensaje)

	defer funcion1() // con defer quiere decir que se ejecuta al final
	funcion2()

	fmt.Println(mensaje)

}
