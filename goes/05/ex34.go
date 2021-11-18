package main

import "fmt"

func main() {
	s := struct {
		nombre     string
		amigos     map[string]int
		bebidasFav []string
	}{
		nombre:     "Eduar",
		amigos:     map[string]int{"carito": 222, "Adriana": 333},
		bebidasFav: []string{"agua", "naranjada", "cerveza"},
	}

	fmt.Println(s.nombre)
	fmt.Println("\tAmigos:")
	for k, v := range s.amigos {
		fmt.Println("\t\t", k, v)
	}

	fmt.Println("\tBebidas favortias:")
	for k, v := range s.bebidasFav {
		fmt.Println("\t\t", k, v)
	}

}
