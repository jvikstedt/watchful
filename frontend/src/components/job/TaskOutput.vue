<template>
  <div>
    {{ output.name }} - <span v-if="resultItem.output">{{ resultItem.output[output.name] | truncate(20) }}</span>
  </div>
</template>

<script>

export default {
  props: ['output', 'resultItem']
}
</script>

<style>
</style>
