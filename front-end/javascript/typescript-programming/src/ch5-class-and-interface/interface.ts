namespace Ch5Interface {
  interface Animal {
    readonly name: string;

    eat(food: string): void;

    sleep(hours: number): void;
  }

  class Dog implements Animal {
    private static type = 'Dog';

    constructor(readonly name: string) {
    }

    eat(food: string): void {
      console.log(`[${ Dog.type }] ${ this.name } eats ${ food }`);
    }

    sleep(hours: number): void {
      console.log(`[${ Dog.type }] ${ this.name } sleeps at ${ hours }`);
    }

    bark() {
      console.log(`${this.name} barks!!!!`);
    }
  }

  function animalAct(animal: Animal, food: string, hours: number) {
    animal.eat(food);
    animal.sleep(hours);
  }

  const dog = new Dog('Cat');
  animalAct(dog, 'snacks', 13);
  dog.bark();

  console.log(typeof dog, typeof Dog);
}