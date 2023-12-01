package contas

import "clientes"

type ContaPoupanca struct {
	Titular clientes.Titular
	NumeroAgencia,
	NumeroConta , 
	Operacao int
	saldo float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso, o novo valor na conta é de:", c.saldo
	} else {
		return "Não é possivel depositar valor negativo, o valor atual é de:", c.saldo
	}	
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}