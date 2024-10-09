class Matematica:
    @staticmethod
    def CalcularFatorial(numero):
        resultado = 1
        for numero in range(1, numero+1):
            resultado *= numero
        print(f"Fatorial de {numero} tem resultado de: {resultado}")        

fatorial = Matematica()
fatorial.CalcularFatorial(7)