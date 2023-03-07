package main

import (
	"fmt"
)

func main() {
	// Bucle infinito
	// for {
	// 	fmt.Println("Hola mundo")
	// }

	// Bucle tipo while
	numbers := 12455
	c := 0
	for numbers > 0 {
		numbers /= 10
		c++
	}

	fmt.Println("Cantidad de dígitos es:", c)

	// Bucle tipo for
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
