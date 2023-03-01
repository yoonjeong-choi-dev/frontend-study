const fs = require('fs');
const express = require('express');

const app = express();

if (!process.env.PORT) {
  throw new Error(
    'Please specify the port number for the server with environment variable PORT.',
  );
}
const port = process.env.PORT;

app.get('/video', (req, res) => {
  const path = './localResource/sample_video.mp4';
  fs.stat(path, (err, stats) => {
    if (err) {
      console.error('An error occurred while reading a local file');
      res.sendStatus(500);
      return;
    }

    res.writeHead(200, {
      'Content-Length': stats.size,
      'Content-Type': 'video/mp4',
    });
    fs.createReadStream(path).pipe(res);
  });
});

app.listen(port, () => {
  console.log(`Streaming Service App listening on port ${port}`);
});
