import http from 'http';

const httpServer = http.createServer((req, res) => {
  console.log(req);

  const url = req.url;
  const method = req.method;
  if (method === 'GET') {
    switch (url) {
      case '/' :
        return res.end('Hello World!');
      case '/a':
        return res.end('Welcome to /a!')
      case '/b':
        return res.end('Welcome to /b!');
      default:
        res.end('Good Bye~');
    }
  }

  if (method === 'POST') {
    let body = [];
    req.on('data', (chunk) => {
      body.push(chunk);
    });
    req.on('end', () => {
      const params = Buffer.concat(body);
      console.log('body', params.toString());
      res.end(`You submitted these parameters: ${params.toString()}`);
    });
  }
});

const port = 8000;
httpServer.listen(port, () => {
  console.log(`Server started on ${port}`);
})