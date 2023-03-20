const http = require('http');
const express = require('express');
const mongodb = require('mongodb');

const { PORT, VIDEO_STORAGE_HOST, VIDEO_STORAGE_PORT, DB_HOST, DB_NAME } =
  process.env;

function connectDB() {
  return mongodb.MongoClient.connect(DB_HOST);
}

function videoStreamHandler(app, dbClient) {
  const db = dbClient.db(DB_NAME);
  const videosCollection = db.collection('videos');

  app.get('/test', (req, res)=>{
    console.log('GET /test', req);
    res.send('Live Reloading Test');
  });

  app.get('/video', (req, res) => {
    const videoId = new mongodb.ObjectId(req.query.id);
    console.log(`Request with Id ${videoId}`);
    videosCollection
      .findOne({ _id: videoId })
      .then(record => {
        if (!record) {
          console.log(`There is no data with id ${videoId}`);
          res.sendStatus(404);
          return;
        }

        console.log(`Start forwarding request with path ${record.videoPath}`);
        const forwardReq = http.request(
          {
            host: VIDEO_STORAGE_HOST,
            port: VIDEO_STORAGE_PORT,
            path: `/video?path=${record.videoPath}`,
            method: 'GET',
            headers: req.headers,
          },
          forwardRes => {
            res.writeHead(forwardRes.statusCode, forwardRes.headers);
            forwardRes.pipe(res);
          },
        );
        req.pipe(forwardReq);
      })
      .catch(err => {
        console.error('Database query failed');
        console.error((err && err.stack) || err);
        req.sendStatus(500);
      });
  });
}

function startServer() {
  return new Promise(resolve => {
    const app = express();
    connectDB().then(client => videoStreamHandler(app, client));
    app.listen(PORT, () => {
      console.log(`Streaming Service App listening on port ${PORT}`);
      resolve();
    });
  });
}

function main() {
  console.log(
    `Forwarding video requests to ${VIDEO_STORAGE_HOST}:${VIDEO_STORAGE_PORT}.`,
  );

  return startServer();
}

main()
  .then(() => console.log('Streaming Service App'))
  .catch(err => {
    console.error('Microservice failed to start');
    console.error((err && err.stack) || err);
  });
