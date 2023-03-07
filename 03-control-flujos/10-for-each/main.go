package main

import "fmt"

func main() {
	names := [...]string{"Juan", "Pedro", "Luis", "Maria", "Jose"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println("names[", i, "] = ", names[i])
	// }

	for indice, name := range names { //indice se puede reemplazar por _
		fmt.Println("names[", indice, "] = ", name)
	}

	for _, name := range names {
		fmt.Println("name = ", name)
	}

}
