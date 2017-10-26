<template>
  <div>
    <select :value="selectedExecutor" @input="setSelectedExecutor($event.target.value)">
      <option value="" selected disabled hidden>-</option>
      <option v-for="executor in executors">{{ executor.identifier }}</option>
    </select>
    <button class="ui button" @click="addTask">Add task</button>
    <div class="ui toggle checkbox">
      <input type="checkbox" name="public" :checked="job.active" @change="updateActive($event.target.checked)">
      <label>On / Off</label>
    </div>
    <button class="ui button" @click="initiateTestRun">Test</button>

    <div class="ui raised segments">
      <div class="ui segment" v-for="task in orderedTasks">
        {{ task.id }}
        {{ task.executor }}
        <i class="close icon" @click="removeTask(task.id)"></i>
        <div v-for="inputID in task.inputs">
          <label :for="'value' + inputID" v-text="getInputByID(inputID).name" />
          <input :id="'value' + inputID" :value="getInputByID(inputID).value" @input="setInputValue({inputID: inputID, value: $event.target.value})" />
          <button class="mini green ui icon button" :disabled="!getInputByID(inputID).changed" @click="saveInput(inputID)">
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
    ...mapMutations('job', [
      'setSelectedExecutor',
      'updateInputValue',
      'setInputValue'
    ]),
    ...mapActions('job', [
      'addTask',
      'removeTask',
      'saveInput',
      'updateActive',
      'initiateTestRun'
    ]),
    getInputByID (id) {
      return this.$store.state.job.inputs[id]
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
    ...mapGetters('job', [
      'orderedTasks'
    ]),
    job () {
      return this.$store.state.job.job
    }
  },

  created () {
    this.$store.dispatch('getExecutors')
    this.$store.dispatch('job/getJob', this.$route.params.id)
    this.$store.dispatch('job/getTasks', this.$route.params.id)
  }
}
</script>

<style>
</style>
