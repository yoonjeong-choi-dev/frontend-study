module.exports = {
  env: {
    browser: true,
    commonjs: true,
    es2021: true,
    jest: true
  },
  extends: [
    'airbnb-base',
    'prettier',
  ],
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
