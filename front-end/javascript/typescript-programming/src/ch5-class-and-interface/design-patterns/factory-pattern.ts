namespace Ch5FactoryPattern {
  // 아래 Shoe 인터페이스는 타입으로서 타입 네임스페이스에 저장
  interface Shoe {
    purpose: string;
  }

  class BalletFlat implements Shoe {
    purpose = 'dancing';
  }

  class Boot implements Shoe {
    purpose = 'woodcutting';
  }

  class Sneaker implements Shoe {
    purpose = 'walking';
  }

  type ShoeType = 'balletFlat' | 'boot' | 'sneaker';

  // 컴패니언 객체 패턴
  // 아래 Shoe 변수는 값으로서 값 네임스페이스에 저장
  let Shoe = {
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
  //shoes.push(Shoe.create('test'));
  shoes.forEach(shoe => console.log(shoe.purpose));
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