module.exports = {
  root: true,

  env: {
    node: true,
  },

  extends: [
    'plugin:vue/essential',
    'plugin:cypress/recommended',
    '@vue/airbnb',
  ],

  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'import/extensions': ['error', 'always',
      {
        js: 'never',
        vue: 'never'
      }
    ],
    'max-len': ['error',
      {
        code: 130, ignoreUrls: true,
      },
    ],
    'comma-dangle': 'off',
    "linebreak-style": 0
  },

  parserOptions: {
    parser: 'babel-eslint',
  },

};
