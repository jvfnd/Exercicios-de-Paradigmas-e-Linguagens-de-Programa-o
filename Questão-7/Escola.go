package main

import (
	"fmt"
)

type Professor struct{
		nome string
		materia string
}

type Escola struct{
		nome string
		professores []Professor
}


func (e *Escola) addProfessor(p Professor){
		e.professores = append(e.professores, p)
}
func (e Escola) ListarProfessores(){
		var i int = 1
		fmt.Printf("\nLista de Professores %s: \n", e.nome)		
		for _, p := range e.professores{
				fmt.Printf("%d - %s Matéria: %s\n", i, p.nome, p.materia)
				i+=1
		}
}



func main() {


	escola1 := Escola{nome:"Escola Boa"}
	escola2 := Escola{nome:"Escola Bacana"}

	professor1 := Professor{nome: "Leonardo", materia: "Matemática"}
	professor2 := Professor{nome: "Hugo", materia: "Ciência"}
	professor3 := Professor{nome: "Fernando", materia: "Química"}

	escola1.addProfessor(professor1)
	escola1.addProfessor(professor2)
	escola1.addProfessor(professor3)

	escola2.addProfessor(professor3)
	escola2.addProfessor(professor2)
	escola2.addProfessor(professor1)

	escola1.ListarProfessores()
	escola2.ListarProfessores()
	
}
