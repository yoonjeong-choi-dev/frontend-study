namespace Ch6FunctionTypeVariance {
  class Animal {
  }

  class Bird extends Animal {
    fly() {
    }
  }

  class Eagle extends Bird {
    hunt() {
    }
  }

  // 1. 인자는 매개변수보다 구체적이어야 함
  // 매개변수 <- 인자 형태로 대입을 함
  // => 본질적으로는 구체적인 타입의 변수를 추상적인 타입의 변수에 할당 가능한가에 대한 질문
  const animal = new Animal();
  const bird = new Bird();
  const eagle = new Eagle();

  function actLikeBird(b: Bird): Bird {
    // 인자로 넘어오는 타입은 매개변수 Bird 타입보다 구체적이어야 함
    // => 구현 코드는 매개변수 타입으로 가정하고 구현되어 있기 때문에 매개변수보다 추상적이면 에러
    b.fly();
    return b;
  }

  //actLikeBird(animal);  // 매개변수가 더 추상적이어서 컴파일 에러
  actLikeBird(bird);
  actLikeBird(eagle);


  // 2. 함수 타입의 매개변수 예제
  // function test(callback: F) => test(g:G) 가능?
  // 매개변수는 F 타입, 인자는 G 타입이라고 가정 <=> F 타입을 쓸 수 있는 곳에 G를 쓸수 있는가
  // 인자는 매개변수보다 구체적 => G 타입은 F 타입보다 구체적
  // G 는 F 보다 추상적인 것들을 인자로 받아 더 구체적인 것을 반환
  // 조건 1: G 함수에 전달한 인자들을 F 도 받을 수 있어야 구현 블록에서 사용 가능
  // 조건 2: G 함수 반환 타입이 F 반환 타입보다 구체적이어야 구현 블록에서 F 반환 타입을 가정하고 사용 가능
  function afterActingLikeBird(act: (b: Bird) => Bird): void {
    // 매개변수 F를 사용하는 구현 부분에서는 Bird 타입만 고려함
    let before = new Bird();
    before.fly();

    // 조건 1 : 인자로 넘어온 act 함수(G)의 인자는 Bird 타입보다 추상적이어야 함
    // act 함수의 인자로 넘기기 위해서
    let after = act(before);

    // 조건 2 : act 함수(G)의 반환값은 Bird 타입보다 구체적이어야 함
    // => 반환 값을 Bird 로 가정하기 때문에 Bird 타입의 모든 기능을 사용 가능해야 함
    after.fly();
  }

  // Example 1 : 매개변수 타입이 추상적
  // => 추상적인 타입을 반환하지 않으면 OK
  //afterActingLikeBird((animal: Animal) => new Animal());
  afterActingLikeBird((animal: Animal) => new Bird());
  afterActingLikeBird((animal: Animal) => new Eagle());

  // Example 2 : 매개변수 타입 동일
  // => 추상적인 타입을 반환하지 않으면 OK
  //afterActingLikeBird((bird: Bird) => new Animal());
  afterActingLikeBird((bird: Bird) => new Bird());
  afterActingLikeBird((bird: Bird) => new Eagle());

  // Example 3 : 매개변수 타입이 구체적
  // => 모든 반환 타입에 대해서 FAIL
  //afterActingLikeBird((eagle: Eagle) => new Animal());
  //afterActingLikeBird((eagle: Eagle) => new Bird());
  //afterActingLikeBird((eagle: Eagle) => new Eagle());
}
