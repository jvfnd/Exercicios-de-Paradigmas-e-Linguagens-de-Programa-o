from typing import Protocol

class Imprimivel(Protocol):
    def Imprimir(self) -> None:
        pass

class Relatorio(Imprimivel):
    def Imprimir(self) -> None:
        print("Imprimindo Relatório")

class Contrato(Imprimivel):
    def Imprimir(self) -> None:
        print("Imprimindo Contrato")

relatorio: Imprimivel = Relatorio()

contrato: Imprimivel = Contrato()

relatorio.Imprimir()

contrato.Imprimir()
