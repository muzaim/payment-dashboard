<script setup lang="ts">
defineProps<{
  merchants: Array<{
    merchant: string
    total: number
  }>
}>()
</script>

<template>
  <VCard class="pa-4">
    <div class="text-subtitle-1 font-weight-medium mb-2 ">
      Top 3 Merchants
    </div>

    <VList density="compact">
      <VListItem
        v-for="(item, i) in merchants"
        :key="i"
      >
        <template #prepend>
          <VAvatar
            size="32"
            color="primary"
            variant="tonal"
          >
            {{ i + 1 }}
          </VAvatar>
        </template>

        <VListItemTitle>
          <b>  {{ item.merchant }}</b>
        </VListItemTitle>

        <VListItemSubtitle>
          {{ item.total }} transactions
        </VListItemSubtitle>

        <template #append>
          <VChip
            size="x-small"
            color="primary"
            variant="tonal"
          >
            {{ item.total }}
          </VChip>
        </template>
      </VListItem>
    </VList>
  </VCard>
</template>
