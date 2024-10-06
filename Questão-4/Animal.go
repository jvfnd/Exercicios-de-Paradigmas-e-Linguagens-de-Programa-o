package main

import (
	"fmt"
)

type Animal struct {
	nome    string
	especie string
}
type Cachorro struct {
	Animal
}
type Gato struct {
	Animal
}

func (a Animal) EmitirSom() {
}
func (c Cachorro) EmitirSom() {
	fmt.Println("AuAu")
}
func (g Gato) EmitirSom() {
	fmt.Printf("Miau")
}

func main() {
	cachorro := new(Cachorro)
	cachorro.nome = "Cão"
	cachorro.especie = "Caninus Felizis"

	gato := new(Gato)
	gato.nome = "Gatinho"
	gato.especie = "Gatitus Felizis"

	fmt.Printf("Nome do Animal: %s\nEspécie do Animal: %s\nSom do Animal: ", gato.nome, gato.especie)
	gato.EmitirSom()
	fmt.Printf("\n\nNome do Animal: %s\nEspécie do Animal: %s\nSom do Animal: ", cachorro.nome, cachorro.especie)
	cachorro.EmitirSom()

}
