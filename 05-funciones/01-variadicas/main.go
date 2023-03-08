package main

import "fmt"

func sumar(nombre string, numbers ...int) (string, int) {

	mensaje := fmt.Sprintf("la suma de %s es: ", nombre)
	var total int
	for _, number := range numbers {
		total += number
	}

	return mensaje, total
}

func main() {
	mensaje, result := sumar("Kevin", 2, 3, 3)
	fmt.Println(mensaje, result)
}
