package main

import "fmt"

func main() {
	//Slicen
	numbers := []int{1, 2, 3}
	fmt.Println(numbers, len(numbers))

	// agregar elementos
	numbers = append(numbers, 4)
	numbers = append(numbers, 5)

	fmt.Println(numbers, len(numbers))

	// sub Slicen
	subSlicen := numbers[2:]

	numbers[0] = 100
	fmt.Println(subSlicen, len(subSlicen))
	fmt.Println(numbers, len(numbers))

	// punteros
	// longitud
	// capacidad

	meses := []string{"Enero", "Febrero", "Marzo"}
	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)

	meses = append(meses, "Abril")
	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)
}
