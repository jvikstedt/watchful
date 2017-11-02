<template>
  <div>
    <select v-model="selectedExecutable">
      <option value="" selected disabled hidden>-</option>
      <option v-for="executable in executables">{{ executable.identifier }}</option>
    </select>
    <button class="ui button" @click="taskCreate">Add task</button>
  </div>
</template>

<script>

export default {
  props: ['executables', 'onTaskAdd'],
  data: () => ({
    selectedExecutable: ''
  }),
  methods: {
    taskCreate () {
      if (this.selectedExecutable) {
        this.onTaskAdd(this.selectedExecutable)
      }
    }
  }
}
</script>

<style>
</style>
