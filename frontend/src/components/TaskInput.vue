<template>
  <div>
    <label :for="'value' + input.id" v-text="input.name" />
    <input :id="'value' + input.id" v-model="currentValue" />
    <button class="mini green ui icon button" :disabled="!changed" @click="onUpdateClick">
      <i class="checkmark icon" />
    </button>
  </div>
</template>

<script>

export default {
  props: ['input', 'onUpdate'],
  data: () => ({
    currentValue: ''
  }),
  methods: {
    onUpdateClick () {
      this.onUpdate({ value: this.currentValue, id: this.input.id })
    }
  },
  computed: {
    changed () {
      return this.input.value !== this.currentValue
    }
  },
  created () {
    this.currentValue = this.input.value
  }
}
</script>

<style>
</style>
