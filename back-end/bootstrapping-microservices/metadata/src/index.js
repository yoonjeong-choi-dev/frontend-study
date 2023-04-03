const express = require('express');
const mongodb = require('mongodb');
const amqp = require('amqplib');

const { PORT, DB_HOST, DB_NAME } = process.env;

function connectDB(dbHost, dbName) {
  return mongodb.MongoClient.connect(dbHost, { useUnifiedTopology: true }).then(
    client => {
      const db = client.db(dbName);
      return {
        db,
        close: () => client.close(),
      };
    },
  );
}

function setupHandler(microservice) {
  const videosCollection = microservice.db.collection('videos');

  microservice.app.get('/videos', (req, res) =>
    videosCollection
      .find()
      .toArray()
      .then(videos => {
        res.json({
          videos,
        });
      })
      .catch(err => {
        console.error('Fail to get videos collection from database');
        console.error((err && err.stack) || err);
        req.sendStatus(500);
      }),
  );
}

function startHTTPServer(dbConn) {
  return new Promise(resolve => {
    const app = express();
    const microservice = {
      app,
      db: dbConn.db,
    };
    setupHandler(microservice);

    const port = PORT || 3000;
    const server = app.listen(port, () => {
      microservice.close = () =>
        // eslint-disable-next-line no-shadow
        new Promise(resolve => {
          server.close(() => {
            // Close the Express server.
            resolve();
          });
        }).then(() => dbConn.close());
      resolve(microservice);
    });
  });
}

function startMicroservice(dbHost, dbName) {
  return connectDB(dbHost, dbName).then(dbConn => startHTTPServer(dbConn));
}

function main() {
  if (!DB_HOST) {
    throw new Error(
      'Please specify the database host using environment variable DB_HOST.',
    );
  }

  if (!DB_NAME) {
    throw new Error(
      'Please specify the database name using environment variable DB_NAME.',
    );
  }

  return startMicroservice(DB_HOST, DB_NAME);
}

if (require.main === module) {
  // Start microservice
  main()
    .then(() => console.log('Metadata Management Service App'))
    .catch(err => {
      console.error('Microservice failed to start');
      console.error((err && err.stack) || err);
    });
} else {
  // For test, export functions to test
  module.exports = {
    startMicroservice,
  };
}
