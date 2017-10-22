import _ from 'lodash'

import api from '@/Api'

export default {
  namespaced: true,
  state: {
    selectedExecutor: '',
    tasksOrder: [],
    tasks: {},
    inputs: {}
  },
  getters: {
    orderedTasks (state) {
      return state.tasksOrder.map(id => state.tasks[id])
    }
  },
  mutations: {
    addTask (state, task) {
      state.tasks = { ...state.tasks, [task.id]: task }
      state.tasksOrder = [ ...state.tasksOrder, task.id ]
    },
    setTasks (state, tasks) {
      state.tasksOrder = tasks.map(t => t.id)
      state.tasks = Object.assign({}, ...tasks.map(t => ({[t['id']]: t})))
    },
    setInputs (state, inputs) {
      state.inputs = Object.assign({}, ...inputs.map(t => ({[t['id']]: t})))
    },
    removeTask (state, task) {
      state.tasks = _.omit(state.tasks, [task.id])
      state.tasksOrder = state.tasksOrder.filter(element => element !== task.id)
    },
    setInputValue (state, payload) {
      const input = { ...state.inputs[payload.inputID], value: payload.value }
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
        const task = await api.post('/tasks', { jobID: 1, executor })
        commit('addTask', task)
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
    }
  }
}
