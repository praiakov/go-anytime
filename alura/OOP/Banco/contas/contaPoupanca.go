package contas

import (
	clientes "example.com/cliente"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valordoSaque float64) string {
	podeSacar := valordoSaque > 0 && valordoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valordoSaque
		return "Saque realizado com sucesso"
	}

	return "saldo insuficiente"
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito

		return "Deposito realizado com sucesso", c.saldo
	}

	return "O valor do deposito é invalido", valorDeposito
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
