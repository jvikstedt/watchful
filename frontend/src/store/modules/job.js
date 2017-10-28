import {
  ERROR_TRIGGERED,
  JOB_FETCH_SUCCESS,
  JOB_UPDATE_ACTIVE_SUCCESS,
  TASK_FETCH_BY_JOB_SUCCESS,
  TASK_CREATE_SUCCESS,
  TASK_DELETE_SUCCESS,
  INPUT_UPDATE_SUCCESS,
  TEST_INITIATE_SUCCESS,
  TEST_POLL_SUCCESS,
  TEST_POLL_ERROR
} from '@/store/types'

import api from '@/Api'

export default {
  state: {
    all: {},
    test: {
      status: 'none',
      uuid: '',
      id: null,
      startedAt: null,
      tries: 0
    }
  },
  mutations: {
    [JOB_FETCH_SUCCESS] (state, job) {
      state.all = { ...state.all, [job.id]: job }
    },
    [JOB_UPDATE_ACTIVE_SUCCESS] (state, job) {
      state.all = { ...state.all, [job.id]: { ...state.all[job.id], active: job.active } }
    },
    [TEST_INITIATE_SUCCESS] (state, uuid) {
      state.test = { status: 'waiting', uuid: uuid, startedAt: Date.now(), tries: 0 }
    },
    [TEST_POLL_SUCCESS] (state, result) {
      state.test = { ...state.test, status: result.status, id: result.id }
    },
    [TEST_POLL_ERROR] (state, error) {
      const timeout = state.test.tries >= 10
      state.test = { ...state.test, tries: state.test.tries + 1, status: timeout ? 'timeout' : state.test.status }
    }
  },
  actions: {
    async updateActive ({ commit, state }, { jobID, active }) {
      try {
        const job = await api.put(`/jobs/${jobID}`, { active })
        commit(JOB_UPDATE_ACTIVE_SUCCESS, job)
      } catch (e) {
        commit(ERROR_TRIGGERED, e)
      }
    },
    async jobFetch ({ commit, state }, jobID) {
      try {
        const job = await api.get(`/jobs/${jobID}`)

        commit(JOB_FETCH_SUCCESS, job)
      } catch (e) {
        commit(ERROR_TRIGGERED, e)
      }
    },
    async taskFetchByJob ({ commit, state }, jobID) {
      try {
        const response = await api.get(`/jobs/${jobID}/tasks`)
        commit(TASK_FETCH_BY_JOB_SUCCESS, response)
      } catch (e) {
        commit(ERROR_TRIGGERED, e)
      }
    },
    async taskCreate ({ commit, state }, executor) {
      try {
        const task = await api.post('/tasks', { jobID: 1, executor })
        commit(TASK_CREATE_SUCCESS, task)
      } catch (e) {
        commit(ERROR_TRIGGERED, e)
      }
    },
    async taskDelete ({ commit, state }, taskID) {
      try {
        const task = await api.delete(`/tasks/${taskID}`)
        commit(TASK_DELETE_SUCCESS, task)
      } catch (e) {
        commit(ERROR_TRIGGERED, e)
      }
    },
    async inputUpdate ({ commit, state, rootState }, { value, id }) {
      try {
        const response = await api.put(`/inputs/${id}`, { value })
        commit(INPUT_UPDATE_SUCCESS, response)
      } catch (e) {
        commit(ERROR_TRIGGERED, e)
      }
    },
    async initiateTestRun ({ dispatch, commit, state }, jobID) {
      try {
        const uuid = await api.post(`/jobs/${jobID}/test_run`, {})
        commit(TEST_INITIATE_SUCCESS, uuid)
        dispatch('pollTest')
      } catch (e) {
        commit(ERROR_TRIGGERED, e)
      }
    },
    async pollTest ({ dispatch, commit, state }) {
      try {
        const response = await api.get(`/results/${state.test.uuid}`)
        commit(TEST_POLL_SUCCESS, response)
        if (response.status !== 'done') {
          setTimeout(function () {
            dispatch('pollTest')
          }, 2000)
        }
      } catch (e) {
        commit(TEST_POLL_ERROR)
        if (state.test.status === 'waiting') {
          setTimeout(function () {
            dispatch('pollTest')
          }, 2000)
        } else {
          commit(ERROR_TRIGGERED, e)
        }
      }
    }
  }
}
