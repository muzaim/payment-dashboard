<script setup lang="ts">
import IllustrationDocument from '@images/illustrations/document.png';

const props = defineProps<{
  modelValue: boolean
  previewUrl: string
}>()

const emit = defineEmits(['update:modelValue'])

const isOpen = computed({
  get: () => props.modelValue,
  set: val => emit('update:modelValue', val),
})

const isImage = computed(() => /\.(jpg|jpeg|png|webp|gif)$/i.test(props.previewUrl))
const isPdf = computed(() => /\.pdf$/i.test(props.previewUrl))

const canPreview = computed(() => isImage.value || isPdf.value)
</script>

<template>
  <VDialog
    v-model="isOpen"
    max-width="90vw"
    scrollable
  >
    <VCard>
      <VCardText class="d-flex align-center overflow-hidden pa-6">
        <VCardTitle class="text-h5">
          Preview Dokumen
        </VCardTitle>
        <VSpacer />
        <VBtn
          type="button"
          class="me-4"
          @click="isOpen = false"
        >
          Back
        </VBtn>
      </VCardText>

      <VCardText
        class="overflow-auto"
        style="max-height: 85vh;"
      >
        <div
          v-if="canPreview"
          class="text-center"
        >
          <img
            v-if="isImage"
            :src="previewUrl"
            alt="Preview"
            class="max-w-full h-auto mx-auto d-block rounded"
          >
          <iframe
            v-else
            :src="previewUrl"
            width="100%"
            style="border: none; height: 80vh;"
          />
        </div>

        <div
          v-else
          class="text-center py-10"
        >
          <VImg
            :width="345"
            :src="IllustrationDocument"
            class="d-none d-sm-block mx-auto"
          />
          <p class="text-subtitle-1 text-medium-emphasis">
            File ini tidak dapat ditampilkan dalam preview.
          </p>
        </div>
      </VCardText>
    </VCard>
  </VDialog>
</template>
