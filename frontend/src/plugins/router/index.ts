import type { App } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import { routes } from './routes'
import { useCookie } from '@/@core/composable/useCookie'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

router.beforeEach((to, from, next) => {
  const token = useCookie('accessToken').value

  if (!token && to.name !== 'login')
    return next({ name: 'login' })

  if (token && to.name === 'login')
    return next({ path: '/' })

  next()
})

export { router }

export default function (app: App) {
  app.use(router)
}
