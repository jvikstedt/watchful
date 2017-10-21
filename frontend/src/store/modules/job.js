import _ from 'lodash'

export default {
  namespaced: true,
  state: {
    tasksOrder: [],
    tasks: {},
    nextID: 0
  },
  getters: {
    orderedTasks (state) {
      return state.tasksOrder.map(id => state.tasks[id])
    },
    tasksBefore: (state, getters) => (task) => {
      const position = state.tasksOrder.indexOf(task.id)
      return state.tasks.filter(e => {
        return state.tasksOrder.indexOf(e.id) < position
      })
    }
  },
  mutations: {
    addTask (state, task = {}) {
      const nextID = state.nextID
      state.nextID = nextID + 1

      state.tasks = { ...state.tasks, [nextID]: { ...task, id: nextID } }
      state.tasksOrder = [ ...state.tasksOrder, nextID ]
    },
    setTask (state, task) {
      state.tasks = { ...state.tasks, [task.id]: task }
    },
    updateTask (state, payload = {}) {
      const task = payload.task
      state.tasks = { ...state.tasks, [task.id]: { ...task, ...payload.attributes } }
    },
    removeTask (state, id) {
      state.tasks = _.omit(state.tasks, [id])
      state.tasksOrder = state.tasksOrder.filter(element => element !== id)
    }
  }
}
