package main

import(
	"fmt"
)
type Calculadora struct{}

func (c Calculadora) soma(numero1 int, numero2 int){
	var resultado int = numero1 + numero2
	fmt.Printf("\nO resultado da soma desses números é: %d", resultado)
}
func (c Calculadora) soma2(numero1 int, numero2 int, numero3 int){
	var resultado int = numero1 + numero2 + numero3
	fmt.Printf("\nO resultado da soma desses números é: %d", resultado)
}
func (c Calculadora) soma3(numero1 int, numero2 int, numero3 int, numero4 int){
	var resultado int = numero1 + numero2 + numero3 + numero4
	fmt.Printf("\nO resultado da soma desses números é: %d", resultado)
}



func main(){

	calculadora := Calculadora{}

	calculadora.soma(1, 2)
	calculadora.soma2(1, 2, 3)
	calculadora.soma3(1, 2, 3, 4)
}