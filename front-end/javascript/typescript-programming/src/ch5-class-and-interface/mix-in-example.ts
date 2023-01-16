namespace Ch5MinInExample {
  // 생성자 함수를 표현하는 타입
  type ClassConstructor<T> = new (...args: any[]) => T;

  // 믹스인 : 클래스 생성자를 인자로 받아 믹스인한 클래스 생성자 반환
  function withSimpleDebug<C extends ClassConstructor<{ getDebugValue(): object }>>(Class: C) {
    // 인자로 받은 클래스를 확장한 익명 클래스 생성자 반환
    return class extends Class {
      constructor(...args: any[]) {
        super(...args);
      }
    };
  }

  class HardToDebugUser {
    constructor(private id: number, private firstName: string, private lastName: string) {
    }

    getDebugValue() {
      return {
        id: this.id,
        name: `${this.firstName} ${this.lastName}`
      };
    }
  }

  // Mixin Class
  let User = withSimpleDebug(HardToDebugUser);
  let user = new User(123, 'Yoonjeong', 'Choi');
  console.log(JSON.stringify(user.getDebugValue()));
}