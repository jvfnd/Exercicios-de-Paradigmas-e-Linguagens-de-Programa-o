public class Main {

    public static void main(String[] args) {
        Produto produto1 = new Produto("Shampoo", 23.50);
        Produto produto2 = new Produto("Condicionador", 20.50);

        System.out.println("Preço total do "+produto1.getNome()+" + "+produto2.getNome()+" é de: "
                +produto1.Somar(produto1, produto2));



    }
}
