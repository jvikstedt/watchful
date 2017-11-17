<template>
  <span>
    Boolean
    <select :value="value" @change="onChange">
      <option value="">Not selected</option>
      <option value="true">True</option>
      <option value="false">False</option>
    </select>
  </span>
</template>

<script>

export default {
  props: ['value'],
  methods: {
    onChange (e) {
      if (e.target.value === 'true') {
        this.$emit('input', true)
      } else if (e.target.value === 'false') {
        this.$emit('input', false)
      } else {
        this.$emit('input', null)
      }
    }
  }
}
</script>

<style>
</style>
