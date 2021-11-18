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

	m := map[string]persona{
		p.apellido: p,
	}

	for _, v := range m {
		fmt.Println(v.nombre)
		fmt.Println(v.apellido)

		for i, v := range v.sabores_helados {
			fmt.Println(" ", i, v)
		}
		fmt.Println("-----------")
	}
}
