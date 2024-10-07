package main

import (
  "fmt"
)


type AnimalInter interface {
  EmitirSom() string
}


type Animal struct {
  Nome    string
  Especie string
}

type Cachorro struct {
  Animal
}

func (c Cachorro) EmitirSom() string {
  return "AuAu"
}

type Gato struct {
  Animal
}

func (g Gato) EmitirSom() string {
  return "Miau"
}

type Vaca struct{
  Animal
}

func (v Vaca) EmitirSom() string {
  return "Muuh"
}

type GrupoDeAnimais struct {
  animais []AnimalInter
}


func (g *GrupoDeAnimais) AdicionarAnimal(animal AnimalInter) {
  g.animais = append(g.animais, animal)
}


func (g GrupoDeAnimais) EmitirSons() {
  for _, animal := range g.animais {
    fmt.Println("Som:", animal.EmitirSom())
  }
}

func main() {
  
  cachorro := Cachorro{Animal{"Cão", "Caninus Felizis"}}
  gato := Gato{Animal{"Gato", "Gatitus Felizis"}}
  vaca := Vaca{Animal{"Vaca", "Mimosa Felizis"}}
  
  grupo := GrupoDeAnimais{}
  grupo.AdicionarAnimal(cachorro)
  grupo.AdicionarAnimal(gato)
  grupo.AdicionarAnimal(vaca)


  grupo.EmitirSons()
}
