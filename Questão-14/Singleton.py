class Configuracao:
    __instance = None
    def __new__(cls):
        if Configuracao.__instance is None:
            Configuracao.__instance  = super().__new__(cls)
        return Configuracao.__instance
    
    def __init__(self):
        self.tema = "Escuro"
        
    
config = Configuracao()
config2 = Configuracao()

config.tema = "Claro"

print (config2.tema)

config.tema = "Escuro Avermelhado"

print (config2.tema)
        
    
        
