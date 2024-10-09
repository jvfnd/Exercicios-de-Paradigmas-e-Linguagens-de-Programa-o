class Produto:
    def __init__(self, nome, preco):
        self.preco = preco
        self.nome = nome
        
    

produto1 = Produto("Shampoo", 23.50)
produto2 = Produto("Condicionador", 20.50)

print(f"Preço total do {produto1.nome} + {produto2.nome} é de: {produto2.preco + produto1.preco}")
        
    