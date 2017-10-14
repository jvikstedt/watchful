<template>
  <div>
    <select v-model="job.executorID">
      <option v-for="(executor, index) in executors">{{ executor.name }}</option>
    </select>
    <div v-if="currentExecutor">
      <div v-for="(take, index) in currentExecutor.takes">
        <label :for="take.name">{{ take.name }}</label>
        <input :id="take.name" v-model="job.takes[take.name]" />
      </div>
    </div>
    <button class="ui button" @click="onSubmit(job)">Submit</button>
  </div>
</template>

<script>
import _ from 'lodash'

export default {
  props: {
    executors: {
      type: Array,
      required: true
    },
    onSubmit: {
      type: Function,
      required: true
    }
  },
  data: () => ({
    job: {
      executorID: '',
      takes: {}
    }
  }),

  methods: {
    resolveType: function (typeID) {
      switch (typeID) {
        case 1:
          return 'text'
        case 2:
          return 'number'
        default:
          return 'text'
      }
    }
  },

  computed: {
    currentExecutor: function () {
      return _.find(this.executors, { name: this.job.executorID })
    }
  }
}
</script>

<style>
</style>
