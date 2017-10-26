import {
  JOB_FETCH_SUCCESS,
  TASK_FETCH_BY_JOB_SUCCESS,
  TASK_CREATE_SUCCESS,
  TASK_DELETE_SUCCESS,
  INPUT_UPDATE_SUCCESS,
  INPUT_SET_VALUE
} from '@/store/types'

import api from '@/Api'

export default {
  state: {
    selectedExecutor: '',
    jobs: {},
    test: {
      status: 'none',
      id: '',
      startedAt: null
    }
  },
  mutations: {
    [JOB_FETCH_SUCCESS] (state, job) {
      state.job = { ...state.jobs, [job.id]: job }
    },
    setSelectedExecutor (state, executor) {
      state.selectedExecutor = executor
    },
    setTest (state, payload) {
      state.test = payload
    }
  },
  actions: {
    async updateActive ({ commit, state }, active) {
      try {
        const job = await api.put(`/jobs/${state.job.id}`, { active })
        commit('setJob', job)
      } catch (e) {
        commit('setFlash', { status: 'error', header: 'Something went wrong!', body: e.toString() }, { root: true })
      }
    },
    async jobFetch ({ commit, state }, jobID) {
      try {
        const job = await api.get(`/jobs/${jobID}`)

        commit(JOB_FETCH_SUCCESS, job)
      } catch (e) {
        commit('setFlash', { status: 'error', header: 'Something went wrong!', body: e.toString() }, { root: true })
      }
    },
    async taskFetchByJob ({ commit, state }, jobID) {
      try {
        const response = await api.get(`/jobs/${jobID}/tasks`)
        commit(TASK_FETCH_BY_JOB_SUCCESS, response)
      } catch (e) {
        commit('setFlash', { status: 'error', header: 'Something went wrong!', body: e.toString() }, { root: true })
      }
    },
    async taskCreate ({ commit, state }) {
      const executor = state.selectedExecutor
      try {
        const task = await api.post('/tasks', { jobID: 1, executor })
        commit(TASK_CREATE_SUCCESS, task)
      } catch (e) {
        commit('setFlash', { status: 'error', header: 'Something went wrong!', body: e.toString() }, { root: true })
      }
    },
    async taskDelete ({ commit, state }, taskID) {
      try {
        const task = await api.delete(`/tasks/${taskID}`)
        commit(TASK_DELETE_SUCCESS, task)
      } catch (e) {
        commit('setFlash', { status: 'error', header: 'Something went wrong!', body: e.toString() }, { root: true })
      }
    },
    async inputUpdate ({ commit, state, rootState }, inputID) {
      try {
        const response = await api.put(`/inputs/${inputID}`, rootState.input.all[inputID])
        commit(INPUT_UPDATE_SUCCESS, response)
      } catch (e) {
        commit('setFlash', { status: 'error', header: 'Something went wrong!', body: e.toString() }, { root: true })
      }
    },
    async inputSetValue ({ commit, state }, inputID) {
      commit(INPUT_SET_VALUE, inputID)
    },
    async initiateTestRun ({ dispatch, commit, state }) {
      try {
        const response = await api.post(`/jobs/${state.job.id}/test_run`, {})
        commit('setTest', { status: 'waiting', id: response, startedAt: Date.now() })
        dispatch('pollTest')
      } catch (e) {
        commit('setFlash', { status: 'error', header: 'Something went wrong!', body: e.toString() }, { root: true })
      }
    },
    async pollTest ({ dispatch, commit, state }) {
      try {
        const response = await api.get(`/results/${state.test.id}`)
        console.log(response)
      } catch (e) {
        commit('setFlash', { status: 'error', header: 'Something went wrong!', body: e.toString() }, { root: true })
      }
      // setTimeout(function () {
      //   dispatch('pollTest')
      // }, 2000)
    }
  }
}
