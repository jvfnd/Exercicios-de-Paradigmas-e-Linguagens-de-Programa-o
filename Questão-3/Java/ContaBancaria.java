public class ContaBancaria {
    private double saldo;
    public String titular;

    //Métodos
    public void depositar(double valor) {
        this.saldo += valor;
    }
    public void sacar(double valor) {
        if (this.saldo < valor) {
            System.out.println("Valor sacado ultrapassa saldo de "+this.saldo+" da sua conta");
        }
        else{
            this.saldo -= valor;
        }
    }
    //Construtor
    public ContaBancaria(double saldo, String titular) {
        this.saldo = saldo;
        this.titular = titular;
    }

    //Get e sets
    public double getSaldo() {
        return saldo;
    }
    public void setSaldo(double saldo) {
        this.saldo = saldo;
    }

    public String getTitular() {
        return titular;
    }

    public void setTitular(String titular) {
        this.titular = titular;
    }
}
