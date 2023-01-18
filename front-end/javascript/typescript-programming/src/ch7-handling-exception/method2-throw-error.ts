namespace Ch7ThrowError {
  const readline = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout,
  });

  function isValidDate(date: Date) {
    return Object.prototype.toString.call(date) === '[object Date]'
      && !Number.isNaN(date.getDate());
  }

  // Custom Exception 정의
  // 특정 예외 처리가 더 명확해 짐
  class InvalidDateFormatError extends RangeError {
  }

  class DateIsInTheFutureError extends RangeError {
  }

  // 런타임 에러를 상황에 맞춰서 throw
  function parse(birthday: string): Date {
    const date = new Date(birthday);
    if (!isValidDate(date)) {
      throw new InvalidDateFormatError('Enter a date in the form YYYY/MM/DD');
    }
    if(date.getTime() > Date.now()) {
      throw new DateIsInTheFutureError('Future birthday date...');
    }
    return date;
  }

  readline.question(`Enter your birthday :`, (input: string) => {

    try {
      const date = parse(input);
      console.log(`Your birthday is ${ date.toISOString() }`);
    } catch (e) {
      // parse 에서 발생한 특정 예외 처리
      // 문제점 : parse 코드나 문서를 보지 않는 이상 어떤 예외가 발생하는지 사용하는 측에서는 알 수가 없음
      if (e instanceof InvalidDateFormatError) {
        console.log(e.message);
      } else if(e instanceof DateIsInTheFutureError) {
        console.log('curious... ', e.message);
      } else {
        console.log('.......?');
      }
    } finally {
      readline.close();
    }
  });
}