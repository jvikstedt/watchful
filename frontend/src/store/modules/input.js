import {
  INPUT_UPDATE_SUCCESS,
  TASK_FETCH_BY_JOB_SUCCESS,
  INPUT_SET_VALUE,
  TASK_CREATE_SUCCESS
} from '@/store/types'

export default {
  state: {
    all: {}
  },
  mutations: {
    [INPUT_UPDATE_SUCCESS] (state, input) {
      state.all = { ...state.all, [input.id]: input }
    },
    [TASK_FETCH_BY_JOB_SUCCESS] (state, tasks) {
      const inputs = [].concat.apply([], tasks.map(t => t.inputs))
      state.all = Object.assign(state.all, ...inputs.map(t => ({[t['id']]: t})))
    },
    [TASK_CREATE_SUCCESS] (state, task) {
      state.all = Object.assign(state.all, ...task.inputs.map(t => ({[t['id']]: t})))
    },
    [INPUT_SET_VALUE] (state, payload) {
      const input = { ...state.all[payload.inputID] }
      if (!input.changed) {
        input.changed = true
        input.oldValue = input.value
      }
      input.value = payload.value
      if (input.value === input.oldValue) {
        input.changed = false
      }
      state.all = { ...state.all, [payload.inputID]: input }
    }
  }
}
