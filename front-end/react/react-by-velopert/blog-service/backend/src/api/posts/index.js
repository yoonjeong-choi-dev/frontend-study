const Router = require('koa-router');

const checkLoggedIn = require('../../middleware/checkLoggedIn');
const postsCtrl = require('./posts.ctrl');

const posts = new Router();

posts.get('/', postsCtrl.list);
posts.post('/', checkLoggedIn, postsCtrl.write);

const post = new Router();
post.get('/', postsCtrl.read);
post.delete('/', checkLoggedIn, postsCtrl.checkPostOwner, postsCtrl.remove);
post.patch('/', checkLoggedIn, postsCtrl.checkPostOwner, postsCtrl.update);

posts.use('/:id', postsCtrl.getPostById, post.routes());
module.exports = posts;
