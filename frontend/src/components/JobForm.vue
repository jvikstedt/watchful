<template>
  <div>
    <button class="ui button" @click="addExecutor">Add executor</button>
    <button class="ui button" @click="addChecker">Add checker</button>

    <div>
      <div v-for="(task, index) in tasks">
        {{ task.type }}
        <div v-if="task.type === 'executor'">
          <select :value="task.identifier" @input="setTaskIdentifier({task: task, identifier: $event.target.value})">
            <option v-for="(executor, index) in executors">{{ executor.identifier }}</option>
          </select>
        </div>

        <div v-if="task.type === 'checker'">
          <select :value="task.identifier" @input="setTaskIdentifier({task: task, identifier: $event.target.value})">
            <option v-for="(checker, index) in checkers">{{ checker.identifier }}</option>
          </select>
        </div>

        <i class="close icon" @click="removeTask(task.id)"></i>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState, mapActions, mapGetters, mapMutations } from 'vuex'

export default {
  props: {
    onSubmit: {
      type: Function,
      required: true
    }
  },

  methods: {
    ...mapActions('job', [
      'addChecker',
      'addExecutor',
      'setTaskIdentifier'
    ]),
    ...mapMutations('job', [
      'removeTask'
    ])
  },

  computed: {
    ...mapState([
      'checkers',
      'executors'
    ]),
    ...mapGetters('job', {
      tasks: 'orderedTasks'
    })
  },

  created () {
    this.$store.dispatch('getCheckers')
    this.$store.dispatch('getExecutors')
  }
}
</script>

<style>
</style>
