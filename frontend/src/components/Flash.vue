<template>
  <transition name="fade">
    <div v-if="flash" class="flash-message">
      <div :class="'ui message ' + flash.status">
        <i class="close icon" @click="removeFlash"></i>
        <div class="header">
          {{ flash.header }}
        </div>
        <p v-if="flash.body">{{ flash.body }}</p>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  computed: {
    flash () {
      return this.$store.state.flash
    }
  },

  methods: {
    removeFlash (event) {
      this.$store.commit('clearFlash')
    }
  }
}
</script>
<style>
  .flash-message {
    margin-bottom: 1em;
  }
  .fade-enter-active, .fade-leave-active {
    transition: opacity .5s
  }
  .fade-enter, .fade-leave-to {
    opacity: 0
  }
</style>
