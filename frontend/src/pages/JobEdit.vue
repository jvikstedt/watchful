<template>
  <div>
    <task-creator :executables="executables" :onTaskAdd="taskCreate" />
    <div class="ui toggle checkbox">
      <input type="checkbox" name="public" :checked="job.active" @change="updateActive({jobID: jobID, active: $event.target.checked})">
      <label>On / Off</label>
    </div>
    <button :class="testBtnClasses" @click="initiateTestRun(jobID)">Test</button>

    <div class="ui raised segments">
      <div class="ui segment" v-for="(task, index) in orderedTasks">
        {{ task.id }}
        {{ task.executable }}
        <i class="close icon" @click="taskDelete(task.id)"></i>
        <task-input v-for="inputID in task.inputs" :key="inputID" :input="getInputByID(inputID)" :onUpdate="inputUpdate" :tasks="orderedTasks.slice(0, index)" />
        <task-output v-for="output in getExecutableByID(task.executable).output" :key="output.name" :output="output" :resultItem="resultItemByTaskID(task.id)" />
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions } from 'vuex'
import TaskCreator from '@/components/TaskCreator'
import TaskInput from '@/components/TaskInput'
import TaskOutput from '@/components/TaskOutput'

export default {
  created () {
    this.$store.dispatch('getExecutables')
    this.$store.dispatch('jobFetch', this.jobID)
    this.$store.dispatch('taskFetchByJob', this.jobID)
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
    }
  },
  computed: {
    executables () {
      return this.$store.state.executables
    },
    orderedTasks () {
      return this.$store.getters.orderedTasks
    },
    job () {
      return this.$store.state.job.all[this.jobID] || {}
    },
    jobID () {
      return this.$route.params.id
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
    TaskOutput
  }
}
</script>

<style>
</style>
