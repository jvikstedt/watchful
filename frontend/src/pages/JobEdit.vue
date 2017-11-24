<template>
  <div>
    <Control v-if="job" :job="job" />
    <TasksEdit :tasks="orderedTasks" />
    <result-list :results="results" />
  </div>
</template>

<script>
import ResultList from '@/components/ResultList'
import Control from '@/components/job/Control'
import TasksEdit from '@/components/job/TasksEdit'

export default {
  created () {
    this.$store.dispatch('getExecutables')
    this.$store.dispatch('jobFetch', this.jobID)
    this.$store.dispatch('taskFetchByJob', this.jobID)
    this.$store.dispatch('resultFetchByJob', this.jobID)
  },
  computed: {
    results () {
      return this.$store.state.job.results
    },
    job () {
      return this.$store.state.job.jobs[this.jobID]
    },
    jobID () {
      return this.$route.params.id
    },
    orderedTasks () {
      return this.$store.getters.orderedTasks
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
