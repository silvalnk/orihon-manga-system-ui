import { defineStore } from 'pinia'

export const useSession = defineStore('session', {
  state: () => ({
    lang: 'en' as 'en' | 'pt-br',
    busy: false,
  }),
  actions: {
    syncLang(lang: string) {
      if (lang === 'en' || lang === 'pt-br') this.lang = lang
    },
    setBusy(busy: boolean) {
      this.busy = busy
    },
  },
})
