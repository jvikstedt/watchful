<template>
  <span>
    Integer
    <input :value="value" @input="onChange" />
  </span>
</template>

<script>

export default {
  props: ['value'],
  methods: {
    onChange ({ target: { value } }) {
      const asInt = parseInt(value)
      if (isNaN(asInt)) {
        this.$emit('input', null)
      } else {
        this.$emit('input', asInt)
      }
    }
  }
}
</script>

<style>
</style>
