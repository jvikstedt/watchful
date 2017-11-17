<template>
  <div>
    {{ input.name }}
    <i class="close icon" @click="onInputDelete(input.id)"></i>
    <component :is="getComponent" v-model="currentValue" />
    <button class="mini green ui icon button" :disabled="!isChanged" @click="onUpdateClick">
      <i class="checkmark icon" />
    </button>
  </div>
</template>

<script>
import IntegerInput from '@/components/inputs/Integer'
import StringInput from '@/components/inputs/String'
import FloatInput from '@/components/inputs/Float'
import BoolInput from '@/components/inputs/Bool'

export default {
  props: ['input', 'onInputUpdate', 'onInputDelete'],
  data: () => ({
    currentValue: null
  }),
  methods: {
    onUpdateClick () {
      if (this.currentValue) {
        this.onInputUpdate({ value: { type: this.input.type, val: this.currentValue } })
      } else {
        this.onInputUpdate({ value: null })
      }
    }
  },
  computed: {
    isChanged () {
      return this.currentValue !== (this.input.value ? this.input.value.val : null)
    },
    getComponent () {
      switch (this.input.type) {
        case 0:
          return IntegerInput
        case 1:
          return StringInput
        case 2:
          return FloatInput
        case 3:
          return BoolInput
      }
    }
  },
  created () {
    if (this.input.value) {
      this.currentValue = this.input.value.val
    }
  },
  components: {
    IntegerInput,
    StringInput,
    FloatInput,
    BoolInput
  }
}
</script>

<style>
</style>
