const express = require('express');
const fs = require('fs');
const path = require('path');

const app = express();

const { PORT } = process.env;

const storagePath = path.join(__dirname, '../tmp/storage');

console.log(`Serving videos from local storage ${storagePath}.`);

// Streaming video
app.get('/video', (req, res) => {
  const videoId = req.query.id;
  const localFilePath = path.join(storagePath, videoId);
  res.sendFile(localFilePath);
});

// Upload video
app.post('/upload', (req, res) => {
  const videoId = req.headers.id;
  const localFilePath = path.join(storagePath, videoId);

  const fileWriteStream = fs.createWriteStream(localFilePath);
  req
    .pipe(fileWriteStream)
    .on('error', err => {
      console.error(
        `Error occurred uploading video ${localFilePath} to stream.`,
      );
      console.error((err && err.stack) || err);

      res.sendStatus(500);
    })
    .on('finish', () => {
      res.sendStatus(200);
    });
});

const port = PORT || 3000;
app.listen(PORT, () => {
  console.log(`Local Storage Service listening on port ${port}`);
});
