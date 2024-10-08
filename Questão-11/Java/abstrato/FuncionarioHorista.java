package abstrato;

public class FuncionarioHorista extends Funcionario{

    public FuncionarioHorista(String nome) {
        super.nome = nome;
    }

    public void CalcularSalario(int horaspordia, double ganhaporhora) {
        double salario = (horaspordia * ganhaporhora) * 30;
        System.out.println("O Funcionário "+FuncionarioHorista.super.nome+" ganha: "+salario);
    }


}
