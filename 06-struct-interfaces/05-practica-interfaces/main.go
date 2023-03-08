package main

import (
	"fmt"
	"math"
)

type Geometrica interface {
	area() float64
	perimetro() float64
}

type Cuadrado struct {
	lado float64
}

type Circulo struct {
	radio float64
}

func (cuadrado *Cuadrado) area() float64 {
	return cuadrado.lado * cuadrado.lado
}

func (cuadrado *Cuadrado) perimetro() float64 {
	return cuadrado.lado * 4
}

func (circulo *Circulo) area() float64 {
	return math.Pi * circulo.radio * circulo.radio
}

func (circulo *Circulo) perimetro() float64 {
	return 2 * math.Pi * circulo.radio
}

func medidas(geometrica Geometrica) {
	fmt.Println(geometrica)
	fmt.Println("Area:", geometrica.area())
	fmt.Println("Perimetro:", geometrica.perimetro())
}

func main() {
	cuadrado := Cuadrado{lado: 4}
	medidas(&cuadrado)
	circulo := Circulo{radio: 10}
	medidas(&circulo)

}
