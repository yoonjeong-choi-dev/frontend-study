// Class Decorator: Read/Write the class definition
// Add a property
function reportableClassDecorator<T extends { new (...args: any[]): any }>(
  constructor: T,
) {
  // 인자로 받은 클래스를 데코레이터 입힌 클래스 객체로 반환
  return class extends constructor {
    // 데코레이터를 적용한 클래스에 프로퍼티 추가
    reportingURL = 'https://github.com/yoonjeong-choi-dev';
  };
}

@reportableClassDecorator
class BugReport {
  type: 'report';
  title: string;

  constructor(title: string) {
    this.title = title;
  }
}

const bug = new BugReport('BUG!!!!');
console.log(bug);

// compile error
// console.log(bug.reportingURL);
