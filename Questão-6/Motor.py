class Motor:
  def ligar(self):
    print("Motor Ligado")
  def desligar(self):
    print("Motor Desligado")

class Carro:
  def __init__(self, marca, modelo):
    self.motor = Motor()
    self.marca = marca
    self.modelo = modelo

  def ligar(self):
    self.motor.ligar()
    print("Carro ligado")
  def desligar(self):
    self.motor.desligar()
    print("Carro Desligado")


carro = Carro("Fiat", "Uno")
carro.ligar()
carro.desligar()
