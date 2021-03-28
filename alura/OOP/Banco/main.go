package main

import (
	"fmt"
)

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	a := ContaCorrente{
		titular:       "James",
		numeroAgencia: 123,
		numeroConta:   12,
		saldo:         14.2,
	}

	fmt.Println(a)
}
