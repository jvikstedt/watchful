import _ from 'lodash'

import api from '@/Api'

export default {
  namespaced: true,
  state: {
    selectedExecutor: '',
    tasksOrder: [],
    tasks: {},
    inputs: {},
    job: {}
  },
  getters: {
    orderedTasks (state) {
      return state.tasksOrder.map(id => state.tasks[id])
    }
  },
  mutations: {
    setTasks (state, tasks) {
      state.tasks = Object.assign(state.tasks, ...tasks.map(t => ({[t['id']]: t})))
      state.tasksOrder = Object.keys(state.tasks).map(k => state.tasks[k].id)
    },
    setJob (state, job) {
      state.job = job
    },
    setInputs (state, inputs) {
      state.inputs = Object.assign(state.inputs, ...inputs.map(t => ({[t['id']]: t})))
    },
    removeTask (state, task) {
      state.tasks = _.omit(state.tasks, [task.id])
      state.tasksOrder = state.tasksOrder.filter(element => element !== task.id)
    },
    setInputValue (state, payload) {
      const input = { ...state.inputs[payload.inputID] }
      if (!input.changed) {
        input.changed = true
        input.oldValue = input.value
      }
      input.value = payload.value
      if (input.value === input.oldValue) {
        input.changed = false
      }
      state.inputs = { ...state.inputs, [payload.inputID]: input }
    },
    setSelectedExecutor (state, executor) {
      state.selectedExecutor = executor
    }
  },
  actions: {
    async addTask ({ commit, state }) {
      const executor = state.selectedExecutor
      try {
        const response = await api.post('/tasks', { jobID: 1, executor })

        const inputs = response.inputs
        const task = { ...response, inputs: response.inputs.map(i => i.id) }

        commit('setTasks', [task])
        commit('setInputs', inputs)
      } catch (e) {
        commit('setFlash', { status: 'error', header: 'Something went wrong!', body: e.toString() }, { root: true })
      }
    },
    async updateActive ({ commit, state }, active) {
      try {
        const job = await api.put(`/jobs/${state.job.id}`, { active })
        commit('setJob', job)
      } catch (e) {
        commit('setFlash', { status: 'error', header: 'Something went wrong!', body: e.toString() }, { root: true })
      }
    },
    async removeTask ({ commit, state }, taskID) {
      try {
        const task = await api.delete(`/tasks/${taskID}`)
        commit('removeTask', task)
      } catch (e) {
        commit('setFlash', { status: 'error', header: 'Something went wrong!', body: e.toString() }, { root: true })
      }
    },
    async getTasks ({ commit, state }, jobID) {
      try {
        const response = await api.get(`/jobs/${jobID}/tasks`)

        const inputs = [].concat.apply([], response.map(t => t.inputs))
        const tasks = response.map(r => ({ ...r, inputs: r.inputs.map(i => i.id) }))

        commit('setTasks', tasks)
        commit('setInputs', inputs)
      } catch (e) {
        commit('setFlash', { status: 'error', header: 'Something went wrong!', body: e.toString() }, { root: true })
      }
    },
    async getJob ({ commit, state }, jobID) {
      try {
        const job = await api.get(`/jobs/${jobID}`)

        commit('setJob', job)
      } catch (e) {
        commit('setFlash', { status: 'error', header: 'Something went wrong!', body: e.toString() }, { root: true })
      }
    },
    async saveInput ({ commit, state }, inputID) {
      try {
        const response = await api.put(`/inputs/${inputID}`, state.inputs[inputID])
        commit('setInputs', [response])
      } catch (e) {
        commit('setFlash', { status: 'error', header: 'Something went wrong!', body: e.toString() }, { root: true })
      }
    }
  }
}
