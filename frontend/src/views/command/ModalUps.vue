<script setup>
import IllustrationWarning from '@images/illustrations/warning.png'
import { useRouter } from 'vue-router'

const props = defineProps({
  visible: {
    type: Boolean,
    default: false,
  },
  desc: {
    type: String,
    default: 'Token autentikasi Anda telah kedaluwarsa. Silakan login kembali untuk melanjutkan penggunaan aplikasi.',
  },
})

const emit = defineEmits(['update:visible'])

const updateModelValue = value => {
  emit('update:visible', value)
}

const router = useRouter()

const confirmModal = () => {
  LogoutGlobal(router)
  updateModelValue(false)
}
</script>

<template>
  <VDialog
    :model-value="props.visible"
    max-width="420"
    persistent
    @update:model-value="updateModelValue"
  >
    <VCard class="py-6 px-4">
      <VCardText class="text-center">
        <VImg
          :width="245"
          :src="IllustrationWarning"
          class="d-none d-sm-block  mx-auto"
        />

        <h2 class="text-h5 font-weight-medium mb-2">
          Sesi Telah Berakhir
        </h2>

        <p class="text-body-2 mb-6">
          {{ desc }}
        </p>

        <VBtn
          color="primary"
          class="px-6"
          @click="confirmModal"
        >
          Login Kembali
        </VBtn>
      </VCardText>
    </VCard>
  </VDialog>
</template>
