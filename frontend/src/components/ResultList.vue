<template>
  <transition-group name="list" tag="div" class="ui middle aligned divided list animated">
    <router-link :to="'/results/' + result.id" class="item" v-for="result in orderedResults" :key="result.id">
      {{ result.createdAt }}
      <i :class="resultStatusClasses(result)" :title="result.status" style="float: right"></i>
      <i v-if="result.testRun" title="Test run" class="large icon user" style="float: right"/>
    </router-link>
  </transition-group>
</template>

<script>
import _ from 'lodash'

export default {
  props: ['results'],
  methods: {
    resultStatusClasses (result) {
      return [
        'large',
        'icon',
        { smile: result.status === 'success' },
        { green: result.status === 'success' },
        { frown: result.status === 'error' },
        { red: result.status === 'error' }
      ]
    }
  },
  computed: {
    orderedResults: function () {
      return _.orderBy(this.results, 'createdAt', 'desc')
    }
  }
}
</script>

<style>
  div.list {
    cursor: pointer;
  }
  .list-enter-active, .list-leave-active {
    transition: all 4s;
  }
  .list-enter, .list-leave-to {
    background: yellow;
  }
</style>
