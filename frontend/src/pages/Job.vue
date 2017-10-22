<template>
  <div>
    <select :value="selectedExecutor" @input="setSelectedExecutor($event.target.value)">
      <option value="" selected disabled hidden>-</option>
      <option v-for="(executor, index) in executors">{{ executor.identifier }}</option>
    </select>
    <button class="ui button" @click="addTask">Add task</button>

    <div class="ui raised segments">
      <div class="ui segment" v-for="(task, _) in orderedTasks">
        {{ task.id }}
        {{ task.executor }}
        <i class="close icon" @click="removeTask(task.id)"></i>
        <div v-for="(inputID, _) in task.inputs">
          <label v-text="getInputByID(inputID).name" />
          <input :value="getInputByID(inputID).value" @input="setInputValue({inputID: inputID, value: $event.target.value})" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState, mapMutations, mapActions, mapGetters } from 'vuex'

export default {
  methods: {
    ...mapMutations('job', [
      'setSelectedExecutor',
      'updateInputValue',
      'setInputValue'
    ]),
    ...mapActions('job', [
      'addTask',
      'removeTask'
    ]),
    getInputByID (id) {
      return this.$store.state.job.inputs[id]
    }
  },

  computed: {
    ...mapState([
      'executors',
      'selectedExecutor'
    ]),
    ...mapGetters('job', [
      'orderedTasks'
    ])
  },

  created () {
    this.$store.dispatch('getExecutors')
    this.$store.dispatch('job/getTasks', this.$route.params.id)
  }
}
</script>

<style>
</style>
