<template>
  <main id="main">
    <Config></Config>
    <Websocket></Websocket>
    <div class="players m-2 flex flex-row">
      <Player v-for="(client, i) in store.clients" :key="i" :client="client"></Player>
    </div>
    <Chat></Chat>
  </main>
</template>

<script setup lang="ts">
import Config from "@/home/Config.vue";
import Chat from "@/home/chat/Chat.vue";
import Websocket from "@/home/Websocket.vue";
import Player from "@/home/Player.vue";
import {mainStore} from "@/plugins/store";

const store = mainStore();
</script>

<style scoped>
.players {
}

.players > * {
}
</style>
