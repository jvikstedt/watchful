<template>
  <div>
    <button class="ui button" @click="addExecutor">Add executor</button>
    <button class="ui button" @click="addChecker">Add checker</button>

    <div>
      <div v-for="(task, index) in tasks">
        {{ task.type }}
        {{ task.id }}
        <div v-if="task.type === 'executor'">
          <select :value="task.identifier" @input="setTaskIdentifier({task: task, identifier: $event.target.value})">
            <option v-for="(executor, index) in executors">{{ executor.identifier }}</option>
          </select>

          <div v-if="task.identifier">
            <div v-for="(take, index) in executorByIdentifier(task.identifier).takes">
              <label v-text="take.name" />
              <input :value="task.takes[take.name]" @input="updateTaskTakeValue({ task: task, takeName: take.name, value: $event.target.value })" />
            </div>
          </div>
        </div>

        <div v-if="task.type === 'checker'">
          <select :value="task.identifier" @input="setTaskIdentifier({task: task, identifier: $event.target.value})">
            <option v-for="(checker, index) in checkers">{{ checker.identifier }}</option>
          </select>
        </div>

        {{ executorsBefore(task) }}

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
      'setTaskIdentifier',
      'updateTaskTakeValue'
    ]),
    ...mapMutations('job', [
      'removeTask'
    ]),
    checkerByIdentifier (identifier) {
      return this.$store.getters.checkerByIdentifier(identifier)
    },
    executorByIdentifier (identifier) {
      return this.$store.getters.executorByIdentifier(identifier)
    },
    executorsBefore (task) {
      return this.$store.getters['job/executorsBefore'](task)
    }
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
