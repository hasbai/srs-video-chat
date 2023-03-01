<template>
  <n-button :type="connected ? 'error' : 'success'" @click="onClick">
    {{ connected ? 'Disconnect' : 'Connect' }}
  </n-button>
</template>

<script setup lang="ts">

import {NButton, useMessage} from "naive-ui";
import {configStore, mainStore} from "@/plugins/store";
import {ref} from "vue";
import {ChatMessage} from "@/models/chat";
import {Client} from "@/models/client";
import Websocket from "@/home/Websocket.vue";

const message = useMessage()

const config = configStore()
const store = mainStore()

let ws: Websocket
let connected = ref(false)

function connect(): Promise<Event> {
  const url = window.location.origin.replace('http', 'ws') + '/ws'
  ws = new WebSocket(
      url + '?room=' + config.room + '&user=' + config.username
  )

  ws.onmessage = (e: MessageEvent) => {
    let payload
    try {
      payload = JSON.parse(e.data)
    } catch (err) {
      console.log(e.data)
      return
    }
    switch (payload.event) {
      case 'chat':
        store.messages.push(Object.assign(new ChatMessage(), payload.data))
        break
      case 'join':
        store.users.push(payload.data)
        break
      case 'leave':
        const client = payload.data as Client
        store.users = store.users.filter(u => u.name !== client.name)
        break
    }
  }
  ws.onclose = (e: CloseEvent) => {
    if (!connected.value) return
    connected.value = false
    console.log('websocket closed')
    message.warning('websocket closed')
  }

  store.sig = ws

  return new Promise((resolve, reject) => {
    ws.onopen = (e: Event) => {
      connected.value = true
      console.log('websocket connected')
      message.success('websocket connected')
      resolve(e)
    }
    ws.onerror = (e: Event) => {
      connected.value = false
      console.log('websocket error')
      message.error('websocket error')
      reject(e)
    }
  })
}

function close() {
  if (ws) {
    ws.onclose?.(new CloseEvent('close'))
    ws.close()
  }
}

async function connectWS() {
  try {
    await connect()
  } catch (e) {
    message.error('websocket connect failed')
  }
}

function onClick() {
  if (connected.value) {
    close()
  } else {
    connectWS()
  }
}

connectWS()

</script>

<style scoped>

</style>