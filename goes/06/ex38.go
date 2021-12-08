package main

import "fmt"

type persona struct {
	nombre   string
	edad     int
	apellido string
}

func (p persona) presentar() {
	fmt.Println("Hola, soy ", p.nombre, p.apellido, "y tengo", p.edad, "años")
}

func main() {
	p := persona{
		nombre:   "james bond",
		edad:     32,
		apellido: "007",
	}

	p.presentar()
}
