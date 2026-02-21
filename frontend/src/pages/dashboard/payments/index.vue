<script setup lang="ts">
import { computed } from 'vue'
import { useCookie } from '@/@core/composable/useCookie'
import { formatIndoTime, formatRupiah, paymentStatusChip } from '@/helpers/formatter'
import { $api } from '@/utils/api'
import ReviewModal from '@/views/dashboard/ReviewModal.vue'

const searchQuery = ref('')
const resData: any = ref([])

const totalData = ref(0)
const itemsPerPage = ref(10)
const page = ref(1)
const isLoading = ref(false)
const userData = useCookie<any>('userData')

const reviewDialog = ref(false)
const selectedPayment = ref<any>(null)

const selectedStatus = ref<string | null>(null)
const selectedSort = ref<string | ''>(null)

const snackbar = ref(false)
const snackbarText = ref('')
const snackbarColor = ref<'success' | 'error'>('success')

const headers = computed(() => {
  const base = [
    { title: 'No', key: 'index', sortable: false },
    { title: 'Payment ID', key: 'payment_id', sortable: false },
    { title: 'Merchant Name', key: 'merchant', sortable: false },
    { title: 'Amount', key: 'amount', sortable: false },
    { title: 'Created', key: 'created_at', sortable: false },
    { title: 'Status', key: 'status', sortable: false },
  ]

  if (userData.value.role === 'operation')
    base.push({ title: 'Action', key: 'action', sortable: false })

  return base
})

const fetchData = async () => {
  try {
    isLoading.value = true

    const params = new URLSearchParams({
      page: page.value.toString(),
      limit: itemsPerPage.value.toString(),
      sort: selectedSort.value,
    })

    if (selectedStatus.value)
      params.append('status', selectedStatus.value)

    const res = await $api(`/dashboard/v1/payments?${params.toString()}`, {
      method: 'GET',
    })

    resData.value = (res.payments ?? []).map((item: any, index: number) => ({
      ...item,
      index: (page.value - 1) * itemsPerPage.value + index + 1,
    }))

    totalData.value = res.meta.total
  }
  finally {
    isLoading.value = false
  }
}

const submitReview = async () => {
  try {
    isLoading.value = true

    await $api(
      `/dashboard/v1/payment/${selectedPayment.value.payment_id}/review`,
      {
        method: 'PUT',
      },
    )

    snackbarText.value = 'Payment berhasil diproses'
    snackbarColor.value = 'success'
    snackbar.value = true

    await fetchData()
  }
  catch (error: any) {
    snackbarText.value
      = error?.response?._data?.message || 'Terjadi kesalahan'
    snackbarColor.value = 'error'
    snackbar.value = true
  }
  finally {
    isLoading.value = false
    reviewDialog.value = false
  }
}

const totalPages = computed(() =>
  Math.ceil(totalData.value / itemsPerPage.value),
)

const statusOptions = [
  { title: 'All', value: null },
  { title: 'Pending', value: 'pending' },
  { title: 'Processing', value: 'processing' },
  { title: 'Completed', value: 'completed' },
  { title: 'Failed', value: 'failed' },
]

const sortOption = [
  { title: 'Default', value: null },
  { title: 'Latest', value: 'latest' },
  { title: 'Earliest', value: 'earliest' },
]

watch(
  [page, itemsPerPage, selectedSort, searchQuery, selectedStatus],
  ([, , , , status], [, , , , prevStatus]) => {
    if (status !== prevStatus)
      page.value = 1

    fetchData()
  },
  { immediate: true },
)
</script>

<template>
  <section>
    <VCard>
      <VCardText class="d-flex flex-wrap pt-0 gap-4 mt-5">
        <div class="me-3 d-flex gap-3">
          <VSelect
            :model-value="itemsPerPage"
            :items="[
              { value: 10, title: '10' },
              { value: 25, title: '25' },
              { value: 50, title: '50' },
              { value: 100, title: '100' },
              { value: 99999999, title: 'All' },
            ]"
            style="inline-size: 6.25rem;"
            @update:model-value="itemsPerPage = $event"
          />
        </div>
        <VSpacer />

        <div class="app-user-search-filter d-flex align-center flex-wrap gap-4">
          <div style="inline-size: 10rem;">
            <VSelect
              v-model="selectedSort"
              :items="sortOption"
              item-title="title"
              item-value="value"
              label="Sort by"
              clearable
              class="mb-4"
              style="max-width: 200px"
            />
          </div>
        </div>

        <div class="app-user-search-filter d-flex align-center flex-wrap gap-4">
          <div style="inline-size: 10rem;">
            <VSelect
              v-model="selectedStatus"
              :items="statusOptions"
              item-title="title"
              item-value="value"
              label="Status"
              clearable
              class="mb-4"
              style="max-width: 200px"
            />
          </div>
        </div>
      </VCardText>

      <VDivider />

      <VDataTable
        v-model:items-per-page="itemsPerPage"
        :items="resData"
        :items-length="totalData"
        :headers="headers"
        class="text-no-wrap"
        :loading="isLoading"
      >
        <template #item.payment_id="{ item }">
          <VChip
            rounded="pill"
            size="small"
            color="primary"
            variant="elevated"
          >
            {{ (item as any).payment_id }}
          </VChip>
        </template>

        <template #item.status="{ item }">
          <VChip
            rounded="pill"
            size="small"
            :color="paymentStatusChip((item as any).status).color"
            variant="elevated"
          >
            {{ paymentStatusChip((item as any).status).label }}
          </VChip>
        </template>

        <template #item.amount="{ item }">
          Rp {{ formatRupiah((item as any).amount) }}
        </template>

        <template #item.created_at="{ item }">
          {{ formatIndoTime((item as any).created_at) }}
        </template>

        <template #item.action="{ item }">
          <VTooltip text="Review Payment">
            <template #activator="{ props }">
              <VBtn
                v-bind="props"
                icon
                size="small"
                color="primary"
                variant="tonal"
                :disabled="['completed', 'failed'].includes((item as any).status)"

                @click="selectedPayment = item as any; reviewDialog = true"
              >
                <VIcon size="18">
                  bx-search
                </VIcon>
              </VBtn>
            </template>
          </VTooltip>
        </template>

        <template #bottom>
          <VPagination
            v-model="page"
            :length="totalPages"
            :total-visible="5"
            show-first-last-page
          />
        </template>
      </VDataTable>
    </VCard>

    <ReviewModal
      v-model="reviewDialog"
      :payment="selectedPayment"
      @submit="submitReview"
    />

    <VSnackbar
      v-model="snackbar"
      :color="snackbarColor"
      transition="scroll-y-reverse-transition"
      location="top end"
      timeout="3000"
    >
      {{ snackbarText }}
    </VSnackbar>
  </section>
</template>
