<script setup lang="ts">
type SummaryKey = 'pending' | 'processing' | 'completed' | 'failed'

defineProps<{
  summary: Record<SummaryKey, number>
}>()

const items = [
  {
    key: 'pending',
    title: 'Pending',
    color: 'warning',
    icon: 'bx-time',
  },
  {
    key: 'processing',
    title: 'Processing',
    color: 'info',
    icon: 'bx-loader',
  },
  {
    key: 'completed',
    title: 'Completed',
    color: 'success',
    icon: 'bx-check',
  },
  {
    key: 'failed',
    title: 'Failed',
    color: 'error',
    icon: 'bx-x',
  },
] as const
</script>

<template>
  <VCol cols="12">
    <VRow>
      <VCol
        v-for="item in items"
        :key="item.key"
        cols="12"
        md="3"
      >
        <VCard class="pa-4">
          <div class="d-flex align-center justify-space-between">
            <div>
              <div class="text-caption text-medium-emphasis">
                {{ item.title }}
              </div>

              <div class="text-h5 font-weight-bold mt-1">
                {{ summary[item.key] }}
              </div>
            </div>

            <VAvatar
              size="36"
              :color="item.color"
              variant="tonal"
            >
              <VIcon size="18">
                {{ item.icon }}
              </VIcon>
            </VAvatar>
          </div>
        </VCard>
      </VCol>
    </VRow>
  </VCol>
</template>
