<template>
  <div>
    <input v-model="job.name" />
    <button class="ui button" @click="onSaveClick">Save</button>
  </div>
</template>

<script>
export default {
  props: ['initialValues'],
  data: () => ({
    job: {}
  }),
  methods: {
    onSaveClick () {
      this.$emit('onSave', this.job)
    }
  },
  created () {
    this.job = this.initialValues || {}
  }
}
</script>

<style>
</style>
