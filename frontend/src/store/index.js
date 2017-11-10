import Vue from 'vue'
import Vuex from 'vuex'

import job from './modules/job'

import api from '@/Api'

import {
  ERROR_TRIGGERED,
  CLEAR_FLASH,
  EXECUTOR_FETCH_ALL_SUCCESS
} from '@/store/types'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    'job': job
  },
  state: {
    executables: {},
    flash: null
  },
  actions: {
    async getExecutables ({ commit }) {
      try {
        const executables = await api.get('/executables')
        commit(EXECUTOR_FETCH_ALL_SUCCESS, executables)
      } catch (e) {
        commit(ERROR_TRIGGERED, e)
      }
    }
  },
  mutations: {
    [EXECUTOR_FETCH_ALL_SUCCESS] (state, executables) {
      state.executables = Object.assign({}, ...executables.map(e => ({[e['identifier']]: e})))
    },
    [ERROR_TRIGGERED] (state, error) {
      state.flash = { status: 'error', header: 'Something went wrong!', body: error.toString() }
    },
    [CLEAR_FLASH] (state) {
      state.flash = null
    }
  }
})
