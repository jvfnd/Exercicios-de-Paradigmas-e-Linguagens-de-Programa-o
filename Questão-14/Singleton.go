package main

import (
	"fmt"
	"sync"
)

type Configuracao struct{
	tema string
}

var instance *Configuracao
var once sync.Once

func GetConfiguracao() *Configuracao{
	once.Do(func(){
		instance = &Configuracao{
			tema: "Escuro",
		}
	})
	return instance
}

func (c *Configuracao) AlterarTema(tema string){
	c.tema = tema

}

func main(){
	config := GetConfiguracao()
	config2 := GetConfiguracao()

	config.AlterarTema("Claro")

	fmt.Println(config2.tema)

	config.AlterarTema("Vermelho Escuro")

	fmt.Println(config2.tema)

}