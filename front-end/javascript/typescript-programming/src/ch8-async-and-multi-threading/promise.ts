import { readFile } from 'fs';

namespace Ch8Promise {
  type Executor<T> = (
    resolve: (result: T) => void,
    reject?: (error: any) => void
  ) => void;

  type Status = 'pending' | 'fulfilled' | 'rejected';

  // TODO : Promise 구현하기
  class Promise<T> {
    private status: Status;
    private executor: Executor<T>;
    private result: T | undefined;

    constructor(private f: Executor<T>) {
      this.executor = f;
      this.status = 'pending';

      // run executor

    }

    private resolve(result: T) {
      this.result = result;
    }

    private reject(error: unknown) {
      this.status = 'rejected';
    }

    then<U>(cb: (result: T) => Promise<U>): Promise<U> {

    }

    catch<U>(cb: (error: unknown) => Promise<U>): Promise<U> {

    }
  }

  function appendAndReadFile(path: string, data: string) {

  }

}