const authCookieKey = 'access_token';

exports.setTokenCookie = (ctx, token) => {
  ctx.cookies.set(authCookieKey, token, {
    maxAge: 1000 * 60 * 60 * 24 * 7,
    httpOnly: true,
  });
};

exports.deleteTokenCookie = (ctx) => {
  ctx.cookies.set(authCookieKey);
};