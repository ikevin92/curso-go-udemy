package main

import "fmt"

func saludar(nombre string) {
	// fmt.Println("Hola Mundo")
	fmt.Println("Hola,", nombre)
}

func sumar(a, b int) int {
	return a + b
}

func main() {
	saludar("kevin")
	result := sumar(10, 20)
	fmt.Println("El resultado es:", result)
}
