package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar(expression string) int {
	//4+4+5+8
	valores := strings.Split(expression, "+")
	var result int

	for _, valor := range valores {
		num, error := strconv.Atoi(valor)
		if error != nil {
			fmt.Println(error)
			fmt.Println("Error: tiene que ingresar un número entero")
			fmt.Println("Ejemplo: 4+4+5+8, solo debes realizar con operador +!!")
			return 0
		} else {
			result += num
		}
	}
	return result
}

func main() {
	var expression string
	var result int

	fmt.Print("=>")
	fmt.Scanln(&expression)
	result = sumar(expression)
	fmt.Printf("=> %d \n", result)
}
