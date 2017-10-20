import Vue from 'vue'
import Vuex from 'vuex'

import api from '@/Api'

import job from './modules/job'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    'job': job
  },
  state: {
    checkers: [],
    executors: [],
    flash: null
  },
  actions: {
    async getCheckers ({ commit }) {
      try {
        const checkers = await api.get('/checkers')
        commit('setCheckers', checkers)
      } catch (e) {
        commit('setFlash', { status: 'error', header: 'Something went wrong!', body: e.toString() })
      }
    },
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
    setCheckers (state, checkers) {
      state.checkers = checkers
    },
    setExecutors (state, executors) {
      state.executors = executors
    },
    setFlash (state, flash) {
      state.flash = flash
    },
    clearFlash (state) {
      state.flash = null
    }
  }
})
