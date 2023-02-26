import { defineStore } from 'pinia'

export const configStore = defineStore('config', {
  state: () => {
    return {
      username: '',
      room: '',
    }
  },
  persist: true,
})
