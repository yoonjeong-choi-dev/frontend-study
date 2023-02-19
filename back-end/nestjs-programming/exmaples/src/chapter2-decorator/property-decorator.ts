// Property Decorator: Read/Edit the property definition
// Add title to the property
function formatTitle(title: string) {
  return function (target: any, propertyKey: string): any {
    let value = target[propertyKey];

    function getter() {
      return `${title} ${value}`;
    }

    function setter(newVal: string) {
      value = newVal;
    }

    // 공식 문서에서는 반환값을 무시한다고 하나 정상 작동
    return {
      get: getter,
      set: setter,
      enumerable: true,
      configurable: true,
    };
  };
}

class Speaker {
  @formatTitle('[name]')
  name: string;

  @formatTitle('[word]')
  word: string;

  constructor(name: string, word: string) {
    this.name = name;
    this.word = word;
  }
}

const s = new Speaker('YJ', 'Hello~');
console.log(s.name);
console.log(s.word);
