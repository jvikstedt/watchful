<template>
  <div>
    <div class="ui card" v-for="(project, index) in projects">
      <div class="content">
        <div class="header">{{ project.name }}</div>
        <div class="description"></div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ['projects']
}
</script>

<style>
</style>
