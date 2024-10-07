class Empregado:
  def __init__(self, nome, cargo, salario):
    self.nome = nome
    self.cargo = cargo
    self.salario = salario

class Empresa:
  def __init__(self, nome):
    self.empregados = []
    self.nome = nome

  def contratar(self, empregado):
    self.empregados.append(empregado)

  def demitir(self, empregado):
    self.empregados.remove(empregado)

  def listarEmpregados(self):
    print(f"Lista de empregados da empresa: {self.nome}")
    i = 1
    for empregado in self.empregados:
      print(f"\n{i} - Nome: {empregado.nome} Cargo: {empregado.cargo} Salário: {empregado.salario}")
      i+=1

empresa = Empresa("Kingsoft")

empregado1 = Empregado("João", "Analista de Sistemas", 5000)
empregado2 = Empregado("Maria", "Gerente de Projetos", 7000)
empregado3 = Empregado("Pedro", "Desenvolvedor Full Stack", 10000)

empresa.contratar(empregado1)
empresa.contratar(empregado2)
empresa.contratar(empregado3)

empresa.listarEmpregados()
