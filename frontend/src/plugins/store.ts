import { defineStore } from 'pinia'
import { ChatMessage } from '@/models/chat'
import { Client } from '@/models/client'

export const configStore = defineStore('config', {
  state: () => {
    return {
      username: '',
      room: '',
    }
  },
  actions: {
    save() {
      localStorage.setItem('config', JSON.stringify(this.$state))
      console.log('config saved')
    },
    load() {
      const config = localStorage.getItem('config')
      if (config) {
        this.$patch(JSON.parse(config))
      }
      console.log('config loaded')
    },
  },
})

export const mainStore = defineStore('main', {
  state: () => {
    return {
      messages: Array<ChatMessage>(),
      users: Array<Client>(),
      sig: null as any as WebSocket,
    }
  },
})
