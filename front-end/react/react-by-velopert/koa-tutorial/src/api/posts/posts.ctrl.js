let postId = 1;

const posts = [
  {
    id: 1,
    title: 'Sample Title',
    body: 'Sample Body',
  },
];

exports.write = (ctx) => {
  const { title, body } = ctx.request.body;
  postId++;

  const post = { id: postId, title, body };
  posts.push(post);
  ctx.body = post;
};

exports.list = (ctx) => {
  ctx.body = posts;
};

exports.read = (ctx) => {
  const { id } = ctx.params;
  const post = posts.find((post) => post.id.toString() === id);

  if (!post) {
    ctx.status = 404;
    ctx.body = {
      message: `no post with id ${ id }`,
    };
    return;
  }

  ctx.body = post;
};

exports.remove = (ctx) => {
  const { id } = ctx.params;
  const idx = posts.findIndex((post) => post.id.toString() === id);

  if (idx === -1) {
    ctx.status = 404;
    ctx.body = {
      message: `no post with id ${ id }`,
    };
    return;
  }

  posts.splice(idx, 1);
  ctx.status = 204;
};

exports.replace = (ctx) => {
  const { id } = ctx.params;
  const idx = posts.findIndex((post) => post.id.toString() === id);
  if (idx === -1) {
    ctx.status = 404;
    ctx.body = {
      message: `no post with id ${ id }`,
    };
    return;
  }

  posts[idx] = {
    id,
    ...ctx.request.body,
  };
  ctx.body = posts[idx];
};

exports.update = (ctx) => {
  const { id } = ctx.params;
  const idx = posts.findIndex((post) => post.id.toString() === id);
  if (idx === -1) {
    ctx.status = 404;
    ctx.body = {
      message: `no post with id ${ id }`,
    };
    return;
  }

  posts[idx] = {
    ...posts[idx],
    ...ctx.request.body,
  };
  ctx.body = posts[idx];
};
