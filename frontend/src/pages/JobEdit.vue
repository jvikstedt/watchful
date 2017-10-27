<template>
  <div>
    <task-creator :executors="executors" :onTaskAdd="taskCreate" />
    <div class="ui toggle checkbox">
      <input type="checkbox" name="public" :checked="job.active" @change="updateActive({jobID: jobID, active: $event.target.checked})">
      <label>On / Off</label>
    </div>
    <button :class="testBtnClasses" @click="initiateTestRun(jobID)">Test</button>

    <div class="ui raised segments">
      <div class="ui segment" v-for="task in orderedTasks">
        {{ task.id }}
        {{ task.executor }}
        <i class="close icon" @click="taskDelete(task.id)"></i>
        <task-input v-for="inputID in task.inputs" :key="inputID" :input="getInputByID(inputID)" :onUpdate="inputUpdate" />
        <task-output v-for="output in getExecutorByID(task.executor).output" :key="output.name" :output="output" />
      </div>
    </div>
  </div>
</template>

<script>
import { mapState, mapActions, mapGetters } from 'vuex'
import TaskCreator from '@/components/TaskCreator'
import TaskInput from '@/components/TaskInput'
import TaskOutput from '@/components/TaskOutput'

export default {
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
    getExecutorByID (id) {
      return this.$store.state.executors[id]
    }
  },

  computed: {
    ...mapState([
      'executors'
    ]),
    ...mapGetters([
      'orderedTasks'
    ]),
    job () {
      return this.$store.state.job.all[this.jobID] || {}
    },
    jobID () {
      return this.$route.params.id
    },
    testBtnClasses () {
      const status = this.$store.state.job.test.status
      return [
        'ui',
        'button',
        {
          loading: status === 'waiting'
        }
      ]
    }
  },

  created () {
    this.$store.dispatch('getExecutors')
    this.$store.dispatch('jobFetch', this.jobID)
    this.$store.dispatch('taskFetchByJob', this.jobID)
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
