<script setup lang="ts">
import { VForm } from 'vuetify/components/VForm'
import { useCookie } from '@/@core/composable/useCookie'
import { $api } from '@/utils/api'
import authV1BottomShape from '@images/svg/auth-v1-bottom-shape.svg?url'
import authV1TopShape from '@images/svg/auth-v1-top-shape.svg?url'

const isPasswordVisible = ref(false)
const isSnackbarVisible = ref(false)

const snackBarText = ref('')
const snackBarColour = ref('')

const refVForm = ref<VForm>()
const isLoading = ref(false)
const route = useRoute()
const router = useRouter()

const credentials = ref({
  email: '',
  password: '',
  remember: false,
})

const login = async () => {
  isLoading.value = true
  try {
    const res = await $api('/dashboard/v1/auth/login', {
      method: 'POST',
      body: {
        email: credentials.value.email,
        password: credentials.value.password,
      },
    })

    useCookie('userData').value = JSON.stringify({
      email: res.email,
      role: res.role,
    })
    useCookie('accessToken').value = res.token

    await nextTick(() => {
      router.replace(route.query.to ? String(route.query.to) : '/')
    })
  }
  catch (err) {
    console.error(err)
    snackBarText.value = 'Failed, Please check your username and password'
    snackBarColour.value = 'error'
    isSnackbarVisible.value = true
    isLoading.value = false

    return false
  }
}

const onSubmit = () => {
  refVForm.value?.validate()
    .then(({ valid: isValid }) => {
      if (isValid)
        login()
    })
}
</script>

<template>
  <div class="auth-wrapper d-flex align-center justify-center pa-4">
    <div class="position-relative my-sm-16">
      <!-- 👉 Top shape -->
      <VImg
        :src="authV1TopShape"
        class="text-primary auth-v1-top-shape d-none d-sm-block"
      />

      <!-- 👉 Bottom shape -->
      <VImg
        :src="authV1BottomShape"
        class="text-primary auth-v1-bottom-shape d-none d-sm-block"
      />

      <!-- 👉 Auth Card -->
      <VCard
        class="auth-card"
        max-width="460"
        :class="$vuetify.display.smAndUp ? 'pa-6' : 'pa-0'"
      >
        <VCardItem class="justify-center">
          <RouterLink
            to="/"
            class="app-logo d-flex align-center gap-2"
          >
            <VImg
              src="/logo/durianpay.png"
              width="32"
              contain
            />

            <h1 class="app-logo-title font-weight-bold">
              Durian<span class="text-primary">pay</span>
            </h1>
          </RouterLink>
        </VCardItem>

        <VCardText>
          <h4 class="text-h4 mb-1">
            Welcome to <span class="app-logo-title font-weight-bold">
              durian<span class="text-primary">pay</span>
            </span>! 👋🏻
          </h4>
          <p class="mb-0">
            Please sign-in to your account and start the adventure
          </p>
        </VCardText>

        <VCardText>
          <VForm
            ref="refVForm"
            @submit.prevent="onSubmit"
          >
            <VRow>
              <!-- email -->
              <VCol cols="12">
                <VTextField
                  v-model="credentials.email"
                  autofocus
                  label="Email"
                  type="email"
                  placeholder="johndoe@email.com"
                />
              </VCol>

              <!-- password -->
              <VCol cols="12">
                <VTextField
                  v-model="credentials.password"
                  label="Password"
                  placeholder="············"
                  :type="isPasswordVisible ? 'text' : 'password'"
                  autocomplete="password"
                  :append-inner-icon="isPasswordVisible ? 'bx-hide' : 'bx-show'"
                  @click:append-inner="isPasswordVisible = !isPasswordVisible"
                />

                <!-- login button -->
                <VBtn
                  block
                  type="submit"
                  class="mt-5"
                >
                  Login
                </VBtn>
              </VCol>
            </VRow>
          </VForm>
        </VCardText>
      </VCard>
    </div>
  </div>

  <!-- Snackbar -->
  <VSnackbar
    v-model="isSnackbarVisible"
    transition="scroll-y-reverse-transition"
    location="top end"
    :timeout="2000"
    :color="snackBarColour"
  >
    {{ snackBarText }}
  </VSnackbar>
</template>

<style lang="scss">
@use "@core/scss/template/pages/page-auth";
</style>
