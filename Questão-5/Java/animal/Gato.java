package animal;

public class Gato extends Animal{
    public Gato(String nome, String especie) {
        super(nome, especie);
    }
    @Override
    public void EmitirSom(){
        System.out.print("Miau");
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
