<template>
  <div>
    <label :for="'value' + input.id" v-text="input.name" />
    <input type="checkbox" v-model="dynamic" />
    <template v-if="!dynamic">
      <input :id="'value' + input.id" v-model="currentValue" />
      <template v-if="input.type === 4">
        <input type="radio" id="string" value="string" v-model="selectedType">
        <label for="string">string</label>
        <input type="radio" id="integer" value="integer" v-model="selectedType">
        <label for="integer">integer</label>
      </template>
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
    selectedType: 'string'
  }),
  methods: {
    onUpdateClick () {
      if (this.dynamic) {
        this.onUpdate({ id: this.input.id, payload: { dynamic: this.dynamic, sourceTaskID: parseInt(this.sourceTaskID), sourceName: this.sourceName } })
      } else {
        let value = this.currentValue
        if (this.input.type === 4 && this.selectedType === 'integer') {
          value = parseInt(value)
        }
        this.onUpdate({ id: this.input.id, payload: { dynamic: this.dynamic, value: value } })
      }
    }
  },
  computed: {
    changed () {
      if (this.dynamic !== this.input.dynamic) { return true }
      if (this.dynamic) {
        return this.input.sourceTaskID !== parseInt(this.sourceTaskID) || this.input.sourceName !== this.sourceName
      } else {
        return this.input.value !== this.currentValue
      }
    }
  },
  created () {
    this.currentValue = this.input.value
    this.dynamic = this.input.dynamic
    this.sourceTaskID = this.input.sourceTaskID
    this.sourceName = this.input.sourceName
  }
}
</script>

<style>
</style>
