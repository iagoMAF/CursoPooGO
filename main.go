package main
// go mod tidy - limpa os pacotes (limpa dps reaj)
// go mod init "nome"

import (
	"fmt"
	"contas"
	// "clientes" 
) 

func PagarBoleto(conta verificaConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificaConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	fmt.Println(contaDoDenis)

	PagarBoleto(&contaDoDenis, 60)
	fmt.Println(contaDoDenis)
}