const express = require('express');
const mongodb = require('mongodb');
const amqp = require('amqplib');
const bodyParser = require('body-parser');

const { PORT, DB_HOST, DB_NAME, RABBIT_HOST } = process.env;

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

function connectRabbit() {
  return amqp.connect(RABBIT_HOST).then(connection => {
    console.log('Connected to RabbitMQ');
    return connection.createChannel();
  });
}

function setupHandler(microservice) {
  const videosCollection = microservice.db.collection('videos');

  microservice.app.get('/videos', (req, res) =>
    videosCollection
      .find()
      .toArray()
      .then(videos => {
        console.log('metadata for videos: ', videos);
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

  microservice.app.get('/video', (req, res) => {
    const videoId = new mongodb.ObjectId(req.query.id);
    return videosCollection
      .findOne({ _id: videoId })
      .then(video => {
        if (!video) {
          res.sendStatus(404);
        } else {
          res.json({ video });
        }
      })
      .catch(err => {
        console.error(`Fail to get video with id ${videoId} from database`);
        console.error((err && err.stack) || err);
        req.sendStatus(500);
      });
  });

  // handling video upload event
  function consumeVideoUploadedMessage(message) {
    console.log('Consume a uploaded message: ', message);
    const parsedMsg = JSON.parse(message.content.toString());

    const videoMeta = {
      _id: new mongodb.ObjectId(parsedMsg.video.id),
      name: parsedMsg.video.name,
    };

    return videosCollection.insertOne(videoMeta).then(() => {
      console.log('Acknowledging message was handled.');

      // 성공적으로 처리한 경우, 메시지 consume
      microservice.messageChannel.ack(message);
    });
  }

  return microservice.messageChannel
    .assertExchange('video-uploaded', 'fanout')
    .then(() => {
      console.log('Get a message exchange');
      return microservice.messageChannel.assertQueue('', { exclusive: true });
    })
    .then(response => {
      const queueName = response.queue;
      return microservice.messageChannel
        .bindQueue(queueName, 'video-uploaded', '')
        .then(() => {
          console.log('Get a binding queue with name "video-uploaded"');
          return microservice.messageChannel.consume(
            queueName,
            consumeVideoUploadedMessage,
          );
        });
    });
}

function startHTTPServer(dbConn, messageChannel) {
  return new Promise(resolve => {
    const app = express();
    const microservice = {
      app,
      db: dbConn.db,
      messageChannel,
    };

    app.use(bodyParser.json());

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

function startMicroservice(dbHost, dbName, rabbitHost) {
  return connectDB(dbHost, dbName).then(dbConn =>
    connectRabbit(rabbitHost).then(messageChannel =>
      startHTTPServer(dbConn, messageChannel),
    ),
  );
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

  if (!RABBIT_HOST) {
    throw new Error(
      'Please specify the RabbitMQ name using environment variable RABBIT_HOST.',
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
