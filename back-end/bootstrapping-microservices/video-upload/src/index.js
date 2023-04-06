const http = require('http');
const express = require('express');
const amqp = require('amqplib');
const mongodb = require('mongodb');

const { PORT, RABBIT_HOST, STORAGE_HOST } = process.env;

if (!RABBIT_HOST) {
  throw new Error(
    'Please specify the RabbitMQ name using environment variable RABBIT_HOST.',
  );
}

if (!STORAGE_HOST) {
  throw new Error(
    'Please specify the video storage service name using environment variable STORAGE_HOST.',
  );
}

function connectRabbit() {
  return amqp.connect(RABBIT_HOST).then(connection => {
    console.log('Connected to RabbitMQ');
    return connection.createChannel();
  });
}

function publishMessage(messageChannel, videoMetadata) {
  console.log('Start to publish message on "video-uploaded" queue');

  const msg = { video: videoMetadata };
  messageChannel.publish(
    'video-uploaded',
    '',
    Buffer.from(JSON.stringify(msg)),
  );
}

function streamToHttpPost(incomingStream, uploadHost, uploadRoute, headers) {
  return new Promise((resolve, reject) => {
    const forwardRequest = http.request({
      host: uploadHost,
      path: uploadRoute,
      method: 'POST',
      headers,
    });

    incomingStream.on('error', reject);
    incomingStream
      .pipe(forwardRequest)
      .on('error', reject)
      .on('end', resolve)
      .on('finish', resolve)
      .on('close', resolve);
  });
}

function setupHandler(app, messageChannel) {
  app.post('/upload', (req, res) => {
    const fileName = req.headers['file-name'];
    const videoId = new mongodb.ObjectId();
    const newHeaders = {
      ...req.headers,
      id: videoId,
    };

    streamToHttpPost(req, STORAGE_HOST, '/upload', newHeaders)
      .then(() => {
        res.sendStatus(200);
      })
      .then(() => {
        publishMessage(messageChannel, {
          id: videoId,
          name: fileName,
        });
      })
      .catch(err => {
        console.error(`Error occurred uploading video ${fileName}`);
        console.error((err && err.stack) || err);
        res.sendStatus(500);
      });
  });
}

function startHTTPServer(messageChannel) {
  return new Promise(resolve => {
    const app = express();

    setupHandler(app, messageChannel);

    const port = (PORT && parseInt(PORT, 10)) || 3000;
    app.listen(port, () => {
      console.log(`Video Upload Service App listening on port ${port}`);
      resolve();
    });
  });
}

function main() {
  return connectRabbit().then(messageChannel =>
    startHTTPServer(messageChannel),
  );
}

main()
  .then(() => console.log('Video Upload Service App'))
  .catch(err => {
    console.error('Microservice failed to start');
    console.error((err && err.stack) || err);
  });
