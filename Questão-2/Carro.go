package main

import (
	"fmt"
)
//Classe Carro Criada
type Carro struct{
	marca string
	modelo string
	ano int
	velocidade int
}

//Método InformaCarro, que retorna os dados do carro
func(c Carro)InformaCarro(){
	fmt.Printf("Marca: %s Modelo: %s Ano: %d\n\n", c.marca, c.modelo, c.ano)
}

//Método Acelera, que aumenta a velocidade do carro
func(c *Carro)Acelerar(){
	c.velocidade += 10
	fmt.Printf("Acelerando para %d km/h\n\n", c.velocidade)
}
//Método Freia, que diminui a velocidade do carro
func(c *Carro)Frear(){
	c.velocidade -= 10
	fmt.Printf("Freando para %d km/h\n\n", c.velocidade)
}

//Método Velocidade, que exibe a velocidade do carro
func(c Carro)MostraVelocidade(){
	fmt.Printf("Velocidade: %dkm/h\n\n", c.velocidade)
}

func main() {
//Objetos da classe carro criados
carro1 := Carro{marca: "Fiat", modelo: "Uno", ano: 2011, velocidade: 0}

//Exibe Detalhes
carro1.InformaCarro()
//Acelera
carro1.Acelerar()
carro1.Acelerar()
//Freia
carro1.Frear()	
//Mostra Velocidade
carro1.MostraVelocidade()
