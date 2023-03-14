const Router = require('koa-router');

const { getRequestInfo } = require('../utils');

const reqInfo = new Router();
reqInfo.get('/', getRequestInfo);
reqInfo.post('/', getRequestInfo);

reqInfo.get('/:id', getRequestInfo);
reqInfo.delete('/:id', getRequestInfo);
reqInfo.put('/:id', getRequestInfo);
reqInfo.patch('/:id', getRequestInfo);

module.exports = reqInfo;