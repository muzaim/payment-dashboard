<script lang="ts" setup>
import VerticalNavSectionTitle from '@/@layouts/components/VerticalNavSectionTitle.vue'
import VerticalNavLink from '@layouts/components/VerticalNavLink.vue'
</script>

<template>
  <!-- 👉 Apps & Pages -->
  <VerticalNavSectionTitle
    :item="{
      heading: 'Dashboard',
    }"
  />
  <VerticalNavLink
    :item="{
      title: 'Dashboard',
      to: '/dashboard',
      icon: 'bx-home-smile',

    }"
  />
  <!-- 👉 Payments -->
  <VerticalNavSectionTitle
    :item="{
      heading: 'Payments',
    }"
  />
  <VerticalNavLink
    :item="{
      title: 'Payments',
      to: '/dashboard/payments',
      icon: 'bx-credit-card',
    }"
  />
</template>
