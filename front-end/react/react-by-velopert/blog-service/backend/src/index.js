require('dotenv').config({ path: '.env.development' });

const Koa = require('koa');
const Router = require('koa-router');
const bodyParser = require('koa-bodyparser');
const mongoose = require('mongoose');

const jwtParsing = require('./middleware/jwt');
const api = require('./api');
const createFakeData = require('./mock/createFakeData');

const { PORT, MONGO_URI } = process.env;

// #####################
// DB Setting
// #####################
mongoose
  .connect(MONGO_URI)
  .then(() => {
    console.log('Connected to MongoDB');
  })
  .catch((err) => {
    console.error('Error for connection MongoDB');
    console.error(err);
  });


const app = new Koa();
const router = new Router();


// #####################
// Middleware Setting
// #####################
app.use(bodyParser());  // router middleware 설정 전에 등록 필요
app.use(jwtParsing);

// 라우터 미들웨어 적용
router.post('/createMock', async (ctx) => {
  try {
    await createFakeData();
    ctx.body = 'Success to Create';
  } catch (e) {
    console.log(e);
    ctx.status = 500;
  }
});
router.use('/api', api.routes());
app.use(router.routes()).use(router.allowedMethods());

const port = PORT || 7166;
app.listen(port, () => {
  console.log(`Listen to port ${ port }`);
});