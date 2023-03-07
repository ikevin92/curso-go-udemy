package main

import "fmt"

func main() {
	if nombre, edad := "Roel", 30; nombre == "Kevin" {
		fmt.Println("Hola,", nombre)
	} else {
		fmt.Printf("Nombre: %s - Edad: %d\n", nombre, edad)
	}

	// Obtener valor de mapa
	users := make(map[string]string)

	users["Kevin"] = "kevin@gmail.com"
	users["Maria"] = "maria@gmail.com"

	// correo, ok := users["Kevin"]

	if correo, ok := users["Kevin"]; ok {
		fmt.Println("Correo:", correo, "OK:", ok)
	} else {
		fmt.Println("No existe el usuario")
	}

	// fmt.Println("Usuarios:", users["Juan"])
}
