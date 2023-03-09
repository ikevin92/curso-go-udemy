package mensajes

import "fmt"

func Hola() { // mayus e publico y minus es provado
	fmt.Println("Hola desde el paquete mensajes")
}

const mensaje = "Hola desde mi constante"

func funcionPrivada() {
	fmt.Println("Hola desde la funcion privada")
}

func Imprimir() {
	fmt.Println(mensaje)
	funcionPrivada()
}
