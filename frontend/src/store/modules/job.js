import _ from 'lodash'

import api from '@/Api'

export default {
  namespaced: true,
  state: {
    selectedExecutor: '',
    tasksOrder: [],
    tasks: {}
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
    removeTask (state, taskID) {
      state.tasks = _.omit(state.tasks, [taskID])
      state.tasksOrder = state.tasksOrder.filter(element => element !== taskID)
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
        await api.delete(`/tasks/${taskID}`)
        commit('removeTask', taskID)
      } catch (e) {
        commit('setFlash', { status: 'error', header: 'Something went wrong!', body: e.toString() }, { root: true })
      }
    },
    async getTasks ({ commit, state }, jobID) {
      try {
        const tasks = await api.get(`/jobs/${jobID}/tasks`)
        commit('setTasks', tasks)
      } catch (e) {
        commit('setFlash', { status: 'error', header: 'Something went wrong!', body: e.toString() }, { root: true })
      }
    }
  }
}
