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

cachorro = Cachorro("Cão", "Canis familiaris")
gato = Gato("Gatinho", "Felis catus")

print(f"\nO Animal {cachorro.nome} da espécie {cachorro.especie} faz:")
cachorro.EmitirSom()
print(f"\nO Animal {gato.nome} da espécie {gato.especie} faz:")
gato.EmitirSom()
