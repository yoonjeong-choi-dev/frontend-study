const fs = require('fs');
const fsPromise = require('fs/promises');

function readWithCallback() {
  const fileName = 'test.txt';
  fs.writeFile(fileName, 'Hello World - Callback', () => {
    fs.readFile(fileName, 'utf8', (err, data) => {
      console.log(data);
    });
  });
}

async function readAsync() {
  const fileName = 'test-async.txt';
  await fsPromise.writeFile(fileName, 'Hello World - Promise');
  const data = await fsPromise.readFile(fileName, 'utf8');
  console.log(data);
}

// readWithCallback();
// readAsync();