package main

import(
	"fmt"
)

type Imprimivel interface{
	Imprimir()
}

type Relatorio struct{
}

type Contrato struct{
}

func (r Relatorio) Imprimir(){
	fmt.Printf("\nImprimindo Relatório")
}
func (c Contrato) Imprimir(){
	fmt.Printf("\nImprimindo Contrato")
}

func main(){
	
	contrato := Contrato{}
	relatorio := Relatorio{}

	contrato.Imprimir()
	relatorio.Imprimir()
}
