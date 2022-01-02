package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre string `json:"Nombre"`
	Edad   int    `json:"Edad"`
}

func main() {
	s := `[{"nombre":"Adriano Praia", "edad" : 20}, {"nombre":"Ana", "edad" : 20}]`

	var personas []persona

	err := json.Unmarshal([]byte(s), &personas)

	if err != nil {
		fmt.Println(err)
	}

	for i, p := range personas {
		fmt.Println("\t Persona #", i+1)
		fmt.Println("\t ", p.Nombre, p.Edad)
	}
}
