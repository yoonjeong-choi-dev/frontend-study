namespace Ch4Generator {
  function* createFibonacciGenerator(): IterableIterator<[number, number]> {
    let cur = 0, next = 1;
    let steps = 0;
    while (true) {
      yield [steps, cur];
      [cur, next] = [next, cur + next];
      steps += 1;
    }
  }

  let fibonacciSeq = createFibonacciGenerator();
  console.log(fibonacciSeq.next());
  console.log(fibonacciSeq.next());
  console.log(fibonacciSeq.next());


  function generateConsecutive(upperBound: number): { [Symbol.iterator](): Generator<number, void, unknown> } {
    return {
      * [Symbol.iterator]() {
        for (let n = 1; n <= upperBound; n++) yield n;
      }
    }
  }

  const upTo10 = generateConsecutive(10);
  const itr = upTo10[Symbol.iterator]();
  console.log(itr.next());
  console.log(itr.next());

  let sum = 0;
  for (let n of upTo10) sum += n;
  console.log('sum :', sum);

  const copies = [...upTo10];
  console.log(copies);
}