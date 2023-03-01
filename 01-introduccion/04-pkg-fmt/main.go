package main

import "fmt"

func main() {

	hola := "hola"
	mundo := "mundo"

	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Kevin"
	edad := 30

	fmt.Printf("Hola, mi nombre es %s y tengo %d años \n", nombre, edad)
	fmt.Printf("Hola, mi nombre es %v y tengo %v años \n", nombre, edad)

	mensaje := fmt.Sprintf("Hola, mi nombre es %s y tengo %d años", nombre, edad)

	fmt.Println(mensaje)

	fmt.Printf("nombre: %T \n", nombre)
	fmt.Printf("edad: %T \n", edad)

	fmt.Print("Ingrese otro nombre: ")
	fmt.Scanln(&nombre)

	fmt.Println("Otro nombre: ", nombre)

	fmt.Println("El nombre ingresado es:", nombre)

}
