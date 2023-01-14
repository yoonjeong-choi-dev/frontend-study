namespace Ch4Exercise5 {
  // Requirement : 타입이 다른 경우 비교하지 않는다
  function is<T>(target: T, ...values: [T, ...T[]]): boolean {
    // target vs ...values with deep comparison
    const targetJson = JSON.stringify(target);
    return values.every((value) => targetJson === JSON.stringify(value));
  }

  console.log(is('string', 'otherstring'));
  console.log(is('string', 'string', 'string'));
  console.log(is(12,12));
  console.log(is(12,42));
  //console.log(is(12,'12'));
  console.log(is([1], [1], [1]));
  console.log(is([1], [1,2,4]));
}