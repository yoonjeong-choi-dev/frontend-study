namespace Ch5Decorator {
  // 생성자 함수를 표현하는 타입
  type ClassConstructor<T> = new (...args: any[]) => T;

  // @serializable 데코레이터 구현
  type Payload = {
    [key: string]: string | number | Payload
  };

  // getPayload(): Payload 메서드가 있는 클래스를 확장한 클래스 반환
  function serializable<T extends ClassConstructor<{
    getPayload(): Payload
  }>>(Class: T) {
    return class extends Class {
      serialize() {
        return JSON.stringify(this.getPayload());
      }
    };
  }


  class UserResponseOrigin {
    constructor(
      private name: string,
      private age: number,
      private city: string = '',
      private phone: string = '') {
    }

    getPayload() {
      return {
        name: this.name,
        age: this.age,
        more: {
          city: this.city,
          phone: this.phone,
        }
      };
    }
  }

  @serializable
  class UserResponse {
    constructor(
      private name: string,
      private age: number,
      private city: string = '',
      private phone: string = '') {
    }

    getPayload() {
      return {
        name: this.name,
        age: this.age,
        more: {
          city: this.city,
          phone: this.phone,
        }
      };
    }
  }

  let res1 = new UserResponse('YJ', 31, 'Seoul', '010-1234-5678');
  let res2 = new UserResponseOrigin('YJ', 31, 'Seoul', '010-1234-5678');
  console.log(res1.getPayload());
  console.log(res2.getPayload());
  // 웹스톰 상 에러 => 무시
  console.log(res1.serialize());
}