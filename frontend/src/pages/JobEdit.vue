<template>
  <div>
    <Control :job="job" />
    <TasksEdit />
    <result-list :results="results" />
  </div>
</template>

<script>
import ResultList from '@/components/ResultList'

import Control from '@/containers/job/Control'
import TasksEdit from '@/containers/job/TasksEdit'

export default {
  created () {
    this.$store.dispatch('getExecutables')
    this.$store.dispatch('jobFetch', this.jobID)
    this.$store.dispatch('taskFetchByJob', this.jobID)
    this.$store.dispatch('resultFetchByJob', this.jobID)
  },
  computed: {
    results () {
      return this.$store.state.result.all || {}
    },
    job () {
      return this.$store.state.job.all[this.jobID] || {}
    },
    jobID () {
      return this.$route.params.id
    }
  },
  components: {
    Control,
    TasksEdit,
    ResultList
  }
}
</script>

<style>
</style>
