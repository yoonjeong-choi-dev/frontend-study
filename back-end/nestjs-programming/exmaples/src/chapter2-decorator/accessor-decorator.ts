// Accessor Decorator: Read/Edit the accessor(getter/setter) definition
// Set getter/setter to enumerable/non-enumerable
function Enumerable(enumerable: boolean) {
  return function (
    target: any,
    propertyKey: string,
    descriptor: PropertyDescriptor,
  ) {
    descriptor.enumerable = enumerable;
  };
}

class User {
  constructor(private name: string) {}

  @Enumerable(true)
  get Name() {
    return this.name;
  }

  // 같은 이름의 accessor 에는 동일한 데코레이터 적용하면 안됨
  @Enumerable(false)
  set setName(name: string) {
    this.name = name;
  }
}

const u = new User('YJ');
console.log(`getName -> ${u.Name}`);

u.setName = 'Yoonjeong Choi';
console.log(`after setter -> ${u.Name}`);

console.log('\nPrint all enumerable property');
for (const key in u) {
  console.log(`${key}: ${u[key]}`);
}
