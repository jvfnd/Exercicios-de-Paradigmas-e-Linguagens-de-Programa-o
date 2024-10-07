import carro.*;



public class Main {




    public static void main(String[] args) {
        Motor motor = new Motor();
        Carro carro = new Carro("Fiat", "Uno", motor);

        carro.Ligar();
        carro.Desligar();

    }
}
