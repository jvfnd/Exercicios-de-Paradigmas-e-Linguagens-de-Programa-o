package main

import(
  "fmt"
)

type ContaBancaria struct{
  Titular string
  saldo float64
}

func (c *ContaBancaria)Depositar(valor float64){
  c.saldo += valor
}

func (c *ContaBancaria)Sacar(valor float64){
  if c.saldo < valor{
    fmt.Println("Saldo insuficiente para sacar\n")
  } else{
    c.saldo -= valor
  }
}

func (c ContaBancaria)MostrarSaldo() float64{
  return c.saldo  
}




func main() {
  conta := ContaBancaria{Titular: "João", saldo: 5000}
  conta.Depositar(1000)
  conta.Sacar(7000)
  conta.Sacar(2000)
  fmt.Printf("Saldo da Conta de %s é de %.2f", conta.Titular, conta.MostrarSaldo())
  
  
}
