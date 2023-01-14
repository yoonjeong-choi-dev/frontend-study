namespace Ch4Generics {
  // 전체 호출 시그니처
  type Filter = {
    <T>(array: Array<T>, filter: (item: T) => boolean) : T[]
  }

  const myFilter: Filter = (array, filter) => {
    return array.filter(filter);
  }

  const arr = [1, 2, 3, 4, 5, 6];
  const filter1 = (num: number) => num % 2 == 0;

  // 타임 검사기가 T 를 number 로 채움
  console.log(myFilter(arr, filter1));

  // myFilter 타입을 통해 필터 함수의 타입이 자동으로 추론됨
  console.log(myFilter(arr, n => n % 2 == 0));

  // 단축 호출 시그니처
  type TFilter<T> = (array: Array<T>, filter: (item: T) => boolean) => T[];


  const numberFilter: TFilter<number> = (array, filter) => {
    return array.filter(filter);
  }

  console.log(numberFilter(arr, filter1));
  console.log(numberFilter(arr, num => num % 2 == 0));


  // 제너릭 타입 2개
  function arrayMap<T, U>(arr: T[], map:(value: T) => U) {
    const result: U[] = arr.map(map);
    return result;
  }

  const myMap = (num: number) => String(num);
  const stringArr = arrayMap(arr, myMap);
  console.log(stringArr);
}
