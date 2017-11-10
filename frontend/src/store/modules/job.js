import _ from 'lodash'

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
  TEST_POLL_ERROR,
  RESULT_FETCH_BY_JOB_SUCCESS
} from '@/store/types'

import api from '@/Api'

const state = {
  jobs: {},
  results: {},
  tasks: {},
  inputs: {},
  taskOrder: [],
  test: {
    status: 'none',
    uuid: '',
    id: null,
    startedAt: null,
    tries: 0
  }
}

const mutations = {
  [JOB_FETCH_SUCCESS] (state, job) {
    state.jobs = { ...state.jobs, [job.id]: job }
  },
  [JOB_UPDATE_ACTIVE_SUCCESS] (state, job) {
    state.jobs = { ...state.jobs, [job.id]: { ...state.jobs[job.id], active: job.active } }
  },
  [TEST_INITIATE_SUCCESS] (state, uuid) {
    state.test = { status: 'waiting', uuid: uuid, startedAt: Date.now(), tries: 0 }
  },
  [TEST_POLL_SUCCESS] (state, result) {
    state.results = { ...state.results, [result.id]: { ...result } }
    state.test = { ...state.test, status: result.status, id: result.id }
  },
  [TEST_POLL_ERROR] (state, error) {
    const timeout = state.test.tries >= 10
    state.test = { ...state.test, tries: state.test.tries + 1, status: timeout ? 'timeout' : state.test.status }
  },
  [RESULT_FETCH_BY_JOB_SUCCESS] (state, results) {
    state.results = _.keyBy(results, 'id')
  },
  [TASK_FETCH_BY_JOB_SUCCESS] (state, tasks) {
    state.tasks = { ...state.tasks, ...(_.keyBy(tasks.map(t => ({ ...t, inputs: t.inputs.map(i => i.id) })), 'id')) }
    state.taskOrder = tasks.map(t => t.id)

    const inputs = [].concat.apply([], tasks.map(t => t.inputs))
    state.inputs = Object.assign(state.inputs, ...inputs.map(t => ({[t['id']]: t})))
  },
  [TASK_CREATE_SUCCESS] (state, task) {
    state.tasks = { ...state.tasks, [task.id]: { ...task, inputs: task.inputs.map(t => t.id) } }
    state.taskOrder = [...state.taskOrder, task.id]
    state.inputs = Object.assign(state.inputs, ...task.inputs.map(t => ({[t['id']]: t})))
  },
  [TASK_DELETE_SUCCESS] (state, task) {
    state.tasks = _.omit(state.tasks, [task.id])
    state.taskOrder = state.taskOrder.filter(id => id !== task.id)
  },
  [INPUT_UPDATE_SUCCESS] (state, input) {
    state.inputs = { ...state.inputs, [input.id]: { ...input } }
  }
}

const getters = {
  testResult (state) {
    return state.results[state.test.id]
  },
  orderedTasks (state) {
    return state.taskOrder.map(id => state.tasks[id])
  }
}

const actions = {
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
  async resultFetchByJob ({ commit, state }, jobID) {
    try {
      const response = await api.get(`/jobs/${jobID}/results`)
      commit(RESULT_FETCH_BY_JOB_SUCCESS, response)
    } catch (e) {
      commit(ERROR_TRIGGERED, e)
    }
  },
  async taskCreate ({ commit, state }, { jobID, executable }) {
    try {
      const task = await api.post('/tasks', { jobID, executable })
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
  async inputUpdate ({ commit, state, rootState }, { id, payload }) {
    try {
      const response = await api.put(`/inputs/${id}`, payload)
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
      if (response.status === 'waiting') {
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

export default {
  state,
  getters,
  mutations,
  actions
}
