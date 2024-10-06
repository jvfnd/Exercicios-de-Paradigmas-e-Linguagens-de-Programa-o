public class Carro {
    private String marca;
    private String modelo;
    private int ano;
    private int velocidade;

//Métodos
    public void InformaCarro(){
        System.out.println("Marca: "+marca+"Modelo: "+modelo+"Ano: "+ano);
    }
    public void Acelerar(){
        this.velocidade = this.velocidade + 10;
        System.out.println("Acelerando para "+this.velocidade+"km/h\n");
    }
    public void Frear(){
        this.velocidade = this.velocidade - 10;
        System.out.println("Freando para "+this.velocidade+"km/h\n");
    }
    public void MostrarVelocidade(){
        System.out.println("Velocidade atual: "+this.velocidade+"km/h\n");
    }
//Construtor
    public Carro(String marca, String modelo, int ano, int velocidade) {
        this.marca = marca;
        this.modelo = modelo;
        this.ano = ano;
        this.velocidade = velocidade;
    }
// Get e Sets
    public String getMarca() {
        return marca;
    }
    public void setMarca(String marca) {
        this.marca = marca;
    }

    public String getModelo() {
        return modelo;
    }

    public void setModelo(String modelo) {
        this.modelo = modelo;
    }

    public int getAno() {
        return ano;
    }

    public void setAno(int ano) {
        this.ano = ano;
    }

    public int getVelocidade() {
        return velocidade;
    }

    public void setVelocidade(int velocidade) {
        this.velocidade = velocidade;
    }
}
