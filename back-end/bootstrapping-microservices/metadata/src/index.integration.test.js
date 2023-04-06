const axios = require('axios');
const mongodb = require('mongodb');

async function loadDatabaseFixture(db, collectionName, data) {
  await db.dropDatabase();
  const collection = db.collection(collectionName);
  return collection.insertMany(data);
}

describe('metadata microservice integration test', () => {
  const BASE_URL = 'http://localhost:3000';
  const DB_HOST = 'mongodb://localhost:7166'; // Set the port in docker compose
  const DB_NAME = 'testdb';
  const RABBIT_HOST = 'amqp://guest:guest@rabbit:5672';

  // ############################
  // Import modules to test
  // ############################
  // eslint-disable-next-line global-require
  const { startMicroservice } = require('./index');

  // ############################
  // Set up HTTP Server
  // ############################
  let microservice;
  beforeAll(async () => {
    microservice = await startMicroservice(DB_HOST, DB_NAME, RABBIT_HOST);
  });
  afterAll(async () => {
    await microservice.close();
  });

  function httpGetHelper(route) {
    const url = `${BASE_URL}${route}`;
    return axios.get(url);
  }

  test('/videos route retrieves data via videos collection', async () => {
    // generate test data
    const data = [];
    for (let i = 0; i < 2; i += 1) {
      const id = new mongodb.ObjectId();
      const path = `test-video-${i}.mp4`;

      data.push({
        _id: id,
        videoPath: path,
      });
    }

    await loadDatabaseFixture(microservice.db, 'videos', data);

    const response = await httpGetHelper('/videos');
    expect(response.status).toEqual(200);

    const { videos } = response.data;
    expect(videos.length).toEqual(data.length);
    for (let i = 0; i < 2; i += 1) {
      // eslint-disable-next-line no-underscore-dangle
      expect(videos[i]._id).toEqual(data[i]._id.toString());
      expect(videos[i].videoPath).toEqual(data[i].videoPath);
    }
  });
});
