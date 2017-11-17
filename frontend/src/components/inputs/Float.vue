<template>
  <span>
    Float
    <input :value="value" @input="onChange" />
  </span>
</template>

<script>

export default {
  props: ['value'],
  methods: {
    onChange ({ target: { value } }) {
      const asFloat = parseFloat(value)
      if (isNaN(asFloat)) {
        this.$emit('input', null)
      } else {
        this.$emit('input', asFloat)
      }
    }
  }
}
</script>

<style>
</style>
