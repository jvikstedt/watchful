import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/pages/Home'
import NewJob from '@/pages/NewJob'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/jobs/new',
      name: 'NewJob',
      component: NewJob
    }
  ]
})
