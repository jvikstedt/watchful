<template>
  <div>
    <project-list :projects="projects" />
  </div>
</template>

<script>
import { EventBus } from '@/EventBus'

import ProjectList from '@/components/ProjectList'

export default {
  data: () => ({
    projects: []
  }),

  components: {
    ProjectList
  },

  async created () {
    try {
      this.projects = await this.api.get('/projects')
    } catch (e) {
      EventBus.$emit('flash', { status: 'error', header: 'Something went wrong!', body: e.toString() })
    }
  }
}
</script>

<style>
</style>
