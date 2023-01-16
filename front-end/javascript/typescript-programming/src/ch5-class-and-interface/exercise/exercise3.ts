namespace Ch5Exercise3 {
  interface Shoe {
    purpose: string;
    print(): void;
  }

  class BalletFlat implements Shoe {
    purpose = 'dancing';

    print() {
      console.log('Dancing with ballet flat shoes');
    }
  }

  class Boot implements Shoe {
    purpose = 'woodcutting';

    print() {
      console.log('cutting woods with boots');
    }
  }

  class Sneaker implements Shoe {
    purpose = 'walking';

    print() {
      console.log('Walking with sneakers');
    }
  }

  type ShoeType = 'balletFlat' | 'boot' | 'sneaker';

  // 오버로드 함수 시그니처
  type ShoeCreator = {
    create(type: 'balletFlat'): BalletFlat;
    create(type: 'boot'): Boot;
    create(type: 'sneaker'): Sneaker;
  }


  let Shoe: ShoeCreator = {
    create(type: ShoeType): Shoe {
      switch (type) {
        case 'balletFlat':
          return new BalletFlat();
        case 'boot':
          return new Boot();
        case 'sneaker':
          return new Sneaker();
      }
    }
  };

  const shoes = [Shoe.create('balletFlat'), Shoe.create('boot'), Shoe.create('sneaker')];
  shoes.forEach(shoe => {
    if(shoe instanceof BalletFlat) {
      console.log(`BalletFlat : ${shoe.purpose}`);
    } else if(shoe instanceof  Boot) {
      console.log(`Boot : ${shoe.purpose}`);
    } else {
      console.log(`Sneaker : ${shoe.purpose}`);
    }
  })
}
