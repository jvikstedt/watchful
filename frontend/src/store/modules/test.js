import _ from 'lodash'

import {
  TEST_POLL_SUCCESS
} from '@/store/types'

export default {
  state: {
    results: {},
    resultItems: {}
  },
  getters: {
    resultItemsGrouped (state) {
      return _.keyBy(Object.values(state.resultItems), (o) => `${o.resultID}:${o.taskID}`)
    }
  },
  mutations: {
    [TEST_POLL_SUCCESS] (state, result) {
      state.results = { ...state.results, [result.id]: { ...result, resultItems: result.resultItems.map(r => r.id) } }
      state.resultItems = { ...Object.assign(state.resultItems, ...result.resultItems.map(r => ({[r['id']]: { ...r, output: JSON.parse(r.output) }}))) }
    }
  }
}
