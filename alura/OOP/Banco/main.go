package main

import (
	"fmt"
	"example.com/conta"
	"example.com/cliente"
)

func main() {
	contaBruno := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome : "Bruno",
			CPF : "013.455.123-12",
			Profissao : "Desenvolvedor"},
			NumeroAgencia: 123,
			NumeroConta: 431,
		}

		contaBruno.Depositar(100)

		fmt.Println(contaBruno)
		fmt.Println(contaBruno.ObterSaldo())
	}

