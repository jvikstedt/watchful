'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  ENDPOINT: JSON.stringify(process.env.ENDPOINT || 'http://0.0.0.0:8000/api/v1')
})
