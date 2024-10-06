class Carro:
  def __init__(self, marca, modelo, ano, velocidade):
    self.marca = marca
    self.modelo = modelo
    self.ano = ano
    self.velocidade = velocidade

  def InformaCarro(self):
    print(f'Marca: {self.marca} Modelo: {self.modelo} Ano: {self.ano}\n')

  def Acelerar(self):
    self.velocidade += 10
    print(f'Acelerando para {self.velocidade}Km/h')

  def Frear(self):
    self.velocidade -= 10
    print(f"Freando para {self.velocidade}Km/h")

  def MostrarVelocidade(self):
    print(f"Velocidade atual: {self.velocidade}Km/h")




carro1 = Carro("Fiat", "Uno", 2011, 0)

carro1.InformaCarro()

carro1.Acelerar()
carro1.Acelerar()
carro1.Frear()
carro1.MostrarVelocidade()
