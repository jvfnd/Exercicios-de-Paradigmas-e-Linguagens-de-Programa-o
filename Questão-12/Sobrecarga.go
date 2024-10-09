package main

import(
	"fmt"
)

type Produto struct {
	nome string
	preco float64
}

func SomadosProdutos(produto1 Produto, produto2 Produto ){
	var resultado float64 = produto1.preco + produto2.preco
	fmt.Printf("Preço total do %s + %s é de: %.2f", produto1.nome, produto2.nome, resultado)
}



func main(){

	produto1 := Produto{nome: "Shampoo", preco: 23.50}
	produto2 := Produto{nome: "Condicionador", preco: 20.50}

	SomadosProdutos(produto1, produto2)
}