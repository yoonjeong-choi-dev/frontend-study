function uploadFile(file) {
  const uploadAPI = '/api/upload';
  fetch(uploadAPI, {
    method: 'POST',
    headers: {
      'File-Name': file.name,
      'Content-Type': file.type,
    },
    body: file,
  })
    .then(() => {
      const resultElements = document.getElementById('results');
      resultElements.innerHTML += `<div>${file.name}</div>`;

      const uploadInput = document.getElementById('uploadInput');
      uploadInput.value = null;
    })
    .catch(err => {
      console.error(`Fail to upload: ${file.name}`);
      console.error(err);

      const resultElements = document.getElementById('results');
      resultElements.innerHTML += `<div>Fail to upload: ${file.name}</div>`;
    });
}

function uploadFiles(files) {
  for (let i = 0; i < files.length; i += 1) {
    uploadFile(files[i]);
  }
}
