public class Configuracao {
    private static Configuracao instancia;

    private Configuracao() {}

    public static Configuracao getInstancia() {
        if (instancia == null) {
            instancia = new Configuracao();
        }
        return instancia;
    }

    private String tema = "Escuro";

    public String getTema() {
        return tema;
    }

    public void setTema(String tema) {
        this.tema = tema;
    }

}
