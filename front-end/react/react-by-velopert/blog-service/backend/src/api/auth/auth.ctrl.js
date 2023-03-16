const joi = require('joi');

const User = require('../../models/user');
const utils = require('./utils');

exports.register = async (ctx) => {
  const schema = joi.object().keys({
    name: joi.string().alphanum().min(3).max(30).required(),
    password: joi.string().min(3).required(),
  });

  const validate = schema.validate(ctx.request.body);
  if (validate.error) {
    ctx.status = 400;
    ctx.body = validate.error;
    return;
  }

  const { name, password } = ctx.request.body;
  try {
    const existUser = await User.findByName(name);
    if (existUser) {
      ctx.status = 409;
      ctx.body = 'user already exists';
      return;
    }

    const user = new User({ name });
    await user.setPassword(password);
    await user.save();

    const token = user.generateToken();
    utils.setTokenCookie(ctx, token);

    ctx.body = user.serialize();
  } catch (e) {
    ctx.throw(500, e);
  }
};

exports.login = async (ctx) => {
  const schema = joi.object().keys({
    name: joi.string().required(),
    password: joi.string().required(),
  });

  const validate = schema.validate(ctx.request.body);
  if (validate.error) {
    ctx.status = 400;
    ctx.body = validate.error;
    return;
  }

  const { name, password } = ctx.request.body;
  try {
    const user = await User.findByName(name);
    if (!user) {
      ctx.status = 401;
      return;
    }

    const valid = await user.checkPassword(password);
    if (!valid) {
      ctx.status = 401;
      return;
    }

    const token = user.generateToken();
    utils.setTokenCookie(ctx, token);

    ctx.body = user.serialize();
  } catch (e) {
    ctx.throw(500, e);
  }
};

exports.check = async (ctx) => {
  const { user } = ctx.state;
  if (!user) {
    ctx.status = 401;
    return;
  }
  ctx.body = user;
};

exports.logout = async (ctx) => {
  utils.deleteTokenCookie(ctx);
  ctx.status = 204;
};
