<template>
  <div>
    <task-creator :executables="executables" :onTaskAdd="onTaskAdd" />
    <div class="ui toggle checkbox">
      <input type="checkbox" name="public" :checked="job.active" @change="onActiveChange($event.target.checked)">
      <label>On / Off</label>
    </div>
    <button :class="testBtnClasses" @click="initiateTestRun(job.id)">Test</button>
  </div>
</template>

<script>
import { mapActions } from 'vuex'
import TaskCreator from '@/components/TaskCreator'

export default {
  props: ['job'],
  methods: {
    ...mapActions([
      'taskCreate',
      'updateActive',
      'initiateTestRun'
    ]),
    onTaskAdd (executable) {
      this.taskCreate({ jobID: this.job.id, executable })
    },
    onActiveChange (active) {
      this.updateActive({jobID: this.job.id, active})
    }
  },
  computed: {
    executables () {
      return this.$store.state.executables
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
    TaskCreator
  }
}
</script>

<style>
</style>
