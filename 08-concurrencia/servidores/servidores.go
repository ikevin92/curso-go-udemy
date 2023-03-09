package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)

	if err != nil {
		canal <- servidor + " no está disponible"
		// fmt.Println(servidor, "no está disponible")
	} else {
		canal <- servidor + " está funcionando correctamente"
		// fmt.Println(servidor, "está funcionando correctamente")
	}
}

func main() {

	inicio := time.Now()

	canal := make(chan string)

	servidores := []string{
		"http://oregoom.com",
		"http://www.udemy.com",
		"http://www.youtube.com",
		"http://www.facebook.com",
		"http://www.google.com",
	}

	for _, servidor := range servidores {
		go revisarServidor(servidor, canal)
	}

	// se imprime el canal de lo que est pasando en el servidor
	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}

	tiempoPaso := time.Since(inicio)
	fmt.Println("Tiempo de ejecución", tiempoPaso)

}
