package main

import "fmt"

func main() {
	a := 20
	b := 10

	/*
		Comentario de bloque
	*/
	// Suma
	result := a + b
	fmt.Println("Suma:", result)

	// Resta
	result = a - b
	fmt.Println("Resta:", result)

	// Multiplicación
	result = a * b
	fmt.Println("Multiplicación:", result)

	// División
	result = a / b
	fmt.Println("División:", result)

	var div float64 = 3.0 / 2.0
	fmt.Println("División:", div)

	// Módulo
	result = a % b
	fmt.Println("Módulo:", result)

	result = 3 % 2
	fmt.Println("Módulo:", result)

}
