class ContaBancaria:
    def __init__(self, saldo, titular):
      self.__saldo = saldo
      self.titular = titular

    def depositar(self, valor):
      self.__saldo += valor

    def sacar(self, valor):
      if self.__saldo < valor:
        print(f"valor sacado ultrapassa o saldo de {self.__saldo} da sua conta\n")
      else:
        self.__saldo -= valor

    def get_saldo(self):
      return self.__saldo


conta = ContaBancaria(5000, "João")
conta.depositar(1000)
conta.sacar(7000)
conta.sacar(2000)
print(f"A conta de {conta.titular} tem saldo de {conta.get_saldo()}")
