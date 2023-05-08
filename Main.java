interface Bird {
    void fly();
}

interface Human {
    void goToWork();
}

class Superhero implements Bird, Human {
    public void fly() {
        System.out.println("Superhero can fly!");
    }

    public void goToWork() {
        System.out.println("Superhero going to work!");
    }
}

class Superman extends Superhero {
    public void fly() {
        System.out.println("I fly and catch Lex Luthor!");
    }    
}

class Summariser {
    public void show(Superhero hero) {
        hero.fly();
        hero.goToWork();
    }
}


public class Main {
    public static void main(String[] args) {
        Summariser summariser = new Summariser();

        Superhero superhero = new Superhero();
        summariser.show(superhero);

        System.out.println("-----");

        Superman superman = new Superman();
        summariser.show(superman);
    }

}
