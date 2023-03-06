package main

import "fmt"

func main() {
	//arrays
	var numbers [5]int
	fmt.Println(numbers)

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30

	fmt.Println(numbers)

	//arrays con valores
	nombres := [3]string{"Juan", "Pedro", "Maria"}
	fmt.Println(nombres)

	colores := [...]string{
		"Rojo",
		"Verde",
		"Azul",
		"Amarillo",
	}

	fmt.Println(colores, len(colores))

	// indices definidos
	monedas := [...]string{
		0: "Dolar",
		2: "Sol",
		3: "Euro",
		5: "Peso",
	}

	fmt.Println(monedas, len(monedas))

	// sub arrays
	// subMonedas := monedas[0:3]
	// subMonedas := monedas[:3]
	subMonedas := monedas[3:]
	fmt.Println(subMonedas, len(subMonedas))

}
