import java.util.Scanner;

public class Main{
    public static void main(String[] args) {
        Calculadora calculadoradesoma = new Calculadora();
        
        calculadoradesoma.soma(1, 2, 3);        

        calculadoradesoma.soma(1, 2);

        calculadoradesoma.soma(1, 2, 3, 4);
    }
}