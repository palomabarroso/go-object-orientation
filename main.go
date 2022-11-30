package main

import (
	"fmt"

	"github.com/palomabarroso/go-object-orientation/account"
	"github.com/palomabarroso/go-object-orientation/client"
)

type verificarConta interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func main() {
	contaDoGuilherme := account.ContaCorrente{
		Titular: client.Titular{
			Nome:      "Guilherme",
			CPF:       "123.111.123.12",
			Profissao: "Desenvolvedor",
		},
		NumeroAgencia: 589,
		NumeroConta:   123456,
	}
	contaDaBruna := account.ContaCorrente{
		Titular: client.Titular{
			Nome:      "Bruna",
			CPF:       "123.111.123.34",
			Profissao: "Desenvolvedor",
		},
		NumeroAgencia: 222,
		NumeroConta:   111222,
	}

	var contaDaCris *account.ContaCorrente
	contaDaCris = new(account.ContaCorrente)
	contaDaCris.Titular = client.Titular{
		Nome:      "Cris",
		CPF:       "123.111.222.33",
		Profissao: "Desenvolvedor Node",
	}

	fmt.Println(*contaDaCris)
	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)

	fmt.Println("Saque")
	fmt.Println(contaDaCris.Saldo())
	fmt.Println(contaDaCris.Sacar(400))
	fmt.Println(contaDaCris.Saldo())

	fmt.Println("Transferência")
	status := contaDaBruna.Transferir(200, contaDaCris)
	fmt.Println(status)
	fmt.Println(*contaDaCris)

	fmt.Println("Pagar Boleto")
	PagarBoleto(contaDaCris, 50)
	PagarBoleto(&contaDaBruna, 50)
}
