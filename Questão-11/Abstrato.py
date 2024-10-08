from abc import ABC

class Funcionario(ABC):
    def __init__(self, nome):
        self.nome = nome
    
    def CalcularSalario(self):
        pass
    
class FuncionarioHorista(Funcionario):
    def CalcularSalario(self, horastrabalhadas, ganhaporhora):
        print(f"O Funcionario ganha {(horastrabalhadas * ganhaporhora) * 30}")
        
class FuncionarioAssalariado(Funcionario):
    def CalcularSalario(self, ganhaporhora):
        print(f"O Funcionario ganha {(ganhaporhora * 8) * 30}")
        

funcionario1 = FuncionarioAssalariado("João")
funcionario2 = FuncionarioHorista("Fernando")

funcionario1.CalcularSalario(20)
funcionario2.CalcularSalario(12, 20)