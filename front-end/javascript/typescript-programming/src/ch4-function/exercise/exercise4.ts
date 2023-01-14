namespace Ch4Exercise4 {
  function call<ArgType extends [unknown, string, ...unknown[]], ReturnType>(
    f: (...args: ArgType) => ReturnType,
    ...args: ArgType
  ): ReturnType {
    return f(...args);
  }

  function testFunc(length: number, value: string): string[] {
    return Array.from({ length }, () => value);
  }

  console.log(call(testFunc, 10, 'hello'));
}
