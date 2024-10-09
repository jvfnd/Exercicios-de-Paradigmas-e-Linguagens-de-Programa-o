public class Matematica {
    public static void CalcularFatorial(int numero) {
        int resultado = 1;
        for (int i = 1; i <= numero; i++) {
            resultado *= i;
        }
        System.out.println("O Fatorial de "+numero+" tem resultado de: "+resultado);
    }
}
