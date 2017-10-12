<template>
  <div id="app">
    <div class="ui fixed inverted menu">
      <div class="ui container">
        <router-link to="/" class="header item">
          Watchful
        </router-link>
        <router-link to="/" class="item">Home</router-link>
      </div>
    </div>

    <div class="ui main text container">
      <Flash />
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import Flash from '@/components/Flash'

export default {
  name: 'app',

  components: {
    Flash
  }
}
</script>

<style>
  body {
    background-color: #FFFFFF;
  }
  .main.container {
    margin-top: 4em;
  }
</style>
