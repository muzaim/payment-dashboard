<script setup lang="ts">
import { formatRupiah, paymentStatusChip } from '@/helpers/formatter'

defineProps<{
  modelValue: boolean
  payment: any
}>()

defineEmits<{
  (e: 'update:modelValue', value: boolean): void
  (e: 'submit', action: 'complete'): void
}>()
</script>

<template>
  <VDialog
    :model-value="modelValue"
    max-width="620"
    @update:model-value="$emit('update:modelValue', $event)"
  >
    <VCard>
      <div class="d-flex justify-space-between align-center pa-4 bg-background">
        <div class="d-flex align-center gap-2 me-4">
          <VIcon
            icon="bx-receipt"
            color="primary"
          />
          <h1 class="text-h6 font-weight-medium text-primary text-lg py-4">
            Review Payments
          </h1>
        </div>
      </div>

      <VCardText v-if="payment">
        <VRow dense>
          <VCol cols="6">
            <div class="text-caption text-medium-emphasis">
              Payment ID
            </div>
            <div class="font-weight-medium">
              {{ payment.payment_id }}
            </div>
          </VCol>

          <VCol cols="6">
            <div class="text-caption text-medium-emphasis">
              Merchant
            </div>
            <div class="font-weight-medium">
              {{ payment.merchant }}
            </div>
          </VCol>

          <VCol cols="6">
            <div class="text-caption text-medium-emphasis">
              Amount
            </div>
            <div class="font-weight-bold text-primary">
              Rp {{ formatRupiah(payment.amount) }}
            </div>
          </VCol>

          <VCol cols="6">
            <div class="text-caption text-medium-emphasis mb-1">
              Status
            </div>

            <VChip
              size="small"
              :color="paymentStatusChip(payment.status).color"
              variant="tonal"
            >
              {{ paymentStatusChip(payment.status).label }}
            </VChip>
          </VCol>
        </VRow>
      </VCardText>

      <VCardActions class="justify-end gap-2">
        <VBtn
          variant="tonal"
          @click="$emit('update:modelValue', false)"
        >
          Cancel
        </VBtn>

        <VBtn
          variant="tonal"

          color="success"
          @click="$emit('submit', 'complete')"
        >
          Complete
        </VBtn>
      </VCardActions>
    </VCard>
  </VDialog>
</template>
