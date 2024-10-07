package empresa;

import java.util.ArrayList;
import java.util.List;

public class Empresa {
    private String nome;
    List<Empregado> empregados = new ArrayList<>();


    public Empresa(String nome) {
        this.nome = nome;
    }

    public void contratarEmpregados(Empregado empregado) {
        empregados.add(empregado);
    }

    public void demitirEmpregado(Empregado empregado) {
        empregados.remove(empregado);
    }
    public void imprimirEmpregados() {
        int i = 1;
        for (Empregado empregado : empregados) {
            System.out.printf("%d - Nome: %s \nCargo: %s \nSalário: %.2f\n\n", i, empregado.getNome(), empregado.getCargo(),
                    empregado.getSalario());
            i++;
        }
    }
}
