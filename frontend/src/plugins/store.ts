import { defineStore } from 'pinia'
import { ChatMessage } from '@/models/chat'
import { Client } from '@/models/client'

export const configStore = defineStore('config', {
  state: () => {
    return {
      username: '',
      room: '',
      webrtcURL: '',
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
      clients: Array<Client>(),
      sig: null as any as WebSocket,
    }
  },
  actions: {
    joinClient(client: object) {
      this.clients.push(client as Client)
    },
    leaveClient(name: string) {
      this.clients = this.clients.filter((c) => c.name !== name)
    },
  },
})
