namespace Ch4CallSignature {
  type Log = (msg: string, userId?: string, date?: string) => void;

  // 문맥접 타입화 : Log 타입임을 명시했기 때문에 매개변수 및 반환 타입 지정할 필요 없음
  const myLogger: Log = (message, userId='anonymous', date) => {
    let time = date ?? new Date().toDateString();
    console.log(`[${ time }] ${ userId }: ${ message }`);
  }

  myLogger('Hello');
  myLogger('Hi~', 'YJ');
  myLogger('Bye', 'Yoonjeong', '2023-01-10');
}