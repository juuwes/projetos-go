package contas

import "github.com/banco/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(saque float64) string {
	liberado := saque > 0 && saque <= c.saldo
	if liberado {
		c.saldo -= saque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}

}

func (c *ContaCorrente) Tranferir(transferencia float64, contaDestino *ContaCorrente) bool {
	if transferencia < c.saldo && transferencia > 0 {
		c.saldo -= transferencia
		contaDestino.Depositar(transferencia)
		return true

	} else {
		return false
	}

}

func (c *ContaCorrente) Depositar(deposito float64) (string, float64) {
	if deposito > 0 {
		c.saldo += deposito
		return "Deposito realizado com sucesso. Valor do deposito:", deposito
	} else {
		return "Valor de deposito invalido:", deposito
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
