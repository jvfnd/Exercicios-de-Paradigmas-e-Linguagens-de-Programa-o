import escola.*;


public class Main {




    public static void main(String[] args) {
        Escola escola1 = new Escola("Escola Boa");
        Escola escola2 = new Escola("Escola Bacana");

        Professor professor1 = new Professor("Leonardo", "Matemática");
        Professor professor2 = new Professor("Hugo", "Ciência");
        Professor professor3 = new Professor("Fernando", "Química");

        escola1.adicionarProfessor(professor1);
        escola1.adicionarProfessor(professor2);
        escola1.adicionarProfessor(professor3);

        escola2.adicionarProfessor(professor1);
        escola2.adicionarProfessor(professor3);
        escola2.adicionarProfessor(professor2);

        escola1.listarProfessores();
        escola2.listarProfessores();
    }
}
