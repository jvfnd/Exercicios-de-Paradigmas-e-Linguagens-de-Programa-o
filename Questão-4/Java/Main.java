import animal.*;

public class Main {
    public static void main(String[] args) {
        Cachorro cachorro = new Cachorro("Cão", "Caninus Felizis");
        Gato gato = new Gato("Gatinho", "Gatitus Felizis");


        System.out.print("\nO Animal "+gato.getNome()+" da espécie "+gato.getEspecie()+" faz: ");
        gato.EmitirSom();
        System.out.print("\nO Animal "+cachorro.getNome()+" da espécie "+cachorro.getEspecie()+" faz: ");
        cachorro.EmitirSom();

    }
}
