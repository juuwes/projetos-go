package main

import (
	"fmt"
	"github.com/banco/contas"
)

func Boleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDenis := contas.ContaPoupanca{}
	contaDenis.Depositar(100)
	contaDenis.Sacar(50)
	Boleto(&contaDenis, 40)
	fmt.Println(contaDenis.ObterSaldo())

	contaLuisa := contas.ContaCorrente{}
	contaLuisa.Depositar(200)
	Boleto(&contaLuisa, 50)
	fmt.Println(contaLuisa.ObterSaldo())

}

/*	clienteBruno := clientes.Titular{"Bruno", "123.123.111-88", "Estudante"}
	contaBruno := contas.ContaCorrente{clienteBruno, 123, 123456, 100}
	fmt.Println(contaBruno)


	/*contaBruno := contas.ContaCorrente{Titular: clientes.Titular{
		Nome: "Bruno",
		CPF: "123.1111.123-12",
		Profissao: "Estudate"},
		NumeroAgencia: 123, NumeroConta: 123456, Saldo: 100}

	fmt.Println(contaBruno)


/*	contaGuilherme := ContaCorrente1{titular: "Guilherme", numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}
	contaBruna := ContaCorrente1{"Bruna", 222, 111238, 200}

	fmt.Println(contaGuilherme)
	fmt.Println(contaBruna)

	var contaCris *ContaCorrente1
	contaCris =
	 new(ContaCorrente1)
	contaCris.titular = "Cris"

	fmt.Println(*contaCris) */
