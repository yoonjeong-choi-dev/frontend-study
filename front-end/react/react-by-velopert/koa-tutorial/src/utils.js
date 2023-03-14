module.exports.getTitle = (title) => {
  return `<h1>${ title }</h1>`;
};

module.exports.getRequestInfo = ctx => {
  ctx.body = {
    method: ctx.method,
    path: ctx.path,
    params: ctx.params,
  }
  console.log('request context', ctx);
}