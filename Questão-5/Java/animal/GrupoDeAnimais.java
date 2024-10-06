package animal;

import java.util.ArrayList;
import java.util.List;

public class GrupoDeAnimais {
    List<Animal> animais = new ArrayList<>();

    //Adicionar Animais no Array
    public void AdicionarAnimais(Animal animal){
        animais.add(animal);
    }

    public void ExibirSomDeAnimais(){
        for(Animal animal : animais){
            System.out.print("\n"+animal.getNome()+" faz: ");
            animal.EmitirSom();
        }
    }
}
