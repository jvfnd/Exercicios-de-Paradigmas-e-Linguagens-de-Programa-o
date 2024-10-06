package animal;

public class Cachorro extends Animal{

    public Cachorro(String nome, String especie) {
        super(nome, especie);
    }
    @Override
    public void EmitirSom(){
        System.out.print("AuAu");
    }

    @Override
    public String getNome() {
        return super.getNome();
    }

    @Override
    public void setNome(String nome) {
        super.setNome(nome);
    }

    @Override
    public String getEspecie() {
        return super.getEspecie();
    }

    @Override
    public void setEspecie(String especie) {
        super.setEspecie(especie);
    }
}
