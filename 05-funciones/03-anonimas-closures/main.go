package main

import (
	"fmt"
	"strings"
)

// Closures
func repeat(n int) func(cadena string) string {
	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}
}

func main() {
	// funcion anonima
	// func() {
	// 	fmt.Println("Hola desde la funcion anonima")
	// }()

	myFunc := func(nombre string) string {
		return fmt.Sprintf("Hola, %s desde la funciona anonima", nombre)
	}

	fmt.Println(myFunc("Jorge"))

	// Closures
	repeat3 := repeat(3)
	fmt.Println(repeat3("Hola "))
	fmt.Println(repeat3("Mundo"))

	repeat5 := repeat(5)
	fmt.Println(repeat5("Hola "))
	fmt.Println(repeat5("Mundo"))

}
