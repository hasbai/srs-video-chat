<template>
  <main id="main">
    <Config></Config>
    <n-button type="success" @click="connectWS" :disabled="sig.connected">
      {{ sig.connected ? 'Connected' : 'Connect' }}
    </n-button>
    <Chat></Chat>
  </main>
</template>


<script setup lang="ts">
import {onMounted, reactive} from "vue"
import {SrsRtcSignaling} from "@/websocket"
import Config from "@/components/Config.vue";
import {NButton, useMessage} from "naive-ui";
import {configStore, mainStore} from "@/plugins/store";
import Chat from "@/components/Chat.vue";
import {Client} from "@/models/client";
import {ChatMessage} from "@/models/chat";

const message = useMessage()

const config = configStore()
const store = mainStore()

const sig = reactive(new SrsRtcSignaling(config.room, config.username))
sig.onmessage = e => {
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
sig.onclose = e => {
  message.warning('websocket closed')
}

store.sig = sig

async function connectWS() {
  sig.close()
  sig.room = config.room
  sig.username = config.username
  try {
    await sig.connect()
  } catch (e) {
    message.error('websocket connect failed')
  }
}

onMounted(async () => {
  console.log('mounted')
  await connectWS()
})

</script>

<style scoped>
</style>
