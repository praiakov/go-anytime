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

func (c *ContaCorrente) Sacar(valordoSaque float64) string {
	podeSacar := valordoSaque > 0 && valordoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valordoSaque
		return "Saque realizado com sucesso"
	}

	return "Saldo insuficiente"
}

func main() {
	contaCris := ContaCorrente{}
	contaCris.titular = "cris"
	contaCris.saldo = 500

	fmt.Println(contaCris.saldo)

	fmt.Println(contaCris.Sacar(-200))
	fmt.Println(contaCris.saldo)
}
