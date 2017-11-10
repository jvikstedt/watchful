<template>
  <div>
    <task-creator :executables="executables" :onTaskAdd="onTaskAdd" />
    <div class="ui toggle checkbox">
      <input type="checkbox" name="public" :checked="job.active" @change="updateActive({jobID: jobID, active: $event.target.checked})">
      <label>On / Off</label>
    </div>
    <button :class="testBtnClasses" @click="initiateTestRun(jobID)">Test</button>

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
    <result-list :results="results" />
  </div>
</template>

<script>
import { mapActions } from 'vuex'
import TaskCreator from '@/components/TaskCreator'
import TaskInput from '@/components/TaskInput'
import TaskOutput from '@/components/TaskOutput'
import ResultList from '@/components/ResultList'

export default {
  created () {
    this.$store.dispatch('getExecutables')
    this.$store.dispatch('jobFetch', this.jobID)
    this.$store.dispatch('taskFetchByJob', this.jobID)
    this.$store.dispatch('resultFetchByJob', this.jobID)
  },
  methods: {
    ...mapActions([
      'taskCreate',
      'taskDelete',
      'inputUpdate',
      'updateActive',
      'initiateTestRun'
    ]),
    getInputByID (id) {
      return this.$store.state.input.all[id] || {}
    },
    getExecutableByID (id) {
      return this.$store.state.executables[id]
    },
    resultItemByTaskID (taskID) {
      return this.$store.getters.resultItemsGrouped[`${this.$store.state.job.test.id}:${taskID}`] || { output: {} }
    },
    onTaskAdd (executable) {
      this.taskCreate({ jobID: this.jobID, executable })
    }
  },
  computed: {
    executables () {
      return this.$store.state.executables
    },
    orderedTasks () {
      return this.$store.getters.orderedTasks
    },
    results () {
      return this.$store.state.result.all || {}
    },
    job () {
      return this.$store.state.job.all[this.jobID] || {}
    },
    jobID () {
      return parseInt(this.$route.params.id)
    },
    result () {
      const resultID = this.$store.state.job.test.id
      return this.$store.state.test.results[resultID] || {}
    },
    testBtnClasses () {
      const status = this.$store.state.job.test.status
      return [
        'ui',
        'button',
        { loading: status === 'waiting' }
      ]
    }
  },
  components: {
    TaskCreator,
    TaskInput,
    TaskOutput,
    ResultList
  }
}
</script>

<style>
  div.error {
    color: red;
  }
</style>
