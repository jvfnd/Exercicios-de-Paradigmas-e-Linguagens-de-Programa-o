class Animal:
  def __init__(self, nome, especie):
    self.nome = nome
    self.especie = especie

  def EmitirSom(self):
    pass


class Cachorro(Animal):
  def EmitirSom(self):
    print("AuAu")

class Gato(Animal):
  def EmitirSom(self):
    print("Miau")

class Vaca(Animal):
  def EmitirSom(self):
    print("Muuh")

class GrupoDeAnimais(Animal):
  def __init__(self):
    self.animais = []
    
  def AdicionarAnimal(self, animal):
    self.animais.append(animal)
  
  def EmitirSomdeAnimais(self):
    for animal in self.animais:
      print(f"{animal.nome} faz:", end=' ')
      animal.EmitirSom()

    
cachorro = Cachorro("Cão", "Canis familiaris")
gato = Gato("Gatinho", "Felis catus")
vaca = Vaca("Vaca", "Mimosa Felizis")


listadeanimais = GrupoDeAnimais()

listadeanimais.AdicionarAnimal(cachorro)
listadeanimais.AdicionarAnimal(gato)
listadeanimais.AdicionarAnimal(vaca)

listadeanimais.EmitirSomdeAnimais()
