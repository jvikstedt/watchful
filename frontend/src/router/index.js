import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/pages/Home'
import JobEdit from '@/pages/JobEdit'

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
      path: '/jobs/:id',
      name: 'Job',
      component: JobEdit
    }
  ]
})
