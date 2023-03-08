package main

import "fmt"

type User struct {
	name     string
	email    string
	isActive bool
}

type Student struct {
	user User
	code string
}

// Relacion de uno a muchos
type Curso struct {
	titulo string
	videos []Video
}

type Video struct {
	titulo string
	curso  Curso
}

func main() {

	// Relacion de uno a uno
	kevin := User{
		name:     "Kevin",
		email:    "kevin@gmail.com",
		isActive: true,
	}

	roel := User{
		name:     "Roel",
		email:    "roel@gmail.com",
		isActive: true,
	}

	kevinStudent := Student{
		user: kevin,
		code: "1234abc",
	}

	fmt.Println("Roel", roel)
	fmt.Println("kevinStudent", kevinStudent)

	// Relacion de uno a muchos
	video1 := Video{titulo: "01-Introducción"}
	video2 := Video{titulo: "02-Instalación"}
	video3 := Video{titulo: "03-Variables"}

	cursoGo := Curso{
		titulo: "Curso de Go",
		videos: []Video{video1, video2, video3},
	}

	video1.curso = cursoGo
	video2.curso = cursoGo
	video3.curso = cursoGo

	fmt.Println(video1.curso.titulo)
	fmt.Println("cursoGo", cursoGo)

	for _, video := range cursoGo.videos {
		fmt.Println(video.titulo)
	}
}
