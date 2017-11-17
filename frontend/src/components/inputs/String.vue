<template>
  <span>
    String
    <input :value="value" @input="onChange" />
  </span>
</template>

<script>

export default {
  props: ['value'],
  methods: {
    onChange ({ target: { value } }) {
      if (value.length > 0) {
        this.$emit('input', value)
      } else {
        this.$emit('input', null)
      }
    }
  }
}
</script>

<style>
</style>
