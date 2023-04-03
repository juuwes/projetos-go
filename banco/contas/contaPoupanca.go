package contas

import "github.com/banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(saque float64) string {
	liberado := saque > 0 && saque <= c.saldo
	if liberado {
		c.saldo -= saque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(deposito float64) (string, float64) {
	if deposito > 0 {
		c.saldo += deposito
		return "Deposito realizado com sucesso. Valor do deposito:", deposito
	} else {
		return "Valor de deposito invalido:", deposito
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
