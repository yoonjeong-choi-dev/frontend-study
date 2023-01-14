namespace Ch4FunctionDefinition {
  console.log('매개변수 선언, 기본/선택 매개변수');
  function logger(message: string, userId = 'anonymous', date?: string) {
    let time = date ?? new Date().toDateString();
    console.log(`[${ time }] ${ userId }: ${ message }`);
  }

  logger('Hello');
  logger('Hi~', 'YJ');
  logger('Bye', 'Yoonjeong', '2023-01-10');

  console.log('나머지 매개변수');
  function sumAndAverage(...numbers: number[]): [number, number] {
    if (numbers.length == 0) return [0, 0];
    let sum = numbers.reduce((acc, num) => acc + num, 0);
    return [sum, sum / numbers.length]
  }

  console.log(sumAndAverage());
  console.log(sumAndAverage(1,2,3,4));

  console.log('this in the function');
  function displayDate(this: Date) {
    console.log(`${this.getDate()} - ${this.getMonth()} - ${this.getFullYear()}`);
  }

  // compile error
  //displayDate();  // TS2684: The 'this' context of type 'void' is not assignable to method's 'this' of type 'Date'.
  //displayDate(new Date());  // TS2554: Expected 0 arguments, but got 1.
  displayDate.apply(new Date());
  displayDate.call(new Date());
}
