package main

import "fmt"

type person struct {
	nombre string
}

func cambiame(p *person) {
	p.nombre = "Miss Moneypenny"
}

func main() {
	p := person{
		nombre: "james bond",
	}
	fmt.Println(p)
	cambiame(&p)
	fmt.Println(p)
}
