import { defineStore } from 'pinia'
import { ChatMessage } from '@/models/chat'
import { SrsRtcSignaling } from '@/websocket'
import { Client } from '@/models/client'

export const configStore = defineStore('config', {
  state: () => {
    return {
      username: '',
      room: '',
    }
  },
  persist: true,
})

export const mainStore = defineStore('main', {
  state: () => {
    return {
      messages: Array<ChatMessage>(),
      users: Array<Client>(),
      sig: new SrsRtcSignaling('', ''),
    }
  },
})
