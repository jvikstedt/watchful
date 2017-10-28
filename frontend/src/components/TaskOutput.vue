<template>
  <div>
    {{ output.name }} - {{ resultItem.output[output.name] }}
  </div>
</template>

<script>

export default {
  props: ['output', 'resultItem']
}
</script>

<style>
</style>
