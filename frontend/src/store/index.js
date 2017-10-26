import Vue from 'vue'
import Vuex from 'vuex'

import api from '@/Api'

import job from './modules/job'
import input from './modules/input'
import task from './modules/task'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    'job': job,
    'input': input,
    'task': task
  },
  state: {
    executors: {},
    flash: null
  },
  actions: {
    async getExecutors ({ commit }) {
      try {
        const executors = await api.get('/executors')
        commit('setExecutors', executors)
      } catch (e) {
        commit('setFlash', { status: 'error', header: 'Something went wrong!', body: e.toString() })
      }
    }
  },
  mutations: {
    setExecutors (state, executors) {
      state.executors = Object.assign({}, ...executors.map(e => ({[e['identifier']]: e})))
    },
    setFlash (state, flash) {
      state.flash = flash
    },
    clearFlash (state) {
      state.flash = null
    }
  }
})
