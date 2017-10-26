import _ from 'lodash'

import {
  TASK_FETCH_BY_JOB_SUCCESS,
  TASK_CREATE_SUCCESS,
  TASK_DELETE_SUCCESS
} from '@/store/types'

export default {
  state: {
    all: {},
    order: []
  },
  getters: {
    orderedTasks (state) {
      return state.order.map(id => state.all[id])
    }
  },
  mutations: {
    [TASK_FETCH_BY_JOB_SUCCESS] (state, tasks) {
      state.all = { ...state.all, ...(_.keyBy(tasks.map(t => ({ ...t, inputs: t.inputs.map(i => i.id) })), 'id')) }
      state.order = tasks.map(t => t.id)
    },
    [TASK_CREATE_SUCCESS] (state, task) {
      state.all = { ...state.all, [task.id]: { ...task, inputs: task.inputs.map(t => t.id) } }
      state.order = [...state.order, task.id]
    },
    [TASK_DELETE_SUCCESS] (state, task) {
      state.all = _.omit(state.all, [task.id])
      state.order = state.order.filter(id => id !== task.id)
    }
  }
}
