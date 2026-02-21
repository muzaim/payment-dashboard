import type { Router } from 'vue-router'
import { useCookie } from '@/@core/composable/useCookie'
import { canNavigate } from '@layouts/plugins/casl'

export const setupGuards = (router: Router) => {
  router.beforeEach(to => {
    if (to.meta.public)
      return

    const isLoggedIn = !!(useCookie('userData').value && useCookie('accessToken').value)

    if (to.meta.unauthenticatedOnly) {
      if (isLoggedIn)
        return '/'
      else
        return
    }

    if (!isLoggedIn) {
      return {
        name: 'login',
        query: {
          ...to.query,
          to: to.fullPath !== '/' ? to.fullPath : undefined,
        },
      }
    }

    const allowedRoutes: any = useCookie('allowedRoutes').value || '[]'

    const routeName = to.name as string

    let isAllowed = allowedRoutes.includes(routeName)

    if (!isAllowed) {
      isAllowed = allowedRoutes.some((prefix: string) => {
        return routeName === prefix || routeName.startsWith(`${prefix}-`)
      })
    }

    if (!isAllowed)
      return { name: 'not-authorized' }

    if (!canNavigate(to))
      return { name: 'not-authorized' }
  })
}
