import contabancaria.*;

public class Main {

    public static void main(String[] args) throws SaldoInsuficiente {

        ContaBancaria conta = new ContaBancaria(5000, "João");

        conta.sacar(7000);
    }
}
