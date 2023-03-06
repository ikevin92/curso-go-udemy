package main

import "fmt"

func main() {
	days := make(map[int]string)
	fmt.Println(days)

	//Agregar datos
	days[1] = "Domingo"
	days[2] = "Lunes"
	days[3] = "Martes"
	days[4] = "Miércoles"
	days[5] = "Jueves"
	days[6] = "Viernes"
	days[7] = "Sábado"

	fmt.Println(days)

	//modificar
	days[7] = "SABADO"
	fmt.Println(days)

	// eliminar
	delete(days, 1)
	fmt.Println(days)

	fmt.Println(days, len(days))

	// NUevo mapa
	students := make(map[string][]int)

	students["Alex"] = []int{13, 15, 16}
	students["Roel"] = []int{14, 13, 17}

	fmt.Println(students)

	fmt.Println(students["Alex"][1])
}
