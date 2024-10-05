package main

import (
	"fmt"
)
//Classe Carro Criada
type Carro struct{
	marca string
	modelo string
	ano int
}


func main() {
//Objetos da classe carro criados
carro1 := Carro{marca: "Fiat", modelo: "Uno", ano: 2011}
carro2 := Carro{marca: "Fiat", modelo: "Palio", ano: 2015}
carro3 := Carro{marca: "Lamborghini", modelo: "Urus", ano: 2023}


//Prints 
	fmt.Printf("Marca: %s Modelo: %s Ano: %d\n\n", carro1.marca, carro1.modelo, carro1.ano)
	fmt.Printf("Marca: %s Modelo: %s Ano: %d\n\n", carro2.marca, carro2.modelo, carro2.ano)
	fmt.Printf("Marca: %s Modelo: %s Ano: %d\n\n", carro3.marca, carro3.modelo, carro3.ano)
}
