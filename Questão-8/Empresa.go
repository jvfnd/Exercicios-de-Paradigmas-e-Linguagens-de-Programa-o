package main

import (
	"fmt"
)

type Empregado struct {
	nome string
	cargo string
	salario float64
}

type Empresa struct{
	nome string
	empregados []Empregado
}

func (e *Empresa) Contratar(empregado Empregado){
	e.empregados = append(e.empregados, empregado)
}
func (e Empresa) Demitir(empregado Empregado){
	for i, emp := range e.empregados{
		if emp.nome == empregado.nome{
			e.empregados = append(e.empregados[:i], e.empregados[i+1:]...)
			return
		}
	}
}

func (e Empresa) ImprimirEmpregados(){
	fmt.Printf("Lista de Empregados da empresa %s\n", e.nome)
	var i = 1
	for _, empregado := range e.empregados{
		fmt.Printf("%d - Nome: %s \nCargo: %s \nSalario: %.2f\n\n", i, empregado.nome, empregado.cargo, empregado.salario)
		i+=1
	}
}


func main() {
	empresa := Empresa{nome: "Kingsoft"}

	empregado1 := Empregado{nome: "João", cargo: "Analista de Sistemas", salario:  5000.00} 
	empregado2 := Empregado{nome: "Maria", cargo: "Gerente de Projetos", salario: 7000.00}
	empregado3 := Empregado{nome: "João", cargo: "Desenvolvedor Full Stack", salario: 10000.00}

	empresa.Contratar(empregado1)
	empresa.Contratar(empregado2)
	empresa.Contratar(empregado3)

	empresa.ImprimirEmpregados()

	
}
