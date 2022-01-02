package main

import (
	"encoding/json"
	"fmt"
)

type usuario struct {
	Nombre string `json:"Nombre"`
	Edad   int    `json:"Edad"`
}

func main() {
	u1 := usuario{
		Nombre: "Adriano Praia",
		Edad:   20,
	}

	u2 := usuario{
		Nombre: "Renan",
		Edad:   12,
	}

	usuarios := []usuario{u1, u2}

	bs, err := json.Marshal(usuarios)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

}
