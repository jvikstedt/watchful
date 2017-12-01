<template>
  <div>
    <task-creator :executables="executables" :onTaskAdd="onTaskAdd" />
    <div class="ui toggle checkbox">
      <input type="checkbox" name="public" :checked="job.active" @change="onActiveChange($event.target.checked)">
      <label>On / Off</label>
    </div>
    <button :class="testBtnClasses" @click="initiateTestRun(job.id)">Test</button>
    <input v-model="cron" @change="onCronChange" />
  </div>
</template>

<script>
import { mapActions } from 'vuex'
import TaskCreator from '@/components/job/TaskCreator'

export default {
  props: ['job'],
  data: () => ({
    cron: ''
  }),
  methods: {
    ...mapActions([
      'taskCreate',
      'updateActive',
      'updateCron',
      'initiateTestRun'
    ]),
    onTaskAdd (executable) {
      this.taskCreate({ jobID: this.job.id, executable })
    },
    onActiveChange (active) {
      this.updateActive({ jobID: this.job.id, active })
    },
    onCronChange (e) {
      this.updateCron({ jobID: this.job.id, cron: e.target.value })
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
  created () {
    this.cron = this.job.cron
  },
  components: {
    TaskCreator
  }
}
</script>

<style>
</style>
