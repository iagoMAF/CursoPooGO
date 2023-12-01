package main
// go mod tidy - limpa os pacotes (limpa dps reaj)

import (
	"fmt"
	"contas"
) 

func main() {
	contaDaSilvia := contas.ContaCorrente{Titular:"Silvia", Saldo:300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo:100}

	status := contaDaSilvia.Transferir(200, &contaDoGustavo)
	fmt.Println(status) 
}