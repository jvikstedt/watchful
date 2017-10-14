<template>
  <div>
    <job-form :executors="executors" :onSubmit="onFormSubmit" />
  </div>
</template>

<script>
import { EventBus } from '@/EventBus'
import JobForm from '@/components/JobForm'

export default {
  data: () => ({
    executors: []
  }),

  components: {
    JobForm
  },

  methods: {
    onFormSubmit: async function (job) {
      try {
        this.job = await this.api.post('/jobs', job)
      } catch (e) {
        EventBus.$emit('flash', { status: 'error', header: 'Something went wrong!', body: e.toString() })
      }
    }
  },

  async created () {
    try {
      this.executors = await this.api.get('/executors')
    } catch (e) {
      EventBus.$emit('flash', { status: 'error', header: 'Something went wrong!', body: e.toString() })
    }
  }
}
</script>

<style>
</style>
