package main

import (
	"fmt"
	"example.com/conta"
)

func main() {
	contaCris := contas.ContaCorrente{Titular: "Cris", Saldo: 300}
	contaBruno := contas.ContaCorrente{Titular: "Bruno", Saldo: 100}

	status := contaCris.Transferir(150, &contaBruno)

	fmt.Println(status)
	fmt.Println(contaCris)
	fmt.Println(contaBruno)
}
