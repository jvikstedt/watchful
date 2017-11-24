<template>
  <div>
    <input v-model="inputName" />
    <select v-model="inputType">
      <option value="0">Integer</option>
      <option value="1">String</option>
      <option value="2">Float</option>
      <option value="3">Bool</option>
      <option value="900">Dynamic</option>
    </select>
    <button class="ui button" @click="inputCreate">Add input</button>
  </div>
</template>

<script>

export default {
  props: ['onInputAdd'],
  data: () => ({
    inputName: '',
    inputType: '0'
  }),
  methods: {
    inputCreate () {
      if (this.inputName.length > 1) {
        this.onInputAdd(this.inputName, parseInt(this.inputType))
      }
    }
  }
}
</script>

<style>
</style>
