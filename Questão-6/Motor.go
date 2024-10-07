package main

import (
	"fmt"
)

type Motor struct{}

func (m Motor) Ligar() {
	fmt.Println("Motor está Ligado")
}
func (m Motor) Desligar() {
	fmt.Println("Motor está Desligado")
}

type Carro struct {
	motor  Motor
	marca  string
	modelo string
}

func (c Carro) Ligar() {
	c.motor.Ligar()
	println("Carro está ligado")
}

func (c Carro) Desligar() {
	c.motor.Desligar()
	println("Carro está desligado")
}

func main() {
	carro := Carro{motor: Motor{}, marca: "Fiat", modelo: "Uno"}
	carro.Ligar()
	carro.Desligar()

}
