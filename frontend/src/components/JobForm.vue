<template>
  <div>
    <button class="ui button" @click="addTask({ inputs: [] })">Add task</button>

    <div class="ui raised segments">
      <div class="ui segment" v-for="(task, index) in tasks">
        {{ task.id }}
        <i class="close icon" @click="removeTask(task.id)"></i>
        <select :value="task.executor" @input="updateTask({task: task, attributes: { executor: $event.target.value }})">
          <option v-for="(executor, index) in executors">{{ executor.identifier }}</option>
        </select>
        <div v-if="task.executor">
          <b>Inputs:</b>
          <div v-for="(input, index) in findExecutor(task.executor).input">
            <label v-text="input.name" />
            <input :value="task.inputs[input.name]" />
          </div>

          <b>Outputs:</b>
          <div v-for="(output, index) in findExecutor(task.executor).output">
            <span v-text="output.name" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState, mapMutations } from 'vuex'

export default {
  props: {
    onSubmit: {
      type: Function,
      required: true
    }
  },

  methods: {
    ...mapMutations('job', [
      'addTask',
      'removeTask',
      'updateTask'
    ]),
    findExecutor (identifier) {
      return this.$store.getters.findExecutor(identifier)
    }
  },

  computed: {
    ...mapState([
      'executors'
    ]),
    ...mapState('job', [
      'tasks'
    ])
  },

  created () {
    this.$store.dispatch('getExecutors')
  }
}
</script>

<style>
</style>
