// Method Decorator: Read/Edit the method definition
// logging & error handling
function HandleError() {
  return function (
    target: any,
    propertyKey: string,
    descriptor: PropertyDescriptor,
  ) {
    console.log('==================================');
    console.log('target:', target);
    console.log('propertyKey:', propertyKey);
    console.log('descriptor:', descriptor);
    console.log('==================================');

    const method = descriptor.value;
    // this binding 을 위해 args 받아옴
    descriptor.value = function (...args: any[]) {
      try {
        // 함수 실행을 데코레이터에서 실행
        method.apply(this, args);
      } catch (e) {
        // error handling
        console.log(`Handle Error :${e.message}`);
      }
    };
  };
}

class Greeter {
  private readonly name: string;

  constructor(name: string) {
    this.name = name;
  }

  @HandleError()
  hello() {
    console.log(`${this.name} says "Hello~~"`);
  }

  @HandleError()
  helloError() {
    throw new Error('Wrong Greet');
  }
}

const greeter = new Greeter('YJ');

greeter.hello();
greeter.helloError();
