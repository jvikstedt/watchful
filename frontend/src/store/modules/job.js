import _ from 'lodash'

import {
  ERROR_TRIGGERED,
  JOB_FETCH_SUCCESS,
  JOB_FETCH_ALL_SUCCESS,
  JOB_UPDATE_ACTIVE_SUCCESS,
  TASK_FETCH_BY_JOB_SUCCESS,
  TASK_SWAP_SEQ_SUCCESS,
  TASK_CREATE_SUCCESS,
  TASK_DELETE_SUCCESS,
  INPUT_UPDATE_SUCCESS,
  INPUT_DELETE_SUCCESS,
  INPUT_CREATE_SUCCESS,
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
  [JOB_FETCH_ALL_SUCCESS] (state, jobs) {
    state.jobs = _.keyBy(jobs, 'id')
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
    state.tasks = { ...state.tasks, ...(_.keyBy(tasks.map(t => ({ ...t, inputs: t.inputs ? t.inputs.map(i => i.id) : [] })), 'id')) }

    const inputs = [].concat.apply([], tasks.map(t => t.inputs ? t.inputs : []))
    state.inputs = Object.assign(state.inputs, ...inputs.map(t => ({[t['id']]: t})))
  },
  [TASK_SWAP_SEQ_SUCCESS] (state, { id1, id2 }) {
    const task1 = state.tasks[id1]
    const task2 = state.tasks[id2]

    state.tasks = { ...state.tasks, [task1.id]: { ...task1, seq: task2.seq }, [task2.id]: { ...task2, seq: task1.seq } }
  },
  [TASK_CREATE_SUCCESS] (state, task) {
    state.tasks = { ...state.tasks, [task.id]: { ...task, inputs: [] } }
  },
  [INPUT_CREATE_SUCCESS] (state, input) {
    const task = state.tasks[input.taskID]
    task.inputs = task.inputs ? [ ...task.inputs, input.id ] : [input.id]

    state.tasks = { ...state.tasks, [task.id]: { ...task } }
    state.inputs = { ...state.inputs, [input.id]: { ...input } }
  },
  [TASK_DELETE_SUCCESS] (state, task) {
    state.tasks = _.omit(state.tasks, [task.id])
  },
  [INPUT_UPDATE_SUCCESS] (state, input) {
    state.inputs = { ...state.inputs, [input.id]: { ...input } }
  },
  [INPUT_DELETE_SUCCESS] (state, input) {
    const task = state.tasks[input.taskID]
    task.inputs = _.remove(task.inputs, id => id !== input.id)

    state.tasks = { ...state.tasks, [task.id]: { ...task } }
    state.inputs = _.omit(state.inputs, [input.id])
  }
}

const getters = {
  testResult (state) {
    return state.results[state.test.id]
  },
  orderedTasks (state) {
    return _.orderBy(state.tasks, 'seq', 'asc')
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
  async jobFetchAll ({ commit, state }) {
    try {
      const response = await api.get(`/jobs`)
      commit(JOB_FETCH_ALL_SUCCESS, response)
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
  async taskSwapSeq ({ commit, state }, { id1, id2 }) {
    try {
      await api.post('/tasks/swap_seq', { id1, id2 })
      commit(TASK_SWAP_SEQ_SUCCESS, { id1, id2 })
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
  async inputCreate ({ commit, state }, { taskID, name, type }) {
    try {
      const response = await api.post('/inputs', { taskID, name, type })
      commit(INPUT_CREATE_SUCCESS, response)
    } catch (e) {
      commit(ERROR_TRIGGERED, e)
    }
  },
  async inputDelete ({ commit, state }, inputID) {
    try {
      const input = await api.delete(`/inputs/${inputID}`)
      commit(INPUT_DELETE_SUCCESS, input)
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
