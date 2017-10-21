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
    executors: [],
    flash: null
  },
  getters: {
    findExecutor: (state, getters) => (identifier) => {
      return state.executors.find(e => e.identifier === identifier)
    }
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
