<template>
  <div>
    <label :for="'value' + input.id" v-text="input.name" />
    <input type="checkbox" v-model="dynamic" />
    <template v-if="!dynamic">
      <input :id="'value' + input.id" v-model="currentValue" />
      <select v-model="type">
        <option :value="0">Integer</option>
        <option :value="1">String</option>
        <option :value="2">Float</option>
      </select>
    </template>
    <template v-if="dynamic">
      <select v-model="sourceTaskID">
        <option value="" selected disabled hidden>-</option>
        <option v-for="task in tasks">{{ task.id }}</option>
      </select>
      <input v-model="sourceName" />
    </template>
    <button class="mini green ui icon button" :disabled="!changed" @click="onUpdateClick">
      <i class="checkmark icon" />
    </button>
  </div>
</template>

<script>

export default {
  props: ['input', 'onUpdate', 'tasks'],
  data: () => ({
    currentValue: '',
    dynamic: false,
    sourceTaskID: null,
    sourceName: '',
    type: 0
  }),
  methods: {
    onUpdateClick () {
      if (this.dynamic) {
        this.onUpdate({ id: this.input.id, payload: { dynamic: this.dynamic, sourceTaskID: parseInt(this.sourceTaskID), sourceName: this.sourceName } })
      } else {
        this.onUpdate({ id: this.input.id, payload: { dynamic: this.dynamic, value: this.currentValueTyped, type: this.type } })
      }
    }
  },
  computed: {
    changed () {
      if (this.dynamic !== this.input.dynamic) { return true }
      if (this.dynamic) {
        return this.input.sourceTaskID !== parseInt(this.sourceTaskID) || this.input.sourceName !== this.sourceName
      } else {
        return this.input.value !== this.currentValueTyped || this.input.type !== this.type
      }
    },
    currentValueTyped () {
      switch (this.type) {
        case 0:
          let i = parseInt(this.currentValue)
          return isNaN(i) ? null : i
        case 1:
          return this.currentValue
        case 2:
          i = parseFloat(this.currentValue)
          return isNaN(i) ? null : i
        default:
          return this.currentValue
      }
    }
  },
  created () {
    this.currentValue = this.input.value
    this.dynamic = this.input.dynamic
    this.sourceTaskID = this.input.sourceTaskID
    this.sourceName = this.input.sourceName
    this.type = this.input.type
  }
}
</script>

<style>
</style>
