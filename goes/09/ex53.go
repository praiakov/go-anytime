package main

import "fmt"

type pesrona struct {
	nombre string
}

func (p *pesrona) hablar() {
	fmt.Println("Hola, soy ", p.nombre)
}

type humanno interface {
	hablar()
}

func algo(h humanno) {
	h.hablar()
}

func main() {
	p1 := pesrona{
		nombre: "Eduar",
	}

	algo(&p1)
}
