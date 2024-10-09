public class Produto {
    private String nome;
    private double preco;

    public double Somar(Produto produto, Produto produto2) {
        double resultado = produto.getPreco() + produto2.getPreco();
        return resultado;
    }

    public Produto(String nome, double preco) {
        this.nome = nome;
        this.preco = preco;
    }
    public String getNome() {
        return nome;
    }
    public void setNome(String nome) {
        this.nome = nome;
    }
    public double getPreco() {
        return preco;
    }
    public void setPreco(double preco) {
        this.preco = preco;
    }
}
