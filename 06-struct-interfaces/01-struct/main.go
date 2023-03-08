package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
}

// métodos
func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s\nEdad: %d\n", p.nombre, p.edad)
}

func (p *Persona) cambiarNombre(nuevoNombre string) {
	p.nombre = nuevoNombre
}

// herencia
type Empleado struct {
	Persona
	puesto string
	sueldo float64
}

func main() {
	p1 := Persona{"Juan", 20}
	// fmt.Println("p1", p1)
	p1.imprimir()

	// p1.nombre = "Kevin"
	p1.cambiarNombre("Kevin")
	p1.imprimir()

	// fmt.Println("p1", p1)

	p2 := Persona{
		nombre: "Pedro",
		edad:   30,
	}
	// fmt.Println("p2", p2)
	p2.imprimir()
	p2.cambiarNombre("Maria")
	p2.imprimir()

	empleado1 := Empleado{
		sueldo: 1000,
		puesto: "Programador",
	}
	empleado1.nombre = "Juan"
	empleado1.edad = 20
	fmt.Println("empleado1", empleado1)
	// empleado1.imprimir()

}
