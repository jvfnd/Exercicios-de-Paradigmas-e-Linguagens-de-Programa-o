package contabancaria;

public class SaldoInsuficiente extends Exception {
    public SaldoInsuficiente() {
        super("Saldo Insuficiente para Sacar");
    }
}
