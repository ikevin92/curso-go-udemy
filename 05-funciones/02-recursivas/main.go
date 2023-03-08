package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	//3 1 2 3 factorial
	// 3! = 3 * 2 * 1
	result := factorial(3)
	fmt.Println("El factorial de 3 es: ", result)
}
