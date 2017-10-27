import Vue from 'vue'
import Vuex from 'vuex'

import job from './modules/job'
import input from './modules/input'
import task from './modules/task'

import api from '@/Api'

import {
  ERROR_TRIGGERED,
  CLEAR_FLASH,
  EXECUTOR_FETCH_ALL_SUCCESS
} from '@/store/types'

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
        commit(EXECUTOR_FETCH_ALL_SUCCESS, executors)
      } catch (e) {
        commit(ERROR_TRIGGERED, e)
      }
    }
  },
  mutations: {
    [EXECUTOR_FETCH_ALL_SUCCESS] (state, executors) {
      state.executors = Object.assign({}, ...executors.map(e => ({[e['identifier']]: e})))
    },
    [ERROR_TRIGGERED] (state, error) {
      state.flash = { status: 'error', header: 'Something went wrong!', body: error.toString() }
    },
    [CLEAR_FLASH] (state) {
      state.flash = null
    }
  }
})
