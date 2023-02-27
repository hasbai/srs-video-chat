<template>
  <main id="main">
    <Config></Config>
    <n-button type="success" @click="connectWS">Connect</n-button>
    <Chat></Chat>
  </main>
</template>


<script setup lang="ts">
import {onMounted} from "vue"
import {SrsRtcSignaling} from "@/websocket"
import Config from "@/components/Config.vue";
import {NButton} from "naive-ui";
import {configStore, mainStore} from "@/plugins/store";
import Chat from "@/components/Chat.vue";
import {Client} from "@/models/client";
import {ChatMessage} from "@/models/chat";

const config = configStore()
const store = mainStore()

const sig = new SrsRtcSignaling(config.room, config.username)
sig.onmessage = e => {
  console.log(e.data)
  const payload = JSON.parse(e.data)
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
store.sig = sig

async function connectWS() {
  sig.close()
  await sig.connect()
}

onMounted(async () => {
  console.log('mounted')
})

</script>

<style scoped>
</style>
