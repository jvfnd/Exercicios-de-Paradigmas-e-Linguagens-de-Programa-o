package main

import(
	"fmt"
)

type Funcionario interface{
	CalcularSalario()
}
type BaseFuncionario struct{
	nome string
}

type FuncionarioHorista struct{
	BaseFuncionario
}
type FuncionarioAssalariado struct{
	BaseFuncionario
}

func (f FuncionarioHorista) CalcularSalario(ganhaporhora float64, horastrabalhadas int){
	var salario float64 = (ganhaporhora * float64(horastrabalhadas)) * 30
	fmt.Printf("\nO Funcionário %s ganha %.2f", f.nome, salario)
}

func (f FuncionarioAssalariado) CalcularSalario(ganhaporhora float64){
	var salario float64 = (ganhaporhora * 8) * 30
	fmt.Printf("\nO Funcionário %s ganha %.2f", f.nome, salario)
}




func main(){
	funcionario1 := FuncionarioHorista{BaseFuncionario: BaseFuncionario{nome:"João"}}
	funcionario2 := FuncionarioAssalariado{BaseFuncionario: BaseFuncionario{nome:"Fernando"}}

	funcionario1.CalcularSalario(20, 12)
	funcionario2.CalcularSalario(20)

}