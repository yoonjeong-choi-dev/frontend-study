const express = require('express');
const azure = require('azure-storage');

const app = express();

const {
  PORT,
  AZURE_STORAGE_ACCOUNT_NAME,
  AZURE_STORAGE_ACCESS_KEY,
  AZURE_STORAGE_CONTAINER_NAME,
} = process.env;

if (!AZURE_STORAGE_ACCOUNT_NAME) {
  throw new Error(
    'Please specify the storage account name for the server with environment variable PORT.',
  );
}
if (!AZURE_STORAGE_ACCESS_KEY) {
  throw new Error(
    'Please specify the storage account name for the server with environment variable PORT.',
  );
}

if (!AZURE_STORAGE_CONTAINER_NAME) {
  throw new Error(
    'Please specify the storage container name for the server with environment variable PORT.',
  );
}

console.log(
  `Serving videos from Azure storage account ${AZURE_STORAGE_ACCOUNT_NAME}.`,
);

function createBlobService() {
  return azure.createBlobService(
    AZURE_STORAGE_ACCOUNT_NAME,
    AZURE_STORAGE_ACCESS_KEY,
  );
}

function uploadStreamToAzure(blobService, incomingStream, mimeType, filePath) {
  console.log('Uploading stream to Azure storage: ', filePath);
  return new Promise((resolve, reject) => {
    const stream = blobService.createWriteStreamToBlockBlob(
      AZURE_STORAGE_CONTAINER_NAME,
      filePath,
      {
        contentSettings: {
          contentType: mimeType,
        },
      },
    );

    stream.on('error', reject);
    incomingStream
      .pipe(stream)
      .on('error', reject)
      .on('end', () => {
        console.log('End to upload');
        resolve();
      })
      .on('finish', () => {
        console.log('Finish to upload');
        resolve();
      })
      .on('close', resolve);
  });
}

function streamVideoFromAzure(blobService, videoId, res) {
  return new Promise((resolve, reject) => {
    blobService.getBlobProperties(
      AZURE_STORAGE_CONTAINER_NAME,
      videoId,
      (err, props) => {
        if (err) {
          reject(err);
          return;
        }

        res.writeHead(200, {
          'Content-Length': props.contentLength,
          'Content-Type': 'video/mp4',
        });

        // stream from azure storage
        blobService.getBlobToStream(
          AZURE_STORAGE_CONTAINER_NAME,
          videoId,
          res,
          streamErr => {
            if (streamErr) {
              reject(err);
            } else {
              resolve();
            }
          },
        );
      },
    );
  });
}

// Streaming video
app.get('/video', (req, res) => {
  const videoId = req.query.id;
  const blobService = createBlobService();

  streamVideoFromAzure(blobService, videoId, res).catch(err => {
    console.error(
      `Error occurred getting video ${AZURE_STORAGE_CONTAINER_NAME}/${videoId} to stream.`,
    );
    console.error((err && err.stack) || err);
    res.sendStatus(500);
  });
});

// Upload video
app.post('/upload', (req, res) => {
  const videoId = req.headers.id;
  const mimeType = req.headers['content-type'];

  const blobService = createBlobService();
  uploadStreamToAzure(blobService, req, mimeType, videoId)
    .then(() => {
      console.log(`success to upload video ${videoId}`);
      res.sendStatus(200);
    })
    .catch(err => {
      console.error(
        `Error occurred uploading video ${AZURE_STORAGE_CONTAINER_NAME}/${videoId}`,
      );
      console.error((err && err.stack) || err);
      res.sendStatus(500);
    });
});

const port = (PORT && parseInt(PORT, 10)) || 3000;
app.listen(PORT, () => {
  console.log(`Azure Storage Service listening on port ${port}`);
});
