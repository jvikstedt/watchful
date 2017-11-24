<template>
  <transition-group name="list" tag="div" class="ui middle aligned divided list animated">
    <router-link :to="'/jobs/' + job.id" class="item" v-for="job in jobs" :key="job.id">
      {{ job.name }} ({{ job.lastRun }})
      <i :class="jobStatusClasses(job.status)" :title="job.status" style="float: right"></i>
    </router-link>
  </transition-group>
</template>

<script>
export default {
  created () {
    this.$store.dispatch('jobFetchAll')
  },
  methods: {
    jobStatusClasses (status) {
      return [
        'large',
        'icon',
        { smile: status === 'success' },
        { green: status === 'success' },
        { frown: status === 'error' },
        { red: status === 'error' }
      ]
    }
  },
  computed: {
    jobs () {
      return this.$store.state.job.jobs
    }
  }
}
</script>

<style>
</style>
