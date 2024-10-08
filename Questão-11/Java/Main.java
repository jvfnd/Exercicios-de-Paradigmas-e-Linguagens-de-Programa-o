import abstrato.*;


public class Main {




    public static void main(String[] args) {

        FuncionarioHorista funcionario1 = new FuncionarioHorista("João");
        FuncionarioAssalariado funcionario2 = new FuncionarioAssalariado("Fernando");

        funcionario1.CalcularSalario(12, 20);
        funcionario2.CalcularSalario(20);

    }
}
