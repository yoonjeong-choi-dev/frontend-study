namespace Ch7ReturnError {
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

  // 런타임 에러를 상황에 맞춰서 반환
  // => 타입스크립트가 명시적으로 반환 타입을 알려줌
  // but 예외를 던지는 경우에는 알려줄 방법이 없음 i.e 자바와 같은 문법을 지원하지 않음
  function parse(birthday: string): Date | InvalidDateFormatError | DateIsInTheFutureError {
    const date = new Date(birthday);
    if (!isValidDate(date)) {
      return new InvalidDateFormatError('Enter a date in the form YYYY/MM/DD');
    }
    if (date.getTime() > Date.now()) {
      return new DateIsInTheFutureError('Future birthday date...');
    }
    return date;
  }

  readline.question(`Enter your birthday :`, (input: string) => {
      // parse 반환값 타입을 타입스크립트가 알고 있어 컴파일 타임에 예외를 포함한 모든 상황 처리 가능
      console.log(input);
      const parseResult = parse(input);

      if (parseResult instanceof InvalidDateFormatError) {
        console.log(parseResult.message);
      } else if (parseResult instanceof DateIsInTheFutureError) {
        console.log('curious... ', parseResult.message);
      } else {
        console.log(`Your birthday is ${ parseResult.toISOString() }`);
      }
      readline.close();
    }
  );
}