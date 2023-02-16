import { readFile } from 'fs';

namespace Ch8Exercise1 {
  // f: 인자 하나(ArgType)를 받아 ResultType 데이터를 반환하는 비동기 함수
  // => 비동기 함수의 인자를 받아 프로미스를 반환하는 함수 반환
  function promisify<ArgType, ResultType>(
    f: (arg: ArgType, cb: (error: unknown, result: ResultType | null) => void) => void
  ): (arg: ArgType) => Promise<ResultType> {
    // 비동기 함수의 인자를 받아 프로미스를 반환하는 함수
    return (arg: ArgType) => new Promise<ResultType>((resolve, reject) => {
      // 비동기 함수 실행
      f(arg, (error, result) => {
        if(error) {
          return reject(error);
        }

        if(result === null) {
          return reject(null);
        }

        return resolve(result);
      })
    });
  }

  // test
  const readFilePromise = promisify(readFile);
  readFilePromise('../example.txt')
    .then(data => console.log(data.toString()))
    .catch(err => console.error(err));
}