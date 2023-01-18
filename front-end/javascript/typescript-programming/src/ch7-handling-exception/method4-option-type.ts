namespace Ch7OptionType {
  // Some<T> 와 None 타입을 공유하는 인터페이스
  // Java Option 타입과 비슷한 역할
  interface Option<T> {
    // None 타입일 수도 있는 Option 에 대해서 연산
    operate<U>(f: (value: T) => None): None;
    operate<U>(f: (value: T) => Option<U>): Option<U>;

    // 옵션의 값을 가져옴
    // None 타입의 경우 value 반환
    getOrElse(value: T): T;
  }

  // 연산이 제대로 성공하여 값이 만들어진 상황 => 해당 값 사용이 가능
  class Some<T> implements Option<T> {
    constructor(private value: T) {
    }

    // 연산자 오버로딩
    operate<U>(f: (value: T) => None): None
    operate<U>(f: (value: T) => Some<T>): Some<T>
    operate<U>(f: (value: T) => Option<U>): Option<U> {
      return f(this.value);
    }

    getOrElse(): T {
      return this.value;
    }
  }

  // 연산이 실패한 상황 : 예외 발생과 비슷한 상황
  // => 연산에 대한 결과 값 사용 불가능
  class None implements Option<never> {
    operate(): None {
      return this;
    }

    getOrElse<U>(value: U): U {
      return value;
    }
  }

  // 컴패니언 패턴 : Option 값에 대한 정의 => Option 타입을 생성
  function Option<T>(value: null | undefined): None
  function Option<T>(value: T): Some<T>
  function Option<T>(value: T): Option<T> {
    if(!value) return new None;
    return new Some(value);
  }

  // 특정 연산(여기서는 flatMap)의 성공/실패와 상관없이 연쇄적으로 연산 수행 가능
  console.log(Option(5));
  console.log(Option(5).operate(n => Option(n*3)));
  console.log(Option(5).operate(n => Option(n*3)).operate(n => new None()));
  console.log(Option(5).operate(n => Option(n*3)).operate(n => new None()).getOrElse(7166));
}