package main

import (
	"fmt"
)

type numero int

var x numero

func main() {
	fmt.Println(x)
	fmt.Printf("Tipo de x é: %T\n", x)

	x = 42
	fmt.Println(x)
}
