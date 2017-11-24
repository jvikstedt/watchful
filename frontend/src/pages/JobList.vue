<template>
  <div>
    <transition-group name="list" tag="div" class="ui middle aligned divided list animated">
      <router-link :to="'/jobs/' + job.id" class="item" v-for="job in jobs" :key="job.id">
        {{ job.name }} ({{ job.lastRun }})
        <i :class="jobStatusClasses(job.status)" :title="job.status" style="float: right"></i>
      </router-link>
    </transition-group>
    <job-form @onSave="jobCreate" />
  </div>
</template>

<script>
import { mapActions } from 'vuex'

import JobForm from '@/components/job/Form'

export default {
  created () {
    this.$store.dispatch('jobFetchAll')
  },
  methods: {
    ...mapActions([
      'jobCreate'
    ]),
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
  },
  components: {
    JobForm
  }
}
</script>

<style>
</style>
