package main

import "fmt"

type persona struct {
	nombre          string
	apellido        string
	sabores_helados []string
}

func main() {
	p := persona{
		nombre:   "eduar",
		apellido: "tua",
		sabores_helados: []string{
			"chocolate", "mantecado",
		},
	}

	fmt.Println(p.nombre, p.apellido)
	for i, v := range p.sabores_helados {
		fmt.Println("\t", i, v)
	}
}
