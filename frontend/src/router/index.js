import Vue from 'vue'
import Router from 'vue-router'
import JobEdit from '@/pages/JobEdit'
import JobList from '@/pages/JobList'
import Result from '@/pages/Result'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'JobList',
      component: JobList
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
