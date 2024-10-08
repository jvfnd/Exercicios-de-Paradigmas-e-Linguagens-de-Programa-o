package abstrato;

public class FuncionarioAssalariado extends Funcionario {

    public FuncionarioAssalariado(String nome) {
        super.nome = nome;
    }

    public void CalcularSalario(double ganhaporhora) {
        double salario = (8 * ganhaporhora) * 30;
        System.out.println("O Funcionário "+FuncionarioAssalariado.super.nome+" ganha: "+salario);
    }
}
