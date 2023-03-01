module.exports = {
  env: {
    browser: true,
    commonjs: true,
    es2021: true,
  },
  extends: 'airbnb-base',
  overrides: [],
  parserOptions: {
    ecmaVersion: 'latest',
  },
  root: true,
  rules: {
    semi: 'error',
    'no-unused-vars': 'off',
  },
};
