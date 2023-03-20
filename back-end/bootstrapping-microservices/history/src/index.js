const express = require('express');

function setupHandler(app) {}

function startServer() {
  return new Promise(resolve => {
    const app = express();
    setupHandler(app);

    const port = (process.env.PORT && parseInt(process.env.PORT, 10)) || 3000;
    app.listen(port, () => {
      resolve();
    });
  });
}

function main() {
  console.log('History Service Live Re-Loading Test');
  return startServer();
}

main()
  .then(() => console.log('History Microservice'))
  .catch(err => {
    console.error('Microservice failed to start');
    console.error((err && err.stack) || err);
  });
