<script setup lang="ts">
import { ref, watch } from 'vue'

const props = defineProps({
  show: Boolean, // Dikontrol dari App
  type: String,
  message: String,
  autoClose: Boolean,
  duration: Number,
})

const emit = defineEmits(['update:modelValue'])
const showAlert = ref(props.show)

// Watch perubahan modelValue agar selalu sinkron
watch(
  () => props.show,
  newValue => {
    showAlert.value = newValue

    if (newValue && props.autoClose) {
      setTimeout(() => {
        showAlert.value = false
        emit('update:modelValue', false) // Emit perubahan ke App
      }, props.duration)
    }
  },
)
</script>

<template>
  <Transition name="fade-slide">
    <VAlert
      v-if="show"
      :type="type"
      class="alert-container"
      border="bottom"
      closable
    >
      <template #text>
        {{ message }}
      </template>
    </VAlert>
  </Transition>
</template>

<style scoped>
.alert-container {
  position: fixed;
  top: 20px;
  left: 50%;
  transform: translateX(-50%);
  width: calc(100% - 48px);
  max-width: 600px;
  z-index: 9999;
  text-align: left;
  padding: 12px;
  border-radius: 8px;
}

/* Animasi masuk & keluar */
.fade-slide-enter-active,
.fade-slide-leave-active {
  transition: all 0.3s ease;
}

.fade-slide-enter-from {
  opacity: 0;
  transform: translate(-50%, -100%);
}

.fade-slide-leave-to {
  opacity: 0;
  transform: translate(-50%, -100%);
}
</style>
