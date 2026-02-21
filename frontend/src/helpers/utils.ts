import { useAlertStore } from '@/utils/store'

export interface AlertConfig {
  message: string
  type?: 'success' | 'error' | 'info' | 'warning'
  autoClose?: boolean
  duration?: number
}

export const configAlert = ({
  message,
  type = 'info', // Default "info"
  autoClose = true,
  duration = 3000, // Default 3 detik
}: AlertConfig) => {
  const alertStore = useAlertStore()

  alertStore.setMessageALert(message)
  alertStore.setType(type)
  alertStore.setAutoClose(autoClose)
  alertStore.setDuration(duration)
  alertStore.setIsALert(true)
}
