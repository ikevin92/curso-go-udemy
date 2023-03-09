package main

import "github.com/donvito/hellomod"

// "paquetes/mensajes"

func main() {
	// mensajes.Hola()
	// mensajes.Imprimir()

	/*
		 	cuadrado1 := figuras.Cuadrado{Lado: 10}
			figuras.Medidas(&cuadrado1)

			circulo1 := figuras.Circulo{Radio: 10}
			figuras.Medidas(&circulo1)
	*/

	/*
		 	p1 := models.Persona{}
			p1.Constructor("Juan", 20)
			fmt.Println("Persona:", p1)

			p1.SetNombre("Pedro")
			fmt.Println("Nombre p1:", p1.GetNombre())
	*/

	hellomod.SayHello()
}
