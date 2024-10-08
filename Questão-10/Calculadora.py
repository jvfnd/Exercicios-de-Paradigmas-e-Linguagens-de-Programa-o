class Calculadora:
    def soma(self, numero1, numero2):
        print(f"O resultado da soma desses números é: {numero1 + numero2}")
        
    def soma2(self, numero1, numero2, numero3):
        print(f"O resultado da soma desses números é: {numero1 + numero2 + numero3}")

    def soma3(self, numero1, numero2, numero3, numero4):
        print(f"O resultado da soma desses números é: {numero1 + numero2 + numero3 + numero4}")
        
        
calculadora = Calculadora()

calculadora.soma(1, 2)
calculadora.soma2(1, 2, 3)
calculadora.soma3(1, 2, 3, 4)