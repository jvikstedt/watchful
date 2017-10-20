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
    },
    executors (state) {
      return Object.keys(state.tasks).map(key => state.tasks[key]).filter(t => t.type === 'executor')
    },
    executorsBefore: (state, getters) => (task) => {
      const position = state.tasksOrder.indexOf(task.id)
      return getters.executors.filter(e => {
        return state.tasksOrder.indexOf(e.id) < position
      })
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
      commit('addTask', { type: 'checker', identifier: '' })
    },
    addExecutor ({ getters, commit }) {
      commit('addTask', { type: 'executor', identifier: '', takes: {} })
    },
    setTaskIdentifier ({ commit }, { task, identifier }) {
      commit('setTask', { ...task, identifier: identifier, takes: [] })
    },
    updateTaskTakeValue ({ commit }, { task, takeName, value }) {
      commit('setTask', { ...task, takes: { ...task.takes, [takeName]: value } })
    }
  }
}
