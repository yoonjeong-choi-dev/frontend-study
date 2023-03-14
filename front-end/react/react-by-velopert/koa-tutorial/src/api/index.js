const Router = require('koa-router');
const api = new Router();

const utils = require('../utils');
const posts = require('./posts');
const info = require('./reqInfo');

api.get('/test', (ctx, next) => {
  ctx.body = utils.getTitle('API Router Test');
  ctx.body += '<p>Success!!</p>';
  next();
});

api.use('/info', info.routes());
api.use('/posts', posts.routes());

module.exports = api;