import _ from 'lodash'

export default {
  namespaced: true,
  state: {
    tasksOrder: [],
    tasks: {},
    nextID: 0
  },
  getters: {
    tasksAmount (state) {
      return Object.keys(state.tasks).length
    },
    orderedTasks (state) {
      return state.tasksOrder.map(id => state.tasks[id])
    }
  },
  mutations: {
    addTask (state, task) {
      const nextID = state.nextID
      state.nextID = nextID + 1

      state.tasks = { ...state.tasks, [nextID]: { ...task, id: nextID } }
      state.tasksOrder = [ ...state.tasksOrder, nextID ]
    },
    setTask (state, task) {
      state.tasks = { ...state.tasks, [task.id]: task }
    },
    removeTask (state, id) {
      state.tasks = _.omit(state.tasks, [id])
      state.tasksOrder = state.tasksOrder.filter(element => element !== id)
    }
  },
  actions: {
    addChecker ({ getters, commit }) {
      commit('addTask', { type: 'checker', commandName: '' })
    },
    addExecutor ({ getters, commit }) {
      commit('addTask', { type: 'executor', commandName: '' })
    },
    setTaskCommandName ({ commit }, { task, name }) {
      commit('setTask', { ...task, commandName: name })
    }
  }
}