import animal.*;



public class Main {




    public static void main(String[] args) {
        Cachorro cachorro = new Cachorro("Cão", "Caninus Felizis");
        Gato gato = new Gato("Gatinho", "Gatitus Felizis");
        Vaca vaca = new Vaca("Vaquinha", "Mimosa Felizis");

        GrupoDeAnimais animais = new GrupoDeAnimais();

        animais.AdicionarAnimais(vaca);
        animais.AdicionarAnimais(cachorro);
        animais.AdicionarAnimais(gato);

        animais.ExibirSomDeAnimais();

    }
}
