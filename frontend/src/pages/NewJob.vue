<template>
  <div>
    <job-form :executors="executors" :checkers="checkers" :onSubmit="onFormSubmit" />
  </div>
</template>

<script>
import { mapState } from 'vuex'

import JobForm from '@/components/JobForm'

export default {
  components: {
    JobForm
  },

  methods: {
    onFormSubmit: function (job) {
      console.log(job)
      // TODO Send job
    }
  },

  computed: mapState([
    'checkers',
    'executors'
  ]),

  created () {
    this.$store.dispatch('getCheckers')
    this.$store.dispatch('getExecutors')
  }
}
</script>

<style>
</style>
