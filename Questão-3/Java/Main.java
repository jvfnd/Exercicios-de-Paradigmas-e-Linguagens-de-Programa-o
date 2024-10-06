
public class Main {
    public static void main(String[] args) {
        ContaBancaria conta = new ContaBancaria(5000, "João");

        conta.depositar(1000);
        conta.sacar(7000);
        conta.sacar(2000);

        System.out.println("A conta de "+conta.titular+" tem saldo de "+conta.getSaldo());
    }
}
