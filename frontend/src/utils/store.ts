export const useAlertStore = defineStore('snack', {
  state: () => ({
    isALert: false,
    message: '',
    autoClose: true,
    duration: 3000,
    type: 'info',
  }),
  actions: {
    setIsALert(value: boolean) {
      this.isALert = value
    },
    setMessageALert(value: string) {
      this.message = value
    },
    setAutoClose(value: boolean) {
      this.autoClose = value
    },
    setDuration(value: number) {
      this.duration = value
    },

    setType(value: string) {
      this.type = value
    },
  },
})
