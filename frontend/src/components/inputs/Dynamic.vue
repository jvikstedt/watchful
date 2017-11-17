<template>
  <span>
    Dynamic
    TaskID:
    <input v-model="taskID" @input="onChange" />
    outputName:
    <input v-model="outputName" @input="onChange" />
  </span>
</template>

<script>

export default {
  props: ['value'],
  data: () => ({
    taskID: null,
    outputName: null
  }),
  methods: {
    onChange () {
      const asInt = parseInt(this.taskID)
      if (this.outputName && this.taskID && !isNaN(asInt)) {
        this.$emit('input', { taskID: asInt, outputName: this.outputName })
      } else {
        this.$emit('input', null)
      }
    }
  },
  created () {
    this.taskID = this.value ? this.value.taskID : null
    this.outputName = this.value ? this.value.outputName : null
  }
}
</script>

<style>
</style>
