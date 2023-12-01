package main
// go mod tidy - limpa os pacotes (limpa dps reaj)
// go mod init "nome"

import (
	"fmt"
	"contas"
	// "clientes"
) 

func main() {
	 
	// contaDoBruno := contas.ContaCorrente{
	// 	Titular: clientes.Titular {
	// 		Nome: "Bruno", 
	// 		CPF: "123", 
	// 		Profissao: "Desenvolvedor"},
	// 	NumeroAgencia: 123, 
	// 	NumeroConta: 123456, 
	// 	Saldo: 100,
	// }

	// clienteBruno := clientes.Titular {"Bruno", "123.123.123.12", "Desenvolvedor GO"}
	// fmt.Println(clienteBruno)

	// contaBruno := contas.ContaCorrente{ clienteBruno, 123, 123456, 200}
	// fmt.Println(contaBruno)

	contaExmplo := contas.ContaCorrente{}
	contaExmplo.Depositar(100)
	
	fmt.Println(contaExmplo)
	fmt.Println(contaExmplo.ObterSaldo())
}