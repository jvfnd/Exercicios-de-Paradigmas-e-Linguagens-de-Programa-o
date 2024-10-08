import imprimir.*;

public class Main{
    public static void main(String[] args) {
        Imprimivel contrato = new Contrato();
        Imprimivel relatorio = new Relatorio();
        

        contrato.Imprimir();
        relatorio.Imprimir();
    }
}