package main

import (
	"fmt"
	"sort"
)

type Usuario struct {
	Nombre string `json:"Nombre"`
	Edad   int    `json:"Edad"`
}

type PorEdad []Usuario

func (a PorEdad) Len() int           { return len(a) }
func (a PorEdad) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PorEdad) Less(i, j int) bool { return a[i].Edad < a[j].Edad }

func main() {
	u1 := Usuario{
		Nombre: "Adriano Praia",
		Edad:   20,
	}

	u2 := Usuario{
		Nombre: "Renan",
		Edad:   12,
	}

	usuarios := []Usuario{u1, u2}

	sort.Sort(PorEdad(usuarios))

	for _, v := range usuarios {
		fmt.Println(v.Nombre, v.Edad)
	}

}
