<template>
  <div>
    {{ output.name }} - {{ resultItem.output[output.name] | truncate(20) }}
  </div>
</template>

<script>

export default {
  props: ['output', 'resultItem']
}
</script>

<style>
</style>
