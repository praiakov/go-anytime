package main

import (
	"fmt"

	contas "example.com/conta"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaBruno := contas.ContaPoupanca{}
	contaBruno.Depositar(100)
	PagarBoleto(&contaBruno, 60)

	contaLuisa := contas.ContaCorrente{}
	contaLuisa.Depositar(150)
	PagarBoleto(&contaLuisa, 87)

	fmt.Println(contaBruno)
	fmt.Println(contaLuisa)

}
