package main

import(
	"fmt"
)

type Matematica struct{}

func (m Matematica) CalcularFatorial(numero int){
	var resultado int = 1	
	for i := 1; i<= numero; i++{
		resultado *= i
	} 
	fmt.Printf("Fatoria de %d tem resultado de: %d", numero, resultado)
}


func main(){
	fatorial := Matematica{}

	fatorial.CalcularFatorial(7)
}