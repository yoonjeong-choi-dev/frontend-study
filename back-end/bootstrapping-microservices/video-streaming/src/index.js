const http = require('http');
const express = require('express');
const mongodb = require('mongodb');
const amqp = require('amqplib');

const {
  PORT,
  VIDEO_STORAGE_HOST,
  VIDEO_STORAGE_PORT,
  DB_HOST,
  DB_NAME,
  RABBIT_HOST,
} = process.env;

function connectDB() {
  return mongodb.MongoClient.connect(DB_HOST);
}

function connectRabbitWithSingleRecipient() {
  return amqp.connect(RABBIT_HOST).then(connection => {
    console.log('Connected to RabbitMQ');
    return connection.createChannel();
  });
}

function connectRabbit() {
  return amqp.connect(RABBIT_HOST).then(connection => {
    console.log('Connected to RabbitMQ');
    return connection
      .createChannel()
      .then(messageChannel =>
        messageChannel
          .assertExchange('viewed', 'fanout')
          .then(() => messageChannel),
      );
  });
}

function sendDirectMessage(message) {
  const postOptions = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
  };
  const requestBody = {
    message,
  };

  // http://{container_name} 는 토커 컴포즈의 DNS가 이름 해석
  const request = http.request('http://history/message', postOptions, res => {
    res.on('data', body => {
      console.log(`BODY: ${body}`);
    });
    res.on('close', () => {
      console.log('Sent a message to history microservice.');
    });
  });
  request.on('error', err => {
    console.error('Error for request');
    console.error((err && err.stack) || err);
  });

  request.write(JSON.stringify(requestBody));
  request.end();
}

function publishMessage(messageChannel, videoPath) {
  console.log('Start to publish message on "viewed" queue');

  const msg = { videoPath };
  messageChannel.publish('', 'viewed', Buffer.from(JSON.stringify(msg)));
}

function videoStreamHandler(app, dbClient, messageChannel) {
  const db = dbClient.db(DB_NAME);
  const videosCollection = db.collection('videos');

  app.get('/reload', (req, res) => {
    console.log('GET /reload');
    res.send('Live Reloading Test');
  });

  app.get('/message', (req, res) => {
    console.log('GET /message');
    const message = req.query.msg;
    sendDirectMessage(message ?? 'empty message');
    res.end();
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

        // Publish the viewed data to message queue
        publishMessage(messageChannel, record.videoPath);
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
    connectDB().then(client => {
      console.log('Success to connect with MongoDB');
      // return connectRabbitWithSingleRecipient().then(messageChannel => {
      return connectRabbit().then(messageChannel => {
        console.log('Success to connect with RabbitMQ');
        return videoStreamHandler(app, client, messageChannel);
      });
    });
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
