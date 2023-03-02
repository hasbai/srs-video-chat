<template>
  <n-button :type="buttonType" :disabled="buttonDisabled" @click="onClick">
    {{ buttonText }}
  </n-button>
</template>

<script setup lang="ts">

import {NButton, useMessage} from "naive-ui";
import {configStore, mainStore} from "@/plugins/store";
import {computed, ref} from "vue";
import {ChatMessage} from "@/models/chat";
import Websocket from "@/home/Websocket.vue";

const message = useMessage()

const config = configStore()
const store = mainStore()

let ws: Websocket

const CLOSED = 0
const CONNECTING = 1
const CONNECTED = 2
let status = ref(CLOSED)

const buttonType = computed(() => {
  switch (status.value) {
    case CLOSED:
      return 'success'
    case CONNECTING:
      return 'info'
    case CONNECTED:
      return 'error'
  }
})
const buttonDisabled = computed(() => {
  return status.value === CONNECTING
})
const buttonText = computed(() => {
  switch (status.value) {
    case CLOSED:
      return 'Connect'
    case CONNECTING:
      return 'Connecting'
    case CONNECTED:
      return 'Disconnect'
  }
})

function connect() {
  status.value = CONNECTING

  const url = window.location.origin.replace('http', 'ws') + '/ws'
  ws = new WebSocket(
      url + '?room=' + config.room + '&user=' + config.username
  )

  ws.onmessage = (e: MessageEvent) => {
    if (status.value === CONNECTING) {
      if (e.data === 'OK') {
        onconnect()
      } else {
        onclose('websocket connect failed: ' + e.data)
      }
      return
    }
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
        store.joinClient(payload.data)
        break
      case 'leave':
        store.leaveClient(payload.data.name)
        break
    }
  }

  ws.onclose = (e: CloseEvent) => {
    onclose('websocket closed')
  }

  ws.onopen = (e: Event) => {
    status.value = CONNECTING
  }

  ws.onerror = (e: Event) => {
    status.value = CLOSED
    console.log('websocket connect failed: unknown')
    message.error('websocket connect failed: unknown')
  }

  store.sig = ws
}

function onclose(msg = '') {
  if (status.value === CLOSED) return
  status.value = CLOSED
  if (msg) {
    console.log(msg)
    message.error(msg)
  }
  store.clients = []
}

function onconnect() {
  status.value = CONNECTED
  console.log('websocket connected')
  message.success('websocket connected')
  store.joinClient({name: config.username})
}

function close() {
  if (ws) {
    ws.onclose?.(new CloseEvent('close'))
    ws.close()
  }
}

function onClick() {
  switch (status.value) {
    case CLOSED:
      connect()
      break
    case CONNECTED:
      close()
      break
  }
}

connect()

</script>

<style scoped>

</style>