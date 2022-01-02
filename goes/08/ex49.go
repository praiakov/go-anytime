package main

import (
	"encoding/json"
	"fmt"
	"os"
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

	err := json.NewEncoder(os.Stdout).Encode(usuarios)

	if err != nil {
		fmt.Println(err)
	}
}
