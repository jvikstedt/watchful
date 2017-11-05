import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/pages/Home'
import JobEdit from '@/pages/JobEdit'
import Result from '@/pages/Result'

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
    },
    {
      path: '/results/:id',
      name: 'Result',
      component: Result
    }
  ]
})
