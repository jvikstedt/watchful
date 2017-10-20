<template>
  <div>
    <job-form :onSubmit="onFormSubmit" />
  </div>
</template>

<script>
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
  }
}
</script>

<style>
</style>
