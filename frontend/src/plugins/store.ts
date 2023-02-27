import { defineStore } from 'pinia'
import { Chat } from '@/models/chat'
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
      messages: Array<Chat>(),
      users: Array<Client>(),
      sig: new SrsRtcSignaling('', ''),
    }
  },
})
