<template>
  <div>
    <select :value="selectedExecutor" @input="setSelectedExecutor($event.target.value)">
      <option value="" selected disabled hidden>-</option>
      <option v-for="executor in executors">{{ executor.identifier }}</option>
    </select>
    <button class="ui button" @click="taskCreate">Add task</button>
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
        <div v-for="inputID in task.inputs">
          <label :for="'value' + inputID" v-text="getInputByID(inputID).name" />
          <input :id="'value' + inputID" :value="getInputByID(inputID).value" @input="inputSetValue({inputID: inputID, value: $event.target.value})" />
          <button class="mini green ui icon button" :disabled="!getInputByID(inputID).changed" @click="inputUpdate(inputID)">
            <i class="checkmark icon" />
          </button>
        </div>
        <div v-for="output in getExecutorByID(task.executor).output">
          {{ output.name }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState, mapMutations, mapActions, mapGetters } from 'vuex'

export default {
  methods: {
    ...mapMutations([
      'setSelectedExecutor'
    ]),
    ...mapActions([
      'taskCreate',
      'taskDelete',
      'inputUpdate',
      'updateActive',
      'initiateTestRun',
      'inputSetValue'
    ]),
    getInputByID (id) {
      return this.$store.state.input.all[id]
    },
    getExecutorByID (id) {
      return this.$store.state.executors[id]
    }
  },

  computed: {
    ...mapState([
      'executors',
      'selectedExecutor'
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
  }
}
</script>

<style>
</style>
