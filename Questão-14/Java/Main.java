public class Main {

    public static void main(String[] args) {

        Configuracao config = Configuracao.getInstancia();
        Configuracao config2 = Configuracao.getInstancia();

        config.setTema("Claro");
        System.out.println(config2.getTema());
        config.setTema("Vermelho Escuro");
        System.out.println(config2.getTema());
    }
}