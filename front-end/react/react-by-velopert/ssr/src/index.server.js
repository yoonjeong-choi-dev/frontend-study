import express from 'express';
import path from 'path';
import fs from 'fs';

import ReactDomServer from 'react-dom/server';
import { StaticRouter } from 'react-router-dom/server';
import { applyMiddleware, createStore } from 'redux';
import { Provider } from 'react-redux';
import { composeWithDevTools } from 'redux-devtools-extension';
import thunk from 'redux-thunk';

import rootReducer from './redux';
import PreloadContext from './lib/PreloadContext';

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


const serverRender = async (req, res, next) => {
  const context = {};

  // 리덕스 설정 추가
  const store = createStore(
    rootReducer,
    composeWithDevTools(applyMiddleware(thunk))
  );

  // 리덕스 상태 초기화 작업을 위한 PreloadContext 상태 초기화
  // : 서버 측에서 렌더링에 필요한 데이터를 준비해서 리덕스 스토어에 미리 저장
  const preloadContext = {
    done: false,
    promises: []
  };

  const jsx = (
    <PreloadContext.Provider value={preloadContext}>
      <Provider store={store}>
        <StaticRouter location={req.url} context={context}>
          <App/>
        </StaticRouter>
      </Provider>
    </PreloadContext.Provider>
  );

  // PreloadContext 컨텍스트를 이용하는 Preloader 컴포넌트 내부 로직(스토어 업데이트)를 위해서
  // 정적인 페이지 렌더링을 1차적으로 함
  ReactDomServer.renderToStaticMarkup(jsx);
  // renderToStaticMarkup 렌더링 과정 중에 preloadContext 에 등록된 프로미스들을 처리
  try {
    await Promise.all(preloadContext.promises);
  } catch (e) {
    // 서버 측 에러로 처리
    return res.status(500);
  }
  preloadContext.done = true;

  // 서버 측에서 렌더링한 결과를 문자열로 변환
  const root = ReactDomServer.renderToString(jsx);

  // 리덕스 초기 상태를 스크립트로 주입하기 위해 문자열로 변환
  // => 브라우저에서는 window.__PRELOADED_STATE__ 객체를 통해서 스토어 초기값 사용
  const stateString = JSON.stringify(store.getState()).replace(/</g, '\\u003c');
  const stateScript = `<script>__YJ_PRELOADED_STATE__ = ${stateString}</script>`;

  res.end(createPage(root, stateScript));
};

// stateScript : 서버에서 렌더링하면서 업데이트된 스토어의 상태 문자열로 주입
function createPage(root, stateScript) {
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
      ${stateScript}
      <script src="${manifest.files['main.js']}"></script>
    </body>
    </html>
      `;
}

app.use(serverRender);

app.listen(7166, () => {
  console.log(`Running with port ${port}`);
});