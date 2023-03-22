const express = require('express');
const mongodb = require('mongodb');
const bodyParser = require('body-parser');
const amqp = require('amqplib');

const { DB_HOST, DB_NAME, RABBIT_HOST } = process.env;

function connectDB(dbHost, dbName) {
  return mongodb.MongoClient.connect(dbHost).then(client => client.db(dbName));
}

function connectRabbit() {
  return amqp.connect(RABBIT_HOST).then(connection => {
    console.log('Connected to RabbitMQ');
    return connection.createChannel();
  });
}

function consumeViewedMessage(dbCollection, messageChannel, message) {
  console.log('Consume message: ', message);
  const parsedMsg = JSON.parse(message.content.toString());
  console.log('Data in message: ', message.content.toString());
  return dbCollection.insertOne({ videoPath: parsedMsg.videoPath }).then(() => {
    messageChannel.ack(message);
  });
}

function setupHandler(app, db, messageChannel) {
  const messageCollection = db.collection('messages');
  const videosCollection = db.collection('videosHistory');
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
        return messageChannel.consume('viewed', msg =>
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
  return connectDB(DB_HOST, DB_NAME).then(db => {
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
