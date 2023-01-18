namespace Ch7ReturnNull {
  const readline = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout,
  });


  function isValidDate(date: Date) {
    return Object.prototype.toString.call(date) === '[object Date]'
          && !Number.isNaN(date.getDate());
  }

  // 예외는 null 반환
  function parse(birthday: string): Date | null {
    const date = new Date(birthday);
    return isValidDate(date) ? date : null;
  }

  readline.question(`Enter your birthday :`, (input: string) => {

    const date = parse(input);

    // null check 통해서 예외처리
    if(date) {
      console.log(`Your birthday is ${date.toISOString()}`);
    } else {
      // 해당 예외의 원인을 알 수 없음
      console.log('Error for parsing date');
    }

    readline.close();
  });
}