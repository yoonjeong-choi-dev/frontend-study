const express = require('express');
const azure = require('azure-storage');

const app = express();

const { PORT, AZURE_STORAGE_ACCOUNT_NAME, AZURE_STORAGE_ACCESS_KEY } =
  process.env;

if (!PORT) {
  throw new Error(
    'Please specify the port number for the server with environment variable PORT.',
  );
}

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

console.log(
  `Serving videos from Azure storage account ${AZURE_STORAGE_ACCOUNT_NAME}.`,
);

function createBlobService() {
  return azure.createBlobService(
    AZURE_STORAGE_ACCOUNT_NAME,
    AZURE_STORAGE_ACCESS_KEY,
  );
}

app.get('/video', (req, res) => {
  const videoPath = req.query.path;
  const blobService = createBlobService();

  const containerName = 'videos';
  blobService.getBlobProperties(containerName, videoPath, (err, props) => {
    if (err) {
      console.error(
        `error for getting properties for video ${containerName}/${videoPath}`,
      );
      console.error((err && err.stack) || err);
      res.sendStatus(500);
      return;
    }

    res.writeHead(200, {
      'Content-Length': props.contentLength,
      'Content-Type': 'video/mp4',
    });

    // stream from azure storage
    blobService.getBlobToStream(containerName, videoPath, res, (streamErr) => {
      if (streamErr) {
        console.error(
          `error for streaming video ${containerName}/${videoPath}`,
        );
        console.error((streamErr && streamErr.stack) || streamErr);
        res.sendStatus(500);
      }
    });
  });
});

app.listen(PORT, () => {
  console.log(`Azure Storage Service listening on port ${PORT}`);
});
