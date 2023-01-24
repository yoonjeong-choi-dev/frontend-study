# Server Side Rendering Example

* `npm run eject` 통해서 웹팩 커스터마이징
* `config/webpack.config.server.js`: ssr 서버를 위한 웹팩 설정
* `scripts/build.server.js`: ssr 서버 빌드 스크립트
* `redux-thunk` 를 이용한 비동기 데이터 렌더링
  * 서버 사이드 렌더링 시, 스토어에 미리 비동기 데이터를 반영하는 작업 필요
  * `PreloadContext` 컴포넌트를 통해 해당 작업을 함
* 서버 사이드 스플리팅은 `Loadable Components` 사용