process.env.BABEL_ENV = 'production';
process.env.NODE_ENV = 'production';

process.on('unhandledRejection', err => {
  throw err;
});

require('../config/env');

const fs = require('fs-extra');
const webpack = require('webpack');
const paths = require('../config/paths');

// 서버 웹팩 설정 파일
const config = require('../config/webpack.config.server');

function build() {
  console.log('Creating server build -------->');
  fs.emptyDirSync(paths.ssrBuild);

  const compiler = webpack(config);
  return new Promise((resolve, reject) => {
    compiler.run((err, stats) => {
      if (err) {
        console.error(err);
        return;
      }
      console.log(stats.toString());
    });
  });
}

build();
