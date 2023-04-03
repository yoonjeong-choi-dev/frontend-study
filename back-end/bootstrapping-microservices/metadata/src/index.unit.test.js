describe('metadata microservice unit test', () => {
  // ############################
  // Set up Mock
  // ############################
  // 해당 함수가 호출되었는지 여부 확인을 위한 Mock Function
  const mockListenFn = jest.fn((port, callback) => callback());
  const mockGetFn = jest.fn();

  // Mock the express module
  jest.doMock('express', () => () => ({
    listen: mockListenFn,
    get: mockGetFn,
  }));

  // Mock database collection.
  const mockVideosCollection = {};

  // Mock Database
  const mockDb = {
    collection: () => mockVideosCollection,
  };

  const mockMongoClient = {
    db: () => mockDb,
  };

  // Mock the mongodb module
  jest.doMock('mongodb', () => ({
    // Mock Mongodb module.
    MongoClient: {
      // Mock MongoClient.
      connect: async () => mockMongoClient,
    },
  }));

  // ############################
  // Import modules to test
  // ############################
  // eslint-disable-next-line global-require
  const { startMicroservice } = require('./index');

  // ############################
  // Test
  // ############################
  test('microservice starts web server on startup', async () => {
    await startMicroservice();

    expect(mockListenFn.mock.calls.length).toEqual(1);
    expect(mockListenFn.mock.calls[0][0]).toEqual(3000);
  });

  test('/videos route is handled', async () => {
    await startMicroservice();

    expect(mockGetFn).toHaveBeenCalled();

    const videosRoute = mockGetFn.mock.calls[0][0];
    expect(videosRoute).toEqual('/videos');
  });

  test('/videos route retrieves data via videos collection', async () => {
    await startMicroservice();

    const mockRequest = {};
    const mockJsonFn = jest.fn();
    const mockResponse = {
      json: mockJsonFn,
    };

    const mockRecord1 = {};
    const mockRecord2 = {};

    // Mock the function for mongo client
    mockVideosCollection.find = () => ({
      toArray: async () => [mockRecord1, mockRecord2],
    });

    // Call router handler
    const videosRouteHandler = mockGetFn.mock.calls[0][1];
    await videosRouteHandler(mockRequest, mockResponse);

    expect(mockJsonFn.mock.calls.length).toEqual(1);
    expect(mockJsonFn.mock.calls[0][0]).toEqual({
      videos: [mockRecord1, mockRecord2],
    });
  });

});
