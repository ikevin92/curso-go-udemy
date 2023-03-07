package main

import "fmt"

func main() {
	// Not
	// ! negación
	fmt.Println(!true)  // false
	fmt.Println(!false) // true

	// And
	// true && true = true
	fmt.Println("true && true = ", true && true)
	// true && false = false
	fmt.Println("true && false = ", true && false)

	// Or
	// true || false = true
	fmt.Println("true || false = ", true || false)
	// false || false = false
	fmt.Println("false || false = ", false || false)

	b := 2
	r := b == 2 && !(4 > b)
	fmt.Println(r)
}
