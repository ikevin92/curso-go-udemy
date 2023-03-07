package main

import "fmt"

func main() {
	var vocal string

	fmt.Print("Ingresa una vocal: ")
	fmt.Scan(&vocal)

	/*
		switch {
		case vocal == "a" || vocal == "A":
			fmt.Println(vocal, "Es una vocal")
		case vocal == "e" || vocal == "E":
			fmt.Println(vocal, "Es una vocal")
		case vocal == "i" || vocal == "I":
			fmt.Println(vocal, "Es una vocal")
		case vocal == "o" || vocal == "O":
			fmt.Println(vocal, "Es una vocal")
		case vocal == "u" || vocal == "U":
			fmt.Println(vocal, "Es una vocal")
		default:
			fmt.Println("No es una vocal")
		}
	*/

	switch vocal {
	case "a", "A", "e", "E", "i", "I", "o", "O", "u", "U":
		fmt.Println(vocal, "Es una vocal")
	default:
		fmt.Println("No es una vocal")
	}
}
