const express = require('express');
const mongodb = require('mongodb');
const bodyParser = require('body-parser');
const amqp = require('amqplib');

const { DB_HOST, DB_NAME, RABBIT_HOST } = process.env;

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

function connectDB() {
  return mongodb.MongoClient.connect(DB_HOST).then(client =>
    client.db(DB_NAME),
  );
}

function connectRabbit() {
  return amqp.connect(RABBIT_HOST).then(connection => {
    console.log('Connected to RabbitMQ');
    return connection.createChannel();
  });
}

function consumeViewedMessage(dbCollection, messageChannel, message) {
  const parsedMsg = JSON.parse(message.content.toString());
  console.log('Data in message: ', message.content.toString());

  return dbCollection
    .insertOne({
      videoId: parsedMsg.video.id,
      watched: new Date(),
    }) // Record the "view" in the database.
    .then(() => {
      console.log('Acknowledging message was handled.');

      // 성공적으로 처리한 경우, 메시지 consume
      messageChannel.ack(message);
    });
}

function setupHandler(app, db, messageChannel) {
  const messageCollection = db.collection('messages');
  const videosCollection = db.collection('videos');
  app.post('/message', (req, res) => {
    const { message } = req.body;
    messageCollection
      .insertOne({ message })
      .then(() => {
        console.log(`Save message '${message}' to history`);
        res.json({
          message: `[Echo] '${message}'`,
        });
      })
      .catch(err => {
        console.error(`Error for saving ${message} to db`);
        console.error((err && err.stack) || err);
        res.sendStatus(500);
      });
  });

  app.get('/videos', (req, res) => {
    videosCollection
      .find()
      .toArray()
      .then(videos => {
        res.json({ videos });
      })
      .catch(err => {
        console.error('Fail to get videos collection from database');
        console.error((err && err.stack) || err);
        req.sendStatus(500);
      });
  });

  // consume message
  // return messageChannel.assertQueue('viewed', {}).then(() => {
  //   console.log('assert "viewed" Queue');
  //   return messageChannel.consume('viewed', msg =>
  //     consumeViewedMessage(videosCollection, messageChannel, msg),
  //   );
  // });

  return messageChannel
    .assertQueue('viewed', 'fanout')
    .then(() => {
      console.log('Get a Message Exchange');
      return messageChannel.assertQueue('', { exclusive: true });
    })
    .then(response => {
      console.log('Get an anonymous queue');
      const queueName = response.queue;
      return messageChannel.bindQueue(queueName, 'viewed', '').then(() => {
        console.log('Get a binding queue with name "viewed"');
        return messageChannel.consume(queueName, msg =>
          consumeViewedMessage(videosCollection, messageChannel, msg),
        );
      });
    });
}

function startServer(db, messageChannel) {
  return new Promise(resolve => {
    const app = express();
    app.use(bodyParser.json());
    setupHandler(app, db, messageChannel);

    const port = (process.env.PORT && parseInt(process.env.PORT, 10)) || 3000;
    app.listen(port, () => {
      resolve();
    });
  });
}

function main() {
  return connectDB().then(db => {
    console.log('Success to connect with MongoDB');

    return connectRabbit().then(messageChannel => {
      console.log('Success to connect with RabbitMQ');
      return startServer(db, messageChannel);
    });
  });
}

main()
  .then(() => console.log('History Microservice'))
  .catch(err => {
    console.error('Microservice failed to start');
    console.error((err && err.stack) || err);
  });
