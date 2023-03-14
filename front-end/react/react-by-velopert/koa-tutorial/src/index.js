const Koa = require('koa');
const Router = require('koa-router');
const bodyParser = require('koa-bodyparser');

const api = require('./api');

const { getTitle } = require('./utils');

const app = new Koa();
const router = new Router();

// #####################
// Router Setting
// #####################
router.get('/', (context, next) => {
  context.body = getTitle('Home');
  next();
});

// URL Parameter Example
router.get('/about/:topic?', (context, next) => {
  context.body = getTitle('About - Parameter');
  const topic = context.params.topic ?? 'NOTHING';
  context.body += `<p>About with ${ topic }!</p>`;
  next();
});

// Query Example
router.get('/hello', (context, next) => {
  context.body = getTitle('Hello - Query');

  const name = context.query.name ?? 'anonymous';
  context.body += `<p>Hello~~ ${ name }!</p>`;
  next();
});

// #####################
// Middleware Setting
// #####################
// router middleware 설정 전에 등록 필요
app.use(bodyParser());

app.use(async (context, next) => {
  console.log('\n\n1st Middleware : url is', context.url);
  await next();
  console.log('1st middleware End');
});

app.use((context, next) => {
  console.log('2nd middleware');
  next().then(() => {
    console.log('2nd middleware End');
  });
});

// 라우터 미들웨어 적용
router.use('/api', api.routes());
app.use(router.routes()).use(router.allowedMethods());

app.use((context) => {
  if (!context.body) {
    context.body = '';
  }
  context.body += 'This is last middleware : Hello Koa Server';
});

const port = 7166;
app.listen(port, () => {
  console.log(`Listen to port ${ port }`);
});