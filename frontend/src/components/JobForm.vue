<template>
  <div>
    <select v-model="selectedExecutor">
      <option v-for="(executor, index) in executors">{{ executor.name }}</option>
    </select>
    <button class="ui button" @click="addExecutor">Add executor</button>

    <select v-model="selectedChecker">
      <option v-for="(checker, index) in checkers">{{ checker.name }}</option>
    </select>
    <button class="ui button" @click="addChecker">Add checker</button>
  </div>
</template>

<script>
export default {
  props: {
    executors: {
      type: Array,
      required: true
    },
    checkers: {
      type: Array,
      required: true
    },
    onSubmit: {
      type: Function,
      required: true
    }
  },

  methods: {
    addChecker () {
      console.log('add checker')
    },
    addExecutor () {
      console.log('add executor')
    }
  }
}
</script>

<style>
</style>
