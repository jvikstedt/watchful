<template>
  <div>
    <select v-model="selectedExecutor">
      <option value="" selected disabled hidden>-</option>
      <option v-for="executor in executors">{{ executor.identifier }}</option>
    </select>
    <button class="ui button" @click="taskCreate">Add task</button>
  </div>
</template>

<script>

export default {
  props: ['executors', 'onTaskAdd'],
  data: () => ({
    selectedExecutor: ''
  }),
  methods: {
    taskCreate () {
      if (this.selectedExecutor) {
        this.onTaskAdd(this.selectedExecutor)
      }
    }
  }
}
</script>

<style>
</style>
