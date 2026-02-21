<script setup lang="ts">
import { $api } from '@/utils/api'
import AnalyticsCongratulations from '@/views/dashboard/AnalyticsCongratulations.vue'
import PaymentSummaryCards from '@/views/dashboard/SummaryPayment.vue'
import TopMerchantCards from '@/views/dashboard/TopMerchantsCard.vue'

const isLoading = ref(false)

const summary = ref({
  pending: 0,
  processing: 0,
  completed: 0,
  failed: 0,
})

const topMerchants = ref([])

const fetchDataSummary = async () => {
  try {
    isLoading.value = true

    const res = await $api('/dashboard/v1/payments/summary', {
      method: 'GET',
    })

    summary.value = res.summary
  }
  finally {
    isLoading.value = false
  }
}

const fetchDataTopMerchants = async () => {
  try {
    isLoading.value = true

    const res = await $api('/dashboard/v1/payments/top-merchants', {
      method: 'GET',
    })

    topMerchants.value = res.data
  }
  finally {
    isLoading.value = false
  }
}

onMounted(async () => {
  await fetchDataSummary()
  await fetchDataTopMerchants()
})
</script>

<template>
  <VRow>
    <PaymentSummaryCards :summary="summary" />

    <VRow>
      <VCol
        cols="12"
        md="6"
      >
        <AnalyticsCongratulations />
      </VCol>

      <VCol
        cols="4"
        md="6"
      >
        <TopMerchantCards :merchants="topMerchants" />
      </VCol>
    </VRow>
  </VRow>
</template>
