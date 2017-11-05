// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'

import store from './store'

require('semantic-ui-css/semantic.css')
require('semantic-ui-css/semantic.js')

Vue.filter('truncate', function (text = '', stop, clamp) {
  if (typeof text === 'string' || text instanceof String) {
    return text.slice(0, stop) + (stop < text.length ? clamp || '...' : '')
  }
  return text
})

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App },
  store
})
