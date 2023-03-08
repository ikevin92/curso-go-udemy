package main

import "fmt"

func modificarValor(cadena *string) {
	fmt.Printf("%p\n", cadena)
	*cadena = "Hola desde la funcion"
}

func main() {

	cadena := "Hola desde el main"
	fmt.Printf("%p\n", &cadena)
	fmt.Println("Antes de llamar a la funcion:", cadena)

	modificarValor(&cadena) // copia
	fmt.Println("Despues de llamar a la funcion:", cadena)

}
