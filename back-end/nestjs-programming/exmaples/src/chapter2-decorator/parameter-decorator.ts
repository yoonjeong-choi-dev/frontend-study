import { BadRequestException } from '@nestjs/common';

// parameter decorator
function MinLength(minLen: number) {
  return function (target: any, propertyKey: string, parameterIndex: number) {
    target.validators = {
      minLength: function (args: string[]) {
        return args[parameterIndex].length >= minLen;
      },
    };
  };
}

// method decorator
function Validate(
  target: any,
  propertyKey: string,
  descriptor: PropertyDescriptor,
) {
  const method = descriptor.value;
  descriptor.value = function (...args) {
    Object.keys(target.validators).forEach((key) => {
      if (!target.validators[key](args)) {
        throw new BadRequestException(`key: ${key}, args: ${args}`);
      }
    });
    method.apply(this, args);
  };
}

class UserService {
  private name: string;

  get Name() {
    return this.name;
  }

  @Validate
  setName(@MinLength(3) name: string) {
    this.name = name;
  }

  print() {
    console.log(`name: ${this.Name}`);
  }
}

const us = new UserService();
us.setName('Yoonjeong');
us.print();

try {
  us.setName('YJ');
  us.print();
} catch (e) {
  console.log('Error: ', e.message);
}
