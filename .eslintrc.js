module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: [
    'plugin:vue/essential',
    '@vue/airbnb',
  ],
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'import/extensions': 'off',
    'max-len': ['error',
      {
        code: 130, ignoreUrls: true
      }
    ],
    'comma-dangle': 'off'
  },
  parserOptions: {
    parser: 'babel-eslint',
  },
};
