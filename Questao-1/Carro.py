class Carro:
  def __init__(self, marca, modelo, ano):
    self.marca = marca
    self.modelo = modelo
    self.ano = ano


  def InformaCarro(self):
    print(f'Marca: {self.marca} Modelo: {self.modelo} Ano: {self.ano}\n')

carro1 = Carro("Fiat", "Uno", 2011)
carro2 = Carro("Fiat", "Palio", 2015)
carro3 = Carro("Lamborghini", "Urus", 2023)

carro1.InformaCarro()
carro2.InformaCarro()
carro3.InformaCarro()
