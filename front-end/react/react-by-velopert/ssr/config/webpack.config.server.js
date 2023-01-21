const nodeExternals = require('webpack-node-externals');
const webpack = require('webpack');

const paths = require('./paths');

// css 관련 파일을 로더에 별도로 설정
const getCSSModuleLocalIdent = require('react-dev-utils/getCSSModuleLocalIdent');
const cssRegex = /\.css$/;
const cssModuleRegex = /\.module\.css$/;
const sassRegex = /\.(scss|sass)$/;
const sassModuleRegex = /\.module\.(scss|sass)$/;

// 환경 변수 주입
// => 현재 환경이 개발환경인지 아닌지 확인 가능
const getClientEnvironment = require('./env');
const env = getClientEnvironment(paths.publicUrlOrPath.slice(0, -1));

module.exports = {
  mode: 'production',
  entry: paths.ssrIndexJs,
  target: 'node',
  output: {
    path: paths.ssrBuild,
    filename: 'server.js',
    chunkFilename: 'js/[name].chunk.js',
    publicPath: paths.publicUrlOrPath,
  },
  module: {
    rules: [
      {
        oneOf: [
          // js 소스 코드 처리 : webpack.config.js line 406 참조
          {
            test: /\.(js|mjs|jsx|ts|tsx)$/,
            include: paths.appSrc,
            loader: require.resolve('babel-loader'),
            options: {
              customize: require.resolve(
                'babel-preset-react-app/webpack-overrides'
              ),
              presets: [
                [
                  require.resolve('babel-preset-react-app'),
                  {
                    runtime: 'automatic',
                  }
                ]
              ],
              plugins: [
                [
                  require.resolve('babel-plugin-named-asset-import'),
                  {
                    loaderMap: {
                      svg: {
                        ReactComponent:
                          '@svgr/webpack?-svgo,+titleProp,+ref![path]',
                      },
                    },
                  },
                ],
              ],
              cacheDirectory: true,
              cacheCompression: true,
              compact: false,
            }
          },

          // CSS 처리
          {
            test: cssRegex,
            exclude: cssModuleRegex,
            loader: require.resolve('css-loader'),
            options: {
              importLoaders: 1,
              modules: {
                // true 옵션을 주어야 output 으로 생성하지 않음
                exportOnlyLocals: true,
              },
            },
          },

          // CSS Module 처리
          {
            test: cssModuleRegex,
            loader: require.resolve('css-loader'),
            options: {
              importLoaders: 1,
              modules: {
                exportOnlyLocals: true,
                getLocalIdent: getCSSModuleLocalIdent,
              }
            }
          },

          // Sass 처리
          {
            test: sassRegex,
            exclude: sassModuleRegex,
            use: [
              {
                loader: require.resolve('css-loader'),
                options: {
                  importLoaders: 3,
                  modules: {
                    exportOnlyLocals: true,
                  }
                }
              },
              require.resolve('sass-loader'),
            ]
          },

          // Sass + CSS Module 처리
          {
            test: sassModuleRegex,
            use: [
              {
                loader: require.resolve('css-loader'),
                options: {
                  importLoaders: 3,
                  modules: {
                    exportOnlyLocals: true,
                    getLocalIdent: getCSSModuleLocalIdent,
                  }
                }
              },
              require.resolve('sass-loader'),
            ]
          },

          // url-loader
          {
            test: [/\.bmp$/, /\.gif$/, /\.jpe?g$/, /\.png$/],
            loader: require.resolve('url-loader'),
            options: {
              // output 으로 생성하지 않음
              emitFile: false,
              limit: 10000,
              // output 생성하지 않기 때문에 경로 설정
              name: 'static/media/[name].[hash:8].[ext]',
            },
          },

          // file-loader
          {
            loader: require.resolve('file-loader'),
            // 자바스크립트 및 html 리소스는 제외
            exclude: [/\.(js|mjs|jsx|ts|tsx)$/, /\.html$/, /\.json$/],
            options: {
              emitFile: false,
              name: 'static/media/[name].[hash:8].[ext]',
            },
          }
        ]
      }
    ]
  },
  resolve: {
    // 번들링 시에 임포트한 라이브러리 함께 번들링
    modules: ['node_modules']
  },
  externals: [
    // 서버를 위한 번들링이므로, 라이브러리들을 번들링할 필요 없음(브라우저가 안쓰고 서버에서 렌더링하기떄문)
    nodeExternals({
      allowlist: [/@babel/],
    })
  ],
  plugins: [
    new webpack.DefinePlugin(env.stringified), // 환경변수를 주입해줍니다.
  ],
};