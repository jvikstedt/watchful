<template>
  <div>
    <div class="ui raised segments">
      <div class="ui segment" v-for="(task, index) in orderedTasks">
        {{ task.id }}
        {{ task.executable }}
        <i v-if="resultItemByTaskID(task.id).status === 'error'" class="frown large icon red" style="float: right"></i>
        <i v-if="resultItemByTaskID(task.id).status === 'success'" class="smile large icon green" style="float: right"></i>
        <i class="close icon" @click="taskDelete(task.id)"></i>
        <task-input v-for="inputID in task.inputs" :key="inputID" :input="getInputByID(inputID)" :onUpdate="inputUpdate" :tasks="orderedTasks.slice(0, index)" />
        <task-output v-for="output in getExecutableByID(task.executable).output" :key="output.name" :output="output" :resultItem="resultItemByTaskID(task.id)" />
        <div class="error">
          {{ resultItemByTaskID(task.id).error }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions } from 'vuex'
import TaskInput from '@/components/TaskInput'
import TaskOutput from '@/components/TaskOutput'

export default {
  methods: {
    ...mapActions([
      'taskDelete',
      'inputUpdate'
    ]),
    getInputByID (id) {
      return this.$store.state.input.all[id] || {}
    },
    getExecutableByID (id) {
      return this.$store.state.executables[id]
    },
    resultItemByTaskID (taskID) {
      return this.$store.getters.resultItemsGrouped[`${this.$store.state.job.test.id}:${taskID}`] || { output: {} }
    }
  },
  computed: {
    executables () {
      return this.$store.state.executables
    },
    orderedTasks () {
      return this.$store.getters.orderedTasks
    }
  },
  components: {
    TaskInput,
    TaskOutput
  }
}
</script>

<style>
  div.error {
    color: red;
  }
</style>
