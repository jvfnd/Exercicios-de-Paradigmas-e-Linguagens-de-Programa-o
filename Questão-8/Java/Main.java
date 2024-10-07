import empresa.*;


public class Main {




    public static void main(String[] args) {
        Empresa empresa = new Empresa("Kingsoft");

        Empregado empregado1 = new Empregado("João", "Analista de Sistemas", 5000);
        Empregado empregado2 = new Empregado("Maria", "Gerente de Projetos", 7000);
        Empregado empregado3 = new Empregado("Pedro", "Desenvolvedor Full Stack", 10000);

        empresa.contratarEmpregados(empregado1);
        empresa.contratarEmpregados(empregado2);
        empresa.contratarEmpregados(empregado3);

        empresa.imprimirEmpregados();

    }
}
