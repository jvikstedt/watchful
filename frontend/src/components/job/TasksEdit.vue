<template>
  <div>
    <div class="ui raised segments">
      <div class="ui segment" v-for="(task, index) in tasks">
        {{ task.id }}
        {{ task.executable }}
        <i v-if="resultItemByTaskID(task.id).status === 'error'" class="frown large icon red" style="float: right"></i>
        <i v-if="resultItemByTaskID(task.id).status === 'success'" class="smile large icon green" style="float: right"></i>
        <i class="close icon" @click="taskDelete(task.id)"></i>
        <i v-if="index > 0" class="angle up icon" @click="up(index)"></i>
        <i v-if="index < tasks.length - 1" class="angle down icon" @click="down(index)"></i>
        <input-creator :onInputAdd="(name, type) => inputCreate({ taskID: task.id, name: name, type: type })" />
        <task-input v-for="inputID in task.inputs" :key="inputID" :input="inputByInputID(inputID)" :onInputUpdate="value => inputUpdate({ id: inputID, payload: value })" :onInputDelete="inputDelete" />
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
import TaskInput from '@/components/job/TaskInput'
import TaskOutput from '@/components/job/TaskOutput'
import InputCreator from '@/components/job/InputCreator'

import _ from 'lodash'

export default {
  props: ['tasks'],
  methods: {
    ...mapActions([
      'taskDelete',
      'inputUpdate',
      'inputCreate',
      'inputDelete',
      'taskSwapSeq'
    ]),
    getInputByID (id) {
      return this.$store.state.job.inputs[id] || {}
    },
    getExecutableByID (id) {
      return this.$store.state.executables[id] || {}
    },
    resultItemByTaskID (taskID) {
      const result = this.$store.getters.testResult || {}
      return _.find(result.resultItems, ri => ri.taskID === taskID) || {}
    },
    up (taskIndex) {
      this.taskSwapSeq({ id1: this.tasks[taskIndex].id, id2: this.tasks[taskIndex - 1].id })
    },
    down (taskIndex) {
      this.taskSwapSeq({ id1: this.tasks[taskIndex].id, id2: this.tasks[taskIndex + 1].id })
    },
    inputByInputID (inputID) {
      return this.$store.state.job.inputs[inputID]
    }
  },
  computed: {
    executables () {
      return this.$store.state.executables
    }
  },
  components: {
    TaskInput,
    TaskOutput,
    InputCreator
  }
}
</script>

<style>
  div.error {
    color: red;
  }
</style>
