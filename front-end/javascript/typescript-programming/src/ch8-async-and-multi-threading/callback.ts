import { readFile } from 'fs';

namespace Ch8Callback {
  type Callback = (err: (NodeJS.ErrnoException | null), data: string) => void;

  const myReadFileCallback: Callback = (error, data) => {
    if (error) {
      console.error('error to read file', error);
      return;
    }
    console.log('success to read:\n', data);
  }

  readFile('./example.txt', { encoding: 'utf-8' }, myReadFileCallback);
}