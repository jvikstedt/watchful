import {
  RESULT_FETCH_BY_JOB_SUCCESS,
  TEST_POLL_SUCCESS
} from '@/store/types'

export default {
  state: {
    all: {}
  },
  mutations: {
    [RESULT_FETCH_BY_JOB_SUCCESS] (state, results) {
      state.all = { ...state.all, ...results }
    },
    [TEST_POLL_SUCCESS] (state, result) {
      state.all = { ...state.all, [result.id]: { ...result, resultItems: result.resultItems.map(r => r.id) } }
    }
  }
}
