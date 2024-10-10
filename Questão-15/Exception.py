class SaldoInsuficienteException(Exception):
    def __init__(self):
        self.mensagem = "Saldo insuficiente para Sacar"
        super().__init__(self.mensagem)

class ContaBancaria:
    def __init__(self, saldo, titular):
      self.__saldo = saldo
      self.titular = titular

    def depositar(self, valor):
      self.__saldo += valor

    def sacar(self, valor):
      if self.__saldo < valor:
        raise SaldoInsuficienteException
      else:
        self.__saldo -= valor

    def get_saldo(self):
      return self.__saldo


conta = ContaBancaria(5000, "João")
conta.sacar(7000)
