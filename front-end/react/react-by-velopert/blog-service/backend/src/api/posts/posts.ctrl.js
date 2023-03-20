const mongoose = require('mongoose');
const joi = require('joi');

const utils = require('./utils');
const Post = require('../../models/post');

const { ObjectId } = mongoose.Types;

exports.getPostById = async (ctx, next) => {
  const { id } = ctx.params;
  if (!ObjectId.isValid(id)) {
    ctx.status = 400;
    return;
  }
  try {
    const post = await Post.findById(id).exec();
    if (!post) {
      ctx.status = 404;
      return;
    }
    ctx.state.post = post;
    return next();
  } catch (e) {
    ctx.throw(500, e);
  }
};

exports.checkPostOwner = (ctx, next) => {
  const { post, user } = ctx.state;
  if (post.user._id.toString() !== user._id) {
    ctx.status = 403;
    return;
  }
  return next();
};


exports.write = async (ctx) => {
  const schema = joi.object().keys({
    title: joi.string().required(),
    body: joi.string().required(),
    tags: joi.array().items(joi.string()).required(),
  });

  const validate = schema.validate(ctx.request.body);
  if (validate.error) {
    ctx.status = 400;
    ctx.body = validate.error;
    return;
  }

  const { title, body, tags } = ctx.request.body;

  const post = new Post({
    title,
    body: utils.removeBadTags(body),
    tags,
    user: ctx.state.user,
  });
  try {
    await post.save();
    ctx.body = post;
  } catch (e) {
    ctx.throw(500, e);
  }
};

// GET ?page=&username=&tag=
const pageSize = 10;
exports.list = async (ctx) => {
  const page = parseInt(ctx.query.page || '1', 10);
  if (page < 1) {
    ctx.status = 400;
    ctx.body = 'page parameter must be positive';
    return;
  }

  const { tag, username } = ctx.query;
  const query = {
    ...(username ? { 'user.name': username } : {}),
    ...(tag ? { tags: tag } : {}),
  };

  console.log(page, query);

  try {
    // sort with last created item
    const posts = await Post.find(query)
      .sort(({ _id: -1 }))
      .limit(pageSize)
      .skip((page - 1) * pageSize)
      .lean()
      .exec();
    const postCount = await Post.countDocuments(query).exec();
    ctx.set('Last-Page', Math.ceil(postCount / pageSize));
    ctx.body = posts
      //.map((post) => post.toJSON()) : lean() 으로 처리
      .map((post) => ({
        ...post,
        body: utils.removeHTMLAndShorten(post.body),
      }));
  } catch (e) {
    ctx.throw(500, e);
  }
};

exports.read = (ctx) => {
  ctx.body = ctx.state.post;
};

exports.remove = async (ctx) => {
  const { id } = ctx.params;
  try {
    await Post.findByIdAndRemove(id).exec();
    ctx.status = 204;
  } catch (e) {
    ctx.throw(500, e);
  }
};


exports.update = async (ctx) => {
  const schema = joi.object().keys({
    title: joi.string(),
    body: joi.string(),
    tags: joi.array().items(joi.string()),
  });

  const validate = schema.validate(ctx.request.body);
  if (validate.error) {
    ctx.status = 400;
    ctx.body = validate.error;
    return;
  }

  const { id } = ctx.params;
  const data = { ...ctx.request.body };
  if (data.body) {
    data.body = utils.removeBadTags(data.body);
  }
  try {
    const post = await Post.findByIdAndUpdate(id, data, {
      new: true,
    }).exec();
    if (!post) {
      ctx.status = 404;
      return;
    }
    ctx.body = post;
  } catch (e) {
    ctx.throw(500, e);
  }
};
