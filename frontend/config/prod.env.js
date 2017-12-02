'use strict'
module.exports = {
  NODE_ENV: '"production"',
  ENDPOINT: JSON.stringify(process.env.ENDPOINT || 'http://0.0.0.0:8000/api/v1')
}
