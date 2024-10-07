package escola;

import java.util.ArrayList;
import java.util.List;

public class Escola {
    List<Professor> professores = new ArrayList<>();
    private String nome;

    public Escola(String nome) {
        this.nome = nome;
    }

    public void adicionarProfessor(Professor professor) {
        professores.add(professor);
    }
    public void removerProfessor(Professor professor) {
        professores.remove(professor);
    }
    public void listarProfessores() {
        System.out.print("Lista de Professores: \n");
        int i = 1;
        for (Professor professor : professores) {
            System.out.printf(i+" "+professor.getNome()+" Matéria: "+professor.getMateria()+"\n");
            i++;
        }
    }

    public String getNome() {
        return nome;
    }

    public void setNome(String nome) {
        this.nome = nome;
    }
}
