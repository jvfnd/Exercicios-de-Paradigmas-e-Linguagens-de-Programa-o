class Escola:
  def __init__(self, nome):
    self.professores = []
    self.nome = nome

  def add_professor(self, professor):
    self.professores.append(professor)

  def Listar_professores(self):
    print(f"\nLista de Professores {self.nome}:")
    i = 1
    for professor in self.professores:
      print(f"{i} - {professor.nome} Matéria: {professor.materia}")
      i += 1

class Professor:
  def __init__(self, nome, materia):
    self.nome = nome
    self.materia = materia

escola1 = Escola("Escola Boa")
escola2 = Escola("Escola Bacana")

professor1 = Professor("Leonardo", "Matemática")
professor2 = Professor("Hugo", "Ciência")
professor3 = Professor("Fernando", "Química")

escola1.add_professor(professor1)
escola1.add_professor(professor2)
escola1.add_professor(professor3)
escola2.add_professor(professor3)
escola2.add_professor(professor2)
escola2.add_professor(professor1)

escola1.Listar_professores()
escola2.Listar_professores()
