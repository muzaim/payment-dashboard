<script setup lang="ts">
import { useCookie } from '@/@core/composable/useCookie'
import avatar1 from '@images/avatars/avatar-1.png'

const router = useRouter()

const userData = useCookie<any>('userData')
const accessToken = useCookie<any>('accessToken')
const userAbilityRules = useCookie<any>('userAbilityRules')

const logout = async () => {
  userData.value = null
  accessToken.value = null
  userAbilityRules.value = null

  await router.push('/login')

  router.go(0)
}
</script>

<template>
  <VBadge
    dot
    location="bottom right"
    offset-x="3"
    offset-y="3"
    color="success"
    bordered
  >
    <VAvatar
      class="cursor-pointer"
      color="primary"
      variant="tonal"
    >
      <VImg :src="avatar1" />

      <!-- SECTION Menu -->
      <VMenu
        activator="parent"
        width="230"
        location="bottom end"
        offset="14px"
      >
        <VList>
          <!-- 👉 User Avatar & Name -->
          <VListItem>
            <template #prepend>
              <VListItemAction start>
                <VBadge
                  dot
                  location="bottom right"
                  offset-x="3"
                  offset-y="3"
                  color="success"
                >
                  <VAvatar
                    color="primary"
                    variant="tonal"
                  >
                    <VImg :src="avatar1" />
                  </VAvatar>
                </VBadge>
              </VListItemAction>
            </template>

            <div class="text-left me-2">
              <div class="text-body-2">
                <span class="font-weight-medium text-primary">
                  {{ userData.email }}
                </span>
              </div>
            </div>
            <VListItemSubtitle>{{ userData.role }}</VListItemSubtitle>
          </VListItem>
          <VDivider class="my-2" />

          <!-- 👉 Logout -->
          <VListItem @click="logout">
            <template #prepend>
              <VIcon
                class="me-2"
                icon="bx-log-out"
                size="22"
              />
            </template>

            <VListItemTitle>Logout</VListItemTitle>
          </VListItem>
        </VList>
      </VMenu>
    </VAvatar>
  </VBadge>
</template>
