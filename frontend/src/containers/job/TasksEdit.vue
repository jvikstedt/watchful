<template>
  <div>
    <div class="ui raised segments">
      <div class="ui segment" v-for="(task, index) in tasks">
        {{ task.id }}
        {{ task.executable }}
        <i v-if="resultItemByTaskID(task.id).status === 'error'" class="frown large icon red" style="float: right"></i>
        <i v-if="resultItemByTaskID(task.id).status === 'success'" class="smile large icon green" style="float: right"></i>
        <i class="close icon" @click="taskDelete(task.id)"></i>
        <task-input v-for="inputID in task.inputs" :key="inputID" :input="getInputByID(inputID)" :onUpdate="inputUpdate" :tasks="tasks.slice(0, index)" />
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

import _ from 'lodash'

export default {
  props: ['tasks'],
  methods: {
    ...mapActions([
      'taskDelete',
      'inputUpdate'
    ]),
    getInputByID (id) {
      return this.$store.state.job.inputs[id] || {}
    },
    getExecutableByID (id) {
      return this.$store.state.executables[id]
    },
    resultItemByTaskID (taskID) {
      const result = this.$store.getters.testResult || {}
      return _.find(result.resultItems, ri => ri.taskID === taskID) || {}
    }
  },
  computed: {
    executables () {
      return this.$store.state.executables
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
