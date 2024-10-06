import Classes.*;

public class Main{
  public static void main(String[] args) {
    Carro carro1 = new Carro("Fiat","Uno",2011);
    Carro carro2 = new Carro("Fiat","Palio",2015);
    Carro carro3 = new Carro("Lamborghini","Urus",2023);
    
    carro1.InformaCarro();
    carro2.InformaCarro();
    carro3.InformaCarro();
  }
}
