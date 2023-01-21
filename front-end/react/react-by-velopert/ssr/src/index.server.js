import express from 'express';
import path from 'path';
import fs from 'fs';

import ReactDomServer from 'react-dom/server';
import { StaticRouter } from 'react-router-dom/server';
import App from './App';

const port = 7166;
const app = express();

// static file server setting
// after building client code...
const manifest = JSON.parse(
  fs.readFileSync(path.resolve('./build/asset-manifest.json'), 'utf-8')
);
const staticServer = express.static(path.resolve('./build'), {
  index: false,
});
app.use(staticServer);


const serverRender = (req, res, next) => {
  const context = {};
  const jsx = (
    <StaticRouter location={req.url} context={context}>
      <App/>
    </StaticRouter>
  );

  // 서버 측에서 렌더링하고 브라우저로 응답 넘김
  const root = ReactDomServer.renderToString(jsx);
  res.end(createPage(root));
};

function createPage(root) {
  return `<!DOCTYPE html>
    <html lang="en">
    <head>
      <meta charset="utf-8" />
      <link rel="shortcut icon" href="/favicon.ico" />
      <meta
        name="viewport"
        content="width=device-width,initial-scale=1,shrink-to-fit=no"
      />
      <meta name="theme-color" content="#000000" />
      <title>SSR Tutorial</title>
      <link href="${manifest.files['main.css']}" rel="stylesheet" />
    </head>
    <body>
      <noscript>You need to enable JavaScript to run this app.</noscript>
      <div id="root">
        ${root}
      </div>
      <script src="${manifest.files['main.js']}"></script>
    </body>
    </html>
      `;
}

app.use(serverRender);

app.listen(7166, () => {
  console.log(`Running with port ${port}`);
});