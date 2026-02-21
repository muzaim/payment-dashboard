import { useCookie } from '@/@core/composable/useCookie'

export const $api = async (url: string, options: any = {}) => {
  const token = useCookie('accessToken').value

  const res = await fetch(`${import.meta.env.VITE_API_URL}${url}`, {
    method: options.method || 'GET',
    headers: {
      'Content-Type': 'application/json',
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
      ...(options.headers || {}),
    },
    body: options.body ? JSON.stringify(options.body) : undefined,
  })

  if (!res.ok)
    throw await res.text()

  return res.json()
}
