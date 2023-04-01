const fs = require('fs');
const express = require('express');

const { PORT } = process.env;

function videoStreamHandler(app) {
  app.get('/test', (req, res) => {
    res.send('Test API');
  });

  app.get('/video', (req, res) => {
    // Route for streaming video.

    const videoPath = './localResource/sample_video.mp4';
    fs.stat(videoPath, (err, stats) => {
      if (err) {
        console.error('Error for reading local video');
        res.sendStatus(500);
        return;
      }

      res.writeHead(200, {
        'Content-Length': stats.size,
        'Content-Type': 'video/mp4',
      });

      fs.createReadStream(videoPath).pipe(res);
    });
  });
}

function startServer() {
  return new Promise((resolve, reject) => { // Wrap in a promise so we can be notified when the server has started.
    const app = express();
    videoStreamHandler(app);

    const port = PORT && parseInt(PORT, 10) || 3000;
    app.listen(port, () => {
      resolve();
    });
  });
}

function main() {
  return startServer();
}

main()
  .then(() => console.log('Streaming Service App'))
  .catch(err => {
    console.error('Microservice failed to start');
    console.error((err && err.stack) || err);
  });
