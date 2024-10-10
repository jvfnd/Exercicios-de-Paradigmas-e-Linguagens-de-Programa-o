package main

import (
	"errors"
	"fmt"
)

type ContaBancaria struct{
  Titular string
  saldo float64
}

func (c *ContaBancaria)Depositar(valor float64){
  c.saldo += valor
}

func (c *ContaBancaria)Sacar(valor float64) {
  if c.saldo < valor{
    fmt.Println(errors.New("Erro: Saldo Insuficiente Para Sacar"))
	return 
  } 
  c.saldo -= valor
  }


func (c ContaBancaria)MostrarSaldo() float64{
  return c.saldo  
}




func main() {
  conta := ContaBancaria{Titular: "João", saldo: 5000}
  conta.Sacar(7000)

  
  
}