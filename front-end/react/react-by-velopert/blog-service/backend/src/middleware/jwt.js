const jwt = require('jsonwebtoken');

const User = require('../models/user');
const authUtils = require('../api/auth/utils');

const jwtParsing = async (ctx, next) => {
  const token = ctx.cookies.get('access_token');
  if (!token) return next();
  try {
    const decoded = jwt.verify(token, process.env.JWT_SECRET_KEY);
    ctx.state.user = {
      _id: decoded._id,
      name: decoded.name,
    };
    console.log(`Request User: ${ decoded.name }`);

    const now = Math.floor(Date.now() / 1000);
    if (decoded.exp - now < 60 * 60 * 24 * 3.5) {
      const user = await User.findById(decoded._id);
      const newToken = user.generateToken();
      console.log('Token is updated: ', newToken);
      authUtils.setTokenCookie(ctx, newToken);
    }

    return next();
  } catch (e) {
    return next();
  }
};

module.exports = jwtParsing;