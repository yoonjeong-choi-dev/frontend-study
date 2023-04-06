const express = require('express');
const path = require('path');
const http = require('http');

const { PORT, STREAMING_HOST, METADATA_HOST, HISTORY_HOST, UPLOAD_HOST } =
  process.env;

if (!STREAMING_HOST) {
  throw new Error(
    'Please specify the streaming service using environment variable STREAMING_HOST.',
  );
}

if (!METADATA_HOST) {
  throw new Error(
    'Please specify the metadata service name using environment variable METADATA_HOST.',
  );
}

if (!HISTORY_HOST) {
  throw new Error(
    'Please specify the history service name using environment variable HISTORY_HOST.',
  );
}

function setupHandlers(app) {
  app.set('views', path.join(__dirname, 'views'));
  app.set('view engine', 'hbs');

  app.use(express.static('public'));

  app.get('/', (req, res) => {
    // Get all videos
    http
      .request(
        {
          host: METADATA_HOST,
          path: '/videos',
          method: 'GET',
        },
        response => {
          let data = '';
          response.on('data', chunk => {
            data += chunk;
          });

          response.on('end', () => {
            res.render('video-list', { videos: JSON.parse(data).videos });
          });

          response.on('error', err => {
            console.error('Failed to get video list.');
            console.error(err || `Status code: ${response.statusCode}`);
            res.sendStatus(500);
          });
        },
      )
      .end();
  });

  app.get('/video', (req, res) => {
    const videoId = req.query.id;
    http
      .request(
        {
          host: METADATA_HOST,
          path: `/video?id=${videoId}`,
          method: 'GET',
        },
        response => {
          let data = '';
          response.on('data', chunk => {
            data += chunk;
          });

          response.on('end', () => {
            const metadata = JSON.parse(data).video;
            const video = {
              metadata,
              url: `/api/video?id=${videoId}`,
            };

            res.render('play-video', { video });
          });

          response.on('error', err => {
            console.error('Failed to get detail data of video.');
            console.error(err || `Status code: ${response.statusCode}`);
            res.sendStatus(500);
          });
        },
      )
      .end();
  });

  app.get('/api/video', (req, res) => {
    const forwardRequest = http.request(
      {
        host: STREAMING_HOST,
        path: `/video?id=${req.query.id}`,
        method: 'GET',
      },
      response => {
        res.writeHeader(response.statusCode, response.headers);
        response.pipe(res);
      },
    );
    req.pipe(forwardRequest);
  });

  app.get('/upload', (req, res) => {
    res.render('upload-video', {});
  });

  app.post('/api/upload', (req, res) => {
    const forwardRequest = http.request(
      {
        host: UPLOAD_HOST,
        path: '/upload',
        method: 'POST',
        headers: req.headers,
      },
      response => {
        res.writeHeader(response.statusCode, response.headers);
        response.pipe(res);
      },
    );

    req.pipe(forwardRequest);
  });

  app.get('/history', (req, res) => {
    http
      .request(
        {
          host: HISTORY_HOST,
          path: '/videos',
          method: 'GET',
        },
        response => {
          let data = '';
          response.on('data', chunk => {
            data += chunk;
          });

          response.on('end', () => {
            res.render('history', {
              videos: JSON.parse(data).videos,
            });
          });

          response.on('error', err => {
            console.error('Failed to get history of videos.');
            console.error(err || `Status code: ${response.statusCode}`);
            res.sendStatus(500);
          });
        },
      )
      .end();
  });
}

function startServer() {
  return new Promise(resolve => {
    const app = express();
    setupHandlers(app);

    const port = (PORT && parseInt(PORT, 10)) || 3000;
    app.listen(port, () => {
      console.log(`Gateway Service App listening on port ${port}`);
      resolve();
    });
  });
}

function main() {
  return startServer();
}

main()
  .then(() => console.log('Gateway Service App'))
  .catch(err => {
    console.error('Microservice failed to start.');
    console.error((err && err.stack) || err);
  });
